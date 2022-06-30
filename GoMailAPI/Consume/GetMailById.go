package Consume

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type MailResponse struct {
	PortofolioMail struct {
		ID            int       `json:"Id"`
		Email         string    `json:"Email"`
		Name          string    `json:"Name"`
		Subject       string    `json:"Subject"`
		Message       string    `json:"Message"`
		AddedDateTime time.Time `json:"AddedDateTime"`
		IPSender      string    `json:"IPSender"`
	} `json:"PortofolioMail"`
	ResponseBody struct {
		CodeMessage string `json:"CodeMessage"`
		Message     string `json:"Message"`
	} `json:"ResponseBody"`
}

func GetMailById(Id int) (item *MailResponse) {

	url := "https://personal-i6nfgh8p.outsystemscloud.com/Septagonal_DB/rest/PortoFolio/GetEmailById?PortofolioEmailId=" + strconv.Itoa(Id)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var Item MailResponse

	json.Unmarshal(body, &Item)
	//fmt.Print(string(body))

	return &Item
}
