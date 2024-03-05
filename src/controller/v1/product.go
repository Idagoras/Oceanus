package v1

import (
	"bluesell/src/alarm"
	"bluesell/src/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddProduct(c *gin.Context) {
	name := c.Query("name")
	str, err := hello(name)
	var res = entity.Result{}
	if err != nil {
		res.SetCode(entity.CODE_ERROR)
		res.SetMessage(err.Error())
		c.JSON(http.StatusOK, res)
		c.Abort()
		return
	}
	res.SetCode(entity.CODE_SUCCESS)
	res.SetMessage(str)
	c.JSON(http.StatusOK, res)
}

func hello(name string) (str string, err error) {
	if name == "" {
		err = alarm.Wechat("name 不能为空")
		return
	}
	str = fmt.Sprintf("hellp: %s", name)
	return
}
