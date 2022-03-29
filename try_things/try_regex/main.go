package main

import (
	"fmt"
	"regexp"
	// "strings"
)

var grafanaCurValueRe = regexp.MustCompile(` value=(.*) ]`)

func GetCurValue(s string) string {
	// `[ metric='' labels={} value=0.4 ]` -> `0.4`
	m := grafanaCurValueRe.FindStringSubmatch(s)
    if len(m)==0 {
        return ""
    }
    fmt.Println(m)
    return m[1]
	// return strings.Split(strings.Split(ret, "=")[1], " ")[0]
}

func main() {
	s1 := `[ metric='' labels={} value=0.4 ]`
	s2 := ""
    s3 := "aaa"
    s4 := `[ metric='' labels={} value= ]`
	fmt.Println(GetCurValue(s1))
    fmt.Println(GetCurValue(s2))
    fmt.Println(GetCurValue(s3))
    fmt.Println(GetCurValue(s4))
}
