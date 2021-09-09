package file

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// CheckIsExist 检查文件是否存在
func CheckIsExist(filename string) bool {
	exist := false
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		exist = true
	}
	return exist
}

// Write 写入文件
func Write(filename string, content string) (err error) {
	var f *os.File
	// 检查文件
	if CheckIsExist(filename) {
		// 如果文件存在, 打开文件
		f, err = os.OpenFile(filename, os.O_WRONLY, 0666)
		if nil != err {
			err = fmt.Errorf("os OpenFile err, err : %w", err)
			return
		}
	} else {
		// 创建文件
		f, err = os.Create(filename)
		if nil != err {
			err = fmt.Errorf("os Create err, err : %w", err)
			return
		}
	}

	// 权限
	if err = f.Chmod(0666); nil != err {
		err = fmt.Errorf("f Chmod err, err : %w", err)
		return
	}
	defer f.Close()

	// 写入文件
	w := bufio.NewWriter(f)
	if _, err = w.WriteString(content); nil != err {
		err = fmt.Errorf("bufio Write err, err: %w", err)
		return
	}

	// 刷新缓存
	if err = w.Flush(); nil != err {
		err = fmt.Errorf("bufio Flush err, err: %w", err)
		return
	}

	return
}

// Read 读取文件
func Read(filename string) (content string, err error) {
	var f *os.File

	// 检查文件
	if !CheckIsExist(filename) {
		err = fmt.Errorf("file Not Exist err, err: %w", err)
		return
	}

	// 如果文件存在, 打开文件
	f, err = os.OpenFile(filename, os.O_RDONLY, 0666)
	if nil != err {
		err = fmt.Errorf("os OpenFile err, err : %w", err)
		return
	}

	// 权限
	if err = f.Chmod(0666); nil != err {
		err = fmt.Errorf("f Chmod err, err : %w", err)
		return
	}
	defer f.Close()

	// // 读取数据
	// r := bufio.NewReader(f)
	// if _, err = r.Read([]byte(content)); nil != err {
	// 	err = fmt.Errorf("bufio Read err, err: %w", err)
	// 	return
	// }
	bc, err := ioutil.ReadAll(f)
	if nil != err {
		err = fmt.Errorf("ioutil Read ReadAll, err: %w", err)
		return
	}
	content = string(bc)

	return
}
