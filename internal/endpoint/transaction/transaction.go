package transaction

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"nam_0515/internal/model"
	error2 "nam_0515/pkg/error"
	"nam_0515/pkg/midleware"
	"nam_0515/pkg/util/request"
	"nam_0515/pkg/util/response"
	"net/http"
	"strconv"
)

type (
	Endpoint struct {
		transactionSvc TransactionService
	}

	TransactionService interface {
		CreateTransaction(ctx context.Context, req *model.CreateTransactionRequest) (*model.Transaction, error)
		ListTransactions(ctx context.Context, accountID string) (*model.ListTransactionResponse, error)
	}
)

func InitTransactionHandler(r *chi.Mux, transactionSvc TransactionService) {
	transactionEndpoint := &Endpoint{
		transactionSvc: transactionSvc,
	}

	r.Route("/api/users/{user_id}/transactions", func(r chi.Router) {
		r.Use(midleware.Auth.ValidateRoleUser)

		r.Get("/", transactionEndpoint.ListTransactions)
		r.Post("/", transactionEndpoint.ListTransactions)
	})
}

func (e *Endpoint) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req model.CreateTransactionRequest
	if err := request.DecodeJSON(ctx, r.Body, &req); err != nil {
		log.Printf("read request body error: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	res, err := e.transactionSvc.CreateTransaction(ctx, &req)
	if err != nil {
		log.Printf("failed to create user: %s \n", err)
		response.Error(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, res)
}

func (e *Endpoint) ListTransactions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		log.Printf("failed to get list transactions: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	userIDCtx := ctx.Value("UserID").(int32)
	if userIDCtx != int32(userID) {
		err := error2.NewXError("user can not access info in other user", http.StatusUnauthorized)
		log.Printf("failed to get list transactions: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	accountIDStr := r.URL.Query().Get("account_id")

	res, err := e.transactionSvc.ListTransactions(ctx, accountIDStr)
	if err != nil {
		log.Printf("failed to get list account: %s \n", err)
		response.Error(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, res.ListTransactions)
}
