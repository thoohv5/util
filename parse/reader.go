package parse

import (
	"bytes"
	"io"
	"io/ioutil"
)

func ReaderToBytes(reader io.ReadCloser) ([]byte, error) {
	// io.ReadCloser to []byte
	return ioutil.ReadAll(reader)
}

func BytesToReaderCloser(bs []byte) (io.ReadCloser, error) {
	return ioutil.NopCloser(bytes.NewBuffer(bs)), nil
}

func BytesToReader(bs []byte) (io.Reader, error) {
	return bytes.NewBuffer(bs), nil
}
