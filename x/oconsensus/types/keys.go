package types

const (
	// ModuleName defines the module name
	ModuleName = "oconsensus"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oconsensus"
)

var (
	ParamsKey = []byte("p_oconsensus")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	CONSENSUS_PASS_SUCCESS      = "pass_to_be_success"
	CONSENSUS_PASS_FAILED       = "pass_to_be_fail"
	CONSENSUS_FAIL_INSUFFICIENT = "insufficient_votes"
)
