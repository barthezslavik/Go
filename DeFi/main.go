package main

import (
	"fmt"
	"math/big"
)

// Account represents a DeFi account
type Account struct {
	Address string
	Balance *big.Int
}

// Transfer transfers funds from one account to another
func (a *Account) Transfer(to *Account, amount *big.Int) error {
	if a.Balance.Cmp(amount) < 0 {
		return fmt.Errorf("insufficient balance")
	}

	a.Balance.Sub(a.Balance, amount)
	to.Balance.Add(to.Balance, amount)
	return nil
}

func main() {
	// Create two accounts
	alice := &Account{
		Address: "alice",
		Balance: big.NewInt(1000),
	}
	bob := &Account{
		Address: "bob",
		Balance: big.NewInt(0),
	}

	// Transfer funds from Alice to Bob
	if err := alice.Transfer(bob, big.NewInt(500)); err != nil {
		fmt.Println("Error transferring funds:", err)
	}

	fmt.Println("Alice's balance:", alice.Balance)
	fmt.Println("Bob's balance:", bob.Balance)
}
