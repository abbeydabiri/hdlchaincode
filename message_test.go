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

func  TestMessageContract(t *testing.T) {
	os.Setenv("MODE","TEST")
	
	assert := assert.New(t)
	uid := uuid.New().String()

	cc, err := contractapi.NewChaincode(new(MessageContract))
	assert.Nil(err, "error should be nil")

	stub := shimtest.NewMockStub("TestStub", cc)
	assert.NotNil(stub, "Stub is nil, TestStub creation failed")

	// - - - test MessageContract:Put function - - - 
	putResp := stub.MockInvoke(uid,[][]byte{
		[]byte("MessageContract:Put"),
		[]byte("1"),
		[]byte("From"),
		[]byte("To"),
		[]byte("Subject"),
		[]byte("Message"),
		[]byte("Message Date"),
	})
	assert.EqualValues(OK, putResp.GetStatus(), putResp.GetMessage())
	

	// - - - test MessageContract:Get function - - - 
	testID := "1"
	getResp := stub.MockInvoke(uid, [][]byte{
		[]byte("MessageContract:Get"),
		[]byte(testID),
	})
	assert.EqualValues(OK, getResp.GetStatus(), getResp.GetMessage())
	assert.NotNil(getResp.Payload, "getResp.Payload should not be nil")
	
	messageObj := new(MessageObj)
	err = json.Unmarshal(getResp.Payload, messageObj)
	assert.Nil(err, "json.Unmarshal error should be nil")
	assert.NotNil(messageObj, "messageObj should not be nil")

	retrievedID := strconv.Itoa(messageObj.MessageID)
	assert.EqualValues(testID, retrievedID, "testID and retrievedID mismatch")
}