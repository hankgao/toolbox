package main

func main() {
	// Give a wallet
	// Give
}

// Dest represents a target that coins will be transferred to
type Dest struct {
	Address string
	Coins   uint64
}

// Dests represents a batch of transfer targets
type Dests []Dest

// CoinSum returns the sum of coins of all transfer coins
func (dests Dests) CoinSum() uint64 {
	var sum = uint64(0)
	for _, d := range dests {
		sum += d.Coins
	}

	return sum
}

// Wallet represent a Skyledger wallet
type Wallet struct {
	// Head {}
	// Entries []{}
}

// Load loads a wallet from a file
func (w Wallet) Load(wf string) error {
	return nil
}

// Balance returns the balance of a wallet
func (w Wallet) Balance() uint64 {
	balance := uint64(0)
	return balance
}

// HasEnoughCoins checks whether a wallet has enough coins to spend
func HasEnoughCoins(w Wallet, ds Dests) bool {
	return w.Balance() >= ds.CoinSum()
}

// PickInputs picks addresses that have enough coins in total
func (w Wallet) PickInputs(coins2Spend uint64) error {
	return nil
}
