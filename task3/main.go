package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type User struct {
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Age       uint8  `json:"age"`
}

var users map[string]User

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		userId := r.URL.Path[1:]

		u, ok := users[userId]
		if ok == false {
			http.Error(w, "User does not exist.", http.StatusNotFound)
		} else {
			b, err := json.Marshal(u)
			if err != nil {
				panic(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(b)
			w.Write([]byte("\n"))
		}
	case "POST":
		userId := r.URL.Path[1:]

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		var newUser User
		err = json.Unmarshal(reqBody, &newUser)
		if err != nil {
			panic(err)
		}
		users[userId] = newUser
		w.WriteHeader(http.StatusCreated)
	case "PUT":
		userId := r.URL.Path[1:]

		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		firstName := r.Form.Get("first-name")
		lastName := r.Form.Get("last-name")
		age := r.Form.Get("age")
		ageNumber, err := strconv.Atoi(age)
		if err != nil {
			panic(err)
		}

		u, ok := users[userId]
		if ok == true {
			if firstName != "" {
				u.FirstName = firstName
			}
			if lastName != "" {
				u.LastName = lastName
			}
			if age != "" {
				u.Age = uint8(ageNumber)
			}
			users[userId] = u
			w.WriteHeader(http.StatusNoContent)
		} else {
			var newUser User
			newUser.FirstName = firstName
			newUser.LastName = lastName
			newUser.Age = uint8(ageNumber)

			users[userId] = newUser
			w.WriteHeader(http.StatusCreated)
		}
	case "DELETE":
		userId := r.URL.Path[1:]

		_, ok := users[userId]
		if ok == false {
			http.Error(w, "User does not exist.", http.StatusNotFound)
		} else {
			delete(users, userId)
			w.WriteHeader(http.StatusNoContent)
		}
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}

func main() {
	users = make(map[string]User)
	http.HandleFunc("/", handler)

	fmt.Printf("Starting server...\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
