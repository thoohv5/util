package parse

import "time"

/*** Parse ***/

// StrToBool str 2 bool
func StrToBool(str string) (b bool, err error) {
	return b, convertAssign(&b, str)
}

// StrToInt str 2 int
func StrToInt(s string) (i int, err error) {
	return i, convertAssign(&i, s)
}

// StrToInt32 str 2 int32
func StrToInt32(s string) (i int32, err error) {
	return i, convertAssign(&i, s)
}

// str 2 uint32
func StrToUin32(s string) (i uint32, err error) {
	return i, convertAssign(&i, s)
}

// StrToInt64 str 2 int64
func StrToInt64(s string) (i int64, err error) {
	return i, convertAssign(&i, s)
}

// StrToUin64 str 2 uint64
func StrToUin64(s string) (i uint64, err error) {
	return i, convertAssign(&i, s)
}

// BoolToUint32 bool 2 uint32
func BoolToUint32(b bool) (i uint32, err error) {
	if b {
		return 1, nil
	}
	return 0, nil
}

// Int32ToByteArr int32 2 []byte]
func Int32ToByteArr(i int32) (bytes []byte, err error) {
	return bytes, convertAssign(&bytes, i)
}

// Uint32ToBool uint32 2 bool
func Uint32ToBool(i uint32) (b bool, err error) {
	return b, convertAssign(&b, i)
}

// Uint32ToStr uint32 2 string
func Uint32ToStr(i uint32) (s string, err error) {
	return s, convertAssign(&s, i)
}

// Uint32ToByteArr uint32 2 []byte]
func Uint32ToByteArr(i uint32) (bytes []byte, err error) {
	return bytes, convertAssign(&bytes, i)
}

// Int64ToStr int64 2 str
func Int64ToStr(i int64) (s string, err error) {
	return s, convertAssign(&s, i)
}

// Int64ToTimestamp int64 2 timestamp
func Int64ToTimestamp(i int64) (string, error) {
	return time.Unix(i, 0).Format("2006-01-02 15:04:05"), nil
}

// Int64ToByteArr int64 to []byte]
func Int64ToByteArr(i int64) (bytes []byte, err error) {
	return bytes, convertAssign(&bytes, i)
}

// Uint64ToStr uint64 to str
func Uint64ToStr(i uint64) (s string, err error) {
	return s, convertAssign(&s, i)
}

// Uint64ToByteArr uint64 to []byte
func Uint64ToByteArr(i uint64) (bytes []byte, err error) {
	return bytes, convertAssign(&bytes, i)
}
