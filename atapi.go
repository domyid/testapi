package atapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/go-querystring/query"
)

type Token struct {
	Key    string
	Values string
}

func GetStructWithToken(token Token, structname interface{}, urltarget string) (result interface{}) {
	client := http.Client{}
	v, _ := query.Values(structname)
	req, err := http.NewRequest("GET", urltarget+"?"+v.Encode(), nil)
	if err != nil {
		log.Fatalf("http.NewRequest Got error %s", err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(token.Key, token.Values)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("client.Do(req) Error occured. Error is: %s", err.Error())
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Read Data data from request.")
	}
	json.Unmarshal([]byte(respBody), &result)
	return
}

func PostStructWithToken(token Token, structname interface{}, urltarget string) (result interface{}) {
	client := http.Client{}
	mJson, _ := json.Marshal(structname)
	req, err := http.NewRequest("POST", urltarget, bytes.NewBuffer(mJson))
	if err != nil {
		log.Fatalf("http.NewRequest Got error %s", err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(token.Key, token.Values)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("client.Do(req) Error occured. Error is: %s", err.Error())
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Read Data data from request.")
	}
	json.Unmarshal([]byte(respBody), &result)
	return
}

func PostStruct(structname interface{}, urltarget string) (result interface{}) {
	mJson, _ := json.Marshal(structname)
	resp, err := http.Post(urltarget, "application/json", bytes.NewBuffer(mJson))
	if err != nil {
		fmt.Println("Could not make POST request to server")
	}
	fmt.Println(resp)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Read Data data from request.")
	}
	json.Unmarshal([]byte(body), &result)
	return result
}

func Get(urltarget string) (result interface{}) {
	resp, err := http.Get(urltarget)
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

func GetStruct(structname interface{}, urltarget string) (result interface{}) {
	v, _ := query.Values(structname)
	resp, err := http.Get(urltarget + "?" + v.Encode())
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
