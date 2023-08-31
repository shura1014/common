package tree

import "testing"

type Address struct {
	Addr string
}

func TestName(t *testing.T) {
	data := NewTreeData[[]Address]("/")
	addresses := []Address{{"127.0.0.1"}}
	data.Put("/upc/x/y", &addresses)
	data.Put("/upc/x/u", &addresses)
	data.Put("/upc/y/x", &addresses)

	addresses = append(addresses, Address{"127.0.0.2"})
	value := data.Get("/upc/x/y")
	t.Logf("%+v", value)
	t.Logf("%+v", value.Data)
	t.Logf("%+v", value.FullName)
	data.Iterator(func(name string, fullName string, data *[]Address) {
		t.Log(name, fullName, data)
	})
}
