package random

import (
	"encoding/binary"
)

var (
	number          = "0123456789"                                                                                                                                   // 数字 10
	letter          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"                                                                                         // 字母 52
	letterAndNumber = "abcdefghijklmnopqrstuvwxyz01234567890123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"                                                                     // 数组与字母 72
	symbols         = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"                                                                                                           // 符号 32
	characters      = "abcdefghijklmnopqrstuvwxyz!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~01234567890123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ" // 字符类型 136
)

func Byte(n int) []byte {
	if n <= 0 {
		return nil
	}
	i := 0
	b := make([]byte, n)
	copy(b[i:], <-bufferChan)
	for {
		// 每从bufferChan取出的是 [4]byte 4个大小
		copy(b[i:], <-bufferChan)
		i += 4
		if i >= n {
			break
		}
	}
	return b
}

// Int 32bit 4 byte == <-bufferChan
// 0~4294967295
// 返回正数 使用无符号int
// [0,max]
func Int(max int) int {
	if max < 0 {
		return max
	}
	n := int(binary.LittleEndian.Uint32(<-bufferChan)) % (max + 1)
	return n
}

// IntRange 返回一定范围的随机数
// 支持负数
// [min,max]
func IntRange(min, max int) int {
	if min >= max {
		return min
	}

	// -1,2 -> [0-3]-1
	if min <= 0 {
		return Int(max+(-min)) - (-min)
	}

	// 1,3 -> [0-2]+1
	return Int(max-min) + min
}

func Symbols(n int) string {
	if n <= 0 {
		return ""
	}

	b := make([]byte, n)
	numberBytes := Byte(n)
	for i := range b {
		b[i] = symbols[numberBytes[i]&31]
	}
	return string(b)
}

// Letter 大小写字母
func Letter(n int) string {
	if n <= 0 {
		return ""
	}

	b := make([]byte, n)
	numberBytes := Byte(n)
	for i := range b {
		b[i] = letter[numberBytes[i]%52]
	}
	return string(b)
}

// W 与正则表达式的W一样，返回数字+字母
// 可以作为随机验证码
func W(n int) string {
	if n <= 0 {
		return ""
	}

	b := make([]byte, n)
	numberBytes := Byte(n)
	for i := range b {
		b[i] = letterAndNumber[numberBytes[i]%72]
	}
	return string(b)
}

// Char 字母+数字+符号
// 可以作为linux主机密码
func Char(n int) string {
	if n <= 0 {
		return ""
	}

	b := make([]byte, n)
	numberBytes := Byte(n)
	for i := range b {
		b[i] = characters[numberBytes[i]%136]
	}
	return string(b)
}

func Number(n int) string {
	if n <= 0 {
		return ""
	}

	b := make([]byte, n)
	numberBytes := Byte(n)
	for i := range b {
		b[i] = number[numberBytes[i]%10]
	}
	return string(b)
}
