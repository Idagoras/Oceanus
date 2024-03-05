package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"oceanus/src/common"
	"oceanus/src/token"
	"strings"
)

const (
	AuthorizationHeaderKey  = "authorization"
	AuthorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

func Authorization(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(AuthorizationHeaderKey)
		fmt.Printf("header:%s\n", authorizationHeader)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			common.AbortedAuthorizationHeaderNotProvided(err, ctx)
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			common.AbortedAuthorizationHeaderFormatUncorrected(err, ctx)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != AuthorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			common.AbortedUnsupportedAuthorizationType(err, ctx)
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			common.AbortedTokenVerfiyFail(err, ctx)
			return
		}
		ctx.Set(AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}
