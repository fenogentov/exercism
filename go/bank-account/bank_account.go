// Package account manages a bank account
package account

import (
	"sync"
)

// Account
type Account struct {
	cach int64
	stat bool
	mu   *sync.RWMutex
}

// Open creates a bank account
func Open(deposit int64) *Account {
	if deposit < 0 {
		return nil
	}
	return &Account{
		cach: deposit,
		stat: true,
		mu:   &sync.RWMutex{},
	}

}

// Close closes bank account
func (account *Account) Close() (payout int64, ok bool) {
	account.mu.Lock()
	defer account.mu.Unlock()
	if !account.stat {
		return 0, false
	}
	p := account.cach
	*account = Account{
		cach: 0,
		stat: false,
		mu:   account.mu,
	}
	return p, true
}

// Balance returns the balance on the bank account
func (account *Account) Balance() (balance int64, ok bool) {
	account.mu.RLock()
	defer account.mu.RUnlock()
	return account.cach, account.stat
}

// Deposit manages the movement of funds on the account
func (account *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	account.mu.Lock()
	defer account.mu.Unlock()
	nb := account.cach + amount
	if nb < 0 || !account.stat {
		return account.cach, false
	}
	account.cach = nb
	return account.cach, true
}
