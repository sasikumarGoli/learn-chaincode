
package main
import (
	"errors"
	
	
	"fmt"

	//"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)




type SampleChaincode struct {
}


//custom data models
type EMPInfo struct {
	EmpId     string  `json:"empid"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
}


type EMPDetails struct {
	MonthlySalary      int `json:"monthlySalary"`
	MonthlyRent        int `json:"monthlyRent"`
	OtherExpenditure   int `json:"otherExpenditure"`
	
}





func CreateEmploye(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
	fmt.Printf("enter to create emp")

	if len(args) < 2 {
		
		return nil, errors.New("Expected atleast two arguments  for creation")
	}

	var EmpId= args[0]
	var EmpDetails = args[1]

	err := stub.PutState(EmpId, []byte(EmpDetails))
	if err != nil {
		
		return nil, err
	}

	fmt.Printf("created Emp")
	return nil, nil

}





func GetEmpDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {


	if len(args) < 1 {
		
		return nil, errors.New("Missing loan emp ID")
	}

	var EmpId = args[0]
	bytes, err := stub.GetState(EmpId)
	if err != nil {
		
		return nil, err
	}
	return bytes, nil
}





func (t *SampleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}





func (t *SampleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "GetEmpDetails" {
		return GetEmpDetails(stub, args)
	}
	return nil, nil
}






func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "CreateLoanApplication" {
		
			return CreateEmploye(stub, args)
		} else {
			return nil, errors.New(" does not have access to create a loan application")
		}
return nil, nil
	}
	







func main() {

	

	err := shim.Start(new(SampleChaincode))
	if err != nil {
		
	} else {
		
	}

}