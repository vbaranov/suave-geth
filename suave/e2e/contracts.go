package e2e

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	MevShareBidContract       = newArtifact("bids.sol/MevShareBidContract.json")
	BundleBidContract         = newArtifact("bids.sol/BundleBidContract.json")
	buildEthBlockContract     = newArtifact("bids.sol/EthBlockBidContract.json")
	ethBlockBidSenderContract = newArtifact("bids.sol/EthBlockBidSenderContract.json")
	suaveLibContract          = newArtifact("SuaveAbi.sol/SuaveAbi.json")
)

func newArtifact(name string) *Artifact {
	// Get the caller's file path.
	_, filename, _, _ := runtime.Caller(1)

	// Resolve the directory of the caller's file.
	callerDir := filepath.Dir(filename)

	// Construct the absolute path to the target file.
	targetFilePath := filepath.Join(callerDir, "../artifacts", name)

	data, err := os.ReadFile(targetFilePath)
	if err != nil {
		panic(fmt.Sprintf("failed to read artifact %s: %v", name, err))
	}

	var artifactObj struct {
		Abi              *abi.ABI `json:"abi"`
		DeployedBytecode string   `json:"deployedBytecode"`
		Bytecode         string   `json:"bytecode"`
	}
	if err := json.Unmarshal(data, &artifactObj); err != nil {
		panic(fmt.Sprintf("failed to unmarshal artifact %s: %v", name, err))
	}

	return &Artifact{
		Abi:          artifactObj.Abi,
		Code:         hexutil.MustDecode(artifactObj.Bytecode),
		DeployedCode: hexutil.MustDecode(artifactObj.DeployedBytecode),
	}
}

type Artifact struct {
	Abi          *abi.ABI
	DeployedCode []byte
	Code         []byte
}
