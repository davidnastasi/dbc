package main

import (
	"dbc/models/blockchain"
	"fmt"
	"strconv"
)

func main() {
	chain := blockchain.CreateBlockChain()
	chain.AddBlock("First")
	chain.AddBlock("Second")
	chain.AddBlock("Third")

	fmt.Println("-------------------")

	for _,block := range chain.Blocks {
		fmt.Printf("Hash:%x\n",block.Hash)
		fmt.Printf("Prev:%x\n",block.PrevHash)
		fmt.Printf("Data:%s\n",string(block.Data))
		fmt.Printf("Nonce:%d\n",block.Nonce)

		pow := blockchain.NewPOW(block)
		fmt.Printf("POW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
		//fmt.Println(e.Data)


	/*for value := range []int{1,2,3} {


		fmt.Println("<------" ,value,"------>")
		target := big.NewInt(1)
		fmt.Println("target ", target)
		target = target.Lsh(target, uint(64-1))

		fmt.Println("lsh    ",target)
		fmt.Println("lsh sz ",unsafe.Sizeof(target))

		//fmt.Println("lsh ",target)
	}*/
}
