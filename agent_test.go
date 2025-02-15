package agent

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/aviate-labs/candid-go/idl"
	"github.com/fxamacker/cbor/v2"
	"github.com/mix-labs/IC-Go/utils"
	"github.com/mix-labs/IC-Go/utils/identity"
	"github.com/mix-labs/IC-Go/utils/principal"
)

// EXT data structure
type supply struct {
	Ok  uint64 `ic:"ok"`
	Err string `ic:"err"`
}
type Time struct {
	Some big.Int `ic:"some"`
	None uint8   `ic:"none"`
}
type listing struct {
	Locked Time                `ic:locked`
	Price  uint64              `ic:"price"`
	Seller principal.Principal `ic:"seller"`
}
type listingTuple struct {
	A uint32  `ic:"0"`
	B listing `ic:"1"`
}
type listings []listingTuple

type TokenIndex uint32
type RegistryTuple struct {
	A TokenIndex `ic:"0"`
	B string     `ic:"1"`
}
type Registrys []RegistryTuple

// PUNK data structure
type principalOp struct {
	Some principal.Principal `ic:"some"`
	None uint8               `ic:"none"`
}
type priceOp struct {
	Some uint64 `ic:"some"`
	None uint8  `ic:"none"`
}

type NULL *uint8

type Operation struct {
	Delist   NULL `ic:"delist"`
	Init     NULL `ic:"init"`
	List     NULL `ic:"list"`
	Mint     NULL `ic:"mint"`
	Purchase NULL `ic:"purchase"`
	Transfer NULL `ic:"transfer"`

	//To formulate a enum struct
	Index string `ic:"EnumIndex"`
}

type transaction struct {
	Caller    principal.Principal `ic:"caller"`
	To        principalOp         `ic:"to"`
	From      principalOp         `ic:"from"`
	Index     big.Int             `ic:"index"`
	Price     priceOp             `ic:"price"`
	Timestamp big.Int             `ic:"timestamp"`
	TokenId   big.Int             `ic:"tokenId"`
	Op        Operation           `ic:"op"`
}

type DetailValue struct {
	I64        int64               `ic:"I64"`
	U64        uint64              `ic:"U64"`
	Vec        []DetailValue       `ic:"Vec"`
	Slice      []uint8             `ic:"Slice"`
	TokenIdU64 uint64              `ic:"TokenIdU64"`
	Text       string              `ic:"Text"`
	True       *uint8              `ic:"True"`
	False      *uint8              `ic:"False"`
	Float      float64             `ic:"Float"`
	Principal  principal.Principal `ic:"Principal"`

	Index string `ic:"EnumIndex"`
}

type Details struct {
	Text  string      `ic:"0"`
	Value DetailValue `ic:"1"`
}
type Event struct {
	Time      uint64              `ic:"time"`
	Operation string              `ic:"operation"`
	Details   []Details           `ic:"details"`
	Caller    principal.Principal `ic:"caller"`
}

type GetTransactionsResponseBorrowed struct {
	Data []Event `ic:"data"`
	Page uint32  `ic:"page"`
}

func TestAgent_QueryRaw(t *testing.T) {
	//EXT canister
	//canisterID := "bzsui-sqaaa-aaaah-qce2a-cai"

	//PUNK canister
	// canisterID := "qfh5c-6aaaa-aaaah-qakeq-cai"
	agent := New(true, "")
	//agent := New(false, "833fe62409237b9d62ec77587520911e9a759cec1d19755b7da901b96dca3d42")
	// agent, err := NewFromPem(false, "./utils/identity/priv.pem")
	// if err != nil {
	// 	t.Log(err)
	// }

	// canisterID := principal.Principal([]byte{0, 0, 0, 0, 0, 240, 17, 32, 1, 1}).Encode()

	// testRec := idl.NewRec(map[string]idl.Type{
	// 	"a": idl.NewVec(idl.Int32()),
	// 	"b": idl.NewVec(idl.Int32()),
	// })
	// testRecValue := map[string]interface{}{
	// 	"a":[]interface{}{},
	// 	"b":[]interface{}{},
	// }

	// arg, _ := idl.Encode([]idl.Type{testRec}, []interface{}{testRecValue})
	// fmt.Println(arg)
	// methodName := "get_transactions"

	// rec := map[string]idl.Type{}
	// rec["page"] = idl.NewOpt(idl.Nat32())
	// rec["witness"] = new(idl.Bool)
	// value := map[string]interface{}{}
	// value["page"] = big.NewInt(1)
	// value["witness"] = false

	// arg, _ := idl.Encode([]idl.Type{idl.NewRec(rec)}, []interface{}{value})
	// fmt.Println(arg, canisterID, methodName)

	// _, result, errMsg, err := agent.Query(canisterID, methodName, arg)
	// myresult := GetTransactionsResponseBorrowed{}
	// utils.Decode(&myresult, result[0])
	//fmt.Println(result)
	//
	////EXT method
	// methodName := "supply"
	////methodName := "listings"
	////methodName := "getRegistry"
	//
	////PUNK method
	//methodName := "getHistoryByIndex"
	//
	////arg, err := idl.Encode([]idl.Type{new(idl.Null)}, []interface{}{nil})
	//arg, err := idl.Encode([]idl.Type{new(idl.Nat)}, []interface{}{big.NewInt(10)})
	//if err != nil {
	//	t.Error(err)
	//}
	// Type, result, errMsg, err := agent.Query(canisterID, methodName, arg)
	//
	////myresult := supply{}
	////myresult := listings{}
	////myresult := Registrys{}
	//
	//myresult := transaction{}
	////fmt.Println(result[0])
	//utils.Decode(&myresult, result[0])

	// t.Log("errMsg:", errMsg, "err:", err, "result:", myresult)
	// fmt.Println(myresult.Data[1])

	methodName := "getAllToken"
	canister := "cuvse-myaaa-aaaan-qas6a-cai"
	arg, _ := idl.Encode([]idl.Type{idl.NewOptionalType(idl.NullType{})}, []interface{}{nil})
	_, result, _, err := agent.Query(canister, methodName, arg)
	if err != nil {
		panic(err)
	}
	fmt.Println(result[0])
}

