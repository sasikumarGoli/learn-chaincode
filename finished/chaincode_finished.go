

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

VehicleId string       `json:"vehicleId"`
VehicleMake string     `json:"vehicleMake"`
VehicleModel string    `json:"vehicleModel"`
VehicleColour string   `json:"vehicleColour"`
VehicleReg string      `json:"vehicleReg"`
VehicleOwner string    `json:"vehicleOwner"`

//addes by Reshma
AuctionID string       `json:"AuctionID"`
SaleYear string        `json:"SaleYear"`
SaleNo string          `json:"SaleNo"`
LaneNo string          `json:"LaneNo"`
RunNo   string         `json:"RunNo"`
PaymentMode string     `json:"PaymentMode"`
VehicleStatus string   `json:"VehicleStatus"`
TitleStatus string     `json:"TitleStatus"`

}




// Init create tables for tests
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// Create table one
	err := createTableOne(stub)
	if err != nil {
		return nil, fmt.Errorf("Error creating table one during init. %s", err)
	}

	//added by Reshma
	err2 := createTableTwo(stub)
	if err2 != nil {
		return nil, fmt.Errorf("Error creating table one during init. %s", err2)
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


//start of table twoinsnertion by Reshma
	   case "insertRowTableTwo":
		if len(args) < 3 {
			return nil, errors.New("insertTableOne failed. Must include 3 column values")
		}
	     var str [] string

		for i := 0; i < len(args); i++ {
		
		str = s.Split(args[i],"-")
		
		
		col1Val := str[0]
		col2Val :=str[1]
		col3Val := str[2]
		col4Val := str[3]
		col5Val := str[4]
		col6Val := str[5]
		col7Val := str[6]
		col8Val := str[7]
		col9Val := str[8]
		col10Val := str[9]
		col11Val := str[10]
		col12Val := str[11]
		col13Val := str[12]
		col14Val := str[13]
		col15Val := str[14]

		var columns []*shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: col1Val}}
		col2 := shim.Column{Value: &shim.Column_String_{String_: col2Val}}
		col3 := shim.Column{Value: &shim.Column_String_{String_: col3Val}}
		col4 := shim.Column{Value: &shim.Column_String_{String_: col4Val}}
	        col5 := shim.Column{Value: &shim.Column_String_{String_: col5Val}}
		col6 := shim.Column{Value: &shim.Column_String_{String_: col6Val}}  		
		col7 := shim.Column{Value: &shim.Column_String_{String_: col7Val}}  		
		col8 := shim.Column{Value: &shim.Column_String_{String_: col8Val}}  		
		col9 := shim.Column{Value: &shim.Column_String_{String_: col9Val}}  		
		col10 := shim.Column{Value: &shim.Column_String_{String_: col10Val}}  		
		col11 := shim.Column{Value: &shim.Column_String_{String_: col11Val}}  		
		col12 := shim.Column{Value: &shim.Column_String_{String_: col12Val}}  		
		col13 := shim.Column{Value: &shim.Column_String_{String_: col13Val}}  		
		col14 := shim.Column{Value: &shim.Column_String_{String_: col14Val}}  		
		col15 := shim.Column{Value: &shim.Column_String_{String_: col15Val}}  		

		columns = append(columns, &col1)
		columns = append(columns, &col2)
		columns = append(columns, &col3)
		columns = append(columns, &col4)
		columns = append(columns, &col5)
		columns = append(columns, &col6)
		columns = append(columns, &col7)
		columns = append(columns, &col8)
		columns = append(columns, &col9)
		columns = append(columns, &col10)
		columns = append(columns, &col11)
		columns = append(columns, &col12)
		columns = append(columns, &col13)
		columns = append(columns, &col14)
		columns = append(columns, &col15)

		row := shim.Row{Columns: columns}
		ok, err2 := stub.InsertRow("tableTwo", row)
		if err2 != nil {
			return nil, fmt.Errorf("insertTableTwo operation failed. %s", err2)
		}
		if !ok {
			return nil, errors.New("insertTableTwo operation failed. Row with given key already exists")
		}
       
	   
	   }


	
