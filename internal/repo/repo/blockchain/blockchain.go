package blockchain

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"nam_0515/internal/model"
	"net/http"
)

type Repository struct {
	Blockchain *model.Blockchain
	PoS        *model.PoS
}

func NewPostgresRepository(blockchain *model.Blockchain, poS *model.PoS, peers []string) *Repository {
	// neu ko fetch block tu mot node thi
	if blockchain == nil {
		blockchain = model.NewBlockchain()
	}

	if poS == nil {
		poS = &model.PoS{}
	}

	return &Repository{
		Blockchain: blockchain,
		PoS:        poS,
	}
}

func (r *Repository) CreateTransaction(ctx context.Context, req *model.Transaction) (*model.Transaction, error) {
	var validatorStr string
	if len(r.PoS.Validators) == 0 {
		err := r.Validate(ctx, req)
		if err != nil {
			return nil, err
		}
	} else {
		validator := r.PoS.SelectValidator()
		err := sendHTTPRequestToValidator(validator.Peer, req)
		if err != nil {
			return nil, err
		}
		validatorStr = validator.Address
	}

	req.BlockIndex = len(r.Blockchain.Blocks)

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	r.Blockchain.AddBlock(jsonData, validatorStr)
	go r.broadcastBlock(jsonData)

	r.Blockchain.UpdateBalance(req)
	return req, nil
}

func (r *Repository) broadcastBlock(data []byte) {
	for _, validator := range r.PoS.Validators {
		http.Post(validator.Peer+"/addBlock", "application/json", bytes.NewBuffer(data))
	}
}

func sendHTTPRequestToValidator(validatorURL string, requestData interface{}) error {
	// Chuyển đổi dữ liệu request thành JSON
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	// Tạo HTTP request
	req, err := http.NewRequest("POST", validatorURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Gửi HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (r *Repository) Validate(ctx context.Context, req *model.Transaction) error {
	return nil
}
