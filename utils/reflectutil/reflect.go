package reflectutil

import "reflect"

func Indirect(value any) reflect.Value {
	var reflectValue reflect.Value
	if v, ok := value.(reflect.Value); ok {
		reflectValue = v
	} else {
		reflectValue = reflect.ValueOf(value)
	}
	return reflect.Indirect(reflectValue)
}

func IndirectType(data any) reflect.Type {
	var reflectType reflect.Type
	if t, ok := data.(reflect.Type); ok {
		reflectType = t
	} else {
		reflectType = reflect.TypeOf(data)
	}
	if reflectType.Kind() == reflect.Pointer {
		reflectType = reflectType.Elem()
	}
	return reflectType
}

func IndirectKind(data any) reflect.Kind {
	return Indirect(data).Kind()
}

func NewData(t reflect.Type) any {
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return reflect.New(t).Interface()
}

func GetName(data any) string {
	dataType := IndirectType(data)
	return dataType.Name()
}

func IsPointer(data any) bool {
	dataType := reflect.TypeOf(data)
	return dataType.Kind() == reflect.Pointer
}

// ConvertType 数组类型转换
func ConvertType(src reflect.Value, srcIndex int, target reflect.Value, targetIndex int) reflect.Value {
	srcVar := src.Index(srcIndex)
	targetType := target.Elem().Field(targetIndex).Type()
	of := reflect.ValueOf(srcVar.Interface())
	covertValue := of.Convert(targetType)
	return covertValue
}

// ValueToInterface 获取实际值
func ValueToInterface(v reflect.Value) (value interface{}, ok bool) {
	if v.IsValid() && v.CanInterface() {
		return v.Interface(), true
	}
	switch v.Kind() {
	case reflect.Bool:
		return v.Bool(), true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint(), true
	case reflect.Float32, reflect.Float64:
		return v.Float(), true
	case reflect.Complex64, reflect.Complex128:
		return v.Complex(), true
	case reflect.String:
		return v.String(), true
	case reflect.Ptr:
		return ValueToInterface(v.Elem())
	case reflect.Interface:
		return ValueToInterface(v.Elem())
	default:
		return nil, false
	}
}
