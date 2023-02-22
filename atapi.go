package atapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

func GetStructWithBearer(tokenbearer string, structname interface{}, urltarget string) (result interface{}) {
	client := http.Client{}
	v, _ := query.Values(structname)
	req, err := http.NewRequest("GET", urltarget+"?"+v.Encode(), nil)
	if err != nil {
		fmt.Println("http.NewRequest Got error : ", err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+tokenbearer)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do(req) Error occured. Error is :", err.Error())
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Read Data data from request.")
	}
	json.Unmarshal([]byte(respBody), &result)
	return
}

func GetStructWithToken(tokenkey string, tokenvalue string, structname interface{}, urltarget string) (result interface{}) {
	client := http.Client{}
	v, _ := query.Values(structname)
	req, err := http.NewRequest("GET", urltarget+"?"+v.Encode(), nil)
	if err != nil {
		fmt.Println("http.NewRequest Got error : ", err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(tokenkey, tokenvalue)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do(req) Error occured. Error is :", err.Error())
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Read Data data from request.")
	}
	json.Unmarshal([]byte(respBody), &result)
	return
}

func PostStructWithBearer(tokenbearer string, structname interface{}, urltarget string) (result interface{}) {
	client := http.Client{}
	mJson, _ := json.Marshal(structname)
	req, err := http.NewRequest("POST", urltarget, bytes.NewBuffer(mJson))
	if err != nil {
		fmt.Println("http.NewRequest Got error :", err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+tokenbearer)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do(req) Error occured. Error is :", err.Error())
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Read Data data from request.")
	}
	json.Unmarshal([]byte(respBody), &result)
	return
}

func PostStructWithToken(tokenkey string, tokenvalue string, structname interface{}, urltarget string) (result interface{}) {
	client := http.Client{}
	mJson, _ := json.Marshal(structname)
	req, err := http.NewRequest("POST", urltarget, bytes.NewBuffer(mJson))
	if err != nil {
		fmt.Println("http.NewRequest Got error :", err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(tokenkey, tokenvalue)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do(req) Error occured. Error is :", err.Error())
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
		fmt.Println(err)
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
		fmt.Println(err)
	}
	fmt.Println(resp)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Read data from Response.")
	}
	json.Unmarshal([]byte(body), &result)
	return result
}
