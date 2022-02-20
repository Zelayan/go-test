package pointers

type Wallet struct {
	balance int
}


func (w Wallet) Deposit (amount int) {

}

func (w Wallet) Balance() int{
	return 0
}

func (w Wallet) Deposit(amount int) {
    w.balance += amount
}

func (w Wallet) Balance() int {
    return w.balance
}