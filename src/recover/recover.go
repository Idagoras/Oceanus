package recover

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"oceanus/src/alarm"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				alarm.Panic(fmt.Sprintf("%s", r))
			}
		}()
		c.Next()
	}
}
