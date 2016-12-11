# About

This is a Golang tool to download videos from bilibili.com

# How does it works

1. You provide the AV_ID of bilibili video, for example: http://www.bilibili.com/video/av7342741/ -> AV_ID: 7342741
2. Refer to bilibilijj.com to get the video link
3. Download the video in sequence or in parallel

# Examples

```
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
    err = bilidownload.DownloadVideoFast(url, fmt.Sprintf("%d_fast.mp4", aid))
    if err != nil {
        log.Printf("DownloadVideoFast failed:%v\n", err)
    }
    err = bilidownload.DownloadVideo(url, fmt.Sprintf("%d.mp4", aid))
    if err != nil {
        log.Printf("DownloadVideoFast failed:%v\n", err)
    }
}
```

# FAQ

1. Why you use bilibilijj.com to get the video link?
   1. bilibili use some signature algorithm to query the video link, it's hard to analysis, and bilibili also change the algorithm sometimes.
   2. bilibilijj.com has made it easy, it will help you to analysis the video links.
