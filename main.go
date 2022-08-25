package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/brianium/mnemonic"
	"github.com/wemeetagain/go-hdwallet"
)

// CreateHash returns random 8 byte array as a hex string
func createHash() string {
	key := make([]byte, 8)

	_, err := rand.Read(key)
	if err != nil {
		// handle error here
		fmt.Println(err)
	}

	str := hex.EncodeToString(key)

	return str
}

func main() {

	// Generate a random 256 bit seed
	seed := createHash()
	mnemonic, _ := mnemonic.New([]byte(seed), mnemonic.English)

	// Create a master private key
	masterprv := hdwallet.MasterKey([]byte(mnemonic.Sentence()))

	// Convert a private key to public key
	masterpub := masterprv.Pub()

	// Get your address
	address := masterpub.Address()

	fmt.Println("Wallet")
	fmt.Printf("Address: %v \n", address)
	fmt.Printf("Private Key: %v \n", masterprv)
	fmt.Printf("Public Key: %v \n", masterpub)
	fmt.Printf("Seed: %v \n", mnemonic.Sentence())

}
