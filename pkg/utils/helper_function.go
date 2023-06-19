package utils

import (
	"crypto/rand"
	"math/big"
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

func Generate6Digit() string {

	min := big.NewInt(100000)
	max := big.NewInt(999999)

	// Generate a random number between min and max
	randomNumber, _ := rand.Int(rand.Reader, new(big.Int).Sub(max, min))
	randomNumber.Add(randomNumber, min)
	return randomNumber.String()
}
