package Consume

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type MailStruct struct {
	ID            int       `json:"Id"`
	Email         string    `json:"Email"`
	Name          string    `json:"Name"`
	Subject       string    `json:"Subject"`
	Message       string    `json:"Message"`
	AddedDateTime time.Time `json:"AddedDateTime"`
	IPSender      string    `json:"IPSender"`
}

type MailsList struct {
	MailStruct []*MailStruct
}

func GetMailList() (list *MailsList) {

	url := "https://personal-i6nfgh8p.outsystemscloud.com/Septagonal_DB/rest/PortoFolio/GetEmailList"
	method := "GET"

	payload := strings.NewReader("{\"Code\": \"BKIPA11\",\"NewhealthStatus\": \"1\"}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

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

	// var prettyJSON bytes.Buffer
	// error := json.Indent(&prettyJSON, body, "", "\t")
	// if error != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	var List MailsList

	json.Unmarshal(body, &List.MailStruct)

	return &List
}
