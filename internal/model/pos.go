package model

import (
	"bytes"
	"errors"
	"math/rand"
	"net/http"
	"sync"
)

type Validator struct {
	Address string
	Stake   int
	Peer    string
}

type PoS struct {
	Validators []Validator
	mux        sync.Mutex
}

type NewStakeRequest struct {
	Address string
	Amount  int
	Peer    string
}

func (pos *PoS) AddStake(blockChain *Blockchain, stake *NewStakeRequest) error {
	pos.mux.Lock()
	blockChain.mux.Lock()
	defer func() {
		pos.mux.Unlock()
		blockChain.mux.Unlock()
	}()

	if blockChain.Accounts[stake.Address].Balance > stake.Amount {
		blockChain.Accounts[stake.Address].Balance = blockChain.Accounts[stake.Address].Balance - stake.Amount
	} else {
		return errors.New("amount not enough")
	}

	for i, v := range pos.Validators {
		if v.Address == stake.Address {
			pos.Validators[i].Stake += stake.Amount
			return nil
		}
	}

	pos.Validators = append(pos.Validators, Validator{stake.Address, stake.Amount, stake.Peer})

	return nil
}

func (pos *PoS) broadcastStake(data []byte) {
	for _, validator := range pos.Validators {
		http.Post(validator.Peer+"/changeStake", "application/json", bytes.NewBuffer(data))
	}
}

func (pos *PoS) getTotalStakes() int {
	pos.mux.Lock()
	defer pos.mux.Unlock()
	total := 0
	for _, v := range pos.Validators {
		total += v.Stake
	}
	return total
}

func (pos *PoS) SelectValidator() *Validator {
	pos.mux.Lock()
	defer pos.mux.Unlock()
	totalStakes := pos.getTotalStakes()
	randomPoint := rand.Intn(totalStakes)
	currentPoint := 0

	for _, v := range pos.Validators {
		currentPoint += v.Stake
		if currentPoint > randomPoint {
			return &v
		}
	}
	return nil
}
