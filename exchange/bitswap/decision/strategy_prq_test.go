package decision

//import (
	//"fmt"
	//"testing"

	//"github.com/ipfs/go-ipfs/exchange/bitswap/wantlist"
	//cid "gx/ipfs/QmNp85zy9RLrQ5oQD4hPyS39ezrrXpcaa7R4Y9kxdWQLLQ/go-cid"
	//u "gx/ipfs/QmSU6eubNdhXjFBJBSksTp8kv8YRub8mGAPv8tVJHmL2EU/go-ipfs-util"
	//"gx/ipfs/QmWRCn8vruNAzHx8i6SAXinuheRitKEGu8c7m26stKvsYx/go-testutil"
	//peer "gx/ipfs/QmXYjuNuxVzXKJCfWasQk1RqkhVLDM9jtUKhqc2WPQmFSB/go-libp2p-peer"
//)

//const peerBurst = 5
//const roundBurst = 50

//func testingPRQ() *strategy_prq {
	//prq := newStrategyPRQ(ReciprocationStrategy)
	//prq.rrq = newRRQueueCustom(peerBurst, roundBurst, ReciprocationStrategy)
	//return prq
//}

//func TestEmptyAllocation(t *testing.T) {
	//prq := testingPRQ()
	//numPeers := 5
	//for i := 0; i < numPeers; i++ {
		//partner: random peer ID
		//partner := testutil.RandPeerIDFatal(t)
		//new CID based on i
		//elcid := cid.NewCidV0(u.Hash([]byte(fmt.Sprint(i))))
		//get receipt for this partner, set its value to 0
		//receipt := newLedger(partner).Receipt()
		//receipt.Value = float64(0)
		//push entry to wantlist. Size doesn't matter
		//prq.Push(&wantlist.Entry{Cid: elcid}, partner, receipt)
	//}

	//prq.initRound()

	//all weights should be 0, so allocations should be empty
	//numAllocs := len((*prq.rrq).allocations)
	//if !(len((*prq.rrq).allocations) == 0) {
		//t.Fatalf("Num peers allocated was %v (should be 0)\n", numAllocs)
	//}
//}

//func TestAllocations(t *testing.T) {
	//prq := testingPRQ()
	//expected := make(map[peer.ID]int)
	//numPeers := 5
	//for i := 0; i < numPeers; i++ {
		//partner: random peer ID
		//partner := testutil.RandPeerIDFatal(t)
		//new CID based on i
		//elcid := cid.NewCidV0(u.Hash([]byte(fmt.Sprint(i))))
		//get receipt for this partner, set its value to i
		//receipt := newLedger(partner).Receipt()
		//receipt.Value = float64(i)
		//push entry to wantlist. Size doesn't matter
		//prq.Push(&wantlist.Entry{Cid: elcid}, partner, receipt)
		//create the expected results on-the-fly, since order is nondeterministic.
		//peers with a receipt.Value of 0 are not allocated
		//if receipt.Value != 0 {
			//expected[partner] = i * 5
		//}
	//}

	//prq.initRound()

	//if len((*prq.rrq).allocations) != len(expected) {
		//t.Fatalf("%d peers should have been allocated for the round.", len(expected))
	//}
	//if !compareAllocations((*prq.rrq).allocations, expected) {
		//t.Fatal("Incorrect round robin allocation.\n")
	//}
//}

//func badAllocations(actual, expected []*RRPeer) {
	//TODO
//}

//func badAllocations(actual []*RRPeer, expected map[peer.ID]int) string {
	//actualIDs := make([]peer.ID, len(actual))
	//expectedIDs := make([]peer.ID, len(expected))

	//message := "Incorrect round robin allocation.\n"
	//message += "Actual:   [ "
	//for i, alloc := range actual {
		//message += fmt.Sprintf("%d ", alloc.allocation)
		//actualIDs[i] = alloc.id
	//}
	//message += fmt.Sprintf("]\nExpected: [ ")
	//for _, alloc := range actual {
		//message += fmt.Sprintf("%d ", expected[alloc.id])
	//}
	//i := 0
	//for id, _ := range expected {
		//expectedIDs[i] = id
		//i++
	//}
	//for _, expectedID := range expectedIDs {
		//found := false
		//for _, actualID := range actualIDs {
			//if expectedID == actualID {
				//found = true
				//break
			//}
		//}
		//if !found {
			//message += fmt.Sprintf("%d ", expected[expectedID])
		//}
	//}
	//message += fmt.Sprintf("]\n")
	//return message
//}

//func compareAllocations(actual []*RRPeer, expected map[peer.ID]int) bool {
	//if len(actual) != len(expected) {
		//return false
	//}
	//for _, alloc := range actual {
		//if alloc.allocation != expected[alloc.id] {
			//return false
		//}
	//}
	//return true
//}
