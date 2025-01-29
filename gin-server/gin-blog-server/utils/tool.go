package utils

import "strconv"

// StringToByte 字符串转字符切片
func StringToByte(numStr string) []byte {
	var numByte []byte
	for _, i := range numStr {
		n, _ := strconv.Atoi(string(i))
		numByte = append(numByte, byte(n))
	}
	return numByte
}

// ByteToString 字符切片转字符串
func ByteToString(numByte []byte) string {
	strBigNum := ""
	for _, i := range numByte {
		strBigNum += strconv.Itoa(int(i))
	}
	return strBigNum
}
