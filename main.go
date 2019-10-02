package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type addressBook struct {
	Firstname string
	Lastname  string
	Code      int
	Phone     string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "Welcome to the HomePage!")
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "home.html")
	default:
		http.Error(w, "403 Forbidden", http.StatusForbidden)
	}
}

func getAddressBookAll(w http.ResponseWriter, r *http.Request) {
	addBook := addressBook{
		Firstname: "Khomdrad",
		Lastname:  "Boontae",
		Code:      19243,
		Phone:     "66-xx-xxx-xxxx",
	}
	json.NewEncoder(w).Encode(addBook) // (1)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HelloWorld!")
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	fmt.Println("Starting the Web Services....")
	fmt.Println("Using port:8080....")

	http.HandleFunc("/getAddress", getAddressBookAll)
	http.HandleFunc("/helloWorld", helloWorld)
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Print(err)
		log.Fatal(err)
		errorHandler(err)
	}
}
