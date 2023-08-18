package tree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	root := Node{Name: "/", Children: make([]*Node, 0)}

	root.Put("/user/Get/:id")
	root.Put("/user/ha/:id/:q")
	root.Put("/user/Put/*")
	root.Put("/user/create/hello")
	root.Put("/order/Get/hello")
	root.Put("/order/Get/ha/**")
	root.Put("/hello")
	root.Put("/product/*/id")

	nodes := [...]string{
		"/user/Get/1",
		"/user/Get/aaa",
		"/user/Get/aaa/11",
		"/user/Put/ssss",
		"/user/create/hello",
		"/user/create",
		"/order/Get/hello",
		"/order/Get/no",
		"/order/Get/ha/dhdlhlfw/ddddd",
		"/user/ha/dhdlhlfw/ddddd",
		"/hello",
		"/product/phone/id",
		"/product/phone/xx",
	}

	for _, node := range nodes {
		p := make(Keys)
		result := root.Get(node, p)
		fmt.Printf("%-30s | %v %v\n", node, result, p)
	}

}
