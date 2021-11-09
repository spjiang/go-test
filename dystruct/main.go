package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"
)

type Age struct {
	Value string
}

type Title struct {
	Value string
}

type Name struct {
	Value string
}

type Addr struct {
	Value string
}

type user struct {
	Name  Name
	Age   Age
	Title Title
	Addr  Addr
}

//type user struct {
//	Name  Name
//	Age   Age
//	Title Title
//	Addr  Addr
//}

func main() {
	var obj interface{}
	// u = new(user)
	obj, _ = StringToStruct("user")
	fieldList := []string{"Age", "Title"}
	MsgList(fieldList, obj)
	fmt.Println(obj)

	// user2 := obj.(*user)
	//
	resByre, resByteErr := json.Marshal(obj)
	if resByteErr != nil {
		fmt.Println("1", resByteErr)
		return
	}
	fmt.Println(string(resByre))

	var newData user
	err := json.Unmarshal(resByre, &newData)
	if err != nil {
		fmt.Println("2", )
		return
	}
	fmt.Println(newData)

}

func GetAge() *Age {
	var res *Age
	res = new(Age)
	res.Value = "18"
	return res
}
func GetTitle() *Title {
	var res *Title
	res = new(Title)
	res.Value = "jsp"
	return res
}
func MsgList(fieldList []string, obj interface{}) {
	for _, field := range fieldList {
		var groupData interface{}
		switch field {
		case "Age":
			groupData = GetAge()
		case "Title":
			groupData = GetTitle()
		}
		Tool(obj, field, groupData)
	}
}

func StringToStruct(name string) (obj interface{}, err error) {
	switch name {
	case "user":
		obj = new(user)
	default:
		return nil, fmt.Errorf("%s is not a known struct name", name)
	}
	return
}

func Tool(obj interface{}, fieldName string, val interface{}) {
	value := reflect.ValueOf(obj)
	if value.Kind() == reflect.Ptr {
		elem := value.Elem()
		fieldObj := elem.FieldByName(fieldName)
		if fieldObj.Kind() != reflect.Struct {
			return
		}
		switch fieldName {
		case "Age":
			*(*Age)(unsafe.Pointer(fieldObj.Addr().Pointer())) = *val.(*Age)
		case "Title":
			*(*Title)(unsafe.Pointer(fieldObj.Addr().Pointer())) = *val.(*Title)
		}
	}
}
