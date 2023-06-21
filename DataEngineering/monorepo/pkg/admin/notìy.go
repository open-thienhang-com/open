package admin

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"api_thienhang_com/pkg/utils"
)

func SendEmail() {
	utils.SendEmail([]string{
		"hangtuanthiendl@gmail.com",
		// "huunghi20061997@gmail.com",
		// "phamvanviencr94@gmail.com",
		// "nhoangvy767@gmail.com",
	}, "me@thienhang.com", utils.E_VERSION)
}

func SendLineNotify(message string) {
	data := url.Values{}
	data.Set("message", message)
	req, err := http.NewRequest("POST", "https://notify-api.line.me/api/notify", strings.NewReader(data.Encode()))
	// Header - API get user information
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer sEO55zpMBLKHoKZu4yV4QL2VHZk8Mp8z1zjzVW2Hy39")
	if err != nil {
		return
	}
	//
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
}

func SendNotification(content string) {
	data := url.Values{}
	data.Set("text", content)
	req, err := http.NewRequest("POST", "https://api.telegram.org/bot5682898713:AAEmUzvzx_ZSkmz8aY8SEtJaS2gb_nvhS6o/sendMessage?chat_id=-880396483&text="+content, strings.NewReader(data.Encode()))
	// Header - API get user information
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer 5682898713:AAEmUzvzx_ZSkmz8aY8SEtJaS2gb_nvhS6o")
	if err != nil {
		fmt.Print(err.Error())
	}
	//
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
}
