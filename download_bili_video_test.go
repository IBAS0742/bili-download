package bilidownload

import (
	"fmt"
	"testing"
)

func Test_getVideoLength(t *testing.T) {
	url := "http://203.116.188.207/ws.acgvideo.com/d/fb/12003545-1hd.mp4?wsTime=1481427406&wsSecret2=0032d05ecf5a252e78705f1d9633592f&oi=2919206336&rate=2400&wshc_tag=0&wsts_tag=584c918f&wsid_tag=74567b4a&wsiphost=ipdbm"
	fmt.Println(getVideoLength(url))
}

func Test_downloadVideoPartial(t *testing.T) {
	url := "http://203.116.188.207/ws.acgvideo.com/d/fb/12003545-1hd.mp4?wsTime=1481427406&wsSecret2=0032d05ecf5a252e78705f1d9633592f&oi=2919206336&rate=2400&wshc_tag=0&wsts_tag=584c918f&wsid_tag=74567b4a&wsiphost=ipdbm"
	buffer, err := downloadVideoPartial(url, 0, 4095)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("bytes length:%v\n", len(buffer))
	}
}

func Test_DownloadVideo(t *testing.T) {
	url := "http://203.116.188.207/ws.acgvideo.com/d/fb/12003545-1hd.mp4?wsTime=1481427406&wsSecret2=0032d05ecf5a252e78705f1d9633592f&oi=2919206336&rate=2400&wshc_tag=0&wsts_tag=584c918f&wsid_tag=74567b4a&wsiphost=ipdbm"
	err := DownloadVideo(url, "test.mp4")
	fmt.Println(err)
}

func Test_DownloadVideoFast(t *testing.T) {
	url := "http://203.116.188.207/ws.acgvideo.com/d/fb/12003545-1hd.mp4?wsTime=1481427406&wsSecret2=0032d05ecf5a252e78705f1d9633592f&oi=2919206336&rate=2400&wshc_tag=0&wsts_tag=584c918f&wsid_tag=74567b4a&wsiphost=ipdbm"
	err := DownloadVideoFast(url, "test_fast.mp4")
	fmt.Println(err)
}
