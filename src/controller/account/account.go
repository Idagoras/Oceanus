package accountController

import (
	"bluesell/src/common"
	"bluesell/src/database"
	"bluesell/src/middleware"
	"bluesell/src/token"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

type createAccountRequest struct {
	Currency string `json:"currency" binding:"required,currency"`
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type listAccountsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func CreateAccount(store database.Store, ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.RetJsonWithSpecificHttpStatusCode(http.StatusBadRequest, "-1", err.Error(), nil, ctx)
		return
	}

	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	arg := database.CreateAccountParams{
		Owner:    authPayload.Username,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := store.CreateAccount(ctx, arg)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			common.PostgresErrorHandle(pqErr, ctx)
			return
		}
	}

	ctx.JSON(http.StatusOK, account)
	return
}

func GetAccount(store database.Store, ctx *gin.Context) {
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		common.RetJsonWithSpecificHttpStatusCode(http.StatusBadRequest, "-1", err.Error(), nil, ctx)
		return
	}

	account, err := store.GetAccount(ctx, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			common.RetJsonWithSpecificHttpStatusCode(http.StatusNotFound, "-1", err.Error(), nil, ctx)
			return
		}

		common.RetJsonWithSpecificHttpStatusCode(http.StatusInternalServerError, "-1", err.Error(), nil, ctx)
		return
	}
	if account.Owner != authPayload.Username {
		common.AccountNotBelongToUserErrorHandle(errors.New("account not belong to user"), ctx)
	}
	ctx.JSON(http.StatusOK, account)
}

func ListAccounts(store database.Store, ctx *gin.Context) {
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	var req listAccountsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		common.RetJsonWithSpecificHttpStatusCode(http.StatusBadRequest, "-1", err.Error(), nil, ctx)
	}

	arg := database.ListAccountsParams{
		Owner:  authPayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := store.ListAccounts(ctx, arg)
	if err != nil {
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				common.RetJsonWithSpecificHttpStatusCode(http.StatusNotFound, "-1", err.Error(), nil, ctx)
				return
			}

			common.RetJsonWithSpecificHttpStatusCode(http.StatusInternalServerError, "-1", err.Error(), nil, ctx)
			return
		}
	}
	ctx.JSON(http.StatusOK, accounts)
}
