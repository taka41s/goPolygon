package users

type Investment struct {
	Name   string
	Amount float64
}

type Investor interface {
	User
	AddInvestment(investment *Investment)
	RemoveInvestment(investment *Investment)
	GetInvestments() []*Investment
}

func (a *Account) GetInvestments() []*Investment {
	return a.Investments
}

func (a *Account) RemoveInvestment(investment *Investment) {
	for i, inv := range a.Investments {
		if inv == investment {
			a.Investments = append(a.Investments[:i], a.Investments[i+1:]...)
			return
		}
	}
}

func (a *Account) AddInvestment(investment *Investment) {
	a.Investments = append(a.Investments, investment)
}