// mypackage.go
package assignment01bca

import ("fmt"
"strings"
"crypto/sha256"
"strconv")

// Nothing fancy here just a function to return hash
func calHash(hashme string) string{
	hash := sha256.Sum256([]byte(hashme))
	return fmt.Sprintf("%x", hash)
}


// This struct will present the whole block in the blockchain
type Block struct{
	Transaction_id int
	Current_Hash string
	Nonce int
	data string
	Previous_Hash string
}

// This struct will hold an array of the blocks
// So basically it is our blockchain
type Blockchain struct{
	Myblockchain []*Block
}


// This function will put the values into the blockchain
func NewBlock(Transaction_id int, Current_Hash string, Nonce int, data string, Previous_Hash string) *Block {

	// Once the block is added to the Chain
	// This is just creating the new object of the block
	block := new(Block)
	block.Transaction_id = Transaction_id
	block.Current_Hash = Current_Hash
	block.Nonce = Nonce
	block.data = data
	block.Previous_Hash = Previous_Hash
	return block

//	A method to add new block. To keep things simple, you could provide a
//	sting of your choice as a transaction (e.g., “bob to alice”). Also, use
//	any integer value as a Nonce. The CreateHash() method will provide you the
//	block Hash value.

}

// This function creates the object of the block and add it to the blockchain
func (chain *Blockchain) AddToBlockchain(Transaction_id int, Current_Hash string, Nonce int, data string, Previous_Hash string) *Block{
	// Now, taking the hash of the block
	blockhash :=  strconv.Itoa(Transaction_id) + Current_Hash + strconv.Itoa(Nonce) + data + Previous_Hash
	// Calculate the hash
	calculated_hash := CalculateHash(blockhash)
	// Create the new hash
	block := NewBlock(Transaction_id, calculated_hash, Nonce, data, Previous_Hash)
	// Now adding the block to the chain
	chain.Myblockchain = append(chain.Myblockchain,block)
	return block
}




//	A method to print all the blocks in a nice format showing block data such
//	as transaction, Nonce, previous hash, current block hash
func ListBlocks(chain *Blockchain) {
	// This will iterate the list inside of the blockchain to print the list
	for i := 0; i < len(chain.Myblockchain); i++ {
		fmt.Printf("\n%s Block %d %s", strings.Repeat("=",25) , i , strings.Repeat("=",25) )
		fmt.Printf("\nID: %v", chain.Myblockchain[i].Transaction_id)
		fmt.Printf("\nNonce: %v", chain.Myblockchain[i].Nonce)
		fmt.Printf("\nData: %v", chain.Myblockchain[i].data)
		fmt.Printf("\nPrevious Hash: %v", chain.Myblockchain[i].Previous_Hash)
		fmt.Printf("\nCureent Hash: %v", chain.Myblockchain[i].Current_Hash)
	}
}

// just calculate the hash and return the string
func CalculateHash (stringToHash string) string  {
	//	function for calculating hash of a block
	hash := calHash(stringToHash)
	//fmt.Printf("\nBlock Hash is %v\n",hash)
	return hash;
	
	}

	
func ChangeBlock(chain *Blockchain , block_no_to_change int) {
//	function to change block transaction of the given block ref
    chain.Myblockchain[block_no_to_change].Transaction_id = block_no_to_change

	// The block will be changed because we are now taking the hash including the currect hash
	blockhash :=  strconv.Itoa(chain.Myblockchain[block_no_to_change].Transaction_id) + chain.Myblockchain[block_no_to_change].Current_Hash + strconv.Itoa(chain.Myblockchain[block_no_to_change].Nonce) + chain.Myblockchain[block_no_to_change].data + chain.Myblockchain[block_no_to_change].Previous_Hash
	// Calculate the new hash
	calculated_hash := CalculateHash(blockhash)
	// Now we need to change the hash of the block
	chain.Myblockchain[block_no_to_change].Current_Hash = calculated_hash
	//fmt.Printf("\n -> " + calculated_hash)
	
}


func VerifyChain(chain *Blockchain) {
	// iternate the whole chain
	for i := 1; i < len(chain.Myblockchain); i++ {
		// check if the previous hash of the block is not equal to to the currect hash of the previous block
		if(chain.Myblockchain[i].Previous_Hash  !=  chain.Myblockchain[i-1].Current_Hash){
			fmt.Printf("\n%s[-] HASH MISMATCH DETECTED [-] %s", strings.Repeat("=",25)  , strings.Repeat("=",25) )

			fmt.Printf("\nThe Block Number " + strconv.Itoa(i) + " needs to be mined")
		}
	}

//	function to verify blockchain in case any changes are made. This can be
//	done in two different ways:
}

