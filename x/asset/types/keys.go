package types

const (
	// ModuleName defines the module name
	ModuleName = "asset"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_asset"
)

var (
	ParamsKey = []byte("p_asset")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	AssetKey      = "Asset/value/"
	AssetCountKey = "Asset/count/"
)

const (
	ValueVotesKey      = "ValueVotes/value/"
	ValueVotesCountKey = "ValueVotes/count/"
)
