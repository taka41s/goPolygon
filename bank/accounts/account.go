package users

import("fmt")

type LoanerAccount struct {
	Name string
	Balance float64
}

type Account struct {
	Name        string
	Balance     float64
	Investments []*Investment
}

type User interface {
   	Deposit(value float64)
	Withdrawal(value float64)
	Transfer(value float64, user *Account)
}

func (p *Account) Withdrawal(value float64) {
	if(p.Balance >= value){
		p.Balance -= value		
	} else {
		fmt.Println("insufficient funds")
	}
}

func (p *Account) Deposit(value float64) {
	if value < 0 {
		fmt.Errorf("deposit value must be positive")
	}

	p.Balance += value

	fmt.Println(p.Name, "Balance:", p.Balance)
}

func (p *Account) Transfer(value float64, user *Account) {
	if value > user.Balance {
		fmt.Errorf("Need more funds")
	}

	user.Balance += value
	p.Balance -= value
}

func (p *LoanerAccount) Withdrawal(value float64) {
	limit := -500.0

	if p.Balance >= value || (p.Balance > limit && p.Balance-value >= limit) {
		p.Balance -= value
	} else {
		fmt.Println("Insufficient funds or exceeding withdrawal limit")
	}

	fmt.Println(p.Balance)
}