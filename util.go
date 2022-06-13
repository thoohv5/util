package util

import (
	"runtime"
	"strconv"
	"strings"
	"unicode"
)

// CamelCase cameCase
func CamelCase(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}

	return string(data[:])
}

// Snake snake
func Snake(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// WhoCalling 获取调用的方法名
func WhoCalling() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

func Split(A string, N string) []int32 {
	a := strings.Split(A, " ")
	n, _ := strconv.Atoi(N) // int 32bit
	b := make([]int32, n)
	for i, v := range a {
		vi, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			// proper err handling
			// either b[i] = -1 (in case positive integers)
		}
		b[i] = int32(vi)
	}
	return b
}
