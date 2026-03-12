package state

import (
	"testing"
	"github.com/ethereum/go-ethereum/common"
)

// TestBSCCreateContractPanic demonstrates a Nil Pointer Dereference
// in the CreateContract function when an address is missing from the state.
func TestBSCCreateContractPanic(t *testing.T) {
	// Initialize a mock state database
	// Note: In a real environment, this state would be populated from the disk/db
	state, _ := New(common.Hash{}, NewDatabaseForTesting())

	// An address that is guaranteed not to be in the database
	addr := common.HexToAddress("0xdeadbeef")

	t.Logf("Triggering CreateContract for non-existent address: %v", addr)

	// This call triggers the panic: runtime error: invalid memory address
	state.CreateContract(addr)
}
