package payment

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/zpmep/hmacutil"
)

type IPayment interface {
	CaptureWallet() (interface{}, error)
	GetInvoiceID() string
	GetAmount() string
}

// define a payload, reference in https://developers.momo.vn/#cong-thanh-toan-momo-phuong-thuc-thanh-toan
type momoPayload struct {
	PartnerCode string `json:"partnerCode"`
	PartnerName string `json:"partnerName"`
	StoreId     string `json:"storeId"`
	RequestID   string `json:"requestId"`
	Amount      string `json:"amount"`
	OrderID     string `json:"orderId"`
	OrderInfo   string `json:"orderInfo"`
	RedirectUrl string `json:"redirectUrl"`
	IpnUrl      string `json:"ipnUrl"`
	ExtraData   string `json:"extraData"`
	RequestType string `json:"requestType"`
	Signature   string `json:"signature"`
	Lang        string `json:"lang"`
}

type zaloPayload struct {
	Amount  string `json:"amount"`
	OrderID string `json:"orderId"`
}

type shopeePayload struct {
}
type onePayPayload struct {
	AgainLink   string `json:"againLink"`
	Title       string `json:"title"`
	AccessCode  string `json:"accessCode"`
	Amount      string `json:"amount"`
	Command     string `json:"command"`
	Currency    string `json:"currency"`
	Locale      string `json:"locale"`
	MerchTxnRef string `json:"merchTxnRef"`
	Merchant    string `json:"merchant"`
	OrderInfo   string `json:"orderInfo"`
	TicketNo    string `json:"ticketNo"`
	Version     string `json:"version"`
	ReturnURL   string `json:"returnURL"`
	Signature   string `json:"signature"  example: "90482b3881bdf863d5f61ace078921bbc6dbb58b2fded35261c71c9af3b1ce4f"`
}

var pubKeyData = []byte(`
-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA49mGWEzZt4KD9O0nsVoG6ibZ6C6rZtkTOCuXuBYoG2eAhFZL6u/A5JRvtSONcficyopXq2DVlY3RdTLq3+MFWF46KBcnWs7ThNZB7Vy1UmhwpG7WOg7YOBJ2BydLFhs5Y+yBMAsEBY9E7BwkLl3S+DFpC7vI0NJWrd8yHCCOfAY5uWGqTp7MNpSzwRu3YCa9jzs9H61EsyEbCeodtsX+VrlT0qhmh9P5P99zLMFuRkYceLpo7B0P7KQ9P3H3IvtkIJ0YnE2ltFRlWAi5bCjEK0qoiF6p4cnlyXPG2xq/KHSdXPKwHLcHulsqtJWK9WFu45iW+sr8ZU5oZU4eg5+BY/IzacVZqqCkLzCjeqy1JHMduBpDmzoR5w7BkOO1clQ1O82NhVYSzfkxFBjPW9AMZACiHBfqgg5ESjKTG3VhgkQpmnfGi0+tnGg56m6yzPl30Vkv1MYnu8eNR7atRpEbcktiBezvFzokvQCFPYKBIME2KC2XRb4B8pDhh/4XxZThl9xquI3itdrPD/rF/G4oUOmkpprF3hryKm3cKajC5iUO9qeAL6Z3qjKBJSa4MfMz87JvKbLOUXYyP+um0ioDz8hpeTsB+mfdx70rGGONtB+PCxAMnhYC+9uYNqAhPDHpSJtWsjUV1x8zf6pQKGpBN5NdH1tzF1lujE77+3VyHzkCAwEAAQ==
-----END PUBLIC KEY-----
`)

