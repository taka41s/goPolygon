package main

import(
	"bank/users"
	"fmt"
)
func main(){
	var investor users.Investor

	investor = &users.Account{"Person 8", 1000, nil}

	investment1 := &users.Investment{"Stocks", 500}
	investor.AddInvestment(investment1)

	investment2 := &users.Investment{"Bonds", 300}
	investor.AddInvestment(investment2)

	fmt.Println("Investments:")
	for _, inv := range investor.GetInvestments() {
		fmt.Printf("%s: $%.2f\n", inv.Name, inv.Amount)
	}

	investor.RemoveInvestment(investment2)

	fmt.Println("Investments:")
	for _, inv := range investor.GetInvestments() {
		fmt.Printf("%s: $%.2f\n", inv.Name, inv.Amount)
	}

	var user1 users.User

	user1 = &users.Account{"Person 1", 0, nil}
	user1.Deposit(200)
}