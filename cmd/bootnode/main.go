// Copyright 2015 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

// bootnode runs a bootstrap node for the TrueChain Discovery Protocol.
package main

import (
	"crypto/ecdsa"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taiyuechain/taiyuechain/cmd/utils"
	"github.com/taiyuechain/taiyuechain/crypto/gm/sm2"

	//"github.com/taiyuechain/taiyuechain/crypto"
	"github.com/taiyuechain/taiyuechain/crypto/taiCrypto"
	"github.com/taiyuechain/taiyuechain/p2p/discover"
	"github.com/taiyuechain/taiyuechain/p2p/discv5"
	"github.com/taiyuechain/taiyuechain/p2p/enode"
	"github.com/taiyuechain/taiyuechain/p2p/nat"
	"github.com/taiyuechain/taiyuechain/p2p/netutil"
	"net"
	"os"
)

func main() {
	var taiprivate taiCrypto.TaiPrivateKey
	var taipublic taiCrypto.TaiPublicKey
	var (
		listenAddr  = flag.String("addr", ":30301", "listen address")
		genKey      = flag.String("genkey", "", "generate a node key")
		writeAddr   = flag.Bool("writeaddress", false, "write out the node's pubkey hash and quit")
		nodeKeyFile = flag.String("nodekey", "", "private key filename")
		nodeKeyHex  = flag.String("nodekeyhex", "", "private key as hex (for testing)")
		natdesc     = flag.String("nat", "none", "port mapping mechanism (any|none|upnp|pmp|extip:<IP>)")
		netrestrict = flag.String("netrestrict", "", "restrict network communication to the given IP networks (CIDR masks)")
		runv5       = flag.Bool("v5", false, "run a v5 topic discovery bootnode")
		verbosity   = flag.Int("verbosity", int(log.LvlInfo), "log verbosity (0-9)")
		vmodule     = flag.String("vmodule", "", "log verbosity pattern")
		//caoliang modify
		//nodeKey *ecdsa.PrivateKey
		nodeKey   *ecdsa.PrivateKey
		smnodeKey *sm2.PrivateKey
		err       error
	)
	flag.Parse()
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(*verbosity))
	glogger.Vmodule(*vmodule)
	log.Root().SetHandler(glogger)

	natm, err := nat.Parse(*natdesc)
	if err != nil {
		utils.Fatalf("-nat: %v", err)
	}
	switch {
	case *genKey != "":
		//caoliang modify
		//nodeKey, err = crypto.GenerateKey()
		nodekeytype, err := taiCrypto.GenPrivKey()
		if taiCrypto.AsymmetricCryptoType == taiCrypto.ASYMMETRICCRYPTOECDSA {
			nodeKey = &nodekeytype.Private
			taiprivate.Private = *nodeKey
		}
		if taiCrypto.AsymmetricCryptoType == taiCrypto.ASYMMETRICCRYPTOSM2 {
			smnodeKey = &nodekeytype.GmPrivate
			taiprivate.GmPrivate = *smnodeKey
		}
		if err != nil {
			utils.Fatalf("could not generate key: %v", err)
		}
		//caolaing modify
		//if err = crypto.SaveECDSA(*genKey, nodeKey); err != nil {

		if err = taiprivate.SaveECDSA(*genKey, taiprivate); err != nil {
			utils.Fatalf("%v", err)
		}
		return

	case *nodeKeyFile == "" && *nodeKeyHex == "":
		utils.Fatalf("Use -nodekey or -nodekeyhex to specify a private key")
	case *nodeKeyFile != "" && *nodeKeyHex != "":
		utils.Fatalf("Options -nodekey and -nodekeyhex are mutually exclusive")
	case *nodeKeyFile != "":
		//caolaing modify
		//if nodeKey, err = crypto.LoadECDSA(*nodeKeyFile); err != nil {
		taiprivate, err := taiprivate.LoadECDSA(*nodeKeyFile)
		if err != nil {
			utils.Fatalf("-nodekey: %v", err)
		}
		nodeKey = &taiprivate.Private
	/*	if nodeKey, err = taiprivate.LoadECDSA(*nodeKeyFile); err != nil {
		utils.Fatalf("-nodekey: %v", err)
	}*/
	case *nodeKeyHex != "":
		//caoliang modify
		//if nodeKey, err = crypto.HexToECDSA(*nodeKeyHex); err != nil {
		taiprivate, err := taiprivate.HexToECDSA(*nodeKeyHex)
		if err != nil {
			utils.Fatalf("-nodekeyhex: %v", err)
		}
		nodeKey = &taiprivate.Private
	}

	if *writeAddr {
		//caoliang modify
		//fmt.Printf("%x\n", crypto.FromECDSAPub(&nodeKey.PublicKey)[1:])
		taipublic.Publickey = nodeKey.PublicKey
		fmt.Printf("%x\n", taipublic.FromECDSAPub(taipublic)[1:])
		//fmt.Printf("%v\n", discover.PubkeyID(&nodeKey.PublicKey))
		os.Exit(0)
	}

	var restrictList *netutil.Netlist
	if *netrestrict != "" {
		restrictList, err = netutil.ParseNetlist(*netrestrict)
		if err != nil {
			utils.Fatalf("-netrestrict: %v", err)
		}
	}

	addr, err := net.ResolveUDPAddr("udp", *listenAddr)
	if err != nil {
		utils.Fatalf("-ResolveUDPAddr: %v", err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		utils.Fatalf("-ListenUDP: %v", err)
	}

	realaddr := conn.LocalAddr().(*net.UDPAddr)
	if natm != nil {
		if !realaddr.IP.IsLoopback() {
			go nat.Map(natm, nil, "udp", realaddr.Port, realaddr.Port, "truechain discovery")
		}
		// TODO: react to external IP changes over time.
		if ext, err := natm.ExternalIP(); err == nil {
			realaddr = &net.UDPAddr{IP: ext, Port: realaddr.Port}
		}
	}

	if *runv5 {
		if _, err := discv5.ListenUDP(&taiprivate, conn, "", restrictList); err != nil {
			utils.Fatalf("%v", err)
		}
	} else {
		db, _ := enode.OpenDB("")
		ln := enode.NewLocalNode(db, &taiprivate)
		cfg := discover.Config{
			PrivateKey:  &taiprivate,
			NetRestrict: restrictList,
		}
		if _, err := discover.ListenUDP(conn, ln, cfg); err != nil {
			utils.Fatalf("%v", err)
		}
	}

	select {}
}
