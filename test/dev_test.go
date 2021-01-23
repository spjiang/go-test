package test

import (
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
	"log"
	"reflect"
	"testing"
)

type A struct {
	Age string `json:"age"`
}
type B struct {
	A
}

func TestStruct(t *testing.T) {
	a := &A{
		Age: "age",
	}
	b := &B{}
	b.A = *a
	by, _ := json.Marshal(b)
	fmt.Println(string(by))
}

type Test struct {
	Name string `json:"name,omitempty"`
	Age  int64  `json:"age,omitempty"`
}

type DataInLine struct {
	Id   string
	Name string
	Age  string
}

func GetFieldName(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		result = append(result, t.Field(i).Name)
	}
	return result
}

func TestStructToSlice(t *testing.T) {
	ff := DataInLine{}
	qie := GetFieldName(ff)
	fmt.Println("打印出来为切片：", qie)
}



func TestSlice(t *testing.T) {
	test := &Test{}
	p := []string{"AAA", "BBB"}
	pbyte, _ := json.Marshal(p)
	test.Name = string(pbyte)
	data, _ := json.Marshal(test)
	fmt.Println(string(data))

	//s:=[]string{
	//	"AAA",
	//	"CCC",
	//}
	//b,_:=json.Marshal(s)
	//fmt.Println(string(b))
}


func TestStructInterface(t *testing.T) {
	type BigDataAPIResponse struct {
		ErrorCode   int64
		Message     string
		Data        interface{}
		RefreshTime int64
	}

	type item struct {
		Name string `json:"name"`
		Age  int64  `json:"age"`
	}
	type DataStruct struct {
		Total int64  `json:"total"`
		List  []item `json:"list"`
	}
	d := &DataStruct{}
	r := &BigDataAPIResponse{}
	//j:=`{"ErrorCode":0,"Message":"test","Data":{"total":10,"list":[{"Name":"spjiang","Age":10},]}}`
	j := `{"ErrorCode":0,"Message":"test","Data":{"total":10,"list":[{"name":"spjiang","age":10}]}}`
	_ = json.Unmarshal([]byte(j), r)
	fmt.Println(*r)
	fmt.Println(r.Data)

	_ = mapstructure.Decode(r.Data, d)
	fmt.Println(d.List[0].Name)
}

