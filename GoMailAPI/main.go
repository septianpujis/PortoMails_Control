package main

import (
	H "GoMailAPI/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", H.IndexFunc)
	http.HandleFunc("/newForm", H.NewForm)
	http.HandleFunc("/submitForm", H.SubmitForm)
	http.HandleFunc("/editForm", H.EditForm)
	http.HandleFunc("/deleteAction", H.DeleteForm)

	http.ListenAndServe(":8080", nil)
}
