package main

import (
	"fmt"
	"reflect"
)

func stringInSlice(value string, slice []string) bool {
	for _, elem := range slice {
		if elem == value {
			return true
		}
	}
	return false
}

func mergeStructs(structs ...interface{}) reflect.Type {
	var structFields []reflect.StructField
	var structFieldNames []string

	for _, item := range structs {
		rt := reflect.TypeOf(item)
		for i := 0; i < rt.NumField(); i++ {
			field := rt.Field(i)
			if !stringInSlice(field.Name, structFieldNames) {
				structFields = append(structFields, field)
				structFieldNames = append(structFieldNames, field.Name)
			}
		}
	}

	return reflect.StructOf(structFields)
}

type User struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
}

type Extras struct {
	Amount int `json:"amount"`
}

func main() {

	var user User = User{
		Name: "abc",
	}
	var extras Extras
	fmt.Print(mergeStructs(user, extras))
}
