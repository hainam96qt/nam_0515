package model

import (
	"sync"
)

type Blockchain struct {
	Blocks       []*Block
	mux          sync.Mutex
	Validators   []*Validator              // danh sach cac validator
	Accounts     map[string]*Account       // Map địa chỉ người dùng với tài khoản
	Transactions map[string][]*Transaction // Mỗi người gửi và người nhận đều được ghi lại transaction đễ truy vấn
}

func (bc *Blockchain) AddBlock(data []byte, validator string) {
	bc.mux.Lock()
	defer bc.mux.Unlock()

	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := createBlock(prevBlock, data, validator)

	// validate block
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	genesisBlock := createGenesisBlock()
	genesisBlock.Hash = calculateHash(genesisBlock)
	return &Blockchain{Blocks: []*Block{genesisBlock}}
}

func (bc *Blockchain) CreateAccount(address string) *Account {
	bc.mux.Lock()
	defer bc.mux.Unlock()

	account := &Account{
		Address: address,
		Balance: 0,
	}

	// TODO remove. To define value transfer for test in internal blockchain
	account.Balance = 100

	bc.Accounts[address] = account
	return account
}

func (bc *Blockchain) GetAccount(address string) *Account {
	return bc.Accounts[address]
}

func (bc *Blockchain) UpdateBalance(tx *Transaction) {
	bc.mux.Lock()
	defer bc.mux.Unlock()

	fromAccount := bc.GetAccount(tx.From)
	toAccount := bc.GetAccount(tx.To)

	if fromAccount != nil {
		fromAccount.Balance -= tx.Amount
	}
	if toAccount != nil {
		toAccount.Balance += tx.Amount
	}
}

func (bc *Blockchain) AddTransactionToList(address string, tx *Transaction) {
	bc.mux.Lock()
	defer bc.mux.Unlock()

	bc.Transactions[address] = append(bc.Transactions[address], tx)
}
