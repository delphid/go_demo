package main

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type S struct {
	K string `validate:"just-print"`
	L string `validate:"required"`
}

func (s S) GetL() string {
	return s.L
}

var validate *validator.Validate

func ValidateSField(fl validator.FieldLevel) bool {
	fmt.Printf("fl: %+v\n", fl)
	fmt.Println(fl.Field())
	return true
}

func ValidateSStruct(sl validator.StructLevel) {
	fmt.Printf("sl: %+v\n", sl.Current())
	s := sl.Current()
	s2 := reflect.ValueOf(S{"ccc", "ddd"})
	fmt.Println(reflect.TypeOf(s), reflect.TypeOf(s2))
	fmt.Println(reflect.ValueOf(s), reflect.ValueOf(s2))
	fmt.Println(s.CanAddr(), s2.CanAddr())
	s.Set(s2)
	v := s.Interface()
	fmt.Println("v: ", v)
	// fmt.Println(v.GetL())
	//s2.Set(s)

}

func main() {
	validate = validator.New()
	validate.RegisterValidation("just-print", ValidateSField)
	// validate.RegisterStructValidation(ValidateSStruct, S{})
	s := &S{"aaa", "bbb"}
	fmt.Printf("s: %+v\n", s)
	errs := validate.Struct(s)
	fmt.Println(errs)
	fmt.Println("s in the end: ", s)
	err := validate.Struct(S{K: "aaa"})
	fmt.Println(err)
	/*
	   a := 1
	   var b int64
	   b = 2
	   reflect.ValueOf(a).SetInt(b)
	   fmt.Println(a)
	*/
}
