package main

import (
	"fmt"
	"log"

	"github.com/pdu/bili-download"
)

func main() {
	aid := 7392271
	url, err := bilidownload.GetBiliLink(aid)
	if err != nil {
		log.Printf("GetBiliLink failed:%v\n", err)
	}
	err = bilidownload.DownloadVideoFast(url, fmt.Sprintf("%d.mp4", aid))
	if err != nil {
		log.Printf("DownloadVideoFast failed:%v\n", err)
	}
}
