package p2

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/nicholas-kebbas/project2PrivateBlockchain/p1"
	"golang.org/x/crypto/sha3"
)

type Block struct {
	Header Header
	Value Value
}

type Header struct {
	Hash       string `json:"hash"`
	TimeStamp  int64 `json:"timestamp"`
	Height     int32 `json:"height"`
	ParentHash string `json:"parentHash"`
	Size       int32 `json:"size"`
}

type Value struct {
	mpt p1.MerklePatriciaTrie
	StringDb map[string]string `json:"mpt"`
}

/**
JSON Representation of a block
 */
type JsonBlock struct {
	Hash       string            `json:"hash"`
	Timestamp  int64             `json:"timeStamp"`
	Height     int32             `json:"height"`
	ParentHash string            `json:"parentHash"`
	Size       int32             `json:"size"`
	MPT        map[string]string `json:"mpt"`
}

/**
This function takes arguments(such as height, parentHash, and value of MPT type)
and forms a block.
 */
func Initial(height int32, timestamp int64, parentHash string, mpt p1.MerklePatriciaTrie) Block {
	hash := deriveHash(height, timestamp, parentHash, mpt)
	size := len([]byte(fmt.Sprintf("%v", mpt)))
	size32 := int32(size)
	newHeader := Header{hash, timestamp, height, parentHash, size32}
	newValue := Value{mpt, mpt.GetStringDb()}
	newBlock := Block{newHeader, newValue}
	return newBlock

}

/** Note that you have to reconstruct an MPT from the JSON string, and use that MPT as the block's value. **/
func (block *Block) DecodeFromJson(jsonString string) Block {
	bytes := []byte(jsonString)
	b := Block{}
	j := JsonBlock{}
	err := json.Unmarshal(bytes, &j)
	if err != nil {
		fmt.Println("Error")
	}
	/* Need to build the MPT to insert into block */
	m := p1.MerklePatriciaTrie{}
	m.Initial()
	for k, v := range j.MPT {
		m.Insert(k, v)
	}

	b.Header.Size = j.Size
	b.Header.Hash = j.Hash
	b.Header.ParentHash = j.ParentHash
	b.Header.TimeStamp = j.Timestamp
	b.Header.Height = j.Height

	b.Value.mpt = m
	b.Value.StringDb = j.MPT

	return b
}

/**
This function encodes a block instance into a JSON format string
 */
func (block *Block) EncodeToJSON() string {
	/* Need to get key, value pairs from within MPT */
	encodedString, err := json.Marshal(struct{
		Header
		Value
	}{block.Header, block.Value})
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(encodedString)
}

/**
Block’s hash is the SHA3-256 encoded value of this string
(note that you have to follow this specific order):

hash_str := string(b.Header.Height) + string(b.Header.Timestamp)
+ b.Header.ParentHash + b.Value.Root + string(b.Header.Size)

 **/
func deriveHash(height int32, timestamp int64, parentHash string, mpt p1.MerklePatriciaTrie) string {
	Sheight := string(height)
	Stimestamp := string(timestamp)
	root := mpt.GetRoot()
	size := len([]byte(fmt.Sprintf("%v", mpt)))
	Ssize := string(size)
	str := Sheight + Stimestamp + parentHash + root + Ssize
	sum := sha3.Sum256([]byte(str))
	slice := sum[:]
	hashString := hex.EncodeToString(slice)
	return hashString
}
