# atapi

An Iteung API Interface HTTTP Method


```go
package main

import (
	"fmt"
	"github.com/aiteung/atapi"
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
	url := "https://apps.dev.rayain.net/api/gateway/auth/login"
	res := atapi.PostStruct(dt, url)
	fmt.Println("TestPostStruct : ", res)
}

func TestRequestStructWithToken(t *testing.T) {
	token := atapi.Token{
		Key:    "token",
		Values: "dsfdsfdsfdsfdsf",
	}
	dt := Sister{
		Id_sdm: "8fe6735c-6e28-43e7-9eb3-3ae092bbcd62",
	}
	url := "https://sister.ulbi.ac.id/ws.php/1.0/dokumen"
	res := atapi.GetStructWithToken("GET", token, dt, url)
	fmt.Println("GetStructWithToken : ", res)
	dta := TestApi{
		Phone:      "+6285155476774",
		Password:   "#P@ssw0rd",
		FirebaseId: "123",
		DeviceId:   "6580fb6e714844ca",
	}
	res = atapi.PostStructWithToken(token, dta, url) 
	fmt.Println("PostStructWithToken : ", res)
}

func TestRequestStruct(t *testing.T) {
	dt := Sister{
		Id_sdm: "8fe6735c-6e28-43e7-9eb3-3ae092bbcd62",
	}
	url := "https://sister.ulbi.ac.id/ws.php/1.0/dokumen"
	res := atapi.Get(url) 
	fmt.Println("Get : ", res)
}

func TestGetStruct(t *testing.T) {
	dt := Sister{
		Id_sdm: "8fe6735c-6e28-43e7-9eb3-3ae092bbcd62",
	}
	url := "https://sister.ulbi.ac.id/ws.php/1.0/dokumen"
	res := atapi.GetStruct(dt, url) 
	fmt.Println("TestGetStruct : ", res)
}
```
## Tagging

```sh
git tag v0.0.1
git push origin --tags
go list -m github.com/aiteung/atapi@v0.0.1
```
