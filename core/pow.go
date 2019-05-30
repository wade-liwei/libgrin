// Copyright 2019 BlockCypher
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import (
	"github.com/blockcypher/libgrin/core/consensus"
	"github.com/blockcypher/libgrin/core/pow"
)

const maxSols uint32 = 10

func createPoWContext(chainType consensus.ChainType, height uint64, edgeBits uint8, proofSize int, nonces []uint64, maxSols uint32) pow.PowContext {
	switch chainType {
	case consensus.Mainnet:
		// Mainnet has Cuckaroo29 for AR and Cuckatoo30+ for AF
		if edgeBits == 29 {
			return pow.NewCuckarooCtx(edgeBits, proofSize, consensus.ChainTypeProofSize(chainType))
		}
		return pow.NewCuckatooCtx(edgeBits, proofSize, consensus.ChainTypeProofSize(chainType), maxSols)
	case consensus.Floonet:
		// Same for Floonet
		if edgeBits == 29 {
			return pow.NewCuckarooCtx(edgeBits, proofSize, consensus.ChainTypeProofSize(chainType))
		}
		return pow.NewCuckatooCtx(edgeBits, proofSize, consensus.ChainTypeProofSize(chainType), maxSols)
	default:
		// Everything else is Cuckatoo only
		return pow.NewCuckatooCtx(edgeBits, proofSize, consensus.ChainTypeProofSize(chainType), maxSols)
	}
}

// VerifySize validates the proof of work of a given header, and that the proof of work
// satisfies the requirements of the header.
func VerifySize(chainType consensus.ChainType, prePoW []uint8, bh *BlockHeader) error {
	ctx := createPoWContext(chainType, bh.Height, bh.PoW.EdgeBits(), len(bh.PoW.Proof.Nonces), bh.PoW.Proof.Nonces, maxSols)
	ctx.SetHeaderNonce(prePoW, nil)
	if err := ctx.Verify(bh.PoW.Proof); err != nil {
		return err
	}
	return nil
}
