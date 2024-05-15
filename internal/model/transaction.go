package model

type CreateTransactionRequest struct {
	From   string `json:"from"`   // accountID nguoi gui
	To     string `json:"to"`     // account id nguoi nhan
	Amount int    `json:"amount"` // so tien
}

type Transaction struct {
	From       string `json:"from"`
	To         string `json:"to"`
	Amount     int    `json:"amount"`
	BlockIndex int    `json:"block_index"`
}

type ListTransactionResponse struct {
	ListTransactions []Transaction
}
