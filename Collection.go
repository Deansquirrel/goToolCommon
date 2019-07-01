package goToolCommon

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

//demo
//sort.Sort(goToolCommon.SortByPinyin(zzList))

type SortByPinyin []string

func (s SortByPinyin) Len() int      { return len(s) }
func (s SortByPinyin) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s SortByPinyin) Less(i, j int) bool {
	a, _ := UTF82GBK(s[i])
	b, _ := UTF82GBK(s[j])
	bLen := len(b)
	for idx, chr := range a {
		if idx > bLen-1 {
			return false
		}
		if chr != b[idx] {
			return chr < b[idx]
		}
	}
	return true
}

//UTF82GBK : transform UTF8 rune into GBK byte array
func UTF82GBK(src string) ([]byte, error) {
	GB18030 := simplifiedchinese.All[0]
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), GB18030.NewEncoder()))
}

////GBK2UTF8 : transform  GBK byte array into UTF8 string
//func GBK2UTF8(src []byte) (string, error) {
//	GB18030 := simplifiedchinese.All[0]
//	bytes, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(src), GB18030.NewDecoder()))
//	return string(bytes), err
//}
