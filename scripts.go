package main

import "encoding/json"

/**Create PO **/
func CreatePO(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Debug("Entering CreatePO")

	if len(args) < 2 {
		logger.Error("Invalid number of args")
		return nil, errors.New("Expected atleast two arguments for Purchase Order creation")
	}

	var poID = args[0]
	var poInput = args[1]

	err := stub.PutState(poID, []byte(poInput))
	if err != nil {
		logger.Error("Could not save Purchase Order to ledger", err)
		return nil, err
	}

	var customEvent = "{eventType: 'CreatePO', description:" + poID + "' Successfully created'}"
	err = stub.SetEvent("evtSender", []byte(customEvent))
	if err != nil {
		return nil, err
	}
	logger.Info("Successfully saved Purchase Order")
	return nil, nil

}

/**
Updates the status of the Purchase Order
**/
func UpdatePOStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Debug("Entering UpdatePOStatus")

	if len(args) < 2 {
		logger.Error("Invalid number of args")
		return nil, errors.New("Expected atleast two arguments for PO Status update")
	}

	var poID = args[0]
	var status = args[1]

	poBytes, err := stub.GetState(poID)
	if err != nil {
		logger.Error("Could not fetch Purchase Order from ledger", err)
		return nil, err
	}
	var PO PurchaseOrder
	err = json.Unmarshal(poBytes, &PO)
	PO.Status = status

	poBytes, err = json.Marshal(&PO)
	if err != nil {
		logger.Error("Could not marshal Purchase Order post update", err)
		return nil, err
	}

	err = stub.PutState(poID, poBytes)
	if err != nil {
		logger.Error("Could not save Purchase Order post update", err)
		return nil, err
	}

	var customEvent = "{eventType: 'POStatusUpdate', description:" + poID + "' Successfully updated status'}"
	err = stub.SetEvent("evtSender", []byte(customEvent))
	if err != nil {
		return nil, err
	}
	logger.Info("Successfully updated Purchase Order Status")
	return nil, nil

}

func GetPurchaseOrder(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Debug("Entering GetPurchaseOrder")

	if len(args) < 1 {
		logger.Error("Invalid number of arguments")
		return nil, errors.New("Missing purchase order ID")
	}

	var poID = args[0]
	bytes, err := stub.GetState(poID)
	if err != nil {
		logger.Error("Could not fetch purchase order with id "+poID+" from ledger", err)
		return nil, err
	}
	return bytes, nil
}

func GetCertAttribute(stub shim.ChaincodeStubInterface, attributeName string) (string, error) {
	logger.Debug("Entering GetCertAttribute")
	attr, err := stub.ReadCertAttribute(attributeName)
	if err != nil {
		return "", errors.New("Couldn't get attribute " + attributeName + ". Error: " + err.Error())
	}
	attrString := string(attr)
	return attrString, nil
}
