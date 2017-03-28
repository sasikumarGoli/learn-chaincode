/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type Emp struct{	
	empId string `json:"empId"`
	name string `json:"name"`
	title string `json:"title"`


}





func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	err := stub.PutState("table_ibminsert", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	
	

	return nil, nil
	}

// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)

	if function == "submitEmp" {
		if len(args) != 3 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 20. Got: %d.", len(args))
		}
		
		empId := args[0]
		name := args[1]
		title := args[2]
		
		
		//insert a row
		
		ok, err := stub.InsertRow("EmpTable", shim.Row{
		Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: empId}},
				&shim.Column{Value: &shim.Column_String_{String_: name}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				}})
	
	if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
	
	}
	fmt.Println("values Inserted in the table: ")
	
	
	
	
	return nil, errors.New("Received unknown function invocation: " + function)
}



// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}





// write - invoke function to write key/value pair
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, value string
	var err error
	fmt.Println("running write()")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	key = args[0] //rename for funsies
	value = args[1]
	err = stub.PutState(key, []byte(value)) //write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// read - query function to read key/value pair
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
	
	
	
	
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting applicationid to query")
	}
	
	fmt.Println("came into read func and geting empid: ")
	
	empId := args[0]



// Get the row pertaining to this applicationId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: empId}}
	columns = append(columns, col1)
	
	
	row, err := stub.GetRow("EmpTable", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the dataaa " + empId + "\"}"
		return nil, errors.New(jsonResp)
	}
	
	fmt.Println("got the row from table: ")
	
	res2E := Emp{}
	
	
	res2E.empId = row.Columns[0].GetString_()
	res2E.name = row.Columns[1].GetString_()
	res2E.title = row.Columns[2].GetString_()
	
	
	mapB, _ := json.Marshal(res2E)
    
	fmt.Println(string(mapB))
	
	return mapB, nil
}
