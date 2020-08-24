// Copyright (c) 2013-2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package ppcutil

const (
	// SatoshiPerPeercent is the number of satoshi in one peercoin cent.
	SatoshiPerBitcent = 1e4 // ppc: 10000 sunny per peercoin cent

	// SatoshiPerpeercoin is the number of satoshi in one peercoin (1 BTC).
	SatoshiPerpeercoin = 1e6 // ppc: 1000000 sunny per peercoin

	// MaxSatoshi is the maximum transaction amount allowed in satoshi.
	MaxSatoshi = 21e6 * SatoshiPerpeercoin // ppc: TODO check value for peercoin
)