func TestAgent_UpdateRaw(t *testing.T) {
	// canisterID := "gvbup-jyaaa-aaaah-qcdwa-cai"
	// agent := New(false, "833fe62409237b9d62ec77587520911e9a759cec1d19755b7da901b96dca3d42")

	// methodName := "transfer"
	// var argType []idl.Type
	// var argValue []interface{}
	// p, _ := principal.Decode("aaaaa-aa")
	// argType = append(argType, new(idl.Principal))
	// argType = append(argType, new(idl.Nat))
	// argValue = append(argValue, p)
	// argValue = append(argValue, big.NewInt(10000000000))
	var myresult uint64
	canisterID := "d24m2-dqaaa-aaaah-aa4zq-cai"
	ag := New(false, "833fe62409237b9d62ec77587520911e9a759cec1d19755b7da901b96dca3d42")
	methodName := "total"
	arg, _ := idl.Encode([]idl.Type{idl.NullType{}}, []interface{}{nil})
	_, result, err := ag.Update(canisterID, methodName, arg, 30)
	if err != nil {
		panic(err)
	}
	utils.Decode(&myresult, result[0])

	// arg, _ := idl.Encode(argType, argValue)
	// _, result, err := agent.Update(canisterID, methodName, arg)

	t.Log("errMsg:", err, "result:", myresult)
}

func TestAgent_GetCanisterModule(t *testing.T) {
	canisterID := "bzsui-sqaaa-aaaah-qce2a-cai"
	agent := New(false, "833fe62409237b9d62ec77587520911e9a759cec1d19755b7da901b96dca3d42")
	result, err := agent.GetCanisterModule(canisterID)
	if err != nil {
		t.Log("err:", err)
	} else if result == nil {
		t.Log("no module")
	} else {
		t.Log("hash:", hex.EncodeToString(result))
	}
}

func TestAgent_GetCanisterControllers(t *testing.T) {
	canisterID := "6b4pv-sqaaa-aaaah-qaava-cai"
	agent := New(false, "833fe62409237b9d62ec77587520911e9a759cec1d19755b7da901b96dca3d42")
	result, err := agent.GetCanisterControllers(canisterID)
	if err != nil {
		t.Log("err:", err)
	} else {
		for _, i := range result {
			t.Log("controller:", i.Encode())
		}
	}
	t.Log(result)
}

func TestPrincipal(t *testing.T) {
	pkBytes, _ := hex.DecodeString("833fe62409237b9d62ec77587520911e9a759cec1d19755b7da901b96dca3d42")
	identity := identity.New(false, pkBytes)
	p := principal.NewSelfAuthenticating(identity.PubKeyBytes())
	t.Log(p.Encode(), len(identity.PubKeyBytes()))
}

func TestCbor(t *testing.T) {
	canisterID, _ := principal.Decode("gvbup-jyaaa-aaaah-qcdwa-cai")
	agent := New(true, "833fe62409237b9d62ec77587520911e9a759cec1d19755b7da901b96dca3d42")

	req := Request{
		Type:          "call",
		Sender:        agent.Sender(),
		IngressExpiry: uint64(agent.getExpiryDate().UnixNano()),
		CanisterID:    canisterID,
		MethodName:    "transfer",
		Arguments:     []byte("i love vivian"),
	}

	envelope := Envelope{
		req,
		[]byte{},
		[]byte{},
	}

	data, _ := cbor.Marshal(envelope)
	resp := new(Envelope)
	cbor.Unmarshal(data, resp)
	t.Log("sender", resp.Content.Sender.Encode())
	t.Log("type", resp.Content.Type)
	t.Log("ingress expiryt", resp.Content.IngressExpiry)
	t.Log("method", resp.Content.MethodName)
	t.Log("arg", resp.Content.Arguments)
	t.Log("canister", resp.Content.CanisterID.Encode())
}

func TestAgent_GetCanisterTime(t *testing.T) {
	canisterID := "b65vx-3qaaa-aaaaa-7777q-cai"
	agent := New(false, "833fe62409237b9d62ec77587520911e9a759cec1d19755b7da901b96dca3d42")
	result, err := agent.GetCanisterTime(canisterID)
	if err != nil {
		t.Log("err:", err)
	} else {
		t.Log("result:", result)

	}
}

func TestAgent_GetCanisterCandid(t *testing.T) {
	canisterID := "oeee4-qaaaa-aaaak-qaaeq-cai"
	agent := New(false, "833fe62409237b9d62ec77587520911e9a759cec1d19755b7da901b96dca3d42")
	arg, _ := idl.Encode([]idl.Type{idl.NullType{}}, []interface{}{nil})
	methodName := "__get_candid_interface_tmp_hack"
	_, result, err := agent.Update(canisterID, methodName, arg, 30)
	if err != nil {
		panic(err)
	}
	if err != nil {
		t.Log("err:", err)
	} else {
		t.Log("result:", result)

	}
}
