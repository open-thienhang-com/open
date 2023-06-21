package notification

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"

	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func sendNotifyFirebase2(title, content, icon, to string) error {
	message := map[string]interface{}{
		"data": map[string]interface{}{
			"notification": map[string]interface{}{
				"title": title,
				"body":  content,
				"icon":  icon,
			},
		},
		"to": to,
	}

	data := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(data)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(message)

	req, err := http.NewRequest(http.MethodPost, "https://fcm.googleapis.com/fcm/send", data)
	if err != nil {
		log.Error(err)
		return errors.New("can not make a request by abnormal reason")
	}
	// Header - API get user information
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", os.Getenv("FIREBASE_FCM"))
	log.Warn(req)
	//
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return errors.New("can not send request to API ")
	}
	defer func() {
		resp.Body.Close()
	}()
	log.Error(resp)
	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusUnauthorized:
		return errors.New("forbidden API")
	default:
		return errors.New("abnormal error wwith api authen")
	}
}

func SendDepositSuccess(amount, to string) {
	log.Info(">>> SendDepositSuccess " + amount + to)
	err := sendNotifyFirebase("Sandexcare.com", "Bạn đã nạp thành công số tiền "+amount, "https://static.wixstatic.com/media/9e8e46_64e6516c73214d429b3afcb683228d35~mv2.jpg/v1/fill/w_1876,h_1139,al_c,q_90/9e8e46_64e6516c73214d429b3afcb683228d35~mv2.webp", to)
	if err != nil {
		fmt.Println(err)
	}
}

