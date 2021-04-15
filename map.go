package functional

import (
	"reflect"
)

type reducingFn func(reflect.Value, reflect.Value) reflect.Value
type xfn func(interface{}) reducingFn

func Mapping(f interface{}) xfn {
	fv := reflect.ValueOf(f)
	return func(f1 interface{}) reducingFn {
		if rf, ok := f1.(reducingFn); ok {
			return func(result, input reflect.Value) reflect.Value {
				return rf(result, fv.Call([]reflect.Value{input})[0])
			}
		}
		f1v := reflect.ValueOf(f1)
		return func(result, input reflect.Value) reflect.Value {
			return f1v.Call([]reflect.Value{result, fv.Call([]reflect.Value{input})[0]})[0]
		}
	}
}

func Filtering(pred interface{}) xfn {
	pv := reflect.ValueOf(pred)
	return func(f1 interface{}) reducingFn {
		if rf, ok := f1.(reducingFn); ok {
			return func(result, input reflect.Value) reflect.Value {
				if pv.Call([]reflect.Value{input})[0].Interface().(bool) {
					return rf(result, input)
				}
				return result
			}
		}
		f1v := reflect.ValueOf(f1)
		return func(result, input reflect.Value) reflect.Value {
			if pv.Call([]reflect.Value{input})[0].Interface().(bool) {
				return f1v.Call([]reflect.Value{result, input})[0]
			}
			return result
		}
	}
}

func Taking(n int) xfn {
	return func(f1 interface{}) reducingFn {
		if rf, ok := f1.(reducingFn); ok {
			return func(result, input reflect.Value) reflect.Value {
				if n > 0 {
					n--
					return rf(result, input)
				}
				return result
			}
		}
		f1v := reflect.ValueOf(f1)
		return func(result, input reflect.Value) reflect.Value {
			if n > 0 {
				n--
				return f1v.Call([]reflect.Value{result, input})[0]
			}
			return result
		}
	}
}

func Comp(fs ...xfn) func(interface{}) reducingFn {
	return func(arg interface{}) reducingFn {
		rf := fs[0](arg)
		for _, f := range fs[1:] {
			rf = f(rf)
		}
		return rf
	}
}

func Reduce(f reducingFn, result, input interface{}) interface{} {
	out := reflect.ValueOf(result)
	iv := reflect.ValueOf(input)
	for i := 0; i < iv.Len(); i++ {
		out = f(out, iv.Index(i))
	}
	return out.Interface()
}
