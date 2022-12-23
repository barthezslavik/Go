package main

import (
	"fmt"
	"math/big"
)

// Market represents a predictive market
type Market struct {
	Event       string
	Outcomes    []string
	Contracts   map[string]*Contract
	MarketPrice *big.Int
}

// Contract represents a contract in a predictive market
type Contract struct {
	Owner   string
	Outcome string
	Price   *big.Int
	Shares  *big.Int
}

// Buy buys a contract in the market
func (m *Market) Buy(c *Contract) error {
	// Check if the contract is available
	if _, ok := m.Contracts[c.Outcome]; !ok {
		return fmt.Errorf("contract not available")
	}

	// Check if the buyer has enough funds
	if m.MarketPrice.Cmp(c.Price) < 0 {
		return fmt.Errorf("insufficient funds")
	}

	// Update the market price
	m.MarketPrice.Sub(m.MarketPrice, c.Price)

	// Update the contract
	contract := m.Contracts[c.Outcome]
	contract.Owner = c.Owner
	contract.Price = c.Price
	contract.Shares.Add(contract.Shares, c.Shares)

	return nil
}

func main() {
	// Create a new market
	market := &Market{
		Event:    "Will it rain tomorrow?",
		Outcomes: []string{"Yes", "No"},
		Contracts: map[string]*Contract{
			"Yes": &Contract{
				Owner:   "",
				Outcome: "Yes",
				Price:   big.NewInt(0),
				Shares:  big.NewInt(0),
			},
			"No": &Contract{
				Owner:   "",
				Outcome: "No",
				Price:   big.NewInt(0),
				Shares:  big.NewInt(0),
			},
		},
		MarketPrice: big.NewInt(0),
	}

	contract := &Contract{
		Owner:   "Alice",
		Outcome: "Yes",
		Price:   big.NewInt(100),
		Shares:  big.NewInt(10),
	}

	// Call the Buy function
	if err := market.Buy(contract); err != nil {
		fmt.Println("Error buying contract:", err)
	}
}
