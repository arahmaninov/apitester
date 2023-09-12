package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
)

// Body of json create request
type Body struct {
	Name string `json:"fileName"`
}


func crud_create() {
	
	url := "http://localhost:8080/photos"
	
	data := Body {
		Name: "mypic.jpg",
	}
	
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
	
	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	fmt.Println("Status:", resp.Status)
}

func crud_get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// if the status code is not 200
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// print the response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func crud_delete() {
	url := "http://localhost:8080/photos/1"
	body := strings.NewReader("This is the request body.")
	
	req, err := http.NewRequest(http.MethodDelete, url, body)
	if err != nil {
		log.Fatal(err)
	}
	
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// url not found or the server is not reachable
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	fmt.Println("Status:", resp.Status)
}

func main() {

	url := "http://localhost:8080/photos"
	
	// Reading input
	fmt.Print("Enter text: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input.")
		fmt.Println("Please try again", err)
		return
	}
	
	// trim input end line according to user's OS
	if runtime.GOOS == "windows" {
		input = strings.TrimSuffix(input, "\r\n")
	} else {
		input = strings.TrimSuffix(input, "\n")
	}
	
	
	// Doing stuff
	if strings.EqualFold(input, "get") {
		crud_get(url)
	} else if strings.EqualFold(input, "delete") {
		crud_delete()
	} else if strings.EqualFold(input, "create") {
		crud_create()
	} else {
		fmt.Println("Something went wrong")
	}

}
