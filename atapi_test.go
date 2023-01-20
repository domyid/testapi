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

func TestAtapi(t *testing.T) {
	dt := TestApi{
		Phone:      "+6285155476774",
		Password:   "#P@ssw0rd",
		FirebaseId: "123",
		DeviceId:   "6580fb6e714844ca",
	}
	url := "https://apps.dev.rayain.net/api/gateway/auth/login"
	res := PostStruct(dt, url)
	fmt.Println("res : ", res)
}
