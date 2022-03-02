package tric

import (
	"fmt"
	"regexp"
)

func Testregexp() {
	lang := "ar-EG"
	re := regexp.MustCompile("[a-z]{2}-[A-Za-z0-9]{2,3}")
	fmt.Println(re.MatchString(lang))

}
