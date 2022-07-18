package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	chanImagesUrls chan string    //storage all Url channel
	chanTask       chan string    //Task index
	wg             sync.WaitGroup //go routine task group
)

//get pagContent
func GetPagContent(url string) (str string) {
	resp, err := http.Get(url)
	HandleErr(err, "http.Get")
	defer resp.Body.Close()

	b, err2 := ioutil.ReadAll(resp.Body)
	HandleErr(err2, "ioutil.ReadAll")

	return string(b)
}

//get all page urls
func GetAllPageUrls(url string) {
	urls := GetOnePageUris(url)

	for _, url := range urls {
		chanImagesUrls <- url
	}

	//identification current go routine whether completion
	//when a task completion, wirte a date
	//monitor go routine completion number
	chanTask <- url

	wg.Done()
}

//get one page urls
func GetOnePageUris(url string) (urls []string) {
	str := GetPagContent(url)
	imgRegx := `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`

	regx := regexp.MustCompile(imgRegx)

	results := regx.FindAllStringSubmatch(str, -1)

	fmt.Println("this page have" + strconv.Itoa(len(results)) + "images")

	for _, rest := range results {
		url := rest[0]
		urls = append(urls, url)
	}
	return urls
}

func CheckTask() {
	count := 1

	for {
		url := <-chanTask
		fmt.Println(url + "completion reptile task")
		if url != "" {
			count++
		}
		if count == 46 {
			close(chanImagesUrls)
			close(chanTask)
			break
		}
	}
	wg.Done()
}
func downImages() {
	for url := range chanImagesUrls {
		filename := getFileName(url)
		ok := DownloadImage(url, filename)
		if ok {
			fmt.Println(filename + "download success")
		} else {
			fmt.Println(filename + "download failed")
		}
	}
	wg.Done()
}

func getFileName(url string) (filename string) {
	i := strings.LastIndex(url, "/")

	timePrefix := strconv.Itoa(int(time.Now().Unix()))

	filename = timePrefix + url[i+1:]

	return filename
}

func DownloadImage(url string, fileName string) (ok bool) {

	resp, err := http.Get(url)
	HandleErr(err, "http.Get")
	defer resp.Body.Close()

	bytes, err2 := ioutil.ReadAll(resp.Body)
	HandleErr(err2, "ioutil.ReadAll")

	filepath := "../images/" + fileName
	err3 := ioutil.WriteFile(filepath, bytes, 0666)
	HandleErr(err3, "ioutil.WriteFile")

	if err3 != nil {
		return false
	} else {
		return true
	}
}
func HandleErr(err error, why string) {
	if err != nil {
		fmt.Println(err, why)
	}
}
func main() {
	// 1.init channel
	chanImagesUrls = make(chan string, 1000000)
	chanTask = make(chan string, 46)

	// 2.reptile routine
	for i := 1; i <= 46; i++ {
		wg.Add(1)
		go GetAllPageUrls("http://so.5tu.cn/pic/gaoqingmeinvtuku-p" + strconv.Itoa(i) + ".html")
	}
	wg.Wait()

	// 3.statistics go routine task number,if go routine task number equle 46,then colse  chanImagesUrls channel
	wg.Add(1)
	go CheckTask()
	wg.Wait()

	// 4.download go routine
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go downImages()
	}
	wg.Wait()
}
