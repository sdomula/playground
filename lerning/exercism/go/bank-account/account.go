package account

import "sync"

const testVersion = 1

// Account represents a bank account
type Account struct {
	mu      sync.Mutex
	balance int64
	closed  bool
}

// Open creates a new bank account with the given `initialDeposit`.
// Open retuns `nil` if the `initialDeposit` is negativ.
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{
		balance: initialDeposit,
		closed:  false,
	}
}

// Close disables an account.
func (a *Account) Close() (payout int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.closed {
		return 0, false
	}
	a.closed = true
	return a.balance, true
}

// Balance returns the accounts current balance.
func (a *Account) Balance() (balance int64, ok bool) {
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

// Deposit adds or removes money from the account.
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.closed {
		return 0, false
	}
	nb := a.balance + amount
	if nb < 0 {
		return a.balance, false
	}
	a.balance = nb
	return a.balance, true
}
