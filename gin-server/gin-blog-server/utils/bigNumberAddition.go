package utils

import "strconv"

// 该函数用于对文章阅览量、文章喜欢量进行字符串大数字累加操作
func BigNumAdd(bigNum string) string {
	// 将字符串转为切片，便于操作
	var num []byte
	for _, i := range bigNum {
		n, _ := strconv.Atoi(string(i))
		num = append(num, byte(n))
	}
	// 记录大数字长度
	bigNumLen := len(num)
	// 定义一个布尔值，决定是否进行加法操作
	var add = true
	for i := bigNumLen - 1; i >= 0; i-- {
		if add {
			temp := int(num[i]) + 1
			if temp > 9 {
				num[i] = byte(temp % 10)
				if i == 0 {
					var t []byte
					t = append(t, 1)
					resualt := make([]byte, len(t)+len(num))
					copy(resualt, t)
					copy(resualt[len(t):], num)
					num = resualt
				}
			} else {
				num[i] = byte(temp)
				add = false
			}
		}
	}

	// 切片转字符串
	strBigNum := ""
	for _, i := range num {
		strBigNum += strconv.Itoa(int(i))
	}
	return strBigNum
}

// 比较器
func cmp(a string, b string) (string, string) {
	if len(a) >= len(b) {
		return a, b
	} else {
		return b, a
	}
}

// 字符串转字符切片
func stringToByte(numStr string) []byte {
	var numByte []byte
	for _, i := range numStr {
		n, _ := strconv.Atoi(string(i))
		numByte = append(numByte, byte(n))
	}
	return numByte
}

// 字符切片转字符串
func byteToString(numByte []byte) string {
	strBigNum := ""
	for _, i := range numByte {
		strBigNum += strconv.Itoa(int(i))
	}
	return strBigNum
}

func BigDecimal(a string, b string) string {
	maxNumStr, minNumStr := cmp(a, b)
	// 特判被加数
	if len(minNumStr) <= 0 {
		return maxNumStr
	}
	var maxNumByte []byte
	var minNumByte []byte
	// 将maxNumStr和minNumStr转为字符切片
	maxNumByte = stringToByte(maxNumStr)
	minNumByte = stringToByte(minNumStr)

	maxLen := len(maxNumByte) - 1
	minLen := len(minNumByte) - 1

	// 进位器
	var bitAdd int = 0
	for {
		tmpNum := int(maxNumByte[maxLen]) + int(minNumByte[minLen]) + bitAdd
		// 进位后解除进位器
		bitAdd = 0
		// 进行位处理
		if tmpNum/10 != 0 {
			bitAdd = 1
			maxNumByte[maxLen] = byte(tmpNum % 10)
		} else {
			maxNumByte[maxLen] = byte(tmpNum)
		}
		maxLen--
		minLen--
		// 当最小字符串遍历结束后 外层结束 开始进位运算
		if minLen < 0 && bitAdd == 1 {
			for {
				if maxLen >= 0 {
					tmpNum = int(maxNumByte[maxLen]) + bitAdd
					if tmpNum/10 != 0 {
						maxNumByte[maxLen] = byte(tmpNum) % 10
					} else {
						maxNumByte[maxLen] = byte(tmpNum)
						return byteToString(maxNumByte)
					}
				} else {
					var n []byte
					n = append(n, 1)
					n = append(n, maxNumByte...)
					return byteToString(n)
				}
				maxLen--
			}
		} else if minLen < 0 && bitAdd == 0 {
			return byteToString(maxNumByte)
		}
	}
}
