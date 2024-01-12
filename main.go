package main

import (
	"banking-management-system/src"
)

func main() {
	CDM := src.Bank("Crédit De Maroc")
	myAccount := src.Account{FullName: "Ismail Boularbah", AccountID: 20022205}
	anotherAccount := src.Account{FullName: "Dustin Rock", AccountID: 19991204}

	CDM.Info()
	CDM.AddAccount(myAccount)
	CDM.AddAccount(anotherAccount)

	myAccount.Deposit(200)
	myAccount.Withdrawal(100)

	anotherAccount.Deposit(500)
	anotherAccount.Withdrawal(300)

	CDM.Transfer(anotherAccount.AccountID, myAccount.AccountID, 200)

	// CDM.RemoveAccount(anotherAccount.AccountID)
	// CDM.ShowAccounts()
	myAccount.ShowInfo()
	anotherAccount.ShowInfo()
}
