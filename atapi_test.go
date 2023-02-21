package atapi

import (
	"fmt"
	"testing"
)

type TestApi struct {
	Phone      string `json:"phoneNumber"`
	Password   string `json:"password"`
	FirebaseId string `json:"firebaseId"`
	DeviceId   string `json:"deviceId"`
}

type Sister struct {
	Id_sdm string `url:"id_sdm"`
}

func TestPostStruct(t *testing.T) {
	dt := TestApi{
		Phone:      "+6285155476774",
		Password:   "#P@ssw0rd",
		FirebaseId: "123",
		DeviceId:   "6580fb6e714844ca",
	}
	url := "https://awangga.requestcatcher.com/"
	res := PostStruct(dt, url)
	fmt.Println("TestPostStruct : ", res)
}

func TestRequestStructWithToken(t *testing.T) {
	token := Token{
		Key:    "token",
		Values: "dsfdsfdsfdsfdsf",
	}
	dt := Sister{
		Id_sdm: "8fe6735c-6e28-43e7-9eb3-3ae092bbcd62",
	}
	url := "https://awangga.requestcatcher.com/"
	res := RequestStructWithToken("GET", token, dt, url) // POST DELETE PUT
	fmt.Println("TestRequestStructWithToken : ", res)
}

func TestRequestStruct(t *testing.T) {
	dt := Sister{
		Id_sdm: "8fe6735c-6e28-43e7-9eb3-3ae092bbcd62",
	}
	url := "https://awangga.requestcatcher.com/"
	res := RequestStruct("GET", dt, url) // POST DELETE PUT
	fmt.Println("TestRequestStruct : ", res)
}

func TestGetStruct(t *testing.T) {
	dt := Sister{
		Id_sdm: "8fe6735c-6e28-43e7-9eb3-3ae092bbcd62",
	}
	url := "https://awangga.requestcatcher.com/"
	res := GetStruct(dt, url) // POST DELETE PUT
	fmt.Println("TestGetStruct : ", res)
}
