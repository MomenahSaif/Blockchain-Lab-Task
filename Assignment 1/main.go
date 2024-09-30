package main

import "fmt"

func main() {
    // Creating the first block (genesis block)
    genesisBlock := NewBlock("genesis to network", 1, "0")
    fmt.Println("Genesis Block Created:")
    ListBlocks()

    // Adding more blocks
    NewBlock("bob to alice", 2, genesisBlock.Hash)
    NewBlock("alice to charlie", 3, genesisBlock.Hash)

    fmt.Println("\nAfter Adding More Blocks:")
    ListBlocks()

    // Modifying a block
    fmt.Println("\nModifying the second block:")
    ChangeBlock(1, "bob to dave")
    ListBlocks()

    // Verifying the integrity of the chain
    fmt.Println("\nVerifying Blockchain Integrity:")
    if VerifyChain() {
        fmt.Println("Blockchain is valid.")
    } else {
        fmt.Println("Blockchain is invalid!")
    }
}
