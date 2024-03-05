package transferController

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"oceanus/src/common"
	"oceanus/src/database"
	"oceanus/src/middleware"
	"oceanus/src/token"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Currency      string `json:"currency" binding:"required,currency"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
}

func createTransfer(store database.Store, ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.RetJsonWithSpecificHttpStatusCode(http.StatusBadRequest, "-1", err.Error(), nil, ctx)
		return
	}
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	fromAccount, valid := validAccount(ctx, req.FromAccountID, req.Currency, store)
	if !valid {
		return
	}
	if authPayload.Username != fromAccount.Owner {
		common.AccountNotBelongToUserErrorHandle(errors.New("from account doesn't belong to the authorization user"), ctx)
		return
	}

	_, valid = validAccount(ctx, req.ToAccountID, req.Currency, store)
	if !valid {
		return
	}

	arg := database.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}
	sqlStore := store.(*database.SQLStore)
	result, err := sqlStore.TransferTx(ctx, arg)
	if err != nil {
		common.RetJsonWithSpecificHttpStatusCode(http.StatusInternalServerError, "-1", err.Error(), nil, ctx)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func validAccount(ctx *gin.Context, accountID int64, currency string, store database.Store) (database.Account, bool) {
	account, err := store.GetAccount(ctx, accountID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			common.RetJsonWithSpecificHttpStatusCode(http.StatusNotFound, "-1", err.Error(), nil, ctx)
			return account, false
		}

		common.RetJsonWithSpecificHttpStatusCode(http.StatusInternalServerError, "-1", err.Error(), nil, ctx)
		return account, false
	}
	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", account.ID, account.Currency, currency)
		common.RetJsonWithSpecificHttpStatusCode(http.StatusBadRequest, "-1", err.Error(), nil, ctx)
		return account, false
	}
	return account, true
}
