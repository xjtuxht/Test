package main

import (
	"log"

	"test/cc"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	Chaincode, err := contractapi.NewChaincode(&cc.SmartContract{})
	if err != nil {
		log.Panicf("Error creating gmbf chaincode: %v", err)
	}

	if err := Chaincode.Start(); err != nil {
		log.Panicf("Error starting gmbf chaincode: %v", err)
	}
}
