package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"time"
	"github.com/gobuffalo/validate/validators"
)
type Transaction struct {
    ID uuid.UUID `json:"id" db:"id"`
		Hide bool `json:"hide" db:"hide"`
    BlockNumber int `json:"block_number" db:"block_number"`
    Timestamp time.Time `json:"timestamp" db:"timestamp"`
    Hash string `json:"hash" db:"hash"`
    Nonce int `json:"nonce" db:"nonce"`
    BlockHash string `json:"block_hash" db:"block_hash"`
    FormAddr string `json:"form_addr" db:"form_addr"`
    ToAddr string `json:"to_addr" db:"to_addr"`
    Value string `json:"value" db:"value"`
    Gas string `json:"gas" db:"gas"`
    GasPrice string `json:"gas_price" db:"gas_price"`
    IsError int `json:"is_error" db:"is_error"`
    Input string `json:"input" db:"input"`
		Message string `json:"message" db:"message"`
    CumulativeGasUsed string `json:"cumulative_gas_used" db:"cumulative_gas_used"`
    GasUsed string `json:"gas_used" db:"gas_used"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (t Transaction) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Transactions is not required by pop and may be deleted
type Transactions []Transaction

// String is not required by pop and may be deleted
func (t Transactions) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Transaction) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.TimeIsPresent{Field: t.Timestamp, Name: "Timestamp"},
		&validators.StringIsPresent{Field: t.Hash, Name: "Hash"},
		&validators.StringIsPresent{Field: t.FormAddr, Name: "FormAddr"},
		&validators.StringIsPresent{Field: t.ToAddr, Name: "ToAddr"},
		&validators.StringIsPresent{Field: t.Value, Name: "Value"},
		&validators.StringIsPresent{Field: t.Gas, Name: "Gas"},
		&validators.StringIsPresent{Field: t.GasPrice, Name: "GasPrice"},
		&validators.StringIsPresent{Field: t.Input, Name: "Input"},
		&validators.StringIsPresent{Field: t.CumulativeGasUsed, Name: "CumulativeGasUsed"},
		&validators.StringIsPresent{Field: t.GasUsed, Name: "GasUsed"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Transaction) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Transaction) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}


// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Transaction) InputString() string {
	result, _ := hexutil.Decode(t.Input)
	return string(result)
}
