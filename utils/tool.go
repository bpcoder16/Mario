package utils

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
)

func GetBodyClone(ctx *gin.Context) (io.ReadCloser, error) {
	r, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, err
	}
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(r))
	return ioutil.NopCloser(bytes.NewReader(r)), nil
}
