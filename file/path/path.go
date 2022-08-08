package file

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
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
		return "", err
	}

	return res, nil
}

// CurrentAbPathByExecutable 当前执行文件绝对路径
func CurrentAbPathByExecutable() (string, error) {

	// 执行文件
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	// 绝对路径
	res, err := filepath.EvalSymlinks(filepath.Dir(exePath))
	if err != nil {
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
		abPath = path.Dir(filename)
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
