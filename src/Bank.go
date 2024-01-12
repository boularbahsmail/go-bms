package src

import (
	"fmt"
	"strings"
)

// Separator
var separator string = "---------------------------------------------------------"

// Bank type
type BankType struct {
	Name          string
	usersAccounts []Account
}

// Bank definition
func Bank(name string) *BankType {
	return &BankType{
		Name: name,
	}
}

// Show bank info
func (b *BankType) Info() {
	fmt.Println(separator)
	fmt.Println("•• Bank : ", b.Name)
	fmt.Println(separator)
}

// FindAccountByID searches for an account by accountID
func (b *BankType) FindAccountByID(accountID int) *Account {
	for account := 0; account < len(b.usersAccounts); account++ {
		if b.usersAccounts[account].AccountID == accountID {
			return &b.usersAccounts[account]
			// Account found, Returning the account
		}
	}
	return nil // Account not found
}

// AddAccount adds account to the usersAccounts list
func (b *BankType) AddAccount(newAccount Account) {
	// Check if the account already exists
	if b.FindAccountByID(newAccount.AccountID) != nil {
		fmt.Println(separator)
		fmt.Println("=> Account with ID :", newAccount.AccountID, "Already exists")
		fmt.Println(separator)
	} else {
		b.usersAccounts = append(b.usersAccounts, newAccount)
		fmt.Println(separator)
		fmt.Println("=> Account :", newAccount.AccountID, "Added successfully!!")
		fmt.Println(separator)
	}

}

// RemoveAccount removes an account from the usersAccounts list by account ID
func (b *BankType) RemoveAccount(accountIDToRemove int) {
	for i, account := range b.usersAccounts {
		if account.AccountID == accountIDToRemove {
			// Remove the account from the slice
			b.usersAccounts = append(b.usersAccounts[:i], b.usersAccounts[i+1:]...)
			fmt.Println(separator)
			fmt.Println("=> Account :", accountIDToRemove, "Removed successfully!!")
			fmt.Println(separator)
		}
	}
}

// ShowAccounts method to display accounts
func (b *BankType) ShowAccounts() {
	fmt.Println("\n-- Accounts List:")
	fmt.Println(separator)
	for account := 0; account < len(b.usersAccounts); account++ {
		anAccount := b.usersAccounts[account]
		fmt.Println("•", anAccount.FullName+" -- ", anAccount.AccountID)
	}
}

// Transfer method to make transactions
func (b *BankType) Transfer(firstAccountID, secondAccountID, amount int) error {
	// Find the accounts
	fromAccount := b.FindAccountByID(firstAccountID)
	toAccount := b.FindAccountByID(secondAccountID)

	if fromAccount == nil || toAccount == nil {
		fmt.Println("One or both accounts not found")
	}

	// Perform the transfer
	if amount > 0 {
		if amount < fromAccount.Balance {
			fromAccount.Withdrawal(amount)
			toAccount.Deposit(amount)
			fmt.Println(separator)
			fmt.Println("Transaction succeeded from", strings.ToUpper(fromAccount.FullName), "to", strings.ToUpper(toAccount.FullName), "||", amount, "Dhs")
			fmt.Println(separator)
		} else {
			fmt.Println("Insufficient balance in", fromAccount.AccountID)
		}
	} else {
		fmt.Println("Unable to deposit negative amount!!")
	}

	return nil
}
