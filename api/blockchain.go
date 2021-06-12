package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ozgurkadarozgur/general-blockchain/blockchain"
)

type BlockchainApi interface {
	MineBlock(*gin.Context)
	GetChain(*gin.Context)
	GetBlock(*gin.Context)
	IsValid(*gin.Context)
}

type blockchainApi struct {
	bc blockchain.Blockchain
}

func NewBlockchainApi() BlockchainApi {
	return &blockchainApi{
		bc: blockchain.NewBlockChain(),
	}
}

func (bcApi *blockchainApi) MineBlock(c *gin.Context) {

	bc := bcApi.bc

	previousBlock := bc.GetPreviousBlock()
	previousProof := previousBlock.Proof

	proof := bc.ProofOfWork(previousProof)

	previousHash := previousBlock.Hash()
	createdBlock := bc.CreateBlock(proof, previousHash)

	c.JSON(200, createdBlock)
}

func (bcApi *blockchainApi) GetChain(c *gin.Context) {

	bc := bcApi.bc

	c.JSON(200, bc.GetChain())
}

func (bcApi *blockchainApi) GetBlock(c *gin.Context) {

	index := c.Param("index")

	intIndex, _ := strconv.Atoi(index)

	bc := bcApi.bc

	block := bc.GetChain()[intIndex]
	c.JSON(200, gin.H{
		"block": block,
		"hash":  block.Hash(),
	})
}

func (bcApi *blockchainApi) IsValid(c *gin.Context) {
	c.JSON(200, gin.H{
		"isValid": bcApi.bc.IsValid(),
	})
}
