package userController

import (
	"bluesell/src/common"
	"bluesell/src/config"
	"bluesell/src/database"
	"bluesell/src/token"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	"time"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type userResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user database.User) userResponse {
	return userResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

func CreateUser(store database.Store, ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.BindJSONErrorHandle(err, ctx)
		return
	}
	hashedPassword, err := common.HashPassword(req.Password)
	if err != nil {
		common.PasswordErrorHandle(err, ctx, common.HashOperationFail)
		return
	}
	arg := database.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := store.CreateUser(ctx, arg)
	if err != nil {
		var pqerr *pq.Error
		if errors.As(err, &pqerr) {
			common.PostgresErrorHandle(pqerr, ctx)
			return
		}
		common.ServerErrorHandle(err, ctx)
		return
	}
	response := newUserResponse(user)
	ctx.JSON(http.StatusOK, response)

}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func LoginUser(store database.Store, ctx *gin.Context, tokenMaker token.Maker, config config.Config) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.BindJSONErrorHandle(err, ctx)
		return
	}
	user, err := store.GetUser(ctx, req.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			common.UserNotFoundErrorHandle(err, ctx)
			return
		}
		common.ServerErrorHandle(err, ctx)
		return
	}

	err = common.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		common.PasswordErrorHandle(err, ctx, common.PasswordIsUncorrected)
		return
	}

	accessToken, err := tokenMaker.CreateToken(user.Username, config.AccessTokenDuartion)
	if err != nil {
		common.ServerErrorHandle(err, ctx)
		return
	}

	response := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, response)
	return
}
