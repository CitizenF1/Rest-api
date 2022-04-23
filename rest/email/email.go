package email

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var emailReg = regexp.MustCompile("(?i)([A-Z0-9._%+-]+@[A-Z0-9.-]+\\.[A-Z]{2,24})")

var iin = regexp.MustCompile("(?i)([0-9]{12})$")

func Email(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	bodyString := string(body)
	var str []string

	// Check Email or iin
	if iin.MatchString(bodyString) {
		str = iin.FindAllString(bodyString, 100)
	} else {
		str = emailReg.FindAllString(bodyString, 100)
	}

	fmt.Fprint(w, str)
}
