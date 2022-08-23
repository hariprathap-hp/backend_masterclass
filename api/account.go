package api

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/hariprathap-hp/backend_masterclass/db/sqlc"
	"github.com/hariprathap-hp/backend_masterclass/token"
	"github.com/lib/pq"
)

type createAccountRequest struct {
	Currency string `json:"currency" binding:"required,oneof=USD EUR INR"`
}

func (server *Server) createAccount(c *gin.Context) {
	var req createAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateAccountParams{
		Owner:    authPayload.Username,
		Balance:  0,
		Currency: req.Currency,
	}
	account, err := server.store.CreateAccount(c, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); !ok {
			log.Println(pqErr.Code.Name())
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(c *gin.Context) {
	var req getAccountRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(c, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	if account.Owner != authPayload.Username {
		err := errors.New("account doesn't belong to the authenticated user")
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, account)
}

type listAccountRequest struct {
	ID        int32 `form:"page_id" binding:"required,min=1"`
	Page_Size int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAccount(c *gin.Context) {
	var req listAccountRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListAccountsParams{
		Owner:  authPayload.Username,
		Limit:  req.Page_Size,
		Offset: (req.ID - 1) * req.Page_Size,
	}

	account, err := server.store.ListAccounts(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, account)
}

type delAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteAccount(c *gin.Context) {
	var req delAccountRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteAccount(c, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, "account deleted")
}

type updateAccountRequest struct {
	ID      int64 `json:"id" binding:"required"`
	Balance int64 `json:"balance" binding:"required,min=1"`
}

func (server *Server) updateAccount(c *gin.Context) {
	var req updateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateAccountParams{
		ID:      req.ID,
		Balance: req.Balance,
	}

	account, err := server.store.UpdateAccount(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, account)
}
