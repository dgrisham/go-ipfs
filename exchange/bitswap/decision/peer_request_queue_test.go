package decision

import (
	"fmt"
	"testing"

	"github.com/ipfs/go-ipfs/exchange/bitswap/wantlist"
	u "gx/ipfs/QmPsAfmDBnZN3kZGSuNwvCNDZiHneERSKmRcFyG3UkvcT3/go-ipfs-util"
	"gx/ipfs/QmeDA8gNhvRTsbrjEieay5wezupJDiky8xvCzDABbsGzmp/go-testutil"
	cid "gx/ipfs/QmeSrf6pzut73u6zLQkRFQ3ygt3k6XFT2kjdYP8Tnkwwyg/go-cid"
)

func TestAllocations(t *testing.T) {
	prq := newPRQ(ReciprocationStrategy)
	blockSize := 10
	for i := 1; i <= 2; i++ {
		partner := testutil.RandPeerIDFatal(t)
		elcid := cid.NewCidV0(u.Hash([]byte(fmt.Sprint(i))))
		receipt := ReceiptFromLedger(newLedger(partner))
		receipt.Value = float64(i)
		prq.Push(&wantlist.Entry{Cid: elcid}, partner, receipt, blockSize)
	}

	prq.initRound()

	for i, alloc := range (*prq.rrq).allocations {
		fmt.Printf("peer %d, allocation: %d\n", i, alloc.allocation)
	}
}

func TestPop(t *testing.T) {
	prq := newPRQ(ReciprocationStrategy)

	for i := 0; i < 5; i++ {
		elcid := cid.NewCidV0(u.Hash([]byte(fmt.Sprint(i))))
		prq.Push(&wantlist.Entry{Cid: elcid}, a)
		prq.Push(&wantlist.Entry{Cid: elcid}, b)
		prq.Push(&wantlist.Entry{Cid: elcid}, c)
		prq.Push(&wantlist.Entry{Cid: elcid}, d)
	}

	blockSize := 10
	for i, peer := range peers {
		for j := 0; j < 100; j++ {
			elcid := cid.NewCidV0(u.Hash([]byte(fmt.Sprint(j))))
			prq.Push(&wantlist.Entry{Cid: elcid}, peer, receipts[i], blockSize)
		}
	}

	prq.initRound()
	j := 0
	for len((*prq.rrq).allocations) > 0 {
		fmt.Printf("Round %d\n-------\n", j)
		d := len((*prq.rrq).allocations)
		for i, alloc := range (*prq.rrq).allocations {
			fmt.Printf("peer %d, allocation: %d\n", (i+(j/d))%d, alloc.allocation)
		}
		prq.Pop()
		fmt.Println()
		j++
	}
	for i, alloc := range (*prq.rrq).allocations {
		fmt.Printf("peer %d, allocation: %d\n", i, alloc.allocation)
	}
}
