package file

import (
	"bufio"
	"os"

	"github.com/pkg/errors"
)

var (
	ErrFileNotExist = errors.New("file Not Exist err")
	ErrFileCreate   = errors.New("file Create err")
	ErrFileOpen     = errors.New("file Open err")
	ErrFileChmod    = errors.New("file Chmod err")
	ErrFileWrite    = errors.New("file Write err")
	ErrFileFlush    = errors.New("file Flush err")
	ErrFileClose    = errors.New("file Close err")
)

// CheckIsExist 检查文件/文件夹是否存在
func CheckIsExist(filOrPath string) bool {
	exist := false
	if _, err := os.Stat(filOrPath); !os.IsNotExist(err) {
		exist = true
	}
	return exist
}

// Write 写入文件
func Write(filename string, content string) (err error) {
	var file *os.File
	// 检查文件
	if CheckIsExist(filename) {
		// 如果文件存在, 打开文件
		file, err = os.OpenFile(filename, os.O_WRONLY, 0666)
		if nil != err {
			err = errors.WithMessagef(ErrFileNotExist, "err:%v", err)
			return
		}
	} else {
		// 创建文件
		file, err = os.Create(filename)
		if nil != err {
			err = errors.WithMessagef(ErrFileCreate, "err:%v", err)
			return
		}
	}

	// 权限
	if err = file.Chmod(0666); nil != err {
		err = errors.WithMessagef(ErrFileChmod, "err:%v", err)
		return
	}

	// 写入文件
	w := bufio.NewWriter(file)
	if _, err = w.WriteString(content); nil != err {
		err = errors.WithMessagef(ErrFileWrite, "err:%v", err)
		return
	}

	// 刷新缓存
	if err = w.Flush(); nil != err {
		err = errors.WithMessagef(ErrFileFlush, "err:%v", err)
		return
	}

	// 文件关闭
	if err = file.Close(); nil != err {
		err = errors.WithMessagef(ErrFileClose, "err:%v", err)
	}

	return
}

// Read 读取文件
func Read(filename string, opt func(file *os.File) error) (err error) {
	var file *os.File

	// 检查文件
	if !CheckIsExist(filename) {
		err = errors.WithMessagef(ErrFileNotExist, "err:%v", err)
		return
	}

	// 如果文件存在, 打开文件
	file, err = os.OpenFile(filename, os.O_RDONLY, 0666)
	if nil != err {
		err = errors.WithMessagef(ErrFileOpen, "err:%v", err)
		return
	}

	// 权限
	if err = file.Chmod(0666); nil != err {
		err = errors.WithMessagef(ErrFileChmod, "err:%v", err)
		return
	}

	// 文件操作
	if err = opt(file); nil != err {
		if fErr := file.Close(); nil != err {
			err = errors.WithMessagef(ErrFileClose, "err:%v", fErr)
		}
		return
	}

	// 文件关闭
	if err = file.Close(); nil != err {
		err = errors.WithMessagef(ErrFileClose, "err:%v", err)
	}

	return
}
