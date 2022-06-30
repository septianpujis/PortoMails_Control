package Consume

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func DeleteFormById(id int) (errmsg string) {
	url := "https://personal-i6nfgh8p.outsystemscloud.com/Septagonal_DB/rest/PortoFolio/DeleteMail?PortofolioEmailId=" + strconv.Itoa(id)
	method := "POST"

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

	var Message ResponseBody

	json.Unmarshal(body, &Message)

	if Message.CodeMessage == "1" {
		return Message.Message
	} else {
		return ""
	}
	//fmt.Println(string(body))
}
