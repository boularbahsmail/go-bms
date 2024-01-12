package src

import "fmt"

// Account type
type Account struct {
	FullName  string
	AccountID int
	Balance   int
}

// Show account info
func (acc *Account) ShowInfo() {
	fmt.Println(separator)
	fmt.Println("Full Name :", acc.FullName,
		"\nAccound ID :", acc.AccountID,
		"\nBalance :", acc.Balance, "Dhs")
	fmt.Println(separator)
}

// Deposit money
func (acc *Account) Deposit(amount int) {
	if amount > 0 {
		acc.Balance += amount
	} else {
		fmt.Println("Unable to deposit negative amount!!")
	}

}

// Withdrawal money
func (acc *Account) Withdrawal(amount int) {
	if amount <= acc.Balance {
		acc.Balance -= amount
	} else {
		fmt.Println("Insufficient balance!!")
	}
}
