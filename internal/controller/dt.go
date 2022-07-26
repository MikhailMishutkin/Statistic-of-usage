package controller

import (
	"io/ioutil"
	"log"
	"net/http"
)

type UseCase interface {
	AskUser() string
}

type Controller struct {
	uc UseCase
}

func New(uc UseCase) *Controller { return &Controller{uc: uc} }

//обработка запроса пользователя
func (c *Controller) UserAnsw() (s string) {
	s = c.uc.AskUser()

	resp, err := http.Get(s)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	s = string(body)
	return s

}
