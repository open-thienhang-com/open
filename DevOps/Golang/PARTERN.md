# Design Pattern
References:
- [ GitHub - tmrts/go-patterns: Curated list of Go design patterns, recipes and idioms](https://github.com/tmrts/go-patterns " GitHub - tmrts/go-patterns: Curated list of Go design patterns, recipes and idioms")
- [ How important are Design Patterns really? - Stack Overflow](https://stackoverflow.com/a/978509 " How important are Design Patterns really? - Stack Overflow")
- [ Evaluating the GO Programming Language with Design Patterns PDF](https://ecs.victoria.ac.nz/foswiki/pub/Main/TechnicalReportSeries/ECSTR11-01.pdf " Evaluating the GO Programming Language with Design Patterns PDF")
- [Tutorials · Software adventures and thoughts](http://blog.ralch.com/tutorial/ "Tutorials · Software adventures and thoughts")
- [Software design pattern - Wikipedia](https://en.wikipedia.org/wiki/Software_design_pattern "Software design pattern - Wikipedia")
- [Guru design pattern](https://refactoring.guru/design-patterns "Guru design pattern")
## Design pattern là gì:
- Định nghĩa: Là giải pháp tổng quát cho những vấn đề mà lập trình viên hay gặp phải. Design pattern không thể được chuyển đổi trực tiếp thành code mà nó chỉ là một khuôn mẫu cho những vấn đề được giải quyết.
- Ưu điểm:
	- Code readability: Giúp chúng ta viết code dễ hiểu hơn bằng việc sử dụng những tên biến liên quan đến những gì chúng ta đang thực hiện.
	- Communication: Giúp cho các lập trình viên có thể dễ dàng tiếp cận trao đổi về giải pháp cho một vấn đề bằng việc sử dụng một design pattern bất kỳ.
	- Code reuse: Code được viết theo design pattern có thể được sử dụng lại nhiều lần.
- Nhược điểm:
	- More complexity: Làm phức tạp hoá quá trình code, khiến lập trình viên vất vả hơn khi phải suy nghĩ nên dùng mẫu thiết kế nào để giải quyết vấn đề của mình khi mà anh ta chưa thực sự viết được 1 dòng code nào trong khi deadlione đang cận kề.
	- Different variants: Có nhiều biến thể cho mỗi mẫu đối tượng khiến cho việc áp dụng các mẫu thiết kế đôi khi không đồng nhất nên gây bất đồng khi làm việc giữa các lập trình viên.
	
## Phân loại
 -  Creational pattern

|  Pattern |
| ------------ |
|  Singletone |
|  Builder |
|Factory Method|

 -  Behavioral pattern

|  Pattern |
| ------------ |
|  Observer |
|  Strategy |
|Itertor|
|State|

 -  Structural pattern
 
|  Pattern |
| ------------ |
|  Adapter |
|  Bridge |
|Composite|
|Proxy|

## Creational Patterns
### Singleton
Chỉ duy nhất một đối tượng của một type bất kỳ được khởi tạo trong suốt quá trình hoạt động của chương trình
- Các ứng dụng:
	- Chúng ta muốn sử dụng lại kết nối đến database khi truy vấn
	- Khi chúng ta mở một kết nối SSH đến một server để thực hiện một số công việc và chúng ta không muốn mở một kết nối khác mà sử dụng lại kết nối trước đó.
	- Khi chúng ta cần hạn chế sự truy cập đến một vài biến số, chúng ta sử dụng mẫu singleton là trung gian để truy cập biến này.
```go
package DB

var (
  once sync.Once
  singleton *db.Connection
)

func GetDBConnection(config *mysql.Config) *db.Connection {
  once.Do(func() {
    singleton  = mysql.Dial(config)
  })
  return singleton
}
```
- Lưu ý: là mẫu thiết kế phổ biến nhất nhưng đôi khi bị làm dụng. Nên tránh việc sử dụng mẫu singleton vì khiến cho việc test cụ thể kà việc tạo mock/stub trở nên khó khăn hơn.
### Builder
- Có thể tạo ra nhiều biến thể khác nhau cho cùng một đối tượng
- Ví dụ: Có 1 anh chàng nọ đang làm việc trong một team phát triển các hệ thống cho một ngân hàng và anh ta được giao nhiệm vụ cần tạo một kiểu dữ liệu BankAccount để chứa thông tin các khác đã mở tài khoản tại ngân hàng này.
```go
type BankAccount struct {
	ownerName string
	ownerIdentificationNumber unit64
	balance int64
}

func NewBankAccount(ownerName string, idNo unit64, balance int64){
	return &BankAccount{ownerName, idNo, balance}
}

```
Rất dễ dàng để mở một tài khoản
```go
var account BankAccount = NewBankAccount("Vuong",123456789,10000)
```
tuy nhiên đời không như mơ, Sếp lớn xỉ vả anh ta vì kiểu BankAccount thiếu chi nhánh ngân hàng và tỉ lệ lãi suất. Thế là anh ta lại chỉnh sửa lại.
```go
type BankAccount struct {
	ownerName string
	ownerIdentificationNumber uint64
	balance int64
	interestRate float64 // ti suat
	branch string // chi nhanh
}

func NewBankAccount(ownerName string, idNo uint64, balance int64, interestRate float64, branch string) {
	return &BankAccount{ownerName, idNo, balance, interestRate, branch}
}
```
Một người khác sử dụng code của anh ta viết để thực hiện task của mình và anh ta đã sử dụng như sau:
```go
var account BankAccount = NewBankAccount("Vuong",1000,123456789,0.8,"Sai Gon")
```
anh này đã vô tình đổi số CMND (123456789) ra sau số dư 1000 và khiến ngân hàng thất thoát tiền. Anh ta vô tình nhưng compiler không phát hiện ra vì cả 2 đều có dữ liệu unit64 và int64. Anh ta đã đổ lỗi cho người thiết kế kiểu dữ liệu BankAccount chính là anh chàng trước đó.
- Cách giải quyết được đặt ra:
	- Tạo các setter: Tuy nhiên người dùng đôi khi sẽ quên gọi các setter này
Anh chàng bị sếp trừ lương và quá uất ức anh sử dụng bí kíp design pattern Builder
```go
type bankAccount struct {
	ownerName	string
	identificationNumber unit64
	branch	string
	balance	int64
}
type BankAccount interface{
	WithDraw(amt unit64)
	Deposit(amt unit64)
	GetBalance() unit64
}

type BackAccountBuilder interface(){
	WithOwnerName(name string) BankAccountBuilder
	WithOwnerIdentity(identificationNumber unit64) BankAccountBuilder
	AtBranch(branch string) BankAccountBuilder
	OpeningBalance(balance unit64) BankAccountBuilder
	Build() BankAccount
}

func (acc *bankAccount) WithDraw(amt unit64){

}

func (acc *bankAccount) Deposit(amt unit64){

}

func (acc *bankAccount) GetBalance() unit64{
	return 0
}
func (acc *bankAccount) WithOwnerName(name string) BankAccountBuilder{
	acc.ownerName = name
	return acc
}
func (acc *bankAccount) WithOwnerIdentity(identificationNumber unit64) BankAccountBuilder{
	acc.identificationNumber = identificationNumber
	return acc
}
func (acc *bankAccount) AtBranch(branch string) BankAccountBuilder {
	acc.branch = branch
	return acc
}

func (acc *bankAccount) OpeningBalance(balance uint64) BankAccountBuilder {
	acc.balance = int64(balance)
	return acc
}

func (acc *bankAccount) Build() BankAccount {
	return acc
}

func NewBankAccountBuilder() BankAccountBuilder {
	return &bankAccount{}
}

func main {
	account := NewBankAccountBuilder().
				WithOwnerName("Vuong").
				WithOwnerIdentity(123456789).
				AtBranch("SaiGon").
				OpeningBalance(1000).Build()
	account.Deposit(10000)
	account.WithDraw(50000)
}
```
### Factory method
- Tạo một đối tượng mà không cần thiết chỉ ra một cách chính xác lớp nào sẽ được tạo bằng cách nhóm các đối tượng liên quan đến nhau và sử dụng 1 đối tượng trung gian để khởi tạo đối tượng cần tạo.
- Ví dụ chúng ta sẽ tạo các phương thức thanh toán cho 1 shop bằng factory method, hình thức thanh toán có thể bằng tiền mặt hoặc bằng thẻ debit
```go
	type PaymentMethod interface{
		Pay(amount float32) string
	}
	type PaymentType int
	const (
		Cash PaymentType = iota
		DebitCard
	)
	type CashPM struct{}
	type DebitCardPM struct{}
	
	func (c *CashPM) Pay(amount float32) string{
		return ""
	}
	
	func (c *DebitCardPM) Pay(amout float32) string{
		reuturn ""
	}
	
	func GetPaymentMethod(t PaymentType) PaymentMethod{
		switch t {
		cast Cash:
			return new(CashPM)
		cast DebitCard:
			return new(DebitCardPM)
		}
	}
	
	//usage
	payment := GetPaymentMethod(DebitCard)
	payment.Pay(20)
	payment := GetPaymentMethod(Cash)
	payment.Pay(50)
```

## Behavioural Pattern
### Observer
- Tạo mối liên hệ one-to-many giữa subject và các observer với nhau (chẳng hạn 1 subject sẽ có thuộc tính là một mảng bao gồm nhiều observer) nên khi trạng thái của subject thay đổi, tất cả subject này sẽ được thông báo và tự động cập nhật.
- Gồm 2 thành phần chính là Subject và Observer
- Sơ đồ:
[![](https://camo.githubusercontent.com/7a4737ac8f5415c24311b777c44bff2aea665d4ad024c636c2b91a0f10974aa7/68747470733a2f2f75706c6f61642e77696b696d656469612e6f72672f77696b6970656469612f636f6d6d6f6e732f7468756d622f612f61382f4f627365727665725f775f7570646174652e7376672f35303070782d4f627365727665725f775f7570646174652e7376672e706e67)](https://camo.githubusercontent.com/7a4737ac8f5415c24311b777c44bff2aea665d4ad024c636c2b91a0f10974aa7/68747470733a2f2f75706c6f61642e77696b696d656469612e6f72672f77696b6970656469612f636f6d6d6f6e732f7468756d622f612f61382f4f627365727665725f775f7570646174652e7376672f35303070782d4f627365727665725f775f7570646174652e7376672e706e67)

- Ví dụ: Giả sử chúng ta có 1 kênh youtube và mỗi khi chúng ta realease 1 video mới, các subscriber đều được thông báo về thông tin của video mới này. Ta coi youtube channel là 1 subject và những subscriber trong channel là các observer. Khi các observer nhận được thông tin về video mới, app của các subscriber sẽ chịu trách nhiệm update lại UI để người dùng có thể click và những video được đăng tải:
- Code:
```go
type Observer interface {
	update(interface{})
}

type Subject interface {
	registerObserver(obs Observer)
	removeObserver(obs Observer)
	notifyObservers()
}

type Video struct {
	title string
}

// YoutubeChannel is a concrete implementation of Subject interface
type YoutubeChannel struct {
	Observers []Observer
	NewVideo  *Video
}

func (yt *YoutubeChannel) registerObserver(obs Observer) {
	yt.Observers = append(yt.Observers, obs)
}

func (yt *YoutubeChannel) removeObserver(obs Observer) {
	//
}

// notify to all observers when a new video is released
func (yt *YoutubeChannel) notifyObservers() {
	for _, obs := range yt.Observers {
		obs.update(yt.NewVideo)
	}
}

func (yt *YoutubeChannel) ReleaseNewVideo(video *Video) {
	yt.NewVideo = video
	yt.notifyObservers()
}

// UserInterface is a concrete implementation of Observer interface
type UserInterface struct {
	UserName string
	Videos   []*Video
}

func (ui *UserInterface) update(video interface{}) {
	ui.Videos = append(ui.Videos, video.(*Video))
	for video := range ui.Videos {
		View.AddChildNode(NewVideoComponent(video))
	}
	fmt.Printf("UI %s - Video: '%s' has just been released\n", ui.UserName, video.(*Video).title)
}

func NewUserInterface(username string) Observer {
	return &UserInterface{UserName: username, Videos: make([]*Video, 0)}
}

// usage

func main() {
	var ytChannel Subject = &YoutubeChannel{}
	ui1 := NewUserInterface("Bob")
	ui2 := NewUserInterface("Peter")
	ytChannel.registerObserver(ui1)
	ytChannel.registerObserver(ui2)
	ytChannel.(*YoutubeChannel).ReleaseNewVideo(&Video{title: "Avatar 2 trailer"})
	ytChannel.(*YoutubeChannel).ReleaseNewVideo(&Video{title: "Avengers Endgame trailer"})
}
```

### Strategy pattern
- Là mẫu thiết kế cho phép chọn thuật toán trong 1 nhóm các thuật toán liên quan đến nhau ngay tại lúc chương trình đang chạy (at runtime) để thực hiện một hoạt động nào đó.
- Giả sử chúng ta cần xây dựng 1 thư viện để mã hoá một đoạn tin bằng các phương pháp asymmetric chúng ta có thể sử dụng 1 trong 2 thuật toán sau: RSA hoặc Elliptic curve
```go
package encryption

type AsymEncryptionStrategy interface {
	Encrypt(data interface{}) (byte[] cipher, error)
}

type EllipticCurvestrategy struct {} 
type RSA struct {}

func (strat *EllipticCurvestrategy) Encrypt(data interface{}) (byte[] cipher, error) {
	// some complex math
	...
	return cipher, err 
}

func (strat *RSAstrategy) Encrypt(data interface{}) (byte[] cipher, error) {
	// some complex math
	...
	return cipher, err 
} 

func encryptMessage(msg string, strat AsymEncryptionStrategy) (byte[] cipher, error) {
	return strat.Encrypt(msg)
}

// usage
msg := "this is a confidential message"
cipher, err := encryptMessage(msg, encryption.EllipticCurvestrategy)
cipher, err := encryptMessage(msg, encryption.RSAstrategy)
```
###Iterator pattern
- Mẫu này được sử dụng để truy cập vào các phần tử của 1 collection(array, map, set) một cách tuần tự mà không cần phải hiểu biết về nó.
- Iterate 1 collection sử dụng callback:

```go
	func iterateEvenNumbers(max int, cb func(n int) error) error {
    if max < 0 {
        return fmt.Errorf("'max' is %d, must be >= 0", max)
    }
    for i := 2; i <= max; i += 2 {
        err := cb(i)
        if err != nil {
            return err
        }
    }
    return nil
}

func printEvenNumbers(max int) {
    err := iterateEvenNumbers(max, func(n int) error {
        fmt.Printf("%d\n", n)
        return nil
    })
    if err != nil {
        log.Fatalf("error: %s\n", err)
    }
}

printEvenNumbers(10)
```
- Iterate với Next()
```go
// EvenNumberIterator generates even numbers
type EvenNumberIterator struct {
    max       int
    currValue int
    err       error
}

// NewEvenNumberIterator creates new number iterator
func NewEvenNumberIterator(max int) *EvenNumberIterator {
    var err error
    if max < 0 {
        err = fmt.Errorf("'max' is %d, should be >= 0", max)
    }
    return &EvenNumberIterator{
        max:       max,
        currValue: 0,
        err:       err,
    }
}

// Next advances to next even number. Returns false on end of iteration.
func (i *EvenNumberIterator) Next() bool {
    if i.err != nil {
        return false
    }
    i.currValue += 2
    return i.currValue <= i.max
}

// Value returns current even number
func (i *EvenNumberIterator) Value() int {
    if i.err != nil || i.currValue > i.max {
        panic("Value is not valid after iterator finished")
    }
    return i.currValue
}

// Err returns iteration error.
func (i *EvenNumberIterator) Err() error {
    return i.err
}

func printEvenNumbers(max int) {
	iter := NewEvenNumberIterator(max)
	for iter.Next() {
		fmt.Printf("n: %d\n", iter.Value())
	}
	if iter.Err() != nil {
		log.Fatalf("error: %s\n", iter.Err())
	}
}

func main() {
	fmt.Printf("Even numbers up to 8:\n")
	printEvenNumbers(8)
	fmt.Printf("Even numbers up to 9:\n")
	printEvenNumbers(9)
	fmt.Printf("Error: even numbers up to -1:\n")
	printEvenNumbers(-1)
}
```
- Pattern này được sử dụng nhiều trong go standard library
	- Rows.Next: iterate các kết quả thu được tử SQL SELECT statement
	- Scanner.Scan: iterate text
	- Decoder.Token: XML parsing
	- Reader.Read: CSV reader
References: [https://blog.kowalczyk.info/article/1Bkr/3-ways-to-iterate-in-go.html](https://blog.kowalczyk.info/article/1Bkr/3-ways-to-iterate-in-go.html "https://blog.kowalczyk.info/article/1Bkr/3-ways-to-iterate-in-go.html")

### State pattern
- Mỗi đối tượng có 1 trạng thái gắn với nó và trạng thái có thể được thay đổi thông qua SetState method.
- Ví dụ: Giả sử điện thoại có 2 trạng thái nhắc nhở: Im lặng hoặc Rung.
- Code 
```go
type MobileAlertState interface {
	alert()
}

type AlertStateContext struct {
	currentState MobileAlertState
}

func NewAlertStateContext() *AlertStateContext {
	return &AlertStateContext{currentState: &Vibration{}}
}

func (ctx *AlertStateContext) SetState(state MobileAlertState) {
	ctx.currentState = state
}

func (ctx *AlertStateContext) Alert() {
	ctx.currentState.alert()
}

type Vibration struct{}

func (v *Vibration) alert() {
	fmt.Println("vibrating....")
}

type Silence struct{}

func (s *Silence) alert() {
	fmt.Println("silent ....")
}

func main() {
	stateContext := NewAlertStateContext()
	stateContext.Alert()
	stateContext.Alert()
	stateContext.Alert()
	stateContext.SetState(&Silence{})
	stateContext.Alert()
	stateContext.Alert()
	stateContext.Alert()
}

// result
vibrating....
vibrating....
vibrating....
silent ....
silent ....
silent ....
```
## Structural Patterns

### Composite

- Cho phép tương tác với tất cả các đối tượng tương tự nhau giống như là một đối tượng đơn hoặc 1 nhóm đối tượng.
- Hình dung: Đối tượng File sẽ là 1 đối tượng đơn nếu bên trong nó không có file nào khác, nhưng đối tượng file sẽ được đối xử giống như 1 collection nếu bên trong nó lại có những File khác.
- Sơ đồ:

[![](https://camo.githubusercontent.com/b3ce771e2f9bac43d884aa30566fe2c2953ea07b081fda5cb260bb22a947f060/68747470733a2f2f75706c6f61642e77696b696d656469612e6f72672f77696b6970656469612f636f6d6d6f6e732f7468756d622f352f35612f436f6d706f736974655f554d4c5f636c6173735f6469616772616d5f25323866697865642532392e7376672f36303070782d436f6d706f736974655f554d4c5f636c6173735f6469616772616d5f25323866697865642532392e7376672e706e67)](https://camo.githubusercontent.com/b3ce771e2f9bac43d884aa30566fe2c2953ea07b081fda5cb260bb22a947f060/68747470733a2f2f75706c6f61642e77696b696d656469612e6f72672f77696b6970656469612f636f6d6d6f6e732f7468756d622f352f35612f436f6d706f736974655f554d4c5f636c6173735f6469616772616d5f25323866697865642532392e7376672f36303070782d436f6d706f736974655f554d4c5f636c6173735f6469616772616d5f25323866697865642532392e7376672e706e67)

- Cấu trúc:
	- Component (thành phần):
		- Khai báo interface hoặc abstract chung cho các thành phần đối tượng.
		- Chứa các method thao tác chung của các thành phần đối tượng.
	- Leaf (Lá):
		- Biểu diễn các đối tượng lá (không có con) trong thành phần đối tượng.
	- Composite (Hỗn hợp):
		- Định nghĩa một thao tác cho các thành phần có thành phần con.
		- Lưu trữ và quản lý các thành phần con

- Ví dụ khác: Khi vẽ, chúng ta được cung cấp các đối tượng Square, Circle, các đối tượng này đều là Shape.. Giả sử chúng ta muốn vẽ nhiều loại hình cùng 1 lúc chúng ta sẽ tạo một layer chứa các Shape này và thực hiện vòng lặp để vẽ chúng. Ở đây composite được hiểu là chúng ta có thể sử dụng các đối tượng Square, Circle riêng biệt nhưng khi cần chúng ta có thể gom chúng thành 1 nhóm.

```go
// Shape is the component
type Shape interface {
	Draw(drawer *Drawer) error
}

// Square and Circle are leaves
type Square struct {
	Location Point
	Size float64
}

func (square *Square) Draw(drawer *Drawer) error {
	return drawer.DrawRect(Rect{
		Location: square.Location,
		Size: Size{
			Height: square.Side,
			Width:  square.Side,
		},
	})
}

type Circle struct {
	Center Point
	Radius float64
}

func (circle *Circle) Draw(drawer *Drawer) error {
	rect := Rect{
		Location: Point{
			X: circle.Center.X - circle.Radius,
			Y: circle.Center.Y - circle.Radius,
		},
		Size: Size{
			Width:  2 * circle.Radius,
			Height: 2 * circle.Radius,
		},
	}

	return drawer.DrawEllipseInRect(rect)
}

// Layer is the composite
type Layer struct {
	Shapes []Shape
}

func (layer *Layer) Draw(drawer *Drawer) error {
	for _, shape := range layer.Shapes {
		if err := shape.Draw(drawer); err != nil {
			return err
		}
		fmt.Println()
	}

	return nil
}

// usage
circle := &photoshop.Circle{
	Center: photoshop.Point{X: 100, Y: 100},
	Radius: 50,
}

square := &photoshop.Square{
	Location: photoshop.Point{X: 50, Y: 50},
	Side:     20,
}

layer := &photoshop.Layer{
	Elements: []photoshop.Shapes{
		circle,
		square,
	},
}

circle.Draw(&photoshop.Drawer{})
square.Draw(&photoshop.Drawer{})
// or
layer.Draw(&photoshop.Drawer{})
```


