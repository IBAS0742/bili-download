package bilidownload

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

// DownloadVideo from an URL to a local file
func DownloadVideo(url string, filePath string) error {
	videoLength, err := getVideoLength(url)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	byteSize := 1024 * 128
	for byteStart := 0; byteStart < videoLength; byteStart += byteSize {
		buffer, err := downloadVideoPartial(url, byteStart, byteStart+byteSize-1)
		if err != nil {
			return err
		}
		_, err = file.Write(buffer)
		if err != nil {
			return errors.New("write file failed")
		}
	}

	return nil
}

// DownloadVideoFast download url to filePath
func DownloadVideoFast(url string, filePath string) error {
	defer func() {
		recover()
	}()

	videoLength, err := getVideoLength(url)
	if err != nil {
		return err
	}

	byteSize := 1024 * 128
	inputChan := make(chan int)
	bufferMap := make(map[int][]byte)
	var bufferLock sync.RWMutex

	insertMap := func(byteStart int, buffer []byte) {
		bufferLock.Lock()
		defer bufferLock.Unlock()
		bufferMap[byteStart] = buffer
	}

	var wg sync.WaitGroup

	threads := 10
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				recover()
			}()

		out:
			for {
				select {
				case byteStart, ok := <-inputChan:
					if !ok {
						break out
					}
					buffer, err := downloadVideoPartial(url, byteStart, byteStart+byteSize-1)
					if err != nil {
						close(inputChan)
						break out
					}
					insertMap(byteStart, buffer)
				}
			}
		}()
	}

	for byteStart := 0; byteStart < videoLength; byteStart += byteSize {
		inputChan <- byteStart
	}
	close(inputChan)
	wg.Wait()

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for byteStart := 0; byteStart < videoLength; byteStart += byteSize {
		if buffer, ok := bufferMap[byteStart]; ok {
			file.Write(buffer)
		} else {
			return errors.New("download failed")
		}
	}

	return nil
}

func getVideoLength(url string) (int, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("http.NewRequest failed,url:%v err:%v\n", url, err)
		return 0, err
	}
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.98 Safari/537.36")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Add("Accept-Encoding", "gzip, deflate, sdch")
	req.Header.Add("Accept-Language", "en-US,en;q=0.8,zh-CN;q=0.6,zh;q=0.4")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("http.client.Do failed,url:%v err:%v\n", url, err)
		return 0, err
	}
	defer resp.Body.Close()

	return strconv.Atoi(resp.Header.Get("Content-Length"))
}

func downloadVideoPartial(url string, byteStart int, byteEnd int) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("http.NewRequest failed,url:%v err:%v\n", url, err)
		return nil, err
	}
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", url)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.98 Safari/537.36")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "identity;q=1, *;q=0")
	req.Header.Add("Accept-Language", "en-US,en;q=0.8,zh-CN;q=0.6,zh;q=0.4")
	req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", byteStart, byteEnd))

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("http.client.Do failed,url:%v err:%v\n", url, err)
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
