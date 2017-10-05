package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var logger = shim.NewLogger("mylogger")

type AgroBizChaincode struct {
}

func (t *AgroBizChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

func (t *AgroBizChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "GetPurchaseOrder" {
		return GetPurchaseOrder(stub, args)
	}
	return nil, nil
}

func (t *AgroBizChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "CreatePO" {
		username, _ := GetCertAttribute(stub, "username")
		role, _ := GetCertAttribute(stub, "role")
		if role == "Customer" {
			return CreatePO(stub, args)
		} else {
			return nil, errors.New(username + " with role " + role + " does not have access to create a purchase order")
		}

	}
	return nil, nil
}

func main() {

	lld, _ := shim.LogLevel("DEBUG")
	fmt.Println(lld)

	logger.SetLevel(lld)
	fmt.Println(logger.IsEnabledFor(lld))

	err := shim.Start(new(AgroBizChaincode))
	if err != nil {
		logger.Error("Could not start AgroBizChainCode")
	} else {
		logger.Info("AgroBizChainCode successfully started")
	}

}
