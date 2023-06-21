package payment

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"api_thienhang_com/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/zpmep/hmacutil"
)

// MomoIPN
type MomoIPN struct {
	PartnerCode  string `json:"partnerCode"  example:"MOMOIOLD20190129"`
	InvoiceId    string `json:"orderId"  example:"01234567890123451633504872421"`
	RequestId    string `json:"requestId" example:"01234567890123451633504872421"`
	Amount       int    `json:"amount"  example:"1000"`
	OrderInfo    string `json:"orderInfo"  example:"Test Thue 1234556"`
	OrderType    string `json:"orderType"  example:"momo_wallet"`
	TransId      int    `json:"transId"  example:"2588659987"`
	ResultCode   int    `json:"resultCode"  example:"0"`
	Message      string `json:"message"  example:"Giao dịch thành công."`
	PayType      string `json:"payType"  example:"qr"`
	ResponseTime int    `json:"responseTime"  example:"1633504902954"`
	ExtraData    string `json:"extraData"  example:"eyJyZXN1bHRfbmFtZXNwYWNlIjoidW1hcmtldCIsImVycm9yIjoiIiwic3RhdGUiOjZ9"`
	Signature    string `json:"signature"  example:"90482b3881bdf863d5f61ace078921bbc6dbb58b2fded35261c71c9af3b1ce4f"`
}

func CheckSignatureMoMo(r *http.Request) (amount int, orderID string, err error) {
	var mm MomoIPN
	if err := json.NewDecoder(r.Body).Decode(&mm); err != nil {
		return 0, "", err
	}
	var rawSignature bytes.Buffer
	rawSignature.WriteString("amount=")
	rawSignature.WriteString(fmt.Sprint(mm.Amount))
	rawSignature.WriteString("&extraData=")
	rawSignature.WriteString(mm.ExtraData)
	rawSignature.WriteString("&message=")
	rawSignature.WriteString(mm.Message)
	rawSignature.WriteString("&orderId=")
	rawSignature.WriteString(mm.InvoiceId)
	rawSignature.WriteString("&orderInfo=")
	rawSignature.WriteString(mm.OrderInfo)
	rawSignature.WriteString("&orderType=")
	rawSignature.WriteString(mm.OrderType)
	rawSignature.WriteString("&partnerCode=")
	rawSignature.WriteString(mm.PartnerCode)
	rawSignature.WriteString("&payType=")
	rawSignature.WriteString(mm.PayType)
	rawSignature.WriteString("&requestId=")
	rawSignature.WriteString(mm.RequestId)
	rawSignature.WriteString("&responseTime=")
	rawSignature.WriteString(fmt.Sprint(mm.ResponseTime))
	rawSignature.WriteString("&resultCode=")
	rawSignature.WriteString(fmt.Sprint(mm.ResultCode))
	rawSignature.WriteString("&transId=")
	rawSignature.WriteString(fmt.Sprint(mm.TransId))

	// Create a new HMAC by defining the hash type and the key (as byte array)
	var secretKey = "fXrgYaYR55WpeSHjZpIMp18gtIW4dDcA"
	hmac := hmac.New(sha256.New, []byte(secretKey))

	// Write Data to it
	hmac.Write(rawSignature.Bytes())

	// Get result and encode as hexadecimal string
	// signature := hex.EncodeToString(hmac.Sum(nil))
	// fmt.Println(signature)
	return mm.Amount, mm.InvoiceId, nil
}

func CheckSignatureZalo(r *http.Request) (amount int, orderID string, err error) {
	defer r.Body.Close()
	var cbdata map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&cbdata)

	requestMac := cbdata["mac"].(string)
	dataStr := cbdata["data"].(string)
	mac := hmacutil.HexStringEncode(hmacutil.SHA256, "eG4r0GcoNtRGbO8", dataStr)
	log.Println("mac =", mac)

	// kiểm tra callback hợp lệ (đến từ ZaloPay server)
	if mac != requestMac {
		return 0, "", errors.New("mac not equal")
	}

	// merchant cập nhật trạng thái cho đơn hàng
	var dataJSON map[string]interface{}
	json.Unmarshal([]byte(dataStr), &dataJSON)
	log.Println("update order's status = success where app_trans_id =", dataJSON["app_trans_id"])

	return 0, "mm.InvoiceId", nil
}