func CreateMomo(id string, am string) IPayment {
	var orderId = id
	var requestId = id
	var partnerCode = "MOMOYSWU20211020"
	var accessKey = "FEbWzxcfpCYM6HBx"
	var secretKey = "fXrgYaYR55WpeSHjZpIMp18gtIW4dDcA"
	var orderInfo = "Thanh toán đơn hàng payment.sandexcare.com"
	var redirectUrl = "sandexcare://deposit"
	var ipnUrl = "https://payment.sandexcare.com/api/v1/ipn/momo"
	var amount = am
	var requestType = "captureWallet"
	var extraData = "" //pass empty value or Encode base64 JsonString

	//build raw signature
	var rawSignature bytes.Buffer
	rawSignature.WriteString("accessKey=")
	rawSignature.WriteString(accessKey)
	rawSignature.WriteString("&amount=")
	rawSignature.WriteString(am)
	rawSignature.WriteString("&extraData=")
	rawSignature.WriteString(extraData)
	rawSignature.WriteString("&ipnUrl=")
	rawSignature.WriteString(ipnUrl)
	rawSignature.WriteString("&orderId=")
	rawSignature.WriteString(orderId)
	rawSignature.WriteString("&orderInfo=")
	rawSignature.WriteString(orderInfo)
	rawSignature.WriteString("&partnerCode=")
	rawSignature.WriteString(partnerCode)
	rawSignature.WriteString("&redirectUrl=")
	rawSignature.WriteString(redirectUrl)
	rawSignature.WriteString("&requestId=")
	rawSignature.WriteString(requestId)
	rawSignature.WriteString("&requestType=")
	rawSignature.WriteString(requestType)
	// Create a new HMAC by defining the hash type and the key (as byte array)
	hmac := hmac.New(sha256.New, []byte(secretKey))

	// Write Data to it
	hmac.Write(rawSignature.Bytes())
	fmt.Println("\nRaw signature: " + rawSignature.String())

	// Get result and encode as hexadecimal string
	signature := hex.EncodeToString(hmac.Sum(nil))
	return &momoPayload{
		//AccessKey:   accessKey,
		Amount:      amount,
		ExtraData:   extraData,
		IpnUrl:      ipnUrl,
		OrderID:     orderId,
		OrderInfo:   orderInfo,
		PartnerCode: partnerCode,
		PartnerName: "sandexcare.com",
		RedirectUrl: redirectUrl,
		RequestID:   requestId,
		RequestType: requestType,
		StoreId:     "N1",
		Signature:   signature,
		Lang:        "en",
	}
}

