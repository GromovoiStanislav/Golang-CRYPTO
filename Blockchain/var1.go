package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "time"
)

type Block struct {
    Index        int
    PreviousHash string
    Timestamp    int64
    Data         string
    Hash         string
}

func calculateHash(block Block) string {
    record := string(block.Index) + block.PreviousHash + string(block.Timestamp) + block.Data
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}

func createBlock(previousBlock Block, data string) Block {
    var newBlock Block

    newBlock.Index = previousBlock.Index + 1
    newBlock.PreviousHash = previousBlock.Hash
    newBlock.Timestamp = time.Now().Unix()
    newBlock.Data = data
    newBlock.Hash = calculateHash(newBlock)

    return newBlock
}

func isBlockValid(newBlock, previousBlock Block) bool {
    if previousBlock.Index+1 != newBlock.Index {
        return false
    }

    if previousBlock.Hash != newBlock.PreviousHash {
        return false
    }

    if calculateHash(newBlock) != newBlock.Hash {
        return false
    }

    return true
}

func getLastBlock(blockchain []Block) Block {
    return blockchain[len(blockchain)-1]
}



func main() {
    genesisBlock := Block{0, "", time.Now().Unix(), "Genesis Block", ""}
    blockchain := []Block{genesisBlock}

    fmt.Println("Genesis Block:")
    fmt.Printf("Index: %d\nPrevious Hash: %s\nTimestamp: %d\nData: %s\nHash: %s\n\n",
        genesisBlock.Index, genesisBlock.PreviousHash, genesisBlock.Timestamp, genesisBlock.Data, genesisBlock.Hash)

    numBlocksToAdd := 5

    for i := 0; i < numBlocksToAdd; i++ {
        data := fmt.Sprintf("This is Block %d", i+1)
        previousBlock := getLastBlock(blockchain)
        newBlock := createBlock(previousBlock, data)

        if isBlockValid(newBlock, previousBlock) {
            blockchain = append(blockchain, newBlock)
            fmt.Printf("Block #%d has been added to the blockchain!\n", newBlock.Index)
            fmt.Printf("PreviousHash: %s\n", newBlock.PreviousHash)
            fmt.Printf("Hash: %s\n\n", newBlock.Hash)
        } else {
            fmt.Println("Block is not valid and will not be added to the blockchain.")
        }
    }

    //fmt.Println("Blockchain all:",blockchain)
    fmt.Println("Last Block:",getLastBlock(blockchain))
}
