package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io/ioutil"
	"net"
	"net/http"
	"net/smtp"
	"os"
	"reflect"
	"strconv"
	"strings"
	"text/template"
	"time"

	models "api_thienhang_com/pkg/entity"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

const (
	E_WELCOME = 0
	E_OTP     = 1
	E_RESET   = 2
	E_VERSION = 3
)

type Response struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}

func GetExpirationTime(hours int) int64 {
	return time.Now().Add(time.Hour * time.Duration(hours)).Unix()
}

func GetDurationByHour(hours int) time.Duration {
	return time.Hour * time.Duration(hours)
}

func ParseTemplate(fileName string, data interface{}) (content string, err error) {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return "", err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func ResponseWithJson(w http.ResponseWriter, status int, object interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var res Response
	if status == http.StatusOK {
		res = Response{
			false,
			"success",
			object,
		}
	} else {
		res = Response{
			true,
			"error",
			object,
		}
	}
	json.NewEncoder(w).Encode(res)
}

func SendEmail(to []string, msg string, eType int) (err error) {
	var t *template.Template
	var title string = "Thông báo"
	switch eType {
	case E_WELCOME:
		title = "Thư chào mừng thienhang.com"
		t, err = template.ParseFiles("template/welcome.html")
		if err != nil {
			logrus.Error(err)
		}
	case E_OTP:
		t, err = template.ParseFiles("template/sendotp.html")
		if err != nil {
			logrus.Error(err)
		}
	case E_RESET:
		t, _ = template.ParseFiles("./././template/resetpassword.html")
	case E_VERSION:
		title = "🐷 Thông báo bản cập nhật mới"
		t, _ = template.ParseFiles("./././template/version.html")
	default:
		title = "test"
		t, err = template.ParseFiles("template/welcome.html")
		if err != nil {
			logrus.Error(err)
		}
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s  \n%s\n\n", title, mimeHeaders)))

	t.Execute(&body, struct {
		Name    string
		Message string
	}{

		Name:    to[0],
		Message: msg,
	})

	auth := smtp.PlainAuth("", "postmaster@mail.thienhang.com", "f815d1178c6b3f30aaa1bed91ea13d2b-78651cec-f6284dac", "smtp.eu.mailgun.org")
	err = smtp.SendMail("smtp.eu.mailgun.org:587", auth, "noreply@thienhang.com", to, body.Bytes())
	if err != nil {
		logrus.Error(err)
		return err
	}
	fmt.Println(auth)
	return nil
}

func GenerateOtp(number string) string {
	return "XXXXX"
}

func GetEnvVar(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(bytes), err
}

func ComparePasswordAndHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}

func ValidateUser(user *models.User) bool {
	// email and phone are empty
	if govalidator.IsNull(user.Email) && govalidator.IsNull(user.PhoneNumber) {
		return false
	}

	if !govalidator.IsEmail(user.Email) {
		return false
	}

	return true
}

func setField(field reflect.Value, defaultVal string) error {

	if !field.CanSet() {
		return errors.New("Can't set value\n")
	}

	switch field.Kind() {

	case reflect.Int:
		if val, err := strconv.ParseInt(defaultVal, 10, 64); err == nil {
			field.Set(reflect.ValueOf(int(val)).Convert(field.Type()))
		}
	case reflect.String:
		field.Set(reflect.ValueOf(defaultVal).Convert(field.Type()))
	}

	return nil
}

func SetDefault(ptr interface{}) error {
	tag := "default"
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		return fmt.Errorf("Not a pointer")
	}

	v := reflect.ValueOf(ptr).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		if defaultVal := t.Field(i).Tag.Get(tag); defaultVal != "-" {
			if err := setField(v.Field(i), defaultVal); err != nil {
				return err
			}

		}
	}
	return nil
}

func GenerateUUID() string {
	return uuid.New().String()
}

func Convert2DBRef(data interface{}) (models.DBRef, error) {
	var inter []interface{}
	dbRef := models.DBRef{}

	jsonbody, _ := json.Marshal(data)
	err := json.Unmarshal(jsonbody, &inter)
	if err != nil {
		return dbRef, err
	}
	for _, r := range inter {
		m, _ := r.(map[string]interface{})
		switch m["Key"] {
		case "ref":
			dbRef.Ref = m["Value"]
		case "id":
			dbRef.ID = m["Value"]
		case "db":
			dbRef.DB = m["Value"]
		}
	}
	return dbRef, nil
}

func LoadJsonFile(filename string) []byte {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Error(err)
	}
	return file
}

func getIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			v := ip.String()
			return v, nil
		}
	}
	return "", errors.New("are you connected to the network")
}

