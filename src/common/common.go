package common

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"oceanus/src/config"
	"sort"
	"strconv"
	"time"
)

const (
	USD                    = "USD"
	EUR                    = "EUR"
	CAD                    = "CAD"
	CNY                    = "CNY"
	AuthorizationHeaderKey = "authorization"
)

func GetTimeUnix() int64 {
	return time.Now().Unix()

}

func CreateSign(params url.Values) string {
	var key []string
	for k := range params {
		if k != "sn" {
			key = append(key, k)
		}
	}
	sort.Strings(key)
	var strBuf *bytes.Buffer
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str := fmt.Sprintf("%v=%v", key[i], params.Get(key[i]))
			strBuf = bytes.NewBufferString(str)
		} else {
			strBuf.WriteString(fmt.Sprintf("&%v=%v", key[i], params.Get(key[i])))
		}
	}
	sign, _ := strBuf.ReadString(' ')
	return MD5(MD5(sign) + MD5(config.APP_NAME+config.APP_SECRET))
}

func VerifySign(c *gin.Context) {
	method := c.Request.Method
	var ts int64
	var sn string
	var req url.Values

	if method == "GET" {
		req = c.Request.URL.Query()
		sn = c.Query("sn")
		ts, _ = strconv.ParseInt(c.Query("ts"), 10, 64)
	} else if method == "POST" {
		req = c.Request.PostForm
		sn = c.PostForm("sn")
		ts, _ = strconv.ParseInt(c.PostForm("ts"), 10, 64)
	} else {
		RetJson("500", "Illegal requests", "", c)
		return
	}

	exp, _ := strconv.ParseInt(config.API_EXPIRY, 10, 64)

	if ts > GetTimeUnix() || GetTimeUnix()-ts >= exp {
		RetJson("500", "Ts Error", "", c)
		return
	}

	if sn == "" || sn != CreateSign(req) {
		RetJson("500", "Sn Error", "", c)
		return
	}
}

func MD5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	hashInBytes := hash.Sum(nil)
	return hex.EncodeToString(hashInBytes)
}

func RetJson(code, message string, data interface{}, c *gin.Context) {
	RetJsonWithSpecificHttpStatusCode(http.StatusOK, code, message, data, c)
}

func RetJsonWithSpecificHttpStatusCode(httpStatusCode int, code, message string, data interface{}, c *gin.Context) {
	c.JSON(httpStatusCode, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
	c.Abort()
}

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, CNY:
		return true
	}
	return false
}
