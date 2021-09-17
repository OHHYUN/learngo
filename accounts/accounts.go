package accounts

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount method
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance return
func (a Account) Balance() int {
	return a.balance
}

// Withdraw from your account
func (a *Account) Withdraw(amount int) {
	a.balance -= amount
}
