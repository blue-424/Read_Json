package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"test/model"
)

func main() {
	jsonFile, err := os.Open("cnf/sql.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened sql.json")
	defer jsonFile.Close()
	data, err := ioutil.ReadFile("cnf/sql.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	var obj model.Configuration
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(obj)
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, user")
	})
	log.Println("Starting server...")
	l, err := net.Listen("tcp", "localhost:8085")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		log.Fatal(http.Serve(l, nil))
	}()
	log.Println("Sending request...")
	res, err := http.Get("http://localhost:8085/hello")
	if err := http.ListenAndServe(":8085", nil); err != nil {
		log.Fatal(err)
	}
	log.Println("Reading response...")
	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		log.Fatal(err)
	}
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}
	bodyString := string(body)
	log.Print(bodyString)
}
