package helper

import "strconv"

func StrToInt(src string) (res int, err error) {
	value, err := StrToInt64(src)
	res = int(value)
	return
}

func StrToInt32(src string) (res int32, err error) {
	value, err := StrToInt64(src)
	res = int32(value)
	return
}

func StrToUInt32(src string) (res uint32, err error) {
	value, err := StrToInt64(src)
	res = uint32(value)
	return
}

func StrToInt64(src string) (int64, error) {
	return strconv.ParseInt(src, 10, 32)
}