package goToolCommon

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//获取当前路径,不含文件名
func GetCurrPath() (path string, err error) {
	path, err = filepath.Abs(filepath.Dir(os.Args[0]))
	return
}

//获取指定路径下的文件和文件夹列表
func GetFolderAndFileList(path string) (folderList []string, fileList []string, err error) {
	b, err := PathExists(path)
	if err != nil {
		return nil, nil, err
	}
	if !b {
		return nil, nil, errors.New("指定的路径[" + path + "]不存在")
	}
	fInfoList, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, nil, err
	}
	for _, fInfo := range fInfoList {
		if fInfo.IsDir() {
			folderList = append(folderList, fInfo.Name())
		} else {
			fileList = append(fileList, fInfo.Name())
		}
	}
	return folderList, fileList, nil
}

//检查路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}

//检查路径是否存在,不存在则创建
func CheckAndCreateFolder(path string) error {
	count := strings.Count(path, GetFolderSplitStr())
	if count > 1 {
		lastChar := strings.LastIndex(path, GetFolderSplitStr())
		subPath := path[0:lastChar]
		err := CheckAndCreateFolder(subPath)
		if err != nil {
			return err
		}
	}
	b, err := PathExists(path)
	if err != nil {
		return err
	}
	if !b {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

//获取文件MD5
func FileMD5(path string) (string, error) {
	b, err := PathExists(path)
	if err != nil {
		return "", err
	}
	if !b {
		return "", errors.New("file is not exists")
	}
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return Md5(fileData), nil
}

//获取Json字符串
func GetJsonStr(v interface{}) (string, error) {
	str, err := json.Marshal(v)
	if err != nil {
		return "", err
	} else {
		return string(str), nil
	}
}

//记录日志,文件名固定为日期
func Log(s string) error {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		dir = ""
	} else {
		dir = dir + GetFolderSplitStr()
	}
	fileName := dir + "" + GetDateStr(time.Now()) + ".log"
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	log.SetOutput(f)
	log.Println(s)
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}

func LogFile(s string, fileName string) error {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		dir = ""
	} else {
		dir = dir + GetFolderSplitStr()
	}
	fileName = dir + fileName
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	_, err = f.WriteString(s + GetWrapStr())
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}
