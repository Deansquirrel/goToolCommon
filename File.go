package goToolCommon

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

//获取当前路径,不含文件名
func GetCurrPath() (path string, err error) {
	path, err = filepath.Abs(filepath.Dir(os.Args[0]))
	return
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

//记录日志,文件名固定为日期
func Log(s string) error {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		dir = ""
	} else {
		dir = dir + "\\"
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
		dir = dir + "\\"
	}
	fileName = dir + fileName
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	_, err = f.WriteString(s + "\n")
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}
