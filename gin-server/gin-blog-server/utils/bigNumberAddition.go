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
