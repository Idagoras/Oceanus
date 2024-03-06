package common

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	"oceanus/src/errno"
)

const (
	HashOperationFail     = "hash operation fail"
	PasswordIsUncorrected = "password is uncorrected"
)

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ServerErrorHandle(err error, ctx *gin.Context) {
	response := ErrorResponse{
		Code:    errno.ErrServer.Code(),
		Message: errno.ErrServer.Message(),
		Data:    errno.ErrServer.Data(),
	}
	ctx.JSON(http.StatusInternalServerError, response)
	return
}

func UnauthorizedErrorHandle(err error, ctx *gin.Context) {
	response := ErrorResponse{
		Code:    errno.ErrUserAuthorizationMissing.Code(),
		Message: errno.ErrUserAuthorizationMissing.Message(),
		Data:    errno.ErrUserAuthorizationMissing.Data(),
	}
	ctx.JSON(http.StatusUnauthorized, response)
	return
}

func PostgresErrorHandle(err *pq.Error, ctx *gin.Context) {
	code := err.Code.Name()
	switch code {
	case "foreign_key_violation", "unique_violation":
		response := ErrorResponse{
			Code:    errno.ErrForbiddenOperation.Code(),
			Message: errno.ErrForbiddenOperation.Message(),
			Data:    errno.ErrForbiddenOperation.Data(),
		}
		ctx.JSON(http.StatusForbidden, response)
	}
	return
}

func BindJSONErrorHandle(err error, ctx *gin.Context) {
	sErr := errno.ErrParam
	response := ErrorResponse{
		Code:    sErr.Code(),
		Message: sErr.Message(),
		Data:    sErr.Data(),
	}
	ctx.JSON(http.StatusBadRequest, response)
}

func PasswordErrorHandle(err error, ctx *gin.Context, errorType string) {
	switch errorType {
	case HashOperationFail:
		response := ErrorResponse{
			Code:    errno.ErrServer.Code(),
			Message: errno.ErrServer.Message(),
			Data:    errno.ErrServer.Data(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	case PasswordIsUncorrected:
		response := ErrorResponse{
			Code:    errno.ErrUserPassword.Code(),
			Message: errno.ErrUserPassword.Message(),
			Data:    errno.ErrUserPassword.Data(),
		}
		ctx.JSON(http.StatusOK, response)
		return
	}
	return

}

func UserNotFoundErrorHandle(err error, ctx *gin.Context) {
	response := ErrorResponse{
		Code:    errno.ErrUserNotExist.Code(),
		Message: errno.ErrUserNotExist.Message(),
		Data:    errno.ErrUserNotExist.Data(),
	}
	ctx.JSON(http.StatusNotFound, response)
	return
}

func AccountNotBelongToUserErrorHandle(err error, ctx *gin.Context) {
	response := ErrorResponse{
		Code:    errno.ErrAccountNotBelongToUser.Code(),
		Message: errno.ErrAccountNotBelongToUser.Message(),
		Data:    errno.ErrAccountNotBelongToUser.Data(),
	}
	ctx.JSON(http.StatusUnauthorized, response)
	return
}

func AbortedAuthorizationHeaderNotProvided(err error, ctx *gin.Context) {
	response := ErrorResponse{
		Code:    errno.ErrUserAuthorizationMissing.Code(),
		Message: errno.ErrUserAuthorizationMissing.Message(),
		Data:    errno.ErrUserAuthorizationMissing.Data(),
	}
	ctx.JSON(http.StatusUnauthorized, response)
	ctx.Abort()
}

func AbortedAuthorizationHeaderFormatUncorrected(err error, ctx *gin.Context) {
	response := ErrorResponse{
		Code:    errno.ErrUserAuthorizationFormat.Code(),
		Message: errno.ErrUserAuthorizationFormat.Message(),
		Data:    errno.ErrUserAuthorizationFormat.Data(),
	}
	ctx.JSON(http.StatusUnauthorized, response)
	ctx.Abort()
}

func AbortedUnsupportedAuthorizationType(err error, ctx *gin.Context) {
	response := ErrorResponse{
		Code:    errno.ErrUserUnsupportedAuthorizationType.Code(),
		Message: errno.ErrUserUnsupportedAuthorizationType.Message(),
		Data:    errno.ErrUserUnsupportedAuthorizationType.Data(),
	}
	ctx.JSON(http.StatusUnauthorized, response)
	ctx.Abort()
}

func AbortedTokenVerfiyFail(err error, ctx *gin.Context) {
	response := ErrorResponse{
		Code:    errno.ErrUserInvalidToken.Code(),
		Message: errno.ErrUserInvalidToken.Message(),
		Data:    errno.ErrUserInvalidToken.Data(),
	}
	ctx.JSON(http.StatusUnauthorized, response)
	ctx.Abort()
}
