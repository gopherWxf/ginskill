package result

import "errors"

type ErrorResult struct {
	data interface{}
	err  error
}

func (this *ErrorResult) UnWrap() interface{} {
	if this.err != nil {
		panic(this.err.Error())
	}
	return this.data
}
func Result(vs ...interface{}) *ErrorResult {
	if len(vs) == 1 {
		if vs[0] == nil {
			return &ErrorResult{nil, nil}
		}
		if err, ok := vs[0].(error); ok {
			return &ErrorResult{nil, err}
		}
	}
	if len(vs) == 2 {
		if vs[1] == nil {
			return &ErrorResult{vs[0], nil}
		}
		if err, ok := vs[1].(error); ok {
			return &ErrorResult{nil, err}
		}
	}
	return &ErrorResult{nil, errors.New("err result format")}
}
