package main

import (
	"fmt"
	"reflect"
)

type TestSt struct {
	Name   string `json:"name_of_employee", "custom":"xyz"`
	Age    int    `json:"age_of_employee"`
	Field3 TestIn
}
type TestIn struct {
	Field4 string   `json:"field_4_tag"`
	Field5 []string `json:"field_5_tag"`
	Field6 TestIn2
}

type TestIn2 struct {
	Field7 int            `json:"field_7_tag"`
	Field8 map[string]int `json:"field_8_tag"`
}

func main() {
	obj := TestSt{
		Name: "Abcd",
		Age:  20,
		Field3: TestIn{
			Field4: "qq",
			Field5: []string{"str1", "str2", "str3"},
			Field6: TestIn2{
				Field7: 20,
				Field8: map[string]int{"abc": 60, "xyz": 80},
			},
		},
	}
	fmt.Printf("\n%+v\n", obj)
	fmt.Println("#################################################")
	//printUsingStruct(obj)
	//printUsingStruct2(obj)
	printUsingInterface(obj)
}

func printUsingStruct(t TestSt) {
	fmt.Println("=================================================")
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s (%s) = %v, tag => %s\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface(), typeOfT.Field(i).Tag)
	}
}

func printUsingStruct2(t TestSt) {
	fmt.Println("=================================================")
	e := reflect.ValueOf(&t).Elem()

	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		varTag := e.Type().Field(i).Tag
		varSize := varType.Size()
		fmt.Printf("%v %v %v %v %v \n", varName, varType, varValue, varSize, varTag)
	}
}

func printUsingInterface(v interface{}) {
	fmt.Println("=================================================")
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)
	fmt.Println(rt)
	if rt.Kind() != reflect.Struct {
		fmt.Printf("\nError : Expecting struct, %s passed\n", rt.Kind())
		return
	}
	for i := 0; i < rv.NumField(); i++ {
		varName := rt.Field(i).Name
		varType := rv.Field(i).Type()
		varValue := rv.Field(i).Interface()
		varTag := rt.Field(i).Tag
		varSize := rv.Field(i).Type().Size()

		if rt.Field(i).Type.Kind().String() == "struct" {
			printUsingInterface(varValue)
		} else {
			fmt.Printf("%d: %s (%s) = %+v, size => %d bytes, tag => %s\n", i, varName, varType, varValue, varSize, varTag)
		}
	}
}

