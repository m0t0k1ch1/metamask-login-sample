package model

import (
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	MyTypedDataType = "string"
	MyTypedDataName = "message"
)

type MyTypedData struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func NewMyTypedData(value string) *MyTypedData {
	return &MyTypedData{
		Type:  MyTypedDataType,
		Name:  MyTypedDataName,
		Value: value,
	}
}

func (data MyTypedData) SignatureHashBytes() []byte {
	return crypto.Keccak256(
		crypto.Keccak256([]byte(data.Type+" "+data.Name)),
		crypto.Keccak256([]byte(data.Value)),
	)
}
