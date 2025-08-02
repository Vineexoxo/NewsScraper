package reflectionhelper

import (
	"reflect"
	"unsafe"
)

func GetFieldValueByIndex[T any](object T, idx int) interface{} {
	v:=reflect.ValueOf(&object).Elem()

	if(v.Kind()==reflect.Ptr){
		val:=v.Elem()
		field:=val.Field(idx)
		if field.CanInterface(){
			return field.Interface()
		}else{
			return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()

		}
	}else{
		val := v
		field := val.Field(idx)
		if field.CanInterface() {
			return field.Interface()
		} else {
			// for all unexported fields (private)
			rs2 := reflect.New(val.Type()).Elem()
			rs2.Set(val)
			val = rs2.Field(idx)
			val = reflect.NewAt(val.Type(), unsafe.Pointer(val.UnsafeAddr())).Elem()

			return val.Interface()
		}
	}
	return nil

}

func GetFieldValue(field reflect.Value) reflect.Value {
	if field.CanInterface(){
		return field
	}else{
		res := reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
		return res
	}
}

func SetFieldValue(field reflect.Value, value interface{}) {
	if field.CanInterface() && field.CanAddr() && field.CanSet() {
		field.Set(reflect.ValueOf(value))
	} else {
		// for all unexported fields (private)
		reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).
			Elem().
			Set(reflect.ValueOf(value))
	}
}
func GetFieldValueFromMethodAndReflectValue(val reflect.Value, name string) reflect.Value {
	if val.Kind() == reflect.Ptr {
		method := val.MethodByName(name)
		if method.Kind() == reflect.Func {
			res := method.Call(nil)
			return res[0]
		}
	} else if val.Kind() == reflect.Struct {
		method := val.MethodByName(name)
		if method.Kind() == reflect.Func {
			res := method.Call(nil)
			return res[0]
		} else {
			//https://www.geeksforgeeks.org/reflect-addr-function-in-golang-with-examples/
			pointerType := val.Addr()
			method := pointerType.MethodByName(name)
			res := method.Call(nil)
			return res[0]
		}
	}

	return *new(reflect.Value)
}
