package main

import (
	//"encoding/json"
	//"errors"
	"fmt"
	//"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)



type SimpleChaincode struct {
}



func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// Create table one
	
	return nil, nil
}




func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

return nil, nil

}


func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {



return nil, nil

}




func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}