//End of table twoinsertion by Reshma	

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
		newApp.VehicleId = row.Columns[0].GetString_()
		newApp.VehicleMake = row.Columns[1].GetString_()
		newApp.VehicleModel = row.Columns[2].GetString_()
		newApp.VehicleColour = row.Columns[3].GetString_()
		newApp.VehicleReg = row.Columns[4].GetString_()
		newApp.VehicleOwner = row.Columns[5].GetString_()
		
		
		
		
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

//added by Reshma

func createTableTwo(stub shim.ChaincodeStubInterface) error {
	// Create table two
	var columnDefsTableTwo []*shim.ColumnDefinition


	columnOneTableTwoDef := shim.ColumnDefinition{Name: "VehicleId",
		Type: shim.ColumnDefinition_STRING, Key: true}
	columnTwoTableTwoDef := shim.ColumnDefinition{Name: "VehicleMake",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnThreeTableTwoDef := shim.ColumnDefinition{Name: "VehicleModel",
		Type: shim.ColumnDefinition_STRING, Key: false}
	
	columnFourTableTwoDef := shim.ColumnDefinition{Name: "VehicleColour",
		Type: shim.ColumnDefinition_STRING, Key: false}
		
	columnFiveTableTwoDef := shim.ColumnDefinition{Name: "VehicleReg",
		Type: shim.ColumnDefinition_STRING, Key: false}	
		
	columnSixTableTwoDef := shim.ColumnDefinition{Name: "VehicleOwner",
		Type: shim.ColumnDefinition_STRING, Key: false}	


	columnSevenTableTwoDef := shim.ColumnDefinition{Name: "AuctionID",
		Type: shim.ColumnDefinition_STRING, Key: true}
	columnEightTableTwoDef := shim.ColumnDefinition{Name: "SaleYear",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnNineTableTwoDef := shim.ColumnDefinition{Name: "SaleNo",
		Type: shim.ColumnDefinition_STRING, Key: false}
	
	columnTenTableTwoDef := shim.ColumnDefinition{Name: "LaneNo",
		Type: shim.ColumnDefinition_STRING, Key: false}
		
	columnElevenTableTwoDef := shim.ColumnDefinition{Name: "RunNo",
		Type: shim.ColumnDefinition_STRING, Key: false}	
		
	columnTwelveTableTwoDef := shim.ColumnDefinition{Name: "PaymentMode",
		Type: shim.ColumnDefinition_STRING, Key: false}	
	
	columnThirteenTableTwoDef := shim.ColumnDefinition{Name: "VehicleStatus",
		Type: shim.ColumnDefinition_STRING, Key: false}	
	columneFourteenTableTwoDef := shim.ColumnDefinition{Name: "TitleStatus",
			Type: shim.ColumnDefinition_STRING, Key: false}	
	

	
		
	columnDefsTableTwo = append(columnDefsTableTwo, &columnOneTableTwoDef)
	columnDefsTableTwo = append(columnDefsTableTwo, &columnTwoTableTwoDef)
	columnDefsTableTwo = append(columnDefsTableTwo, &columnThreeTableTwoDef)
	columnDefsTableTwo = append(columnDefsTableTwo, &columnFourTableTwoDef)
	columnDefsTableTwo = append(columnDefsTableTwo, &columnFiveTableTwoDef)
	columnDefsTableTwo = append(columnDefsTableTwo, &columnSixTableTwoDef)
	columnDefsTableTwo = append(columnDefsTableTwo, &columnSevenTableTwoDef)
	columnDefsTableTwo = append(columnDefsTableTwo, &columnEightTableTwoDef)
	columnDefsTableTwo = append(columnDefsTableTwo, &columnNineTableTwoDef)
	columnDefsTableTwo = append(columnDefsTableTwo, &columnTenTableTwoDef)
	columnDefsTableTwo = append(columnDefsTableTwo, &columnElevenTableTwoDef)
	columnDefsTableTwo = append(columnDefsTableTwo, &columnTwelveTableTwoDef)
	columnDefsTableTwo = append(columnDefsTableTwo, &columnThirteenTableTwoDef)
	columnDefsTableTwo = append(columnDefsTableTwo, &columneFourteenTableTwoDef)

	
	
	return stub.CreateTable("tableTwo", columnDefsTableTwo)
	
	
}


