package identity

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"github.com/mix-labs/IC-Go/utils/principal"
	"hash/crc32"
)

func PrincipalToAccount(pp principal.Principal) string{

	vPp := pp[:]
	hash := sha256.New224()
	hash.Write([]byte("\x0Aaccount-id") )
	hash.Write(vPp)
	subAcc := [32]byte{}
	hash.Write(subAcc[:])
	hashSum := hash.Sum(nil)


	cs := make([]byte, 4)
	binary.BigEndian.PutUint32(cs, crc32.ChecksumIEEE(hashSum[:]))

	vResult := append(cs, hashSum[:]...)
	strAccount := hex.EncodeToString(vResult)
	//println("strAccount:",strAccount)
	return strAccount
}

func GenPrincipal(privKey ed25519.PrivateKey) (principal.Principal,error){

	pubKey := privKey.Public()
	asnKeyBytes, _ := MarshalEd25519PublicKey(pubKey)
	pp := principal.NewSelfAuthenticating(asnKeyBytes)
	return  pp,nil

}

func GenPrivateKey() (ed25519.PublicKey,ed25519.PrivateKey,error){
	pubKey,privKey,err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil,nil,err
	}

	return pubKey, privKey,nil
}

