package json

import "fmt"

type someService interface {
	sendJsonData()
}

type jsonService struct {
}

func (js *jsonService) sendJsonData() {
	fmt.Println("send json data")
}
