# Golang

## Tìm hiểu về Graceful Shutdown, Graceful Shutdown trong Golang

### Graceful Shutdown

Hãy tưởng tượng rằng, chúng ta đang có một web service đang tiếp nhận yêu cầu (request) của client để truy xuất dữ liệu từ database, vì dữ liệu truy xuất lớn nên phản hồi (response) phải mất một thời gian mới  truy xuất xong. Trong khi đó anh em lại muốn tắt service đó đi để bảo trì hệ thống hoặc triển khai (deploy) mới, bằng thao tác kill ứng dụng web service đang chạy, có thể bằng câu lệnh stop docker, câu lệnh kill process bằng PID hay Ctril + C chúng ta vẫn thường hay dùng v.v. Ngay lập tức những yêu cầu mà service xử lý chưa xong bị buộc ngưng giữ chừng.

Ngoài ra những kết nối khác như kết nối với database không được kiểm soát và đóng lại lại đúng cách gây hao tổn tài nguyên của server. Những điều nói trên làm chúng ta phải dừng lại và suy nghĩ về nó.

Điều chúng ta muốn khi service bị buộc dừng thì nó sẽ:
 1. Không đón nhận những request mới.
 2. Xử lý và phản hổi những request cũ.
 3. Cuối cùng là đóng kết nối, sau đó mới dừng service hoàn toàn.

### Áp dụng Graceful Shutdown trong Golang

Để thiết lập Graceful Shutdown cho ứng dụng web service của chúng ta, anh em cần phải thực hiện các bước sau:
 1. Cho ứng dụng của chúng ta chạy background trong 1 Goroutine.
 2. Thiết lập một channel để lắng nghe tín hiệu dừng từ hệ điều hành, ở đây chúng ta lưu ý 2 tín hiệu (signal) là SIGINT (The interrupt signal là loại tính hiệu khi user nhấn ctrl C để kết thúc chương trình) và SIGTERM (The termination signal là loại tín hiệu khi một ứng dụng muốn dừng tiến trình của một ứng dụng khác, như từ một câu lệnh stop của Docker hoặc câu lệnh delete pods của Kubernetes)
 3. Thiết lập khoản thời gian timeout, để dừng hoàn toàn ứng dụng và đóng tất cả kết nối.
 
 ### Folder structure
 
 Chúng ta sẽ có folder structure như sau, tùy vào cách xây dựng structure của project anh em như thế nào anh em có thể refactor theo ý mình muốn.
 
 ![](https://lh5.googleusercontent.com/apcVHp1mjD2Uc8wagrPJacLYTDP4RgqjgW-NVvEz0a4MDJGIOWarioRiP-0BZ3EgR3hEQRuVTWCqcNWn-Nmd5brHyjpaDYwYtBxd5Qtf4XhnTdzAe0D973zId408AOuDczBomQmZ)
 
*** response.go***

File này chứa các hàm tiện ích hỗ trợ cho việc phản hồi cho người dùng kết quả dạng JSON.

```go
	package response

	import (
		"encoding/json"
		"net/http"
	)
	
	func ResponseWithError(response http.ResponseWriter, statusCode int, msg string){
			ResponseWithJSON(response, statusCode, map[string]string{
				"error": msg,
			})
	}
	
	func ResponseWithJSON(response http.ResponseWriter, statusCode int, data interface{}{
		result, _ := json.Marshal(data)
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(statusCode)
		response.Write(result)
	}
```

***handle.go***

Khi yêu cầu được gửi đến api/test-graceful-Shutdown thì hàm testGracefulShutDown bên trong  xử lý nhiều công việc, sau 10 giây mới trả kết quả về cho client.

```go
	package handler

	import (
		"log"
		"net/http"
		"time"

		"graceful-shutdown/pkg/utils/response"

		"github.com/gorilla/mux"
	)
	
	func testGracefulShutDown(res http.ResponseWriter, req *http.Request){
		time.Sleep(10 * time.Second)
		log.Println("testGracefulShutDown job completed")
		res.ResponseWithJSON(res,200,map[string]interface{}{"status":"completed"})
	}
	
	func New(r *mux.Router){
		r.HandleFunc("/test-graceful-shutdown", testGracefulShutDown).Methods(http.MethodGet)
	}
```

***router.go***

```go
	package router

	import (
		"graceful-shutdown/internal/handler"

		"github.com/gorilla/mux"
	)

	func New() *mux.Router {
		router := mux.NewRouter()
		apiV1Router := router.PathPrefix("/api/v1").Subrouter()
		handler.New(apiV1Router)
		return router
	}
```

***app.go***

Chúng ta đang thực hiện việc tạo ra một ứng dụng http server có hàm Start và Stop

```go
	package app

import (
	"context"
	"graceful-shutdown/internal/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	app struct {
		Server *http.Server
	}

	App interface {
			Start() error
			Stop(ctx context.Context) error
		}
	)

	const (
		ADDR = ":8081"
	)

	func New() App {
		router := mux.NewRouter()
		apiV1Router := router.PathPrefix("/api/v1").Subrouter()
		handler.New(apiV1Router)

		httpServer := &http.Server{
			Addr:    ADDR,
			Handler: router,
		}
		return app{
			Server: httpServer,
		}
	}

	func (a app) Start() error {
		log.Printf("Server is listening at %s", ADDR)
		return a.Server.ListenAndServe()
	}

	func (a app) Stop(ctx context.Context) error {
		return a.Server.Shutdown(ctx)
	}
```

***main.go***

```go
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"graceful-shutdown/internal/app"
)

func main(){
	timeWait := 15 * time.Second
	application := app.New()
	signChan := make(chan os.Signal,1)
	// 1. Cho ứng dụng của chúng ta chạy background trong 1 Goroutine
	go func() {
		if err := application.Start(); err != nil {
			log.Printf("%v",err.Error())
		}
	}()
	
	// 2. Thiết lập một channel để lắng nghe tín hiệu dừng từ hệ điều hành,
	// Ở đây chúng ta lưu ý 2 tín hiệu (signal) là SIGNINT và SIGNTERM
	signal.Notify(signal, os.Interrupt, syscall.SIGNTERM)
	<- signChan
	log.Println("Shutting Down")
	
	// 3. Thiết lập một khoản thời gian (Timeout) để dừng hoàn toàn ứng dụng và đóng tất cả kết nối.
	ctx, cancel := context.WithTimeout(context.Background(),timeWait)
	defer func(){
		log.Println("Close another connection")
		cancel()
	}()
	
	log.Println("Stop http server")
	if err := application.Stop(ctx); err == context.DeadlineExceeded {
		log.Print("Halted active connections")
	}
	close(signChan)
	log.Printf("Completed")x
}
```

##[ 12 Packages và Libraries của Go cực kì mạnh mẽ mà bạn cần phải biết](https://topdev.vn/blog/12-packages-va-libraries-cua-go-cuc-ki-manh-me-ma-ban-can-phai-biet " 12 Packages và Libraries của Go cực kì mạnh mẽ mà bạn cần phải biết")


