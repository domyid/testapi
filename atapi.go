package atapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func PostStruct(structname interface{}, url string) (result interface{}) {
	mJson, _ := json.Marshal(structname)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(mJson))
	if err != nil {
		fmt.Println("Could not make POST request to whatsauth")
	}
	fmt.Println(resp)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Read Data data from request.")
	}
	json.Unmarshal([]byte(body), &result)
	return result
}

func Get(url string) (result interface{}) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Read data from Response.")
	}
	json.Unmarshal([]byte(body), &result)
	return result
}
