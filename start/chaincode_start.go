package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	//"github.com/hyperledger/fabric/core/crypto/primitives"
)


type EMP struct {

}


type EmpDetail struct{	
	EmpId string `json:"empId"`
	Desgn string `json:"desg"`
	Title string `json:"title"`
	FirstName string `json:"firstName"`
	
	}
	
	







func (t *EMP) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Check if table already exists
	_, err := stub.GetTable("EmpTable")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("EmpTable", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "empId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "desgn", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "title", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "firstName", Type: shim.ColumnDefinition_STRING, Key: false},
		
	})
	if err != nil {
		return nil, errors.New("Failed creating EmpTable.")
	}
	
	return nil, nil
}





 func (t *EMP) getEmp(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting applicationid to query")
	}

	empId := args[0]
	

	// Get the row pertaining to this applicationId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: empId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("EmpTable", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + empId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + empId + "\"}"
		return nil, errors.New(jsonResp)
	}
	
	
	res2E := EmpDetail{}
	
	res2E.EmpId = row.Columns[0].GetString_()
	res2E.Desgn = row.Columns[1].GetString_()
	res2E.Title = row.Columns[2].GetString_()
	res2E.FirstName = row.Columns[3].GetString_()
	
	mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil
	
}




// Invoke invokes the chaincode
func (t *EMP) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "EmpSubmit" {
		if len(args) != 20 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 20. Got: %d.", len(args))
		}

		empId := args[0]
		desgn := args[1]
		title := args[2]
		firstName := args[3]
		

		// Insert a row
		ok, err := stub.InsertRow("EmpTable", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: empId}},
				&shim.Column{Value: &shim.Column_String_{String_: desgn}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}

		return nil, err
	} 

	return nil, errors.New("Invalid invoke function name.")

}



func (t *EMP) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "getEmp" {
		if len(args) != 1 {
			return nil, errors.New("Incorrect number of arguments. Expecting applicationid to query")
		}
		t := EMP{}
		return t.getEmp(stub, args)		
	}
	
	return nil, nil
}




func main() {
	
	err := shim.Start(new(EMP))
	if err != nil {
		fmt.Printf("Error starting HDFC: %s", err)
	}
} 








	
	