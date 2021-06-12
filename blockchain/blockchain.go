package blockchain

import (
	"math"
	"strconv"
	"time"
)

type Blockchain interface {
	CreateBlock(proof uint64, previousHash string) Block
	GetPreviousBlock() Block
	ProofOfWork(previousProof uint64) uint64
	IsValid() bool
	GetChain() []Block
}

type blockchain struct {
	Chain []Block
}

func NewBlockChain() Blockchain {
	bc := &blockchain{
		Chain: []Block{},
	}
	bc.CreateBlock(1, "0")
	return bc
}

func (bc *blockchain) CreateBlock(proof uint64, previousHash string) Block {
	if previousHash == "" {
		previousHash = "0"
	}

	block := Block{
		Index:        uint64(len(bc.Chain) + 1),
		Timestamp:    time.Now().Unix(),
		Proof:        proof,
		PreviousHash: previousHash,
	}

	bc.Chain = append(bc.Chain, block)

	return block
}

func (bc *blockchain) GetPreviousBlock() Block {
	lastIndex := len(bc.Chain) - 1
	return bc.Chain[lastIndex]
}

func (bc *blockchain) ProofOfWork(previousProof uint64) uint64 {
	var newProof uint64
	newProof = 1
	isValidProof := false
	for !isValidProof {
		hashOperationValue := generateProofHash(newProof, previousProof)
		if isValidProofCondition(hashOperationValue) {
			isValidProof = true
		} else {
			newProof += 1
		}
	}
	return newProof

}

func generateProofHash(newProof uint64, previousProof uint64) string {
	result := math.Sqrt(float64(newProof)) - math.Sqrt(float64(previousProof))
	strResult := strconv.FormatUint(uint64(result), 10)
	return EncodeSha256([]byte(strResult))
}

func isValidProofCondition(hash string) bool {
	return hash[:4] == "728b"
}

func (bc *blockchain) IsValid() bool {
	previousBlock := bc.Chain[0]

	for i := 1; i < len(bc.Chain); i++ {
		currentBlock := bc.Chain[i]
		if currentBlock.PreviousHash != previousBlock.Hash() {
			return false
		}
		previousProof := previousBlock.Proof
		currentProof := currentBlock.Proof
		hashOperationValue := generateProofHash(currentProof, previousProof)
		if !isValidProofCondition(hashOperationValue) {
			return false
		}
		previousBlock = currentBlock
	}
	return true
}

func (bc *blockchain) GetChain() []Block {
	return bc.Chain
}
