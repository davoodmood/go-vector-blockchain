package blockchain

type BlockChain struct {
	Blocks []*Block // prototyping sufficent
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// func (b *Block) DeriveHash() {
// 	info := bytes.Join([][]byte{b.Data, b.PrevHash, 0}, []byte{})
// 	hash := sha256.Sum256(info) // Temporary Encryption model for prototyping
// 	b.Hash = hash[:]
// }

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockchain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
