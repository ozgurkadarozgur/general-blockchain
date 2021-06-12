package blockchain

import (
	"bytes"
	"encoding/json"
)

type Block struct {
	Index        uint64
	Timestamp    int64
	Proof        uint64
	PreviousHash string
}

func (block *Block) Hash() string {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(block)
	serialized := string(buffer.Bytes())
	return EncodeSha256([]byte(serialized))
}
