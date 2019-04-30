/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Trade Finance Use Case - WORK IN  PROGRESS
 */

package main


import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}


// Define the letter of credit
type LetterOfCredit struct {
	InvoiceId string `json:"invoiceId"`
	InvoiceDate string `json:"invoiceDate"`
	Supplier string `json:"supplier"`
	Customer string `json:"customer"`
	PaymentTerms string `json:"paymentTerms"`
	Amount int `json:"amount"`
	Notes string `json:"notes"`
	Status string `json:"status"`
}


func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "issueInvoice" {
		return s.issueInvoice(APIstub, args)
	} else if function == "acceptInvoice" {
		return s.acceptInvoice(APIstub, args)
	} else if function == "payInvoice" {
		return s.payInvoice(APIstub, args)
	}else if function == "getInvoice" {
		return s.getInvoice(APIstub, args)
	}else if function == "getInvoiceHistory" {
		return s.getInvoiceHistory(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}





// This function is initiate by Buyer 
func (s *SmartContract) issueInvoice(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	invoiceId := args[0];
	invoiceDate := args[1];
	supplier := args[2];
	customer := args[3];
	paymentTerms := args[4];
	notes := args[6];
	amount, err := strconv.Atoi(args[5]);
	if err != nil {
		return shim.Error("Not able to parse Amount")
	}


	LC := LetterOfCredit{InvoiceId: invoiceId, InvoiceDate: invoiceDate, Supplier: supplier, Customer: customer, PaymentTerms: paymentTerms, Amount: amount, Notes: notes, Status: "Issued"}
	LCBytes, err := json.Marshal(LC)

    APIstub.PutState(invoiceId,LCBytes)
	fmt.Println("Invoice Issued -> ", LC)


	return shim.Success(nil)
}

// This function is initiate by Seller
func (s *SmartContract) acceptInvoice(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	invoiceId := args[0];
	// if err != nil {
	// 	return shim.Error("No Amount")
	// }

	LCAsBytes, _ := APIstub.GetState(invoiceId)

	var lc LetterOfCredit

	err := json.Unmarshal(LCAsBytes, &lc)

	if err != nil {
		return shim.Error("Issue with Invoice json unmarshaling")
	}


	LC := LetterOfCredit{InvoiceId: lc.InvoiceId, InvoiceDate: lc.InvoiceDate, Supplier: lc.Supplier, Customer: lc.Customer, PaymentTerms: lc.PaymentTerms, Amount: lc.Amount, Notes: lc.Notes, Status: "Accepted"}
	LCBytes, err := json.Marshal(LC)

	if err != nil {
		return shim.Error("Issue with Invoice json marshaling")
	}

    APIstub.PutState(lc.InvoiceId,LCBytes)
	fmt.Println("Invoice Accepted -> ", LC)


	return shim.Success(nil)
}

func (s *SmartContract) payInvoice(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	invoiceId := args[0];

	LCAsBytes, _ := APIstub.GetState(invoiceId)

	var lc LetterOfCredit

	err := json.Unmarshal(LCAsBytes, &lc)

	if err != nil {
		return shim.Error("Issue with Invoice json unmarshaling")
	}


	LC := LetterOfCredit{InvoiceId: lc.InvoiceId, InvoiceDate: lc.InvoiceDate, Supplier: lc.Supplier, Customer: lc.Customer, PaymentTerms: lc.PaymentTerms, Amount: lc.Amount, Notes: lc.Notes, Status: "Paid"}
	LCBytes, err := json.Marshal(LC)

	if err != nil {
		return shim.Error("Issue with Invoice json marshaling")
	}

    APIstub.PutState(lc.InvoiceId,LCBytes)
	fmt.Println("Invoice Paid -> ", LC)



	return shim.Success(nil)
}

func (s *SmartContract) getInvoice(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	invoiceId := args[0];
	// if err != nil {
	// 	return shim.Error("No Amount")
	// }

	LCAsBytes, _ := APIstub.GetState(invoiceId)

	return shim.Success(LCAsBytes)
}

func (s *SmartContract) getInvoiceHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	invoiceId := args[0];

	resultsIterator, err := APIstub.GetHistoryForKey(invoiceId)
	if err != nil {
		return shim.Error("Error retrieving Invoice history.")
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error("Error retrieving Invoice history.")
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getInvoiceHistory returning:\n%s\n", buffer.String())


	return shim.Success(buffer.Bytes())
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
