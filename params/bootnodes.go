// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
<<<<<<< HEAD
    "enode://8679de5282f746f632f5743962f498dc84e2e4bafc085d1e6e91f24d99af1bd9afebf2ac3f1780b0b866464f74b9635405c53880a178b2a6172dab80b602dfc1@37.143.13.112:30303",   // RU1
    "enode://ca5829c681a166713c871ee278bca126e10ebddb5e85e3e3df7563d8387a3e42fe7d412cd34e249f59feaa4181a0119cb52670b1eb89b3e5131a116f7eb337c8@5.9.253.133:30303",     // DE1
=======
    "enode://8914b4b9db0783dc3ade8d50d9936958c04d94e6b71bef56043877df23891a48a13a11cd903ae1882b24fe502e99fb38e58b6c398082bd4f09df39800afd7f96@37.143.13.112:30303", // RU1
    "enode://181a30249b92cca560401ce3d255c0db0e4a168e9cdcfb58e8d60b1617e96ac7d1f30825bcfd5defce3e018b07901a6af013c35505c6879c52c8d2e0b75f103d@5.9.253.133:30303",   // DE1
>>>>>>> 4b93c5d02433a14c917ad358b2c888913c90cc66
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
var RinkebyV5Bootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
