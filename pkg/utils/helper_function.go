package utils

import (
	"strconv"

	"github.com/valyala/fasthttp"
)

func ReadInt(r *fasthttp.Request, param string, v int64) (int64, error) {
	p := r.URI().QueryArgs().Peek(param)
	if p == nil {
		return v, nil
	}

	return strconv.ParseInt(string(p), 10, 64)
}
