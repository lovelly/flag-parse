package parse

import (
	"strconv"
	"strings"

	"github.com/kballard/go-shellquote"
)

type Result struct {
	data map[string][]string
}

func (r *Result) GetStringSlice(key string) (res []string) {
	for _, v := range r.data[key] {
		res = append(res, v)
	}
	return
}

func (r *Result) GetIntSlice(key string) (res []int) {
	for _, v := range r.data[key] {
		v1, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			continue
		}
		res = append(res, int(v1))
	}
	return
}

func (r *Result) GetFloatSlice(key string) (res []float64) {
	for _, v := range r.data[key] {
		v1, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		}
		res = append(res, v1)
	}
	return
}

func (r *Result) GetString(key string) string {
	v := r.data[key]
	if len(v) < 1 {
		return ""
	}
	return v[0]
}

func (r *Result) GetParams() map[string][]string {
	return  r.data
}

func (r *Result) GetInt(key string) int {
	v := r.data[key]
	if len(v) < 1 {
		return 0
	}
	v1, err := strconv.ParseInt(v[0], 10, 64)
	if err != nil {
		return 0
	}
	return int(v1)
}

func (r *Result) GetFloat(key string) float64 {
	v := r.data[key]
	if len(v) < 1 {
		return 0
	}
	v1, err := strconv.ParseFloat(v[0], 64)
	if err != nil {
		return 0
	}
	return v1
}

func (r *Result) GetBool(key string) bool {
	v, ok := r.data[key]
	if len(v) < 1 {
		return ok
	}
	v1, err := strconv.ParseBool(v[0])
	if err != nil {
		return false
	}
	return v1
}

func (r *Result) IsExist(key string) bool {
	_, ok := r.data[key]
	return ok
}
func ParseArgs(args string) (*Result, error) {
	list, err := shellquote.Split(args)
	if err != nil {
		return nil, err
	}
	r := &Result{data: make(map[string][]string)}
	length := len(list)
	for i := 0; i < length; i++ {
		v := list[i]
		if v[0] != '-' {
			continue
		}
		k := strings.TrimLeft(v, "-")
		datas := make([]string, 0)
		for {
			if i+1 >= length {
				r.data[k] = append(r.data[k], datas...)
				break
			}
			if list[i+1][0] == '-' {
				r.data[k] = append(r.data[k], datas...)
				break
			}
			i++
			datas = append(datas, list[i])
		}
	}
	return r, nil
}
