package model

type CreateTransactionRequest struct {
	To            string        `json:"to"`     // account id nguoi nhan
	Amount        int           `json:"amount"` // so tien
	PublicKeyJSON PublicKeyJSON `json:"public_key_json"`
	Signature     string        `json:"signature"`
}

type Transaction struct {
	From       string `json:"from"`
	To         string `json:"to"`
	Amount     int    `json:"amount"`
	BlockIndex int    `json:"block_index"`
}

type ListTransactionResponse struct {
	Transactions []*Transaction
}

type PublicKeyJSON struct {
	X string `json:"x"`
	Y string `json:"y"`
}

type DataForSign struct {
	To     string `json:"to"`     // account id nguoi nhan
	Amount int    `json:"amount"` // so tien
}

type CreateTransactionValidateRequest struct {
	From          string        `json:"from"`
	To            string        `json:"to"`     // account id nguoi nhan
	Amount        int           `json:"amount"` // so tien
	PublicKeyJSON PublicKeyJSON `json:"public_key_json"`
	Signature     string        `json:"signature"`
}
