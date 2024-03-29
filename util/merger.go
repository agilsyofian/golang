package util

import "reflect"

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		if v.IsNil() {
			return true
		}
		return isEmptyValue(v.Elem())
	case reflect.Func:
		return v.IsNil()
	case reflect.Invalid:
		return true
	}
	return false
}

func MergerOverwrite(dst, src interface{}) {
	vDst := reflect.ValueOf(dst)
	vSrc := reflect.ValueOf(src)

	// We check if its a pointer to dereference it.
	if vDst.Kind() == reflect.Ptr {
		vDst = vDst.Elem()
	}
	if vSrc.Kind() == reflect.Ptr {
		vSrc = vSrc.Elem()
	}

	for i, n := 0, vDst.NumField(); i < n; i++ {
		for j, m := 0, vSrc.NumField(); j < m; j++ {
			nameFieldSrc := vSrc.Type().Field(j).Name
			if nameFieldSrc == vDst.Type().Field(i).Name && nameFieldSrc != "Model" && !isEmptyValue(vSrc.Field(j)) {
				vDst.Field(i).Set(vSrc.Field(j))
			}
		}

	}
}
