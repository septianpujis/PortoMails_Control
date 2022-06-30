package handlers

import (
	C "GoMailAPI/Consume"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type M map[string]interface{}

func IndexFunc(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/index.html")
	checkError(err)

	err = t.Execute(w, C.GetMailList())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("form").ParseFiles("templates/form.html"))

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func SubmitForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		tmpl := template.Must(template.New("result").ParseFiles("templates/form.html"))

		var err = r.ParseForm()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id := 0
		i, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			id = 0
		} else {
			id = i
		}

		if id != 0 {
			if err := C.DeleteFormById(id); err != "" {
				http.Error(w, err, http.StatusInternalServerError)
			}
		}
		name := r.FormValue("name")
		email := r.FormValue("email")
		subject := r.FormValue("subject")
		message := r.FormValue("message")

		errmsg := ""
		if err := C.PostMail(email, name, subject, message); err != "" {
			errmsg = err
		}

		data := M{"errmsg": errmsg}
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "", http.StatusBadRequest)
	}

}

func EditForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := 0
		var form *C.MailResponse
		if i := r.FormValue("id"); i != "" {
			i, err := strconv.Atoi(r.FormValue("id"))
			if err != nil {
				id = 0
			} else {
				id = i
			}
		}

		form = C.GetMailById(id)

		tmpl := template.Must(template.New("form").ParseFiles("templates/form.html"))

		if err := tmpl.Execute(w, form); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	} else {
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func DeleteForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		tmpl := template.Must(template.New("result").ParseFiles("templates/form.html"))

		var err = r.ParseForm()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id := 0
		i, err := strconv.Atoi(r.FormValue("Id"))
		if err != nil {
			id = 0
		} else {
			id = i
		}
		errmsg := ""
		if id != 0 {
			if err := C.DeleteFormById(id); err != "" {
				errmsg = err
			}
		} else {
			errmsg = "Id not valid"
		}

		data := M{"errmsg": errmsg}
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

	}

	http.Error(w, "", http.StatusBadRequest)
}
