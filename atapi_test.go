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
	Id_sdm string `url:"id_sdm" json:"id_sdm"`
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
	fmt.Println("PostStruct : ", res)
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
	res := GetStructWithToken(token, dt, url)
	fmt.Println("GetStructWithToken : ", res)
	dta := TestApi{
		Phone:      "+6285155476774",
		Password:   "#P@ssw0rd",
		FirebaseId: "123",
		DeviceId:   "6580fb6e714844ca",
	}
	res = PostStructWithToken(token, dta, url)
	fmt.Println("PostStructWithToken : ", res)
}

func TestGet(t *testing.T) {
	url := "https://awangga.requestcatcher.com/"
	res := Get(url)
	fmt.Println("Get : ", res)
}

func TestGetStruct(t *testing.T) {
	dt := Sister{
		Id_sdm: "8fe6735c-6e28-43e7-9eb3-3ae092bbcd62",
	}
	url := "https://awangga.requestcatcher.com/"
	res := GetStruct(dt, url)
	fmt.Println("GetStruct : ", res)
}
