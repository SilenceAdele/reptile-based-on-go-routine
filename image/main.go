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

//get image
func GetImage(url string) {
	pageStr := GetPageContent(url)
	//define emial regexp
	imgRegx := `"https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))"`
	regx := regexp.MustCompile(imgRegx)

	//get result
	result := regx.FindAllStringSubmatch(pageStr, -1)

	//loop email
	for _, res := range result {
		fmt.Println(res[0])
	}
}
func DownloadImage(url string, fileName string) (ok bool) {

	resp, err := http.Get(url)
	HandleErr(err, "http.Get")
	defer resp.Body.Close()

	bytes, err2 := ioutil.ReadAll(resp.Body)
	HandleErr(err2, "ioutil.ReadAll")

	filepath := "../a/" + fileName
	err3 := ioutil.WriteFile(filepath, bytes, 0666)
	HandleErr(err3, "ioutil.WriteFile")

	if err3 != nil {
		return false
	} else {
		return true
	}
}

//handle exception
func HandleErr(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}
func main() {
	//GetImage("http://www.netbian.com/mei/index_5.htm")
	DownloadImage("http://img.netbian.com/file/newc/fc760712977e9efde640af6cfa661c78.jpg", "1.jpg")
}