var (
	Err_ABNORMAL             = errors.New("lỗi không xác định từ hệ thống. Anh chị vui lòng liên lạc lại với CSKH sandexcare.com")
	ErrAmountOut             = errors.New("amount out")
	ErrSignature             = errors.New("invalid signature")
	ErrSignature1            = errors.New("Giao dịch không thành công, Ngân hàng phát hành thẻ không cấp phép cho giao dịch hoặc thẻ chưa được kích hoạt dịch vụ thanh toán trên Internet. Vui lòng liên hệ ngân hàng theo số điện thoại sau mặt thẻ được hỗ trợ chi tiết.")
	ErrSignature2            = errors.New("Giao dịch không thành công, Ngân hàng phát hành thẻ từ chối cấp phép cho giao dịch. Vui lòng liên hệ ngân hàng theo số điện thoại sau mặt thẻ để biết chính xác nguyên nhân Ngân hàng từ chối.")
	ErrSignature3            = errors.New("Giao dịch không thành công, Cổng thanh toán không nhận được kết quả trả về từ ngân hàng phát hành thẻ. Vui lòng liên hệ với ngân hàng theo số điện thoại sau mặt thẻ để biết chính xác trạng thái giao dịch và thực hiện thanh toán lại")
	ErrSignature4            = errors.New("Giao dịch không thành công do thẻ hết hạn sử dụng hoặc nhập sai thông tin tháng/ năm hết hạn của thẻ. Vui lòng kiểm tra lại thông tin và thanh toán lại")
	ErrSignature5            = errors.New("Giao dịch không thành công, Thẻ không đủ hạn mức hoặc tài khoản không đủ số dư để thanh toán. Vui lòng kiểm tra lại thông tin và thanh toán lại")
	ErrSignature6            = errors.New("Giao dịch không thành công, Quá trình xử lý giao dịch phát sinh lỗi từ ngân hàng phát hành thẻ. Vui lòng liên hệ ngân hàng theo số điện thoại sau mặt thẻ được hỗ trợ chi tiết.")
	ErrSignature7            = errors.New("Giao dịch không thành công, Đã có lỗi phát sinh trong quá trình xử lý giao dịch. Vui lòng thực hiện thanh toán lại.")
	ErrSignature8            = errors.New("Giao dịch không thành công. Số thẻ không đúng. Vui lòng kiểm tra và thực hiện thanh toán lại")
	ErrSignature9            = errors.New("Giao dịch không thành công. Tên chủ thẻ không đúng. Vui lòng kiểm tra và thực hiện thanh toán lại")
	ErrSignature10           = errors.New("Giao dịch không thành công. Thẻ hết hạn/Thẻ bị khóa. Vui lòng kiểm tra và thực hiện thanh toán lại")
	ErrSignature11           = errors.New("Giao dịch không thành công. Thẻ chưa đăng ký sử dụng dịch vụ thanh toán trên Internet. Vui lòng liên hê ngân hàng theo số điện thoại sau mặt thẻ để được hỗ trợ.")
	ErrSignature12           = errors.New("Giao dịch không thành công. Ngày phát hành/Hết hạn không đúng. Vui lòng kiểm tra và thực hiện thanh toán lại")
	ErrSignature13           = errors.New("Giao dịch không thành công. thẻ/ tài khoản đã vượt quá hạn mức thanh toán. Vui lòng kiểm tra và thực hiện thanh toán lại")
	ErrSignature21           = errors.New("Giao dịch không thành công. Số tiền không đủ để thanh toán. Vui lòng kiểm tra và thực hiện thanh toán lại")
	ErrSignature22           = errors.New("Giao dịch không thành công. Thông tin tài khoản không đúng. Vui lòng kiểm tra và thực hiện thanh toán lại")
	ErrSignature23           = errors.New("Giao dịch không thành công. Tài khoản bị khóa. Vui lòng liên hê ngân hàng theo số điện thoại sau mặt thẻ để được hỗ trợ")
	ErrSignature24           = errors.New("Giao dịch không thành công. Thông tin thẻ không đúng. Vui lòng kiểm tra và thực hiện thanh toán lại")
	ErrSignature25           = errors.New("Giao dịch không thành công. OTP không đúng. Vui lòng kiểm tra và thực hiện thanh toán lại")
	ErrSignature253          = errors.New("Giao dịch không thành công. Quá thời gian thanh toán. Vui lòng thực hiện thanh toán lại")
	ErrSignature99           = errors.New("Giao dịch không thành công. Người sử dụng hủy giao dịch")
	ErrSignatureB            = errors.New("Giao dịch không thành công do không xác thực được 3D-Secure. Vui lòng liên hệ ngân hàng theo số điện thoại sau mặt thẻ được hỗ trợ chi tiết.")
	ErrSignatureE            = errors.New("Giao dịch không thành công do nhập sai CSC (Card Security Card) hoặc ngân hàng từ chối cấp phép cho giao dịch. Vui lòng liên hệ ngân hàng theo số điện thoại sau mặt thẻ được hỗ trợ chi tiết.")
	ErrSignatureF            = errors.New("Giao dịch không thành công do không xác thực được 3D-Secure. Vui lòng liên hệ ngân hàng theo số điện thoại sau mặt thẻ được hỗ trợ chi tiết.")
	ErrSignatureZ            = errors.New("Giao dịch không thành công do vi phạm quy định của hệ thống. Vui lòng liên hệ với OnePAY để được hỗ trợ (Hotline: 1900 633 927)")
	ErrSignatureOther        = errors.New("Giao dịch không thành công. Vui lòng liên hệ với OnePAY để được hỗ trợ (Hotline: 1900 633 927)")
	Err_STRING_METHOD        = "Anh chị vui lòng lựa chọn phương thức thanh toán hợp lệ."
	Err_STRING_TOKEN         = "Phiên làm việc của anh chị đã hết hạn hoặc tài khoản của anh chị đã bị khoá. Anh chị vui lòng đăng nhập lại hoặc liên hệ bộ phận chăm sóc khách hàng của Sandexcare để được hỗ trợ."
	Err_STRING_RANGE_DEPOSIT = "Định dạng số tiền thanh toán của anh chị chưa hợp lệ. Số tiền thanh toán phải lớn hơn hoặc bằng 50.000 đồng, bội số của 10.000"
	// Err_STRING_METHOD =
	// Err_STRING_METHOD =
	Err_Unauthorization = ""
	Err_Signature       = "Chữ kí chưa hợp lệ"
)
