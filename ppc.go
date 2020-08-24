// Copyright (c) 2014-2014 PPCD developers.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package ppcutil

import (
	"bytes"
	"github.com/btcsuite/btclog"
	"github.com/ppcsuite/ppcd/wire"
	"time"
)

// TxOffsetUnknown is the value returned for a transaction offset that is unknown.
// This is typically because the transaction has not been inserted into a block
// yet.
const TxOffsetUnknown = uint32(0)

// KernelStakeModifierUnknown is the value returned for a block kernel stake
// modifier that is unknown.
// This is typically because the block has not been used for minting yet.
const KernelStakeModifierUnknown = uint64(0)

// Offset returns the saved offset of the transaction within a block.  This value
// will be TxOffsetUnknown if it hasn't already explicitly been set.
func (t *Tx) Offset() uint32 {
	return t.txOffset
}

// SetOffset sets the offset of the transaction in within a block.
func (t *Tx) SetOffset(offset uint32) {
	t.txOffset = offset
}

func (b *Block) Meta() *wire.Meta {
	if b.meta != nil {
		return b.meta
	}
	b.meta = new(wire.Meta)
	return b.meta
}

// MetaToBytes serializes block meta data to byte array
func (b *Block) MetaToBytes() ([]byte, error) {
	// Return the cached serialized bytes if it has already been generated.
	if len(b.serializedMeta) != 0 {
		return b.serializedMeta, nil
	}

	// Serialize the Meta.
	var w bytes.Buffer
	err := b.Meta().Serialize(&w)
	if err != nil {
		return nil, err
	}
	serializedMeta := w.Bytes()

	// Cache the serialized bytes and return them.
	b.serializedMeta = serializedMeta
	return serializedMeta, nil
}

// MetaFromBytes deserializes block meta data from byte array
func (b *Block) MetaFromBytes(serializedMeta []byte) error {
	mr := bytes.NewReader(serializedMeta)
	err := b.Meta().Deserialize(mr)
	if err != nil {
		return err
	}
	b.serializedMeta = serializedMeta
	return nil
}

// NewBlock returns a new instance of a peercoin block given an underlying
// wire.MsgBlock.  See Block.
func NewBlockWithMetas(msgBlock *wire.MsgBlock, meta *wire.Meta) *Block {
	return &Block{
		msgBlock:    msgBlock,
		blockHeight: BlockHeightUnknown,
		meta:        meta,
	}
}

// https://github.com/ppcoin/ppcoin/blob/v0.4.0ppc/src/main.h#L962
// ppc: two types of block: proof-of-work or proof-of-stake
func (block *Block) IsProofOfStake() bool {
	return block.msgBlock.IsProofOfStake()
}

func Now() time.Time {
	return time.Now()
}

func TimeTrack(log btclog.Logger, start time.Time, name string) {
	elapsed := time.Since(start)
	log.Tracef("%s took %s", name, elapsed)
}

func Slice(args ...interface{}) []interface{} {
	return args
}