func SendDepositError(to string) {
	log.Info(">>> SendDepositError " + to)
	sendNotifyFirebase("Sandexcare.com", "Nạp tiền chưa thành công vui lòng thử lại.", "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAANcAAADqCAMAAAAGRyD0AAABX1BMVEX////65TrWxjMREiRr04Qjd0YAAADLy80jVzFt1oYgdETn7O9WmFsdcUJr04Vu2If/5jXZxi9i0oYAABMAABv55Chcv3YLDCAAABcme0lTtG/54yNlzH/55TL+/vg6klk9l1sAbjWsqKX+/fH787JMq2lavXVGomMxiFLaxS/Uwx/898r67Hv56FP9++n66m389b/78J1oaXFDRE+Rkpl7e4P56Fb9+dn55kTp30HG2VSA1Xr67okAai3a5eSQzmoxflF5pYulwbOozFlZo2Pl244cTCosYTfz782czWXYykbr5ari2Hw1N0QdHi5cXWX68qj68Jbp40LD2Veo2mfS4FBak3C1yr+Qs6FpnX+L13fGxT530HpBhV2qw7jf6OnCyEZJh089dkbe0WdGY0qbqpk0YD97l3+TvJl9x42GjoNzqH2dn5hacF0OTiJwgXGxylHi2YGcnKS3t7opKjcbGT2sAAAVTklEQVR4nO2diV/T2NrHpU0LSU1K6ZqSrtjFsohtAzg6CAyKo7ZVZgZwQe6M9973fe87472W/v+f92xJTtIsp3v17c9RaOtgvzzrec5JuHNnoYUWWmihhRZaaKGFFlro/6myG1DZWb+N8Si7s3nv4daD7aVMXFNm6eDB1sN7mzvfJiIgeri1F19fj2eAlkwCT8AX9rYebu7M+n0OpP17W9uZdSuOVRAvs711b3/Wb5dJ2d0nB8DbPJgoOOCZTzbn3Ck3drfi66xIFFx8fWt3bl0yu7sFv/nDCfyfW5uzJrDTztul+OqQUBra0tt5M9rmgyHczwZt/cE8GW13byxUmGxvd9Y4RLvbQ0eVLVl8ex7Idg/iY4TCWj+YNdnm3vrYqRDZ3izjbGNrrB5IKxPf2pgV1puJUWGyNzOh2t8ef2CZFd+bQev4dqLGwsrE306ZavLGwopvT7UDuTe2OuylzPq9qVFlt6ZjLKz41pQWMTsH0zIWVuZgKr64O4WEYQGLT6H9eDiZBsNd8YeTxnoyCywA9mSiVNkH08wYJrAHE8we2b1ph5ahzN7EwDa2Z4cFwSbUCG9MOb/3gW1PBGzWWLCQTQAsO1MnJGDjj7FZpgxD4wd7MA9YMN2PF+vJrOqWVeMt0A/nBWtpaX2M04Hd2TRP9lofWxO8Mx+xpSkzrmXLwaxJzMocjAdra77MBcC2xoF1b35yhqb4GGYeO/OHBcBGD7E5aJ9stD0q1tt5NNfog4H9ubQWUGa0ETdTt7s62r7yUMrsjYL1hsELzz/+/MOj6bON0k9tMGCt/rAGBNlWV6fqtPHhF5ksFXn11zWO46S1NQmzTc1uw1fnfaZc+FjiiIjdpuWT8WFTB9Na8gCai+N5jCbpPjl5tsyQa0ym1cnqb5ArXCkp4TCP6bBP/jJ5tvXh9taZ5k84vPiAKAq5UoLYTZKgT3KAbbI+OVTXwbaYXP0IwyshBgRBEEUZsoU1p4Q++XSSbEMtMbeZvjQKLz4tCvkcMBlgEyAbHwZovBZvTycWb0MYbJetMUThxecEmQtziVI1gNjEgOGT0CuBT07GbkNsjLHNC0n1koVqGPKFpUSlKkPDYZ+UCJsEDQdyyfmY2QbvpjbZzLWKqldCECrEOCAnKolKThD1eJPCdA14+sv5OH1y4BrGuPg/RywlUSjznC6Y74FPCshuAvZJAqf55LjYBm06dtgma7h68XlB4KziUbwRn4RsCl27Adt4fHJ9sJXzQzZzmcPLhk1JlHJUvHGaT3Lj8snMYAtM1q/6HoaXIurh1Y+GfBLHmyDIVVADgCStL+Ew2yhwg2AxZo2lc1K9xLITF7EbD+0WCARgzFVpn5RG9cn4IM0UY9YgzaFdeFkkge6xnAsgAaeEPqmQvgTXgKHZBskcG4xfkzSHILxczQUNlsjLghDQBXNJtQL7EpInSS4ZwiUH2MRk7DVI9XILL/ym+URVpKACut3EKupLjBqwhuw20Lp7gCaReXLNEl4Sr+RtqCifrECf1P72wD7J7ohZViytegUUF2uFy7KVqlCgwFANAD6prwMkabAakGHl2mTc7gLhBf2w6hpeoBmxGuno8ErHyqVzsoDhqpWytvDmUby9R3bzhGNeXj5htBcJL8EtvPiKFUsoHIcOdYOJeVwDyCInVykrPNWYsNSADOvWLBuV1hy6hlfYjCUUCkcXF5ehqx911xQrYbQOKOW1vgTYDfTKuHajeAN2y7ixMe6HsY2h6ObQMbwAtCmujk4OQ1CHJxcFjUwswR4MrmgUYDcBsQmQjc6T2Ccd3gjj7grrhheuXjC8bJtDWHETZqzjkKHDd5ovignN3KCcaWwCYlN4apHDvT+3DxDG1SVrll+1aQ55yeiRwKOqCesqZNKl5pwy/a2gfFJAPqkY85IP9jvCjJmedcfrALwLa3jlQeKulon9zKlQeBey6PDICDGz+wKfTOfwTEEIoDwJmT/6fPbvhGnMscNGZVQvUR/48jIEEcQ0BlVMebCAQ+sKfHx3jD6/1ELMJkBBcIWVdF7zSfhPSE8jDmBMJwRYe3kcXmGqOdRyuiCj92nJhRcI67hwHQqBvPgOkF3rXA6FAtktL4MqB19f+w/gsgVj6ukZl5Q24cXLWibABpNprMIJcr0CcEdYvYTA4ZXRdQQkWy7MlhDEEvp6H3w+ezCmK3SGbg4l3e8QKl82J8NLyHVSKHwCf6AnqO5K81xbsJIoJuAnHyOIK9L/TlgSR3aPEauvOZQk3V7w+wtf6uP6VIDhFbAKO5oDVw779drT+z4nMIbje6xNL1l70c1hOCcaeUAKm9yQZPkb+QiEV78Elw6TTE/WfsL2sgPLeHMxTqJAeEEuxRTzCl464rytWBrDa1y03lHNIeWICUewhCig8Fr74PM5gTFMpVib+SWtelGhAdePIkkklvDSHDF0Q8LLwuUYYCi80Iv3fY5gDAmRsYuybw55rlxK44ajf31yTcrxsQ2Xy0ArJ6LpCapeTmAMB4sY0jwcjKHw4vm+tZe+yOAr1tWkcEHArgp9y2ch7+iHfeGF9Mj0jhimiF5p/vz8t19+/fXpe/RvwrWXQ9Mb7uMKFDSwS+srLlw24dUPxrAE89hUXl1DklAhtYSX2XD9XIHC9SHpDK2jAUcuLbxgc+gC5r3V7HGIcvUjbLqN9y46rb3suPQmMXRZsOGyazrCOYGE130rGP1OvZeWHm6IuycOf3dBcyg7uCH4RttxHd2AhIhjzMwF8gbPK4pkNRsfwBFsCa8+ME8ujzS/+hR1hTnU2Uhug/l037wGCFQvnBZvjkzPgw4FFImAKFcsRtOmJ9KHPiy6VVwfmQu1GYpI3rpLY2ety8gNT0InP+JV87HJnKDJLKMho1i1fHcE0hzaYFFgnlxZL65f0OIf9XMg5TmPNjjexg9Rc/jjIe4UTVz6Yk2kPYAP50XUHNqElwls3auR2vDi+gG38TDM+6sXLRDxFBAJr1BIJo2HKXOAEqW5LT0V4HjJvnr1ga17Dem92kPcZ3AK5FHcyqkpwC5OftTCq2DHJZaocke7gBZeax8cuCJj4lo6X5OMNy6WndIhR48BgPtdgpUkDK9PBUHoT4iiYrRdMsUFD4aUrc2hncVG58LnvJCjgfDKp41RWJ/BtAWYAJcoh8cFAUTWuwJeOIcuaDfMgWyo+WEubPoSuHq9d+aKjIdrKaxn4qoYgGM+JzawgCfvmzQZx3DtdYSxTIsV2FGADIGwzJnIK7w0ME8ur7xBnzcsV6p4y0CQARvHW9nCmsGoeSjpN0IXgslcyDZgjSPKpoWYNj1xDC8NzDMfeuX5paX3RuXkwwpgI2eHIJtkZpPImy9cX1pGh6boIiYKK+lSmaO/gESmJ/3NoVmPvLm86jJpOHTjwNEztBsZPefTCconjZRYODaRHVuSofb3td0Gw+A2ay87MO9+w2tZSY5sVIxzXejsEGDTzrPBeCOv8XkNTChcX2nN/MkRXbJdZjbacHLtJ+e0geXdH3q5IWk4BDFXTickii0sleGZL90n8RmUqoEAll+HF+8u5IKp0xCqtn28Hl4obfzNw14+nyeX15iNNByg1Ti9e/fuabrfboLOBvdWDbDCFapeAZME2aUTA80hCi/Oi8r30pPL87AybDgkONSDXHcxG203jvbJXFrR26nC4Z3jgEWWNtfihmQ46Va9sCKfPbk8N2HPte2Gu5ROQS7kKJ/kqDxpULyTrVgufRiHqpfT2sus2jNPLs9LN/Dua8XMZWM3UgMcjzgAKjnt0oZxkt4ceoZX7ZUnl+ec7eAjh1bD4u9WMMIW1hI9lSftqMSK4n6qSpv9S55Zo+a9Y7nvVcAyj/FA9Izm+mNlZUV74o9cmhq34TwZsLAJcGXsZizOmP17h5ev5n1s1LORwhOOhLhCcZ2tIOFnzp4/XwFsRtNIfFI22ODRKN45vxNVGZpDwuV9RirrdVUvmXAIK3+YrIVEQWI2k08mcA2AXHl0DtEdy2gOvbB8PobbIXglejLhCFBcZxrXH6ZHwG5nubQUNnyS5xJ6zxXIlRTJaY3D0VtrnlQMad470ZOGQ14507lWTFy69TDbc+STiE3Sa4BxNtaRDZYSji28It5p3jshkoZDfs7ERez2B7KbvtK2nkU3rv0wxN4csqR5huM2+JAo3EVMpE9Pac97jhLH71Yu9NLfi/8gPsnrPqmfRUfXfmgmtYaXd3NYYzkPsOGVOLSGQ0aHEXgJ4IlnEE2zoB3Xyj+LxZXnX/7rv/9H4qg8ic+i48Upfe2HEV4MzaHPx3Rk1KvzPcBXfIHiojcX8DyoCDt8ZL8zW7Bi8Z8rX/5VLBb/8b/WPFnKBczXRwCPDZOtNYbw8rGkjTt33nolDmSvkigYS3a4n1eCyQF+g5XS2Vk/2pdi8e8rK38CruJfKE9yNJsE2ETKJ8M8DC+OMbxY0ob3wRR0yZf1eB5djcLVgCzLoFHC7vkFc/1Z/OvLvyFWkSSTs5yxAEVfAPokSSa5UppMflnC6wUTV9aLCzYcfFlw3voyJqJyQHj36UL8AvSv4l8Y699f9GSycorO0vMaGYw37foIsueMT9m4czHepMgjwFDDwSvaKRg70avkE7Dyv7y6urq8WV5evi0W/6SSJMqfp6cJqoah82w5mCfx4Ne7i2KqylAeSxWt4XDZI6LOHWo7ecYhNuieOP6MCng3bTrfB30yh481Sj97xRdT9YLyaOlRwyHxstvAxTjydWTG+qT5Jwy/0unpqYEGfdIYs4VJt++0k0JxMV9S5H4CETccYderHIyxhvzpxrDVyTU1swHxA6cGSoKUB43N3FdJj724vGcbmtwzPbkqpSq4zSYUfWe8ELg+Pjn5dHISOiyYx4b6tQBo99XgO01Q6dVrJsqa5Rkc8ZxcbSPbI2GDKfSWfwEqdEiPN2wOiYJ8CLxPAXi//56HBRo978nFfili1uMaX1SYXTcrrWCQLRS6trMWp/EkEulSJZ+rgtInoFF9Hp5/lXiP5Re7G3retwy9m5LoeMaBKG8+O39ype+VC0IZdBTwP9hdlioYBjXB+NoH+J9ctT2QYjUXazaE8rgJ0WMONxwJd64w6Blog+nRJVSJbWRkGEGjATVLrlZzlVK6rEg86bM8Gg72bAjlumjGDUfC7TAAAVMqgt1ph4CsmwbjAJo8pAGJnoM89Nd1bziYizKW680BVp/isw4uDQc0qATjhitZZ6HkEiIUQcg2CQXOVHneMvCA1zvAy2/Wfnbjqg14Twc3NyRbKvaHUnjsPwqKmzyMGxMO8DSZ8rSwfhWpjiOh69vgv/Dx/dOff/rP31zDy3vDwSy3zIEnHDy9FQwXmHCJqWg5TRD7PC1HPI3X79RB43Do3hZr3EdC88EXuX//fgTINboGyRpQbvcG1CYcQhVmNJSikW2AcYw8QHtaWVFscXh0XSXG0Wh8kQjCcbeSroFv3uM2lvqNbBXJaZjTAoI5p8GkRmwDrwvto6E97bFum0jEyzZ25mLvNXSDuTjiI+2IuV5viKflEY2ETcNj/0J/8BSOydOGoaG4hrj/nMvBUbKlQnCIbUCKDltzGkeucgXiNZqXKHDuj0CjK/J6cCy3kxwH+OYvMG5IgjZnaGggeMGWtEZ5GoibyFhoDA1jLjeDoS0ViTPjSBLlaY/pnDZAGhgM6+swWC4pcZU6wwEjSDLnND1FTwKGUmTIuzs63vsQNRySRqPZxqdltQnjaBoiGWI5Tn5Xf5V0mvuj5rThNfS9sh23IA58U3E0Vw3calByHnTMkggpMsh60iqXvZVZcw2X4zU53zb1YMZYwyYNrKzzzecezZTr5Yg32HcZTc0yawwwhHKQyw18Zwc2ohciudyobVZYI+VCTW4rzFmBjeUO9C4znNnkjkFnNU5y+aEVswAbR3BhuUwTp587aoNNDN3k9iN8po0VGbVy0Zqj3DFa/2SVy9Xa022oGPf+meVya7Np5o5xpUJD9+ahoRplzeUklx+EOLWV/wSwQKfoDDYlrLEVLrNcLPYtY92588YRbApJcTJOiOWcPCaeFCeJBdK9Y4GecO4Yf4I3az/jRDZJqkhkzOW4X84/pXhyWLWXU/hJxdkthyCbWO6oDbMdNISc8v2EcsdkMwatTYcgm0TuiPgmHlqGNh7Ym2z8YLXPU/rp5kT37E02ZqpIZGo+qGlnz9ZkY8WqfZ7KT2y3yNZkY8wdNd/UjYW18WS9n2xcYJHa1wn9HGkGbdo443hyR+3lFNOgjXaX+gYEYwCr+SbcDjLozYGVbHRbzSiwzMpayUZrqEC6mG7JclZ298CUQYbPHZHay7H9uN6xaPNBJm6gDRlitcjn2WYLO+083DaMNgQYMNWrWZRhBu0/iWtoA0PVvr6Yl7Cy0+aTA+SQA+QOwPTy6/z5X5/23zw4WI+zWQww+T6/GnkTfFrK7r/ZelQDco6zCECq+V6/2pln77NV9sWrZ59fQpNAQKIafuR7+fnZq7kOKC9ld/ZfbL569ezr19evX3/9+uzZq80X+9+ekRZaaKGFFlpooYUWWmihhRZy0PL3qTvB71N3/N+nFlzflghXkvz2Ux/9/ljMn6Qewd/Gw/kW5ko2kv5kr40/b7XJa6l6PdVr6CSdmL/daX0jYJgr1qnHUs1oNOVPRYNqLwY+S0aDxW73tl4MRoPBZDIYjN0Egy218U1x+dvNaFtVm82g2lTVbq/ZVBvdznI9GuyE6s3l5UbrcLmx3LjptaZpLxADOD7An8k2fpSEgUE+Q58kk+22H/6RhGETa7dNXNFmu15XY3W1Hgyq/qY/WFc70dbyTbPTbS8HG7e3jWA0tAzCa3pYyUZHbffAG21EezF/R200Ur1Uu60C72qlWuD5XrTeUTuNZq/bajQbzU6zo4IPxKEIV6zTgS/U1UYsqra70RR4EIsGo8tqt3Eb7N0up6KAqxWbGhUQcCD1tt5s3ILv+G2j2bpVVfBouRNUO2pdbXbBh6DaqKvBbqvTBZ916912t66maC5/rAjM1G71uu1ep16vt5udTqyrNpd7N61lFXzFZbW13AtFp8mVAiFRBwboQopWXe02wcdGUe01Gyq0QqvbUOutZhegqbfgdUCt1jvNqJmr3o71VDXZVuvRThB6QCvZaKqxaKeRataj0U6zVw/Wp5s12v5UI9nrtZONWA/9bkXBo14r1Yg2Uu2ev9VuNVrRtr8RjYLwijZi/hb8w8QFK1MyFUO/wa8UjM9kKgWfh0/BZ2IzqF4oMfipX0ntsfYQ/SW//kF7h995v/HdacH1ben/ANhFuEZq1iKQAAAAAElFTkSuQmCC", to)
}

var opt = option.WithCredentialsFile("sandexcareapp-firebase-adminsdk-32yn0-344a638234.json")
var config = &firebase.Config{ProjectID: "sandexcareapp"}
var appFB, _ = firebase.NewApp(context.Background(), config, opt)
var ctx = context.Background()
var client, _ = appFB.Messaging(ctx)

func sendNotifyFirebase(title, content, icon, to string) error {
	// This registration token comes from the client FCM SDKs.
	registrationToken := to

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"title":   title,
			"content": content,
			"icon":    icon,
			"time":    time.Now().String(),
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Error(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
	// [END send_to_token_golang]
	return nil
}
