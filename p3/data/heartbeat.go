package data

import (
	"github.com/nicholas-kebbas/cs686-blockchain-p3-nicholas-kebbas/p1"
	"math/rand"
)

/* Heartbeat is the JSON representation of the data we need to send to the other blockchains
 */
type HeartBeatData struct {
	IfNewBlock  bool   `json:"ifNewBlock"`
	Id          int32  `json:"id"`
	BlockJson   string `json:"blockJson"`
	PeerMapJson string `json:"peerMapJson"`
	Addr        string `json:"addr"`
	Hops        int32  `json:"hops"`
}

func NewHeartBeatData(ifNewBlock bool, id int32, blockJson string, peerMapJson string, addr string) HeartBeatData {
	heartBeatData := HeartBeatData{ifNewBlock,
		id,
		blockJson,
		peerMapJson,
		addr,
		3}
	return heartBeatData

}

/* Create a new instance of HeartBeatData, then decide whether to create a new block and send it to other peers.
*/
func PrepareHeartBeatData(sbc *SyncBlockChain, selfId int32, peerMapJson string, addr string) HeartBeatData {
	heartBeatData := NewHeartBeatData(true, selfId, " ", peerMapJson, addr)

	/* Randomly decide whether to create new block and send to peers. */
	if rand2() == true {
		/* Just get the first one in that array for now since we don't know what to do w/ forks */
		mpt := p1.MerklePatriciaTrie{}
		mpt.Initial()
		mpt.Insert("nick", "kebbas")
		newBlock := sbc.GenBlock(mpt)
		heartBeatData.BlockJson = newBlock.EncodeToJSON()
		return heartBeatData
	}
	return heartBeatData
}

func rand2() bool {
	return rand.Int31()&0x01 == 0
}