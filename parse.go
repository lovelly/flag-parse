package parse

import (
	"github.com/kballard/go-shellquote"
	"strconv"
	"strings"
)

const (
	NoKeyValue = "#nokeyvalue-parse_"
)

type Result struct {
	data map[string]string
	sortKeys []string
}

func (r *Result) GetString(key string) string {
	v := r.data[key]
	if len(v) < 1 {
		return ""
	}
	return v
}


func (r *Result) GetInt(key string) int {
	v := r.data[key]
	if len(v) < 1 {
		return 0
	}
	v1, err := strconv.ParseInt(v, 10, 64)
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
	v1, err := strconv.ParseFloat(v, 64)
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
	v1, err := strconv.ParseBool(v)
	if err != nil {
		return false
	}
	return v1
}

func (r *Result) IsExist(key string) bool {
	_, ok := r.data[key]
	return ok
}

func (r *Result) DelKey(key string)  {
	delete(r.data, key)
}

func (r *Result)String() string {
	var list []string
	for _, k := range r.sortKeys {
		data, ok := r.data[k]
		if !ok {
			continue
		}

		if !strings.Contains(k,NoKeyValue) {
			list = append(list, "-"+k)
		}

		if data != "" {
			list = append(list, data)
		}
	}
	return shellquote.Join(list...)
}

func ParseArgs(cmdStr string) (*Result, error) {
	list, err :=  shellquote.Split(cmdStr)
	if err != nil {
		return nil, err
	}
	r := &Result{data: make(map[string]string)}
	length := len(list)
	for i := 0; i < length; i++ {
		v := list[i]
		if v[0] != '-' {
			nk :=  NoKeyValue+strconv.Itoa(i)
			r.sortKeys = append(r.sortKeys,nk)
			r.data[nk] = v
			continue
		}
		k := strings.TrimLeft(v, "-")
		r.sortKeys = append(r.sortKeys, k)
		if i+1>= length {
			r.data[k] = ""
			break
		}

		if list[i+1][0] == '-' {
			r.data[k] = ""
			continue
		}

		i++
		r.data[k] = list[i]
	}
	return r, nil
}
