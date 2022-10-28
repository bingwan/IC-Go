package agent

import (
	"encoding/hex"
	"github.com/aviate-labs/candid-go/idl"
	"testing"
)

func TestAgent_getDecimals(t *testing.T) {
	agent := New(true, "")
	methodName := "decimals"
	canister := "ryjl3-tyaaa-aaaaa-aaaba-cai"
	arg, _ := idl.Encode([]idl.Type{new(idl.NullType)}, []any{nil})
	_, result, _, err := agent.Query(canister, methodName, arg)

	//308955842 -> 8
	if err != nil {
		//panic(err)
	}
	_ = result
}

func TestAgent_getSymbol(t *testing.T) {
	agent := New(true, "")
	methodName := "symbol"
	canister := "ryjl3-tyaaa-aaaaa-aaaba-cai"
	arg, _ := idl.Encode([]idl.Type{new(idl.NullType)}, []any{nil})
	_, result, _, err := agent.Query(canister, methodName, arg)

	//4007505752 -> ICP
	if err != nil {
		//panic(rrr)
	}
	_ = result
}

func TestAgent_getName(t *testing.T) {
	agent := New(true, "")
	methodName := "name"
	canister := "ryjl3-tyaaa-aaaaa-aaaba-cai"
	arg, _ := idl.Encode([]idl.Type{new(idl.NullType)}, []any{nil})
	_, result, _, err := agent.Query(canister, methodName, arg)

	//1224700491 -> Internet Computer
	if err != nil {
		//panic(err)
	}
	_ = result

}

func TestAgent_getBalance(t *testing.T) {
	agent := New(true, "")
	methodName := "account_balance"
	canister := "ryjl3-tyaaa-aaaaa-aaaba-cai"

	strAccount := "883eef7c44be51afe4a4420d4df4beff708f3cf2f5de5efcc9f58680bb0f3690"
	vAccount, _ := hex.DecodeString(strAccount)

	vValue := make([]any, len(vAccount))
	for nIndex := 0; nIndex < len(vAccount); nIndex++ {
		byteOne := vAccount[nIndex]
		bigIntOne := uint8(byteOne)
		vValue[nIndex] = bigIntOne
	}

	rec := map[string]idl.Type{}
	rec["account"] = idl.NewVectorType(idl.Nat8Type())
	value := map[string]any{}
	value["account"] = vValue
	varg, err := idl.Encode([]idl.Type{idl.NewRecordType(rec)}, []any{value})

	if err != nil {
		println(" idl.Encode:", err.Error())
	}
	println("send arg:", string(varg))
	//5035232 -> 46167779
	_, result, _, err := agent.Query(canister, methodName, varg)
	if err != nil {
		//panic(err)
	}
	_ = result
}

// (record {start:nat64; length:nat64})
func TestAgent_getBlock(t *testing.T) {
	agent := New(true, "")
	methodName := "query_blocks"
	canister := "ryjl3-tyaaa-aaaaa-aaaba-cai"

	rec := map[string]idl.Type{}
	rec["start"] = idl.Nat64Type()
	rec["length"] = idl.Nat64Type()

	value := map[string]any{}
	value["start"] = uint64(4734875)
	value["length"] = uint64(1)

	varg, _ := idl.Encode([]idl.Type{idl.NewRecordType(rec)}, []any{value})

	println("send arg:", string(varg))
	_, result, _, err := agent.Query(canister, methodName, varg)

	if err != nil {
		//panic(err)
	}
	_ = result
}

func TestAgent_getTransfer_fee(t *testing.T) {
	agent := New(true, "")
	methodName := "transfer_fee"
	canister := "ryjl3-tyaaa-aaaaa-aaaba-cai"

	rec := map[string]idl.Type{}
	//value := map[string]any{}
	arg, _ := idl.Encode([]idl.Type{idl.NewRecordType(rec)}, []any{nil})

	_, result, _, err := agent.Query(canister, methodName, arg)

	if err != nil {
		//panic(err)
	}
	_ = result

}
