package tric

import (
	"fmt"
	"regexp"
)

func Testregexp() {
	//lang := "ar-EG"
	//re := regexp.MustCompile("[a-z]{2}-[A-Za-z0-9]{2,3}")
	//fmt.Println(re.MatchString(lang))
	var myRegexp *regexp.Regexp = regexp.MustCompile("[a-z]{2}[.](driver)[\\w.]*")
	ns := "mx.driver.finalcial"
	fmt.Println(myRegexp.MatchString(ns))

}
