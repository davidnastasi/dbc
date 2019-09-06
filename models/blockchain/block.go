package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

// Block basic structure
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce	  int
}

// DeriveHash is the join of hash
/*func (b *Block) DeriveHash() {
	// TODO TEST: info := [][]byte{b.Data, b.PrevHash}
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:] // TODO TEST: copy(b.Hash, hash)
}*/

// createBlock create a new Block with the previus hash
func createBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewPOW(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// Genesis return the initial Block
func genesis() *Block {
	return createBlock("Genesis", []byte{})
}

func (block *Block) Serialize() []byte  {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	if err := encoder.Encode(block); err != nil {
		log.Panic(err)
	}
	return res.Bytes()
}

func  Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewBuffer(data))
	if err := decoder.Decode(&block); err != nil {
		log.Panic(err)
	}
	return &block
}




