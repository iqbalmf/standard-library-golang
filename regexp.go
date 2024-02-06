package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text = "banana burger soup"
	var regex, err = regexp.Compile(`[a-z]+`)
	if err != nil {
		fmt.Println(err.Error())
	}
	var res1 = regex.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)
	var regex2 *regexp.Regexp = regexp.MustCompile(`e[(a-z)o]`)
	var regex3 *regexp.Regexp = regexp.MustCompile(`i[(a-z)l]`)
	fmt.Println(regex2.MatchString("eko"))
	fmt.Println(regex2.MatchString("edo"))
	fmt.Println(regex2.MatchString("eKo"))
	fmt.Println(regex2.FindAllString("eko ego egi edo e1o elo eTo", 10))
	fmt.Println(regex3.MatchString("iqbal"))
	fmt.Println(regex3.FindAllString("iqbal iobel", 5))
}
