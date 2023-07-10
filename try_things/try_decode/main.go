package main

import (
    "encoding/json"
    "fmt"
    "github.com/go-playground/validator/v10"
    "strings"
)

type S struct {
    A string
    B string `validate:"required"`
}

type ConditionReq struct {
    ID        uint   `json:"id"`
    Key       string `json:"key" binding:""`
    Value     string `json:"value" binding:""`
    MatchType string `json:"match_type" validate:"oneof='' 'unknown' 'equal' 'not equal' 'contain' 'not contain' 'regex' 'not regex'"`
}

func main() {
    r := strings.NewReader(`{"a": "b"}`)
    d := json.NewDecoder(r)
    var s S
    d.Decode(&s)
    fmt.Println(d)
    fmt.Printf("%+v\n", s)
    validate := validator.New()
    //validate.RegisterValidation("just-print", ValidateSField)
    err := validate.Struct(s)
    fmt.Println(err)
    r2 := strings.NewReader(`{"key":"app", "value":"a.", "match_type":"regexx"}`)
    d2 := json.NewDecoder(r2)
    var c ConditionReq
    d2.Decode(&c)
    fmt.Printf("%+v\n", c)
    err = validate.Struct(c)
    fmt.Println(err)
}
