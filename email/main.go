package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//get web page content
func GetPageContent(url string) (str string) {
	resp, err := http.Get(url)
	HandleErr(err, "http.Get url")
	defer resp.Body.Close()

	//read page content
	b, err2 := ioutil.ReadAll(resp.Body)
	HandleErr(err2, "ioutil.ReadAll")

	//byte to string
	pageStr := string(b)
	return pageStr
}

//get emial
func GetEmailContent(url string) {
	pageStr := GetPageContent(url)

	//define emial regexp
	emialRegx := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	regx := regexp.MustCompile(emialRegx)

	//get result
	result := regx.FindAllStringSubmatch(pageStr, -1)

	//loop email
	for _, res := range result {
		fmt.Println(res[0])
	}
}

//handle exception
func HandleErr(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}
func main() {
	//GetEmail()
	GetEmailContent("https://xxxxxxx")
}
