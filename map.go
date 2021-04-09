package functional

import "reflect"

type reducingFn func(interface{}, interface{}) interface{}
type xfn func(interface{}) reducingFn

func Mapping(f interface{}) xfn {
	fv := reflect.ValueOf(f)
	return func(f1 interface{}) reducingFn {
		f1v := reflect.ValueOf(f1)
		return func(result, input interface{}) interface{} {
			rv := reflect.ValueOf(result)
			iv := reflect.ValueOf(input)
			return f1v.Call([]reflect.Value{rv, fv.Call([]reflect.Value{iv})[0]})[0].Interface()
		}
	}
}

func Filtering(pred interface{}) xfn {
	pv := reflect.ValueOf(pred)
	return func(f1 interface{}) reducingFn {
		f1v := reflect.ValueOf(f1)
		return func(result, input interface{}) interface{} {
			rv := reflect.ValueOf(result)
			iv := reflect.ValueOf(input)
			if pv.Call([]reflect.Value{iv})[0].Interface().(bool) {
				return f1v.Call([]reflect.Value{rv, iv})[0].Interface()
			}
			return result
		}
	}
}

func Taking(n int) xfn {
	return func(f1 interface{}) reducingFn {
		f1v := reflect.ValueOf(f1)
		return func(result, input interface{}) interface{} {
			rv := reflect.ValueOf(result)
			iv := reflect.ValueOf(input)
			if n > 0 {
				n--
				return f1v.Call([]reflect.Value{rv, iv})[0].Interface()
			}
			return result
		}
	}
}

func Comp(fs ...interface{}) func(...interface{}) interface{} {
	return func(args ...interface{}) interface{} {
		in := make([]reflect.Value, len(args))
		for i, a := range args {
			in[i] = reflect.ValueOf(a)
		}
		for i := len(fs) - 1; i >= 0; i-- {
			f := fs[i]
			fv := reflect.ValueOf(f)
			in = fv.Call(in)
		}
		return in[0].Interface()
	}
}

func Reduce(f, result, input interface{}) interface{} {
	out := reflect.ValueOf(result)
	iv := reflect.ValueOf(input)
	fv := reflect.ValueOf(f)
	for i := 0; i < iv.Len(); i++ {
		out = fv.Call([]reflect.Value{out, iv.Index(i)})[0]
	}
	return out.Interface()
}
