

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	//"strconv"
      s "strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}



type tabone struct {

vehicleId string       `json:"VehicleId"`
vehicleMake string     `json:"VehicleMake"`
vehicleModel string    `json:"VehicleModel"`
vehicleColour string   `json:"VehicleColour"`
vehicleReg string      `json:"vehicleReg"`
vehicleOwner string    `json:"VehicleOwner"`
}




// Init create tables for tests
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// Create table one
	err := createTableOne(stub)
	if err != nil {
		return nil, fmt.Errorf("Error creating table one during init. %s", err)
	}

	

	return nil, nil
}


func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	switch function {

	case "insertRowTableOne":
		if len(args) < 3 {
			return nil, errors.New("insertTableOne failed. Must include 3 column values")
		}
     var str [] string
		
		for i := 0; i < len(args); i++ {
		
		str = s.Split(args[i],"-")
		
		
		col1Val := str[0]
		//col2Int, err := strconv.ParseInt(str[1], 10, 32)
		col2Val :=str[1]
		col3Val := str[2]
		col4Val := str[3]
		col5Val := str[4]
		col6Val := str[5]
		
		
	//	col2Val := int32(col2Int)
	//	col3Int, err := strconv.ParseInt(str[2], 10, 32)
		
		//col3Val := int32(col3Int)

		var columns []*shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: col1Val}}
		col2 := shim.Column{Value: &shim.Column_String_{String_: col2Val}}
		col3 := shim.Column{Value: &shim.Column_String_{String_: col3Val}}
		col4 := shim.Column{Value: &shim.Column_String_{String_: col4Val}}
        col5 := shim.Column{Value: &shim.Column_String_{String_: col5Val}}
        col6 := shim.Column{Value: &shim.Column_String_{String_: col6Val}}  		
		columns = append(columns, &col1)
		columns = append(columns, &col2)
		columns = append(columns, &col3)
		columns = append(columns, &col4)
		columns = append(columns, &col5)
		columns = append(columns, &col6)

		row := shim.Row{Columns: columns}
		ok, err := stub.InsertRow("tableOne", row)
		if err != nil {
			return nil, fmt.Errorf("insertTableOne operation failed. %s", err)
		}
		if !ok {
			return nil, errors.New("insertTableOne operation failed. Row with given key already exists")
		}
       
	   
	   }

	
	

	default:
		return nil, errors.New("Unsupported operation")
	}
	return nil, nil
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	switch function {

	case "getRowTableOne":
		if len(args) < 1 {
			return nil, errors.New("getRowTableOne failed. Must include 1 key value")
		}

		col1Val := args[0]
		var columns []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: col1Val}}
		columns = append(columns, col1)

		row, err := stub.GetRow("tableOne", columns)
		if err != nil {
			return nil, fmt.Errorf("getRowTableOne operation failed. %s", err)
		}

		rowString := fmt.Sprintf("%s", row)
		return []byte(rowString), nil
		
	
	case "getRowsTableOne":
		
		var columns []shim.Column
		rows, err := stub.GetRows("tableOne", columns)
		if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
	
	res2E:= []*tabone{}
	
	for row := range rows {
		newApp:= new(tabone)
		newApp.vehicleId = row.Columns[0].GetString_()
		newApp.vehicleMake = row.Columns[1].GetString_()
		newApp.vehicleModel = row.Columns[2].GetString_()
		newApp.vehicleColour = row.Columns[3].GetString_()
		newApp.vehicleReg = row.Columns[4].GetString_()
		newApp.vehicleOwner = row.Columns[5].GetString_()
		
		
		
		
		fmt.Println("printing test value ----"+row.Columns[0].GetString_())
		fmt.Println("printing test value ----"+row.Columns[1].GetString_())
		fmt.Println("printing test value ----"+row.Columns[2].GetString_())
		//fmt.Println("printing test value *****"+newApp.ColOneTableOne)
		//fmt.Println("printing test value *****"+newApp.ColTwoTableOne)
		//fmt.Println("printing test value *******"+newApp.ColThreeTableOne)
		
		res2E=append(res2E,newApp)
		
	}
	res2F, _ := json.Marshal(res2E)
	
	return res2F, nil
	
		
	default:
		return nil, errors.New("Unsupported operation")
	}
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

func createTableOne(stub shim.ChaincodeStubInterface) error {
	// Create table one
	var columnDefsTableOne []*shim.ColumnDefinition
	columnOneTableOneDef := shim.ColumnDefinition{Name: "VehicleId",
		Type: shim.ColumnDefinition_STRING, Key: true}
	columnTwoTableOneDef := shim.ColumnDefinition{Name: "VehicleMake",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnThreeTableOneDef := shim.ColumnDefinition{Name: "VehicleModel",
		Type: shim.ColumnDefinition_STRING, Key: false}
	
	columnFourTableOneDef := shim.ColumnDefinition{Name: "VehicleColour",
		Type: shim.ColumnDefinition_STRING, Key: false}
		
	columnFiveTableOneDef := shim.ColumnDefinition{Name: "VehicleReg",
		Type: shim.ColumnDefinition_STRING, Key: false}	
		
	columnSixTableOneDef := shim.ColumnDefinition{Name: "VehicleOwner",
		Type: shim.ColumnDefinition_STRING, Key: false}	
	
	
	
	
	columnDefsTableOne = append(columnDefsTableOne, &columnOneTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnTwoTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnThreeTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnFourTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnFiveTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnSixTableOneDef)
	
	
	
	
	
	return stub.CreateTable("tableOne", columnDefsTableOne)
	
	
}

