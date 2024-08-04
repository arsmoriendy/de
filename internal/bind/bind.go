package bind

import "reflect"

// Unsafe and probably slow, but it calls function f with arguments args and returns a callback function that returns the values
func Bind[T any](f interface{}, args ...interface{}) func() (T, error) {
	rf := reflect.ValueOf(f)

	// slow
	in := make([]reflect.Value, rf.Type().NumIn())
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	res := rf.Call(in)
	return func() (T, error) {
		return res[0].Interface().(T), res[1].Interface().(error)
	}
}
