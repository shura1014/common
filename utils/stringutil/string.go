package stringutil

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type IString interface {
	String() string
}

// Underline 驼峰转下划线
func Underline(name string) string {
	var names = name[:]
	lastIndex := 0
	var sb strings.Builder
	for index, value := range names {
		if value >= 65 && value <= 90 {

			// 大写字母
			if index == 0 {
				continue
			}
			sb.WriteString(name[lastIndex:index])
			sb.WriteString("_")
			lastIndex = index
		}
	}
	sb.WriteString(name[lastIndex:])
	return sb.String()
}

func LowUnderline(name string) string {
	return strings.ToLower(Underline(name))
}

func UpperUnderline(name string) string {
	return strings.ToUpper(Underline(name))
}

func Join(one any, two ...any) string {
	s := ToString(one)
	for _, ss := range two {
		s += ToString(ss)
	}
	return s
}

func IsArray(a []string, s string) bool {
	for _, v := range a {
		if s == v {
			return true
		}
	}
	return false
}

func AnyToByte(data any) []byte {
	return []byte(ToString(data))
}

func ToString(data any) string {
	if data == nil {
		return ""
	}

	switch value := data.(type) {
	case int:
		return strconv.Itoa(value)
	case int8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.FormatInt(value, 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	case string:
		return value
	case []byte:
		return string(value)
	case time.Time:
		if value.IsZero() {
			return ""
		}
		return value.String()
	case *time.Time:
		if value == nil {
			return ""
		}
		return value.String()
	default:
		if value == nil {
			return ""
		}
		if f, ok := value.(IString); ok {
			// If the variable implements the String() interface,
			// then use that interface to perform the conversion
			return f.String()
		}
		if f, ok := value.(error); ok {
			// If the variable implements the Error() interface,
			// then use that interface to perform the conversion
			return f.Error()
		}
		rv := reflect.ValueOf(value)
		kind := rv.Kind()

		switch kind {
		case reflect.Chan,
			reflect.Map,
			reflect.Slice,
			reflect.Func,
			reflect.Ptr,
			reflect.Interface,
			reflect.UnsafePointer:
			if rv.IsNil() {
				return ""
			}
		case reflect.String:
			return rv.String()
		}
		if kind == reflect.Ptr {
			return ToString(rv.Elem().Interface())
		}

		if jsonContent, err := json.Marshal(value); err != nil {
			return fmt.Sprint(value)
		} else {
			return string(jsonContent)
		}
	}

}

// IsNumeric 是否是数字类型
func IsNumeric(s string) bool {
	var (
		dotCount = 0
		length   = len(s)
	)
	if length == 0 {
		return false
	}
	for i := 0; i < length; i++ {
		if s[i] == '-' && i == 0 {
			continue
		}
		if s[i] == '.' {
			dotCount++
			if i > 0 && i < length-1 {
				continue
			} else {
				return false
			}
		}
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return dotCount <= 1
}
