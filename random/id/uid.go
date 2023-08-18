package id

import (
	"github.com/shura1014/common/hash"
	"github.com/shura1014/common/random"
	"github.com/shura1014/common/type/atom"
	"github.com/shura1014/common/utils/iputil"
	"os"
	"strconv"
	"time"
)

// Mac地址(7) Pid(3) 当前时间(12) 计数器(3) 随机数(7) 7+4+12+3+7=32
// Why
// Mac 因为几十个mac hash后就已经占用7个字节了 但是即便是上千个地址hash也不会超过7字节
// Pid 64位操作系统默认 32768 这个值不建议修改

var (
	// 计数器
	counter atom.Uint32
	// 7个字节
	macStr = []byte{'0', '0', '0', '0', '0', '0', '0'}
	// 3个字节
	pidStr = []byte{'0', '0', '0'}
)

const (
	// zzz 36进制 0-36*36*36-1 0-z|0-z|0-z
	sequenceMax   = uint32(46655)
	randomStrBase = "abcdefghijklm0123456789nopqrstuvwxyz"
)

func init() {
	copy(macStr, GetMacStr())
	s := strconv.FormatInt(int64(os.Getpid())%46655, 36)
	copy(pidStr, s)

}

// GetCounter 计数器+1 36进制 占用3个字节
func GetCounter() []byte {
	b := []byte{'0', '0', '0'}
	s := strconv.FormatUint(uint64(counter.Add(1)%sequenceMax), 36)
	copy(b, s)
	return b
}

func Str() string {
	b := make([]byte, 32)
	nanoStr := strconv.FormatInt(time.Now().UnixNano(), 36)
	copy(b, macStr)
	copy(b[7:], pidStr)
	copy(b[10:], nanoStr)
	copy(b[22:], GetCounter())
	copy(b[25:], getRandomStr(7))
	return string(b)
}

// 占用7字节随机数
func getRandomStr(n int) []byte {
	if n <= 0 {
		return []byte{}
	}
	var (
		b           = make([]byte, n)
		numberBytes = random.Byte(n)
	)
	for i := range b {
		b[i] = randomStrBase[numberBytes[i]%36]
	}
	return b
}

// GetMacStr mac占用7字节
func GetMacStr() string {
	macs, _ := iputil.GetMacArray()
	if len(macs) > 0 {
		var macAddrBytes []byte
		for _, mac := range macs {
			macAddrBytes = append(macAddrBytes, []byte(mac)...)
		}
		return strconv.FormatUint(uint64(hash.BKDRHash32(macAddrBytes)), 36)
	}
	return "0000000"
}
