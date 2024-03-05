package router

import (
	"github.com/gin-gonic/gin"
	"net/url"
	"oceanus/src/common"
	"strconv"
)

func SignDemo(c *gin.Context) {
	ts := strconv.FormatInt(common.GetTimeUnix(), 10)
	res := map[string]interface{}{}
	params := url.Values{
		"name":  []string{"a"},
		"price": []string{"10"},
		"ts":    []string{ts},
	}
	res["sn"] = common.CreateSign(params)
	res["ts"] = ts
	common.RetJson("200", "", res, c)
}
