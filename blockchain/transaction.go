package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

// Transaction is a struct made of inputs and outputs
// No valuable/sensitive information is stored inside transactions
type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

// TxOutput is an indivisible struct that contains a value of tokens and a key to unlock these tokens
type TxOutput struct {
	Value  int
	PubKey string
}

// TxInput is a reference to a previous output
// ID references the transaction the ouput is inside
// Out is the index of the output: transaction 'A' could contain 3 outputs. An Out of 2 means the second one is being referenced
// Sig is the PubKey + account name
type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

// SetID encodes a transaction object, hashes the encoded bytes, and uses this hash as the ID of the transaction
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// CoinbaseTx is the first transaction in the chain
// There is only one input and one output
func CoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", to)
	}

	txin := TxInput{[]byte{}, -1, data}
	txout := TxOutput{100, to}

	tx := Transaction{nil, []TxInput{txin}, []TxOutput{txout}}
	tx.SetID()

	return &tx
}

// IsCoinbase returns true if the transaction is the Coinbase transaction
func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}

// CanUnlock returns true if the account owns the information inside the Output referenced by the current Input
func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

// CanBeUnlocked returns true if the account owns the information inside the Output
func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
