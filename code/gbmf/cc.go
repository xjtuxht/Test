package main

import (
	"encoding/json"
	"math"
	mrand "math/rand"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

type User struct {
	ID          int
	Location    string
	Gender      string
	Age         string //ciph
	Phonenumber string
}

type GlobalParams struct {
	secureparam int
	rangeparam  int
}

// InitLedger adds a base set of Users to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface, secureparam int, rangeparam int) (string, error) {
	/*users := []User{
		{ID: "1", Location: "China", Gender: "Male", Age: "30"},
		{ID: "2", Location: "China", Gender: "Male", Age: "40"},
		{ID: "3", Location: "China", Gender: "Female", Age: "50"},
		{ID: "4", Location: "Shaanxi", Gender: "Male", Age: "20"},
		{ID: "5", Location: "Shaanxi", Gender: "Male", Age: "25"},
		{ID: "6", Location: "Shaanxi", Gender: "Male", Age: "37"},
	}
	for _, asset := range assets {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}*/
	mkey := strconv.Itoa(Genmkey(secureparam))

	return nil
}

// TransferAsset updates the owner field of asset with given id in world state.
func (s *SmartContract) Genglobalparam(ctx contractapi.TransactionContextInterface, secureparam int, rangeparam int) (string, error) {
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	asset.Owner = newOwner
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

func Genmkey(digit int) int {
	mrand.Seed(time.Now().UnixNano())
	k := float64(digit) //security param
	Range := int(math.Pow(2, k))
	mkey := mrand.Intn(Range)
	return mkey
}
