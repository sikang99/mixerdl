package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson" // get json element quickly
)

func main() {

	url := flag.String("url", "", "url of reddit post you want to download")
	flag.Parse()

	if *url == "" {
		fmt.Println("Error: -url flag is missing. Read the README.md for more details.")
		os.Exit(1)
	}

	vod := strings.Split(*url, "vod=")[1]
	api := "https://mixer.com/api/v2/vods/" + vod

	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error: API request failed.")
		log.Fatal("Do: ", err)
		return
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)                                                        // data in json
	mp4_url := strings.Split(gjson.Get(string(data), "contentLocators.1.uri").String(), "?")[0] // drop the string after ? in url

	DownloadFile(mp4_url, "./")
	os.Rename("source.mp4", vod+".mp4")
}

func PrintDownloadPercent(done chan int64, path string, total int64) {

	for {
		select {
		case <-done:
			return
		default:
			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}

			fi, err := file.Stat()
			if err != nil {
				log.Fatal(err)
			}

			size := fi.Size()
			if size == 0 {
				size = 1
			}

			var percent float64 = float64(size) / float64(total) * 100
			fmt.Printf("%.0f%%\n", percent)
		}

		time.Sleep(time.Second)
	}
}

func DownloadFile(url string, dest string) {

	file := path.Base(url)
	log.Printf("Downloading file %s from %s\n", file, url)

	var path bytes.Buffer
	path.WriteString(dest)
	path.WriteString("/")
	path.WriteString(file)

	start := time.Now()

	out, err := os.Create(path.String())
	if err != nil {
		fmt.Println(path.String())
		panic(err)
	}
	defer out.Close()

	headResp, err := http.Head(url)
	if err != nil {
		panic(err)
	}
	defer headResp.Body.Close()

	size, err := strconv.Atoi(headResp.Header.Get("Content-Length"))
	if err != nil {
		panic(err)
	}

	done := make(chan int64)

	go PrintDownloadPercent(done, path.String(), int64(size))

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}

	done <- n

	elapsed := time.Since(start)
	log.Printf("Download completed in %s", elapsed)
}
