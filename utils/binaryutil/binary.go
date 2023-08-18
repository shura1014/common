package binaryutil

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/shura1014/common/clog"
	"github.com/shura1014/common/goerr"
	"math"
)

func Encode(values ...any) []byte {
	buf := new(bytes.Buffer)
	for i := 0; i < len(values); i++ {
		if values[i] == nil {
			return buf.Bytes()
		}
		switch value := values[i].(type) {
		case int:
			buf.Write(EncodeInt(value))
		case int8:
			buf.Write(EncodeInt8(value))
		case int16:
			buf.Write(EncodeInt16(value))
		case int32:
			buf.Write(EncodeInt32(value))
		case int64:
			buf.Write(EncodeInt64(value))
		case uint:
			buf.Write(EncodeUint(value))
		case uint8:
			buf.Write(EncodeUint8(value))
		case uint16:
			buf.Write(EncodeUint16(value))
		case uint32:
			buf.Write(EncodeUint32(value))
		case uint64:
			buf.Write(EncodeUint64(value))
		case bool:
			buf.Write(EncodeBool(value))
		case string:
			buf.Write(EncodeString(value))
		case []byte:
			buf.Write(value)
		case float32:
			buf.Write(EncodeFloat32(value))
		case float64:
			buf.Write(EncodeFloat64(value))

		default:
			if err := binary.Write(buf, binary.LittleEndian, value); err != nil {
				clog.Error(goerr.Wrap(err))
				buf.Write(EncodeString(fmt.Sprintf("%v", value)))
			}
		}
	}
	return buf.Bytes()
}

func EncodeByLength(length int, values ...any) []byte {
	b := Encode(values...)
	if len(b) < length {
		b = append(b, make([]byte, length-len(b))...)
	} else if len(b) > length {
		b = b[0:length]
	}
	return b
}

func Decode(b []byte, values ...any) error {
	var (
		err error
		buf = bytes.NewBuffer(b)
	)
	for i := 0; i < len(values); i++ {
		if err = binary.Read(buf, binary.LittleEndian, values[i]); err != nil {
			err = goerr.Wrap(err)
			return err
		}
	}
	return nil
}

func EncodeString(s string) []byte {
	return []byte(s)
}

func DecodeToString(b []byte) string {
	return string(b)
}

func EncodeBool(b bool) []byte {
	if b {
		return []byte{1}
	} else {
		return []byte{0}
	}
}

func EncodeInt(i int) []byte {
	if i <= math.MaxInt8 {
		return EncodeInt8(int8(i))
	} else if i <= math.MaxInt16 {
		return EncodeInt16(int16(i))
	} else if i <= math.MaxInt32 {
		return EncodeInt32(int32(i))
	} else {
		return EncodeInt64(int64(i))
	}
}

func EncodeUint(i uint) []byte {
	if i <= math.MaxUint8 {
		return EncodeUint8(uint8(i))
	} else if i <= math.MaxUint16 {
		return EncodeUint16(uint16(i))
	} else if i <= math.MaxUint32 {
		return EncodeUint32(uint32(i))
	} else {
		return EncodeUint64(uint64(i))
	}
}

func EncodeInt8(i int8) []byte {
	return []byte{byte(i)}
}

func EncodeUint8(i uint8) []byte {
	return []byte{i}
}

func EncodeInt16(i int16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(i))
	return b
}

func EncodeUint16(i uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, i)
	return b
}

func EncodeInt32(i int32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(i))
	return b
}

func EncodeUint32(i uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, i)
	return b
}

func EncodeInt64(i int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))
	return b
}

func EncodeUint64(i uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)
	return b
}

func EncodeFloat32(f float32) []byte {
	bits := math.Float32bits(f)
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, bits)
	return b
}

func EncodeFloat64(f float64) []byte {
	bits := math.Float64bits(f)
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, bits)
	return b
}

func DecodeToInt(b []byte) int {
	if len(b) < 2 {
		return int(DecodeToUint8(b))
	} else if len(b) < 3 {
		return int(DecodeToUint16(b))
	} else if len(b) < 5 {
		return int(DecodeToUint32(b))
	} else {
		return int(DecodeToUint64(b))
	}
}

func DecodeToUint(b []byte) uint {
	if len(b) < 2 {
		return uint(DecodeToUint8(b))
	} else if len(b) < 3 {
		return uint(DecodeToUint16(b))
	} else if len(b) < 5 {
		return uint(DecodeToUint32(b))
	} else {
		return uint(DecodeToUint64(b))
	}
}

func DecodeToBool(b []byte) bool {
	if len(b) == 0 {
		return false
	}
	if bytes.Equal(b, make([]byte, len(b))) {
		return false
	}
	return true
}

func DecodeToInt8(b []byte) int8 {
	if len(b) == 0 {
		panic(`empty slice given`)
	}
	return int8(b[0])
}

func DecodeToUint8(b []byte) uint8 {
	if len(b) == 0 {
		panic(`empty slice given`)
	}
	return b[0]
}

func DecodeToInt16(b []byte) int16 {
	return int16(binary.LittleEndian.Uint16(FillUpSize(b, 2)))
}

func DecodeToUint16(b []byte) uint16 {
	return binary.LittleEndian.Uint16(FillUpSize(b, 2))
}

func DecodeToInt32(b []byte) int32 {
	return int32(binary.LittleEndian.Uint32(FillUpSize(b, 4)))
}

func DecodeToUint32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(FillUpSize(b, 4))
}

func DecodeToInt64(b []byte) int64 {
	return int64(binary.LittleEndian.Uint64(FillUpSize(b, 8)))
}

func DecodeToUint64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(FillUpSize(b, 8))
}

func DecodeToFloat32(b []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(FillUpSize(b, 4)))
}

func DecodeToFloat64(b []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(FillUpSize(b, 8)))
}

func FillUpSize(b []byte, l int) []byte {
	if len(b) >= l {
		return b[:l]
	}
	c := make([]byte, l)
	copy(c, b)
	return c
}