func (payload *momoPayload) CaptureWallet() (interface{}, error) {
	var jsonPayload []byte
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := http.Post(os.Getenv("MOMO_CREATE"), "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil || resp.StatusCode != 200 {
		log.Error(err)
		return nil, errors.New("momo no availabe")
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result["payUrl"], nil
}

func (payload *momoPayload) GetInvoiceID() string {
	return payload.OrderID
}

func (payload *momoPayload) GetAmount() string {
	return payload.Amount
}

func CreateZalo(orderId string, amount string) IPayment {
	return &zaloPayload{
		Amount:  amount,
		OrderID: orderId,
	}
}

func (payload *zaloPayload) CaptureWallet() (interface{}, error) {
	type object map[string]interface{}

	rand.Seed(time.Now().UnixNano())
	transID := rand.Intn(1000000) // Generate random trans id
	embedData, _ := json.Marshal(object{})
	items, _ := json.Marshal([]object{})
	// request data
	params := make(url.Values)
	params.Add("app_id", "2553")
	params.Add("amount", "100000") //payload.Amount)
	params.Add("app_user", "user123")
	params.Add("embed_data", string(embedData))
	params.Add("item", string(items))
	params.Add("description", "Sandexcare.com - Payment for the order #"+strconv.Itoa(transID))
	params.Add("bank_code", "zalopayapp")

	now := time.Now()
	params.Add("app_time", strconv.FormatInt(now.UnixNano()/int64(time.Millisecond), 10)) // miliseconds

	params.Add("app_trans_id", fmt.Sprintf("%02d%02d%02d_%v", now.Year()%100, int(now.Month()), now.Day(), transID)) // translation missing: vi.docs.shared.sample_code.comments.app_trans_id

	// appid|app_trans_id|appuser|amount|apptime|embeddata|item
	data := fmt.Sprintf("%v|%v|%v|%v|%v|%v|%v", params.Get("app_id"), params.Get("app_trans_id"), params.Get("app_user"),
		params.Get("amount"), params.Get("app_time"), params.Get("embed_data"), params.Get("item"))
	params.Add("mac", hmacutil.HexStringEncode(hmacutil.SHA256, "PcY4iZIKFCIdgZvA6ueMcMHHUbRLYjPL", data))
	log.Error(params)
	resp, err := http.PostForm("https://sb-openapi.zalopay.vn/v2/create", params)
	log.Error(resp)
	if err != nil || resp.StatusCode != 200 {
		return nil, errors.New("zalopay is not available")
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result["order_url"], nil
}

func (payload *zaloPayload) GetInvoiceID() string {
	return payload.OrderID
}

func (payload *zaloPayload) GetAmount() string {
	return payload.Amount
}

func CreateOnePay(invoiceID string, am string, ip string) IPayment {
	var title = "sandexcare.com"
	var againLink = "payment.sandexcare.com"
	var accessCode = "C0905B01"
	// var accessCode = "6BEB2566" //test OP
	// var accessCode = "6BEB2546" //test OP
	var command = "pay"
	var currency = "VND"
	var locale = "vn"
	var merchTxnRef = invoiceID
	// var merchant = "TESTONEPAY25" //test OP
	// var merchant = "TESTONEPAY" //test OP
	var merchant = "OP_SANDEX"
	var orderInfo = "sandexcare.com"
	var ticketNo = ip
	var version = "2"
	var returnURL = "https://payment.sandexcare.com/api/v1/ipn/onepay"
	// var returnURL = "http://localhost:8001/api/v1/ipn/onepay" //test local
	// var returnURL = "sandexcare://onepay/"
	var amount = am

	// secretKey, _ := hex.DecodeString("6D0870CDE5F24F34F3915FB0045120D6") //test OP
	// secretKey, _ := hex.DecodeString("6D0870CDE5F24F34F3915FB0045120DB") //test OP
	secretKey, _ := hex.DecodeString("B157D0AB54E32DF09156BF5E4D7E9988")
	//build raw signature
	var rawSignature bytes.Buffer
	rawSignature.WriteString("vpc_AccessCode=")
	rawSignature.WriteString(accessCode)
	rawSignature.WriteString("&vpc_Amount=")
	rawSignature.WriteString(am)
	rawSignature.WriteString("&vpc_Command=")
	rawSignature.WriteString(command)
	rawSignature.WriteString("&vpc_Currency=")
	rawSignature.WriteString(currency)
	rawSignature.WriteString("&vpc_Locale=")
	rawSignature.WriteString(locale)
	rawSignature.WriteString("&vpc_MerchTxnRef=")
	rawSignature.WriteString(merchTxnRef)
	rawSignature.WriteString("&vpc_Merchant=")
	rawSignature.WriteString(merchant)
	rawSignature.WriteString("&vpc_OrderInfo=")
	rawSignature.WriteString(orderInfo)
	rawSignature.WriteString("&vpc_ReturnURL=")
	rawSignature.WriteString(returnURL)
	rawSignature.WriteString("&vpc_TicketNo=")
	rawSignature.WriteString(ticketNo)
	rawSignature.WriteString("&vpc_Version=")
	rawSignature.WriteString(version)

	hmac := hmac.New(sha256.New, []byte(secretKey))
	hmac.Write(rawSignature.Bytes())
	signature := hex.EncodeToString(hmac.Sum(nil))

	return &onePayPayload{
		AgainLink:   againLink,
		Title:       title,
		AccessCode:  accessCode,
		Amount:      amount,
		Command:     command,
		Currency:    currency,
		Locale:      locale,
		MerchTxnRef: merchTxnRef,
		Merchant:    merchant,
		OrderInfo:   orderInfo,
		TicketNo:    ticketNo,
		Version:     version,
		ReturnURL:   returnURL,
		Signature:   signature,
	}
}
func (payload *onePayPayload) GetAmount() string {
	return payload.Amount
}

func (payload *onePayPayload) CaptureWallet() (interface{}, error) {
	log.Info("CaptureWallet")
	// var endpoint = os.Getenv("ONEPAY_URL")
	var endpoint = "https://onepay.vn/paygate/vpcpay.op" //Hòa thêm
	// var endpoint = "https://mtf.onepay.vn/paygate/vpcpay.op" //test OP
	log.Error(os.Getenv("ONEPAY_URL"))
	req, err := http.NewRequest("GET", endpoint, nil)
	log.Error(req)
	if err != nil {
		log.Error(err)
		return "", err
	}
	// defer req.Body.Close()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	parm := req.URL.Query()
	parm.Add("vpc_AccessCode", payload.AccessCode)
	parm.Add("vpc_Amount", payload.Amount)
	parm.Add("vpc_Command", payload.Command)
	parm.Add("vpc_Currency", payload.Currency)
	parm.Add("vpc_Locale", payload.Locale)
	parm.Add("vpc_MerchTxnRef", payload.MerchTxnRef)
	parm.Add("vpc_Merchant", payload.Merchant)
	parm.Add("vpc_OrderInfo", payload.OrderInfo)
	parm.Add("vpc_ReturnURL", payload.ReturnURL)
	parm.Add("vpc_TicketNo", payload.TicketNo)
	parm.Add("vpc_Version", payload.Version)
	parm.Add("vpc_SecureHash", payload.Signature)
	parm.Add("AgainLink", payload.AgainLink)
	parm.Add("Title", payload.Title)
	req.URL.RawQuery = parm.Encode()
	log.Info(req.URL.String())
	return req.URL.String(), nil
}

func (payload *onePayPayload) GetInvoiceID() string {
	return payload.MerchTxnRef
}

type Payment struct {
	OrderInfo string `json:"order_info"`
	Amount    string `json:"amount"`
	BankCode  string `json:"bank_code"`
	OrderType string `json:"order_type"`
	VnpUrl    string
	ReturnUrl string
	IpAddr    string
	SecretKey string
	TmnCode   string
}

func (pay Payment) GeneratePaymentUrl() string {
	currentDate := time.Now()
	m := map[string]string{}

	m["vnp_Version"] = "2"
	m["vnp_Command"] = "pay"
	m["vnp_TmnCode"] = pay.TmnCode
	m["vnp_Amount"] = pay.Amount
	m["vnp_CreateDate"] = currentDate.Format("20060102150405")
	m["vnp_CurrCode"] = "VND"
	m["vnp_IpAddr"] = pay.IpAddr
	m["vnp_Locale"] = "vn"
	m["vnp_OrderInfo"] = pay.OrderInfo
	m["vnp_OrderType"] = pay.OrderType
	m["vnp_ReturnUrl"] = pay.ReturnUrl
	m["vnp_TxnRef"] = currentDate.Format("150405")
	m["vnp_BankCode"] = pay.BankCode
	m["vnp_ExpireDate"] = currentDate.Format("20060102150405")

	req, _ := http.NewRequest("GET", pay.VnpUrl, nil)
	q := req.URL.Query()

	for key := range m {
		q.Add(key, m[key])
	}
	keys := []string{}
	for k := range q {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	qData := ""
	for _, key := range keys {
		qData += key + "=" + strings.Join(q[key], "") + "&"
	}

	signData := pay.SecretKey + qData

	lengSignData := len(signData)
	signData = signData[:lengSignData-1]

	hash := sha256.New()
	hash.Write([]byte(signData))
	hashed := hex.EncodeToString(hash.Sum(nil))

	q.Add("vnp_SecureHashType", "SHA256")
	q.Add("vnp_SecureHash", hashed)

	req.URL.RawQuery = q.Encode()

	return req.URL.String()
}
