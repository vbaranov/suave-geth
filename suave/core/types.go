package suave

import (
	"context"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type Bytes = hexutil.Bytes
type BidId = [16]byte

type Bid = types.Bid

var ConfStoreAllowedAny common.Address = common.HexToAddress("0x42")

type ConfidentialStoreBackend interface {
	Initialize(bid Bid, key string, value []byte) (Bid, error)
	Store(bidId BidId, caller common.Address, key string, value []byte) (Bid, error)
	Retrieve(bid BidId, caller common.Address, key string) ([]byte, error)
}

type MempoolBackend interface {
	SubmitBid(Bid) error
	FetchBidById(BidId) (Bid, error)
	FetchBidsByProtocolAndBlock(blockNumber uint64, namespace string) []Bid
}

type ConfidentialEthBackend interface {
	BuildEthBlock(ctx context.Context, args *BuildBlockArgs, txs types.Transactions) (*engine.ExecutionPayloadEnvelope, error)
	BuildEthBlockFromBundles(ctx context.Context, args *BuildBlockArgs, bundles []types.SBundle) (*engine.ExecutionPayloadEnvelope, error)
}

type BuildBlockArgs = types.BuildBlockArgs
