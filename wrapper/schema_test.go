package wrapper

import "os"


var mockRPCURL = os.Getenv("RPC_URL")
var mockContractAddress = "0xE4F771f86B34BF7B323d9130c385117Ec39377c3" // before transition

// func TestVerifyState(t *testing.T) {

// 	stateResult, err := VerifyState(context.Background(), mockRPCURL, mockContractAddress, mockGenesisID, mockGenesisState)
// 	assert.Nil(t, err)
// 	assert.Equal(t, true, stateResult.Latest)
// }

// func TestVerifyInvalidRPC(t *testing.T) {

// 	invalidURL := "test://invalidurl1234.com"
// 	_, err := VerifyState(context.Background(), invalidURL, mockContractAddress, mockGenesisID, mockGenesisState)
// 	assert.NotNil(t, err)
// 	assert.Contains(t, err.Error(), errRPCClientCreationMessage)

// 	invalidURL = "http://invalidurl1234.com"
// 	_, err = VerifyState(context.Background(), invalidURL, mockContractAddress, mockGenesisID, mockGenesisState)

// 	assert.NotNil(t, err)
// 	assert.Contains(t, err.Error(), "no such host")

// }

// func TestVerifyGenesisState(t *testing.T) {

// 	stateResult, err := VerifyState(context.Background(), mockRPCURL, mockContractAddress, mockGenesisID, mockGenesisState)
// 	assert.Nil(t, err)
// 	assert.Equal(t, true, stateResult.Latest)

// }

// func TestVerifyGenesisStateWrongID(t *testing.T) {

// 	wrongID, _ := new(big.Int).SetString("26592849444054787445766572449338308165040390141345377877344569181291872256", 10)
// 	_, err := VerifyState(context.Background(), mockRPCURL, mockContractAddress, wrongID, mockGenesisState)
// 	assert.NotNil(t, err)
// 	assert.Error(t, err, "ID from genesis state (11A2HgCZ1pUcY8HoNDMjNWEBQXZdUnL3YVnVCUvR5s) and provided (118cr7d17eL2sSYk5hrMBo9MKJrWGD5RrFgsqXupGE) don't match")

// }

// func TestVerifyPublishedLatestState(t *testing.T) {

// 	stateResult, err := VerifyState(context.Background(), mockRPCURL, mockContractAddress, mockIDForPublishedLatestState, mockPublishedLatestState)
// 	assert.Nil(t, err)
// 	assert.Equal(t, true, stateResult.Latest)
// }

// func TestVerifyStateTransitionCheck(t *testing.T) {

// 	// latest state - equal
// 	stateResult1, err := VerifyState(context.Background(), mockRPCURL, mockContractAddressForTransitionTest, mockIDForTransitionTest, mockGenesisSecondStateForTransitionTest)
// 	assert.Nil(t, err)
// 	assert.Equal(t, true, stateResult1.Latest)

// 	// latest state - not equal
// 	stateResult2, err := VerifyState(context.Background(), mockRPCURL, mockContractAddressForTransitionTest, mockIDForTransitionTest, mockGenesisFistStateForTransitionTest)
// 	assert.Nil(t, err)
// 	assert.Equal(t, false, stateResult2.Latest)
// 	assert.NotEqual(t, 0, stateResult2.TransitionTimestamp)

// }

