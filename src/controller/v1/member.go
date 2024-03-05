package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oceanus/src/entity"
)

func AddMember(c *gin.Context) {
	res := entity.Result{}
	mem := entity.Member{}

	if err := c.ShouldBind(&mem); err != nil {
		res.SetCode(entity.CODE_ERROR).SetMessage(err.Error())
		c.JSON(http.StatusForbidden, res)
		c.Abort()
		return
	}

	data := map[string]interface{}{
		"name": mem.Name,
		"age":  mem.Age,
	}
	res.SetCode(entity.CODE_ERROR).SetData(data)
	c.JSON(http.StatusOK, res)
}
