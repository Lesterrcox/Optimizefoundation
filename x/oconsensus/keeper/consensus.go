package keeper

import (
	"sort"

	"optimizeglobalcoin/x/oconsensus/types"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// **************************************************************************************************************************************
// RunConsensus implements consensus algorithm.
// This function receives the payload and run consensus algorithm to generate the final result (final validated value with consensus status)
// It counts the frequency of each voted value so that it can generate the final validated value.
// It also check the voting power of all voters and it should exceed the quorem.
//
// Parameters:
// 1. Voter : The address that is used to submit the vote
// 2. Result: True or false value which shows the result is succeed or failed.
// 3. Value: The number/string which needs to be validated.
// 4. ErrReson: If it is failed result, the reason string about it.
//
// Responses:
//   - Validate value: The final determiend value.
//   - Consensus result:
//   - CONSENSUS_FAIL_INSUFFICIENT: Doesn't have enough votes.
//   - CONSENSUS_PASS_SUCCESS: Have enough votes and the result is succeed
//   - CONSENSUS_PASS_FAILED: Have enough votes and the result is failed.
//   - Error reason: The error string it is failed votes.
//
// Note:
// This is very important function as it will be used in different parts that needs to run consensus
// It is being used from different module such as chains, transaction, tss, observers.
// Each module calls this function in their BeginBlocker or EndBlocker of the module ABCI.
// TODO: Should take care of 5-7 seats by refering to validator module parameter.
// **************************************************************************************************************************************
func (k Keeper) RunConsensus(ctx sdk.Context, payloads types.ConsensusPayload) (string, string, string) {
	totalBonded, err := k.stakingKeeper.TotalBondedTokens(ctx)
	if err != nil {
		// returns the validated value, consensus result and error reason
		return "", types.CONSENSUS_FAIL_INSUFFICIENT, "not able to get total bonding power"
	}

	positivePower := math.ZeroInt()
	negativePower := math.ZeroInt()

	// Variables to count each voted value frequency.
	// This variable is used to get the maximum freqent value
	// i.e. in pool balance vote, it can be the amount of token per pool
	// If we receive payload {{voter: voter1, result: true, value: 2},{voter: voter2, result: true, value: 3},{voter: voter3, result: true, value: 3}}
	// Frequency variable will be like frequency[2]=1, frequency[3]=2, and from this we might use the value `3` as final validated value.
	// As we might have succeed and failed votes, we have separated variables for that.
	valueFrequencySucess := make(map[string]int)
	valueFrequencyFail := make(map[string]int)
	errFrequency := make(map[string]int)

	// Final validated value
	// i.e. in transaction observation vote, this will be the amount of token transferred.
	// transactions might get failed or suceeded. so we need 2 values for each case.
	validatedValueSucceed := ""
	validatedValueFailed := ""
	errReaon := ""
	for _, payload := range payloads.Payloads {
		delAddress, err := sdk.AccAddressFromBech32(payload.Voter)
		if err != nil {
			continue
		}
		bonded, err := k.stakingKeeper.GetDelegatorBonded(ctx, delAddress)
		if err != nil {
			continue
		}

		// Sum the voting power of each voter
		if payload.Result {
			positivePower = positivePower.Add(bonded)

			// Collect succeed result
			if payload.Value != "" {
				valueFrequencySucess[payload.Value]++
			}
		} else {
			negativePower = negativePower.Add(bonded)

			// Collect failed result
			if payload.Value != "" {
				valueFrequencyFail[payload.Value]++
			}
		}

		// Ignore empty string
		if payload.ErrReason != "" {
			errFrequency[payload.ErrReason]++
		}
	}

	// Determine the value with the maximum count
	maxCount := 0
	sortKeys := k.SortKeys(valueFrequencySucess)
	// Iterate in sorted order
	for _, key := range sortKeys {
		count := valueFrequencySucess[key]
		if count > maxCount {
			maxCount = count
			validatedValueSucceed = key
		}
	}

	maxCount = 0
	sortKeys = k.SortKeys(valueFrequencyFail)
	// Iterate in sorted order
	for _, key := range sortKeys {
		count := valueFrequencyFail[key]
		if count > maxCount {
			maxCount = count
			validatedValueFailed = key
		}
	}

	maxCount = 0
	sortKeys = k.SortKeys(errFrequency)
	// Determine the error string with the maximum count
	for _, key := range sortKeys {
		count := errFrequency[key]
		if count > maxCount {
			maxCount = count
			errReaon = key
		}
	}

	consensusResult := types.CONSENSUS_FAIL_INSUFFICIENT
	resValue := ""

	// Calculate the quoreum = total voting power * 2 / 3
	quoreum := totalBonded.Mul(math.NewInt(2))
	quoreum = quoreum.Quo(math.NewInt(3))

	// calculate the total bonding power of the voters
	totalVotedPower := positivePower.Add(negativePower)

	// if positive power is exceeding than negative and quoreum of total bonding power
	if positivePower.GT(negativePower) && totalVotedPower.GTE(quoreum) {
		consensusResult = types.CONSENSUS_PASS_SUCCESS
		resValue = validatedValueSucceed
		errReaon = ""
	} else if negativePower.GT(positivePower) && totalVotedPower.GTE(quoreum) {
		// if negative power is exceeding than positive and quoreum of total bonding power
		consensusResult = types.CONSENSUS_PASS_FAILED
		resValue = validatedValueFailed
	}

	// returns the validated value, consensus result and error reason
	return resValue, consensusResult, errReaon
}

// Sort keys of map data to ensure consistency in consensus
func (k Keeper) SortKeys(data map[string]int) []string {
	// Extract keys
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}

	// Sort keys
	sort.Strings(keys)

	return keys
}
