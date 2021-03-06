bitcoind
===========
A Golang client library wrapping the bitcoind JSON RPC API

forked from https://github.com/toorop/go-bitcoind
with below feature updated by Shiming:
1. better error handling



Installation
-----
	$ go get github.com/huahuayu/go-bitcoind


Usage
----

	package main

	import (
		"github.com/huahuayu/go-bitcoind"
		"log"
	)

	const (
		SERVER_HOST        = "You server host"
		SERVER_PORT        = port (int)
		USER               = "user"
		PASSWD             = "passwd"
		USESSL             = false
		WALLET_PASSPHRASE  = "WalletPassphrase"
	)

	func main() {
		bc, err := bitcoind.New(SERVER_HOST, SERVER_PORT, USER, PASSWD, USESSL)
		if err != nil {
			log.Fatalln(err)
		}

		//walletpassphrase
		err = bc.WalletPassphrase(WALLET_PASSPHRASE, 3600)
		log.Println(err)

		// backupwallet
		err = bc.BackupWallet("/tmp/wallet.dat")
		log.Println(err)


		// dumpprivkey
		privKey, err := bc.DumpPrivKey("1KU5DX7jKECLxh1nYhmQ7CahY7GMNMVLP3")
		log.Println(err, privKey)

	}
	
Mores examples in example.go (in examples folder) 

Documentation
-----
Click on the button below to access the full documentation:

[![GoDoc](https://godoc.org/github.com/toorop/go-bitcoind?status.png)](https://godoc.org/github.com/toorop/go-bitcoind)	



Unit tests
----
[![Build Status](https://travis-ci.org/Toorop/go-bitcoind.svg)](https://travis-ci.org/toorop/go-bitcoind)

More than 100 unit tests are made.

To run tests:

	$ go get github.com/onsi/ginkgo/ginkgo
	$ go get github.com/onsi/gomega
	$ ginkgo

	Running Suite: Bitcoind Suite	
	=============================
	Random Seed: 1401120770
	Will run 112 of 112 specs

	•••••••••••••••••••••••••••••••••••
	Ran 112 of 112 Specs in 0.001 seconds
	SUCCESS! -- 112 Passed | 0 Failed | 0 Pending | 0 Skipped PASS

	Ginkgo ran in 10.856335553s
	Test Suite Passed
 



Todo
-----
* GetBlockTemplate
* sendrawtransaction
* signrawtransaction
* submitblock

##### Note on SSL support 

Note on ssl support : bitcoind library doesn't verify the server's certificate chain. That means that it accepts any certificate presented by the server and any host name in that certificate. In this mode, TLS is susceptible to man-in-the-middle attacks.
