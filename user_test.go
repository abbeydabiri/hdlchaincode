package main

import (
	"os"
	"testing"
	"encoding/json"
	"strconv"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func  TestUserContract(t *testing.T) {
	os.Setenv("MODE","TEST")
	
	assert := assert.New(t)
	uid := uuid.New().String()

	cc, err := contractapi.NewChaincode(new(UserContract))
	assert.Nil(err, "error should be nil")

	stub := shimtest.NewMockStub("TestStub", cc)
	assert.NotNil(stub, "Stub is nil, TestStub creation failed")

	// - - - test UserContract:Put function - - - 
	putResp := stub.MockInvoke(uid,[][]byte{
		[]byte("UserContract:Put"),
		[]byte("1"),
		[]byte("UserType"),
		[]byte("AccountStatus"),
		[]byte("UserCategory"),
		[]byte("FirstName"),
		[]byte("LastName"),
		[]byte("Email"),
		[]byte("Phone"),
		[]byte("Password"),
		[]byte("RegDate"),
	})
	assert.EqualValues(OK, putResp.GetStatus(), putResp.GetMessage())
	

	// - - - test UserContract:Get function - - - 
	testID := "1"
	getResp := stub.MockInvoke(uid, [][]byte{
		[]byte("UserContract:Get"),
		[]byte(testID),
	})
	assert.EqualValues(OK, getResp.GetStatus(), getResp.GetMessage())
	assert.NotNil(getResp.Payload, "getResp.Payload should not be nil")
	
	userObj := new(UserObj)
	err = json.Unmarshal(getResp.Payload, userObj)
	assert.Nil(err, "json.Unmarshal error should be nil")
	assert.NotNil(userObj, "userObj should not be nil")

	retrievedID := strconv.Itoa(userObj.UserID)
	assert.EqualValues(testID, retrievedID, "testID and retrievedID mismatch")
}