package byteutil

import (
	"encoding/json"
	"github.com/shura1014/common/utils/binaryutil"
	"github.com/shura1014/common/utils/reflectutil"
	"reflect"
	"unsafe"
)

func ToByte(data any) []byte {
	if data == nil {
		return nil
	}

	switch value := data.(type) {
	case string:
		return []byte(value)
	case []byte:
		return value
	default:
		v := reflectutil.Indirect(data)
		switch v.Kind() {
		case reflect.Map:
			bytes, _ := json.Marshal(data)
			return bytes
		default:
			return binaryutil.Encode(data)
		}
	}
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
