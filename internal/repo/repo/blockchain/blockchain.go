package blockchain

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"nam_0515/internal/model"
	"net/http"
)

type Repository struct {
	Blockchain *model.Blockchain
	PoS        *model.PoS
}

func NewBlockChainRepository(blockchain *model.Blockchain, poS *model.PoS, peers []string) *Repository {
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

func (r *Repository) ListTransactions(ctx context.Context, address string) ([]*model.Transaction, error) {
	if r.Blockchain.Transactions[address] == nil {
		return nil, errors.New("address not existed")
	}
	return r.Blockchain.Transactions[address], nil
}

func (r *Repository) CreateAccount(address string) *model.Account {
	account := r.Blockchain.CreateAccount(address)
	return account
}

func (r *Repository) CreateTransaction(ctx context.Context, req *model.Transaction, publicKey model.PublicKeyJSON, signature string) (*model.Transaction, error) {
	validateRequest := model.CreateTransactionValidateRequest{
		From:          req.From,
		To:            req.To,
		Amount:        req.Amount,
		PublicKeyJSON: publicKey,
		Signature:     signature,
	}
	var validatorStr string
	if len(r.PoS.Validators) == 0 {
		err := r.Validate(ctx, &validateRequest)
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

	//Add history
	r.Blockchain.AddTransactionToList(req.From, req)
	r.Blockchain.AddTransactionToList(req.To, req)

	go r.broadcastBlock(jsonData)

	r.Blockchain.UpdateBalance(req)
	err = r.PoS.AddStake(r.Blockchain, &model.NewStakeRequest{
		Address: validatorStr,
		Amount:  1, // pay for validator
	})
	if err != nil {
		// TODO something here
	}
	return req, nil
}

func (r *Repository) broadcastBlock(data []byte) {
	for _, validator := range r.PoS.Validators {
		http.Post(validator.Peer+"/addBlock", "application/json", bytes.NewBuffer(data))
	}
}

func sendHTTPRequestToValidator(validatorPeer string, requestData interface{}) error {
	// Chuyển đổi dữ liệu request thành JSON
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	// Tạo HTTP request
	req, err := http.NewRequest("POST", validatorPeer+"/validate", bytes.NewBuffer(requestBody))
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

func (r *Repository) Validate(ctx context.Context, req *model.CreateTransactionValidateRequest) error {
	if r.Blockchain.Accounts[req.From].Balance < req.Amount {
		return errors.New("balance not enough")
	}

	b, err := json.Marshal(&model.DataForSign{
		To:     req.To,
		Amount: req.Amount,
	})
	if err != nil {
		return err
	}

	publicKey, err := jsonToPublicKey(req.PublicKeyJSON)
	if err != nil {
		return err
	}

	// xac thuc chu ki
	if !verifySignature(publicKey, b, req.Signature) {
		return errors.New("fail to validate transaction")
	}
	return nil
}

// TODO sign contract. only running in client
func signMessage(privateKey *ecdsa.PrivateKey, data model.DataForSign) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(b)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return "", err
	}
	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)
	return hex.EncodeToString(signature), nil
}

// Xác thực chữ ký của một thông điệp bằng public key
func verifySignature(userPublicKey *ecdsa.PublicKey, data []byte, signature string) bool {
	sigBytes, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}
	r := big.Int{}
	s := big.Int{}
	r.SetBytes(sigBytes[:len(sigBytes)/2])
	s.SetBytes(sigBytes[len(sigBytes)/2:])

	hash := sha256.Sum256(data)
	return ecdsa.Verify(userPublicKey, hash[:], &r, &s)
}

func jsonToPublicKey(pubKeyJSON model.PublicKeyJSON) (*ecdsa.PublicKey, error) {
	xBytes, err := hex.DecodeString(pubKeyJSON.X)
	if err != nil {
		return nil, fmt.Errorf("failed to decode X: %v", err)
	}
	yBytes, err := hex.DecodeString(pubKeyJSON.Y)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Y: %v", err)
	}

	pubKey := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     new(big.Int).SetBytes(xBytes),
		Y:     new(big.Int).SetBytes(yBytes),
	}

	return pubKey, nil
}