func CheckSignatureOnePay(r *http.Request) (amount int, orderID string, err error) {
	//
	keys := []string{}
	for k := range r.URL.Query() {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	//
	result := r.URL.Query().Get("vpc_TxnResponseCode")
	// if result != "0" {
	// 	return 0, "", utils.ErrSignature
	// }
	if result == "1" {
		return 0, "", utils.ErrSignature1
	}
	if result == "2" {
		return 0, "", utils.ErrSignature2
	}
	if result == "3" {
		return 0, "", utils.ErrSignature3
	}
	if result == "4" {
		return 0, "", utils.ErrSignature4
	}
	if result == "5" {
		return 0, "", utils.ErrSignature5
	}
	if result == "6" {
		return 0, "", utils.ErrSignature6
	}
	if result == "7" {
		return 0, "", utils.ErrSignature7
	}
	if result == "8" {
		return 0, "", utils.ErrSignature8
	}
	if result == "9" {
		return 0, "", utils.ErrSignature9
	}
	if result == "10" {
		return 0, "", utils.ErrSignature10
	}
	if result == "11" {
		return 0, "", utils.ErrSignature11
	}
	if result == "12" {
		return 0, "", utils.ErrSignature12
	}
	if result == "13" {
		return 0, "", utils.ErrSignature13
	}
	if result == "21" {
		return 0, "", utils.ErrSignature21
	}
	if result == "22" {
		return 0, "", utils.ErrSignature22
	}
	if result == "23" {
		return 0, "", utils.ErrSignature23
	}
	if result == "24" {
		return 0, "", utils.ErrSignature24
	}
	if result == "25" {
		return 0, "", utils.ErrSignature25
	}
	if result == "253" {
		return 0, "", utils.ErrSignature253
	}
	if result == "99" {
		return 0, "", utils.ErrSignature99
	}
	if result == "B" {
		return 0, "", utils.ErrSignatureB
	}
	if result == "E" {
		return 0, "", utils.ErrSignatureE
	}
	if result == "F" {
		return 0, "", utils.ErrSignatureF
	}
	if result == "Z" {
		return 0, "", utils.ErrSignatureZ
	}
	if result == "0" {
		qData := ""
		for _, k := range keys {
			if k != "vpc_SecureHash" {
				if qData == "" {
					qData += k + "=" + strings.Join(r.URL.Query()[k], "")
					continue
				}
				qData += "&" + k + "=" + strings.Join(r.URL.Query()[k], "")
			}
		}
		//
		secretKey, _ := hex.DecodeString("B157D0AB54E32DF09156BF5E4D7E9988")
		// secretKey, _ := hex.DecodeString("6D0870CDE5F24F34F3915FB0045120DB") //test OP
		// secretKey, _ := hex.DecodeString("6D0870CDE5F24F34F3915FB0045120D6") //test OP
		hmac := hmac.New(sha256.New, []byte(secretKey))

		hmac.Write([]byte(qData))
		//
		signature := hex.EncodeToString(hmac.Sum(nil))
		if strings.ToUpper(signature) == r.URL.Query()["vpc_SecureHash"][0] {
			am, err := strconv.Atoi(r.URL.Query()["vpc_Amount"][0])
			if err != nil {
				return -1, "", err
			}
			return am, r.URL.Query()["vpc_MerchTxnRef"][0], nil
		} else {
			// return 0, "", utils.ErrSignature
			return -1, "", utils.ErrSignatureOther //return -1 khi sai hash
		}
	} else {
		return 0, "", utils.ErrSignatureOther
	}
}
