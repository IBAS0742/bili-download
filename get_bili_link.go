package bilidownload

import (
	"compress/gzip"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// GetBiliLink return the bilibili video link for specific aid
func GetBiliLink(aid int) (string, error) {
	link1, err := getBiliJJLink1(aid)
	if err != nil {
		return "", err
	}
	link2, err := getBiliJJLink2(link1)
	if err != nil {
		return "", err
	}
	link3, err := getBiliJJLink3(link2)
	if err != nil {
		return "", err
	}
	link4, err := getBiliJJLink3(link3)
	if err != nil {
		return "", err
	}
	return link4, nil
}

func getBiliJJLink3(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("http.NewRequest failed,url:%v err:%v\n", url, err)
		return "", err
	}
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cache-Control", "max-age=0")
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
		return "", err
	}
	defer resp.Body.Close()

	return resp.Header.Get("Location"), nil
}

func getBiliJJLink2(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("http.NewRequest failed,url:%v err:%v\n", url, err)
		return "", err
	}
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.98 Safari/537.36")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Add("Accept-Encoding", "gzip, deflate, sdch")
	req.Header.Add("Accept-Language", "en-US,en;q=0.8,zh-CN;q=0.6,zh;q=0.4")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("http.client.Do failed,url:%v err:%v\n", url, err)
		return "", err
	}
	defer resp.Body.Close()

	z := html.NewTokenizer(resp.Body)
	lastURL, err := "", errors.New("JJLink2 resource not found in html")
out:
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			break out
		case html.StartTagToken:
			t := z.Token()
			found, url := false, ""
			for _, attr := range t.Attr {
				if attr.Key == "download" && strings.HasSuffix(attr.Val, "mp4") {
					found = true
				} else if attr.Key == "href" {
					url = fmt.Sprintf("http://www.bilibilijj.com/%s", attr.Val)
				}
			}
			if found && len(url) > 0 {
				lastURL, err = url, nil
				break out
			}
		}
	}

	return lastURL, err
}

func getBiliJJLink1(aid int) (string, error) {
	url := fmt.Sprintf("http://www.bilibilijj.com/video/av%d", aid)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("http.NewRequest failed,url:%v err:%v\n", url, err)
		return "", err
	}
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.98 Safari/537.36")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Add("Accept-Encoding", "gzip, deflate, sdch")
	req.Header.Add("Accept-Language", "en-US,en;q=0.8,zh-CN;q=0.6,zh;q=0.4")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("http.client.Do failed,url:%v err:%v\n", url, err)
		return "", err
	}
	defer resp.Body.Close()

	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		log.Printf("gzip.NewReader failed,url:%v err:%v\n", url, err)
		return "", err
	}

	z := html.NewTokenizer(reader)
	lastURL, err := "", errors.New("JJLink1 resource not found in html")
out:
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			break out
		case html.StartTagToken:
			t := z.Token()
			found, url := false, ""
			for _, attr := range t.Attr {
				if attr.Key == "title" && strings.HasPrefix(attr.Val, "MP4下载") {
					found = true
				} else if attr.Key == "href" {
					url = fmt.Sprintf("http://www.bilibilijj.com/%s", attr.Val)
				}
			}
			if found && len(url) > 0 {
				lastURL, err = url, nil
				break out
			}
		}
	}

	return lastURL, err
}
