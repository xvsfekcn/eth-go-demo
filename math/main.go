package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	conn, err := ethclient.Dial("http://127.0.0.1:8551")
	if err != nil {
		log.Fatalf("Failed to connect the eth client:%s", err)
	}
	defer conn.Close()

	mathInstance, err := NewMath(common.HexToAddress("0xF7b9c47F39986400E9be9863c4821FD915786fEb"), conn)
	if err != nil {
		log.Fatalf("Failed to NewMath:%s", err)
	}

	val, err := mathInstance.Add(nil, big.NewInt(1), big.NewInt(2))
	if err != nil {
		log.Fatalf("can not call method:%s", err)
	}
	log.Printf("Success:%v", val)
}
