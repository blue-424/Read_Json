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

	//0000000000000//

	//resp, err := http.Get("https://example.com/api/endpoint")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(body))

	//00000000000//

	//jsonData, err := json.Marshal(model.Supervisors{})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//resp, err := http.Post("http://127.0.0.1:8085/hello", "application/json", bytes.NewBuffer(jsonData))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer resp.Body.Close()
	//fmt.Println(string(body))

	///////////////
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
		//IP//
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Error retrieving IP address", http.StatusInternalServerError)
			return
		}
		if ip != "127.0.0.1" && ip != "::1" {
			http.Error(w, "You are not authorized to access", http.StatusForbidden)
			return
		}
		//method//
		if r.Method != http.MethodPost {
			http.Error(w, "The request method is not allowed", http.StatusMethodNotAllowed)
			return
		}
		//header//
		contentType := r.Header.Get("Content-Type")
		if contentType != "" {
			fmt.Println("Content-Type:", contentType)
		} else {
			fmt.Println("Content-Type header is not present")
		}
		//0//
		data := `{"function": "sum", "parameters": {}}`

		var p model.Payload
		if err := json.Unmarshal([]byte(data), &p); err != nil {
			fmt.Println("خطا در رمزگشایی JSON:", err)
			return
		}

		if p.Function == "" {
			fmt.Println("فیلد Function خالی است.")
		} else {
			fmt.Println("Function:", p.Function)
		}

		if len(p.Parameters) == 0 {
			fmt.Println("فیلد Parameters خالی است.")
		} else {
			fmt.Println("Parameters:", p.Parameters)
		}
		//0//
		//pay := model.Payload{"Function": "function1", "Parameters": func() {}}
		//fieldName := "Function"
		//value := reflect.ValueOf(pay)
		//field := value.FieldByName(fieldName)
		//if field.IsValid() {
		//	fmt.Printf("The struct has the field '%s'.\n", fieldName)
		//} else {
		//	fmt.Printf("The struct does not have the field '%s'.\n", fieldName)
		//}
		//0//
		fmt.Fprint(w, "Welcome!")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "خطا در خواندن درخواست", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		fmt.Println(string(body))
		fmt.Fprint(w, " Hello, user")
	})
	log.Println("Starting server...")
	l, err := net.Listen("tcp", "localhost:8085")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(l, nil))
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
