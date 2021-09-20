package request

import (
	"io/ioutil"
	"log"
	"net/http"
)

//Пакет предназначен для чтения URL источникa в байтах
func MakeRequest(s string) []byte {
	resp, err := http.Get(s)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}
