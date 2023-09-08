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

//func TestPostStruct(t *testing.T) {
//	dt := TestApi{
//		Phone:      "+6285155476774",
//		Password:   "#P@ssw0rd",
//		FirebaseId: "123",
//		DeviceId:   "6580fb6e714844ca",
//	}
//	url := "https://awangga.requestcatcher.com/"
//	res := PostStruct(dt, url)
//	fmt.Println("PostStruct : ", res)
//}
//
//func TestRequestStructWithToken(t *testing.T) {
//	dt := Sister{
//		Id_sdm: "8fe6735c-6e28-43e7-9eb3-3ae092bbcd62",
//	}
//	url := "https://awangga.requestcatcher.com/"
//	res := GetStructWithToken("token", "dsfdsfdsfdsfdsf", dt, url)
//	fmt.Println("GetStructWithToken : ", res)
//	dta := TestApi{
//		Phone:      "+6285155476774",
//		Password:   "#P@ssw0rd",
//		FirebaseId: "123",
//		DeviceId:   "6580fb6e714844ca",
//	}
//	res = PostStructWithToken("Login", "dsfdsfdsfdsfdsf", dta, url)
//	fmt.Println("PostStructWithToken : ", res)
//	res = PostStructWithBearer("dsfdsfdsfdsfdsf", dta, url)
//	fmt.Println("PostStructWithToken : ", res)
//	res = GetStructWithBearer("dsfdsfdsfdsfdsf", dta, url)
//	fmt.Println("PostStructWithToken : ", res)
//}

type Data struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func TestGet(t *testing.T) {
	myData, err := Get[Data]("https://dog.ceo/api/breeds/image/random")
	// data = message.Data
	fmt.Println(myData.Message, err)
}

//func TestGetStruct(t *testing.T) {
//	dt := Sister{
//		Id_sdm: "8fe6735c-6e28-43e7-9eb3-3ae092bbcd62",
//	}
//	url := "https://awangga.requestcatcher.com/"
//	res := GetStruct(dt, url)
//	fmt.Println("GetStruct : ", res)
//}
