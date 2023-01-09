package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	txt := "0,1,2,3,4,5,6,7,8,9,a,b,c,d,e,f,g"

	regex, _ := regexp.Compile("9")
	fmt.Println(regex.MatchString(txt))     // true
	fmt.Println(regex.FindString(txt))      //9
	fmt.Println(regex.FindStringIndex(txt)) //[18 19]

	lestas, _ := regexp.Compile("[a-g]")
	fmt.Println(lestas.FindAllString(txt, 20)) //[a b c d e f g]

	nums, _ := regexp.Compile("[0-9]")
	fmt.Println(nums.FindAllString(txt, 20)) //[0 1 2 3 4 5 6 7 8 9]

	result := lestas.ReplaceAllFunc([]byte(txt), bytes.ToUpper)
	fmt.Println(string(result)) //0,1,2,3,4,5,6,7,8,9,A,B,C,D,E,F,G

	fmt.Println(lestas.ReplaceAllString(txt, "opa")) //0,1,2,3,4,5,6,7,8,9,opa,opa,opa,opa,opa,opa,opa

}
