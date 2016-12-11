package bilidownload

import (
	"fmt"
	"testing"
)

func Test_GetBiliJJLink(t *testing.T) {
	aidList := []int{7432517, 7358886, 7400169, 7360018, 7344077, 7337530, 7334937, 7369792,
		5832292, 7350781, 7335306, 7326683, 7405749, 7378418, 7426923, 7345930, 7399315, 7392190,
		7342741, 7387725, 7392271, 7351363, 7388751, 7323087, 7405417, 7369524, 4629524, 7365299,
		7359882}
	for _, aid := range aidList {
		url, err := GetBiliLink(aid)
		fmt.Printf("aid:%v url:%v err:%v\n", aid, url, err)
	}
}

func Test_getBiliJJLink1(t *testing.T) {
	aidList := []int{7432517, 7358886, 7400169, 7360018, 7344077, 7337530, 7334937, 7369792,
		5832292, 7350781, 7335306, 7326683, 7405749, 7378418, 7426923, 7345930, 7399315, 7392190,
		7342741, 7387725, 7392271, 7351363, 7388751, 7323087, 7405417, 7369524, 4629524, 7365299,
		7359882}
	// aidList := []int{7432517}
	for _, aid := range aidList {
		url, err := getBiliJJLink1(aid)
		fmt.Printf("aid:%v url:%v err:%v\n", aid, url, err)
	}
}

func Test_getBiliJJLink2(t *testing.T) {
	urlList := []string{"http://www.bilibilijj.com/Files/DownLoad/12077060.mp4"}
	for _, url := range urlList {
		url2, err := getBiliJJLink2(url)
		fmt.Printf("url:%v url2:%v err:%v\n", url, url2, err)
	}
}

func Test_getBiliJJLink3(t *testing.T) {
	urlList := []string{"http://www.bilibilijj.com//FreeDown/DownLoad/1481356299/mp4/12077060.9AE458B1977C1EF7457650E769637D70"}
	for _, url := range urlList {
		fmt.Println(getBiliJJLink3(url))
	}
}
