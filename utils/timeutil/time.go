package timeutil

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	YYYY           = "2006"
	MM             = "01"
	DD             = "02"
	HH             = "15"
	hh             = "03"
	mm             = "04"
	ss             = "05"
	YYYYMMDDHHmmss = fmt.Sprintf("%s-%s-%s %s:%s:%s", YYYY, MM, DD, HH, mm, ss)
	YYYYMMDDhhmmss = fmt.Sprintf("%s-%s-%s %s:%s:%s", YYYY, MM, DD, hh, mm, ss)
	DefaultFormat  = YYYYMMDDHHmmss

	formatMap = map[byte]string{
		'Y': YYYY,
		'M': MM,
		'D': DD,
		'H': HH,
		'h': hh,
		'm': mm,
		's': ss,
	}
)

func Parse(f string) string {
	if f == "" {
		return ""
	}
	builder := strings.Builder{}

	for i := 0; i < len(f); i++ {
		if t, ok := formatMap[f[i]]; ok {
			builder.WriteString(t)
		} else {
			builder.WriteByte(f[i])
		}
	}
	return builder.String()
}

func Now() string {
	return time.Now().Format(DefaultFormat)
}

func Month() int {
	return int(time.Now().Month())
}

func NowFormat(format string) string {
	return time.Now().Format(format)
}

func Seconds() int64 {
	return time.Now().UnixNano() / 1e9
}

// MilliSeconds 毫秒
func MilliSeconds() int64 {
	return time.Now().UnixNano() / 1e6
}

// MicroSeconds 微秒
func MicroSeconds() int64 {
	return time.Now().UnixNano() / 1e3
}

// NanoSeconds 纳秒
func NanoSeconds() int64 {
	return time.Now().UnixNano()
}

func Convert(timeString string) time.Duration {
	timeInt, _ := strconv.Atoi(timeString)
	return time.Duration(timeInt)
}

// TimeStamp 毫秒时间戳
func TimeStamp() int64 {
	return MilliSeconds()
}

// TimeStampToDateString 时间戳转字符串日期
// timestamp 毫秒
func TimeStampToDateString(timestamp int64) string {
	return time.Unix(timestamp, 1_000).Format(DefaultFormat)
}
