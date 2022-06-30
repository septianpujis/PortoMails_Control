package Consume

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ResponseBody struct {
	CodeMessage string `json:"CodeMessage"`
	Message     string `json:"Message"`
}

func PostMail(email string, name string, subject string, message string) (errmsg string) {

	var Message ResponseBody

	url := "https://personal-i6nfgh8p.outsystemscloud.com/Septagonal_DB/rest/PortoFolio/PostEmail"
	method := "POST"

	input := "{ \"Email\": \"" + email + "\", \"Name\" : \"" + name + "\", \"Subject\": \"" + subject + "\", \"Message\": \"" + message + "\"}"

	payload := strings.NewReader(input)

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

	json.Unmarshal(body, &Message)

	if Message.CodeMessage == "1" {
		return Message.Message
	} else {
		return ""
	}
	//fmt.Println(string(body))
}
