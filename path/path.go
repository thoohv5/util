package path

import (
	"fmt"
	"io/ioutil"
	"os"
	spath "path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrPathRead       = errors.New("path read err")
	ErrFileDelete     = errors.New("file delete err")
	ErrFileExecutable = errors.New("file executable err")
	ErrFileSymlinks   = errors.New("file symlinks err")
)

// CurrentAbPath 当前绝对路径
func CurrentAbPath(opts ...Option) (string, error) {

	// 参数拆解
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}

	// 当前执行文件绝对路径
	dir, err := CurrentAbPathByExecutable()
	if nil != err {
		return "", err
	}

	// 临时目录
	tDir, err := tmpDir()
	if nil != err {
		return "", nil
	}

	// 是否为go run场景
	if strings.Contains(dir, tDir) {
		// 当前执行文件绝对路径
		return CurrentAbPathByCaller(opts...), nil
	}

	return dir, nil
}

// 获取系统临时目录，兼容go run
func tmpDir() (string, error) {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}

	// 绝对路径
	res, err := filepath.EvalSymlinks(dir)
	if nil != err {
		err = errors.WithMessagef(ErrFileSymlinks, "err:%v", err)
		return "", err
	}

	return res, nil
}

// CurrentAbPathByExecutable 当前执行文件绝对路径
func CurrentAbPathByExecutable() (string, error) {

	// 执行文件
	exePath, err := os.Executable()
	if err != nil {
		err = errors.WithMessagef(ErrFileExecutable, "err:%v", err)
		return "", err
	}

	// 绝对路径
	res, err := filepath.EvalSymlinks(filepath.Dir(exePath))
	if err != nil {
		err = errors.WithMessagef(ErrFileSymlinks, "err:%v", err)
		return "", err
	}

	return res, nil
}

// CurrentAbPathByCaller 当前执行文件绝对路径
func CurrentAbPathByCaller(opts ...Option) string {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	var abPath string
	_, filename, _, ok := runtime.Caller(o.skip)
	if ok {
		abPath = spath.Dir(filename)
	}
	return abPath
}

// AbPath 获取绝对路径
func AbPath(file string, opts ...Option) string {

	// 参数解析
	o := &options{
		file: file,
		skip: 2,
	}
	for _, opt := range opts {
		opt(o)
	}

	// 绝对路径
	return fmt.Sprintf("%s/%s", CurrentAbPathByCaller(opts...), o.file)
}

// Rm 删除文件夹，不包括当前文件夹
func Rm(path string) (err error) {
	dir, err := ioutil.ReadDir("/tmp")
	if nil != err {
		err = errors.WithMessagef(ErrPathRead, "err:%v", err)
		return
	}
	for _, d := range dir {
		if err = os.RemoveAll(spath.Join([]string{"tmp", d.Name()}...)); nil != err {
			err = errors.WithMessagef(ErrFileDelete, "err:%v", err)
			return
		}
	}
	return
}
