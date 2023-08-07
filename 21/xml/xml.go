package xml

import "fmt"

type xmlService struct {
}

func (x *xmlService) toJson() {
	fmt.Println("convert to json")
}

type xmlAdapter struct {
	xmlService
}

func (x *xmlAdapter) sendJsonData() {
	x.xmlService.toJson()
}
