package jsoniter

import (
	"fmt"
	//"github.com/mymmsc/go-ctp/encoding/json/extra"
	"testing"
)

func TestT1(t *testing.T) {
	//extra.SupportPrivateFields()
	type TestObject struct {
		Field1 string `json:"field1"`
		Field2 string `json:"field2"`
		Field3 int64  `json:"field3,string"`
		Field4 int    `json:"field4,string"`
		Field5 int    `json:"field5"`
	}
	var obj []TestObject
	//var json = ConfigCompatibleWithStandardLibrary
	UnmarshalFromString(`[{Field2:"a12",field1:"Hello",field3:"12345678901234567","field4":"12","field5":5}]`, &obj)

	//should.Equal("Hello", obj.field1)
	fmt.Println(obj)
	Unmarshal([]byte(`[{Field2:"a12",field1:"Hello",field3:"12345678901234567","field4":"12","field5":5}]`), &obj)

	//should.Equal("Hello", obj.field1)
	fmt.Println(obj)
}
