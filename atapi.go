package atapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

func GetWithBearer[T any](tokenbearer string, urltarget string) (result T) {
	client := http.Client{}
	req, err := http.NewRequest("GET", urltarget, nil)
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

func GetImageWithBearer(tokenbearer string, urltarget string) (respBody []byte) {
	client := http.Client{}
	req, err := http.NewRequest("GET", urltarget, nil)
	if err != nil {
		fmt.Println("http.NewRequest Got error : ", err.Error())
	}
	req.Header.Add("Content-Type", "image/jpeg")
	req.Header.Add("Authorization", "Bearer "+tokenbearer)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do(req) Error occured. Error is :", err.Error())
	}
	defer resp.Body.Close()
	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("io.ReadAll(resp.Body) Error occured. Error is :", err.Error())
	}
	return
}

func GetStructWithBearer[T any](tokenbearer string, structname interface{}, urltarget string) (result T) {
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

func GetStructWithToken[T any](tokenkey string, tokenvalue string, structname interface{}, urltarget string) (result T) {
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

func PutStructWithBearer[T any](tokenbearer string, structname interface{}, urltarget string) (result T) {
	client := http.Client{}
	mJson, _ := json.Marshal(structname)
	req, err := http.NewRequest("PUT", urltarget, bytes.NewBuffer(mJson))
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

func PostStructWithBearer[T any](tokenbearer string, structname interface{}, urltarget string) (result T) {
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

func PostStructWithToken[T any](tokenkey string, tokenvalue string, structname interface{}, urltarget string) (result T) {
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

func PostStruct[T any](structname interface{}, urltarget string) (result T) {
	mJson, _ := json.Marshal(structname)
	resp, err := http.Post(urltarget, "application/json", bytes.NewBuffer(mJson))
	if err != nil {
		fmt.Println("Could not make POST request to server : ", err.Error())
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Read Data data from request : ", err.Error())
	}
	json.Unmarshal([]byte(body), &result)
	return result
}

func Get[T any](urltarget string) (result T) {
	resp, err := http.Get(urltarget)
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Read data from Response.")
	}
	json.Unmarshal([]byte(body), &result)
	return result
}

func GetStruct[T any](structname interface{}, urltarget string) (result T) {
	v, _ := query.Values(structname)
	resp, err := http.Get(urltarget + "?" + v.Encode())
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Read data from Response.")
	}
	json.Unmarshal([]byte(body), &result)
	return result
}
