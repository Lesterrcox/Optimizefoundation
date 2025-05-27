package types

const (
	// ModuleName defines the module name
	ModuleName = "validator"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_validator"
)

var (
	ParamsKey = []byte("p_validator")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ValidatorKey      = "Validator/value/"
	ValidatorCountKey = "Validator/count/"
)
