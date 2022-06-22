package identity

import (
	"crypto"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"

	"crypto/ed25519"
)

var errEd25519WrongID = errors.New("incorrect object identifier")
var errEd25519WrongKeyType = errors.New("incorrect key type")

// ed25519OID is the OID for the Ed25519 signature scheme: see
// https://datatracker.ietf.org/doc/draft-ietf-curdle-pkix-04.
var ed25519OID = asn1.ObjectIdentifier{1, 3, 101, 112}

// subjectPublicKeyInfo reflects the ASN.1 object defined in the X.509 standard.
//
// This is defined in crypto/x509 as "publicKeyInfo".
type subjectPublicKeyInfo struct {
	Algorithm pkix.AlgorithmIdentifier
	PublicKey asn1.BitString
}

// MarshalEd25519PublicKey creates a DER-encoded SubjectPublicKeyInfo for an
// ed25519 public key, as defined in
// https://tools.ietf.org/html/draft-ietf-curdle-pkix-04. This is analogous to
// MarshalPKIXPublicKey in crypto/x509, which doesn't currently support Ed25519.
func MarshalEd25519PublicKey(pk crypto.PublicKey) ([]byte, error) {
	pub, ok := pk.(ed25519.PublicKey)
	if !ok {
		return nil, errEd25519WrongKeyType
	}

	spki := subjectPublicKeyInfo{
		Algorithm: pkix.AlgorithmIdentifier{
			Algorithm: ed25519OID,
		},
		PublicKey: asn1.BitString{
			BitLength: len(pub) * 8,
			Bytes:     pub,
		},
	}
	return asn1.Marshal(spki)
}

// ParseEd25519PublicKey returns the Ed25519 public key encoded by the input.
func ParseEd25519PublicKey(der []byte) (crypto.PublicKey, error) {
	var spki subjectPublicKeyInfo
	if rest, err := asn1.Unmarshal(der, &spki); err != nil {
		return nil, err
	} else if len(rest) > 0 {
		return nil, errors.New("SubjectPublicKeyInfo too long")
	}

	if !spki.Algorithm.Algorithm.Equal(ed25519OID) {
		return nil, errEd25519WrongID
	}

	if spki.PublicKey.BitLength != ed25519.PublicKeySize*8 {
		return nil, errors.New("SubjectPublicKeyInfo PublicKey length mismatch")
	}

	return ed25519.PublicKey(spki.PublicKey.Bytes), nil
}

// oneAsymmetricKey reflects the ASN.1 structure for storing private keys in
// https://tools.ietf.org/html/draft-ietf-curdle-pkix-04, excluding the optional
// fields, which we don't use here.
//
// This is identical to pkcs8 in crypto/x509.
type oneAsymmetricKey struct {
	Version    int
	Algorithm  pkix.AlgorithmIdentifier
	PrivateKey []byte
}

// curvePrivateKey is the innter type of the PrivateKey field of
// oneAsymmetricKey.
type curvePrivateKey []byte

// MarshalEd25519PrivateKey returns a DER encoding of the input private key as
// specified in https://tools.ietf.org/html/draft-ietf-curdle-pkix-04.
func MarshalEd25519PrivateKey(sk crypto.PrivateKey) ([]byte, error) {
	priv, ok := sk.(ed25519.PrivateKey)
	if !ok {
		return nil, errEd25519WrongKeyType
	}

	// Marshal the innter CurvePrivateKey.
	curvePrivateKey, err := asn1.Marshal(priv.Seed())
	if err != nil {
		return nil, err
	}

	// Marshal the OneAsymmetricKey.
	asym := oneAsymmetricKey{
		Version: 0,
		Algorithm: pkix.AlgorithmIdentifier{
			Algorithm: ed25519OID,
		},
		PrivateKey: curvePrivateKey,
	}
	return asn1.Marshal(asym)
}

// ParseEd25519PrivateKey returns the Ed25519 private key encoded by the input.
func ParseEd25519PrivateKey(der []byte) (crypto.PrivateKey, error) {
	asym := new(oneAsymmetricKey)
	if rest, err := asn1.Unmarshal(der, asym); err != nil {
		return nil, err
	} else if len(rest) > 0 {
		return nil, errors.New("OneAsymmetricKey too long")
	}

	// Check that the key type is correct.
	if !asym.Algorithm.Algorithm.Equal(ed25519OID) {
		return nil, errEd25519WrongID
	}

	// Unmarshal the inner CurvePrivateKey.
	seed := new(curvePrivateKey)
	if rest, err := asn1.Unmarshal(asym.PrivateKey, seed); err != nil {
		return nil, err
	} else if len(rest) > 0 {
		return nil, errors.New("CurvePrivateKey too long")
	}

	return ed25519.NewKeyFromSeed(*seed), nil
}

func ToPem(sk crypto.PrivateKey, path string) error {
	var (
		err   error
		b     []byte
		block *pem.Block
		priv  ed25519.PrivateKey
	)

	priv, ok := sk.(ed25519.PrivateKey)
	if !ok {
		return errEd25519WrongKeyType
	}
	b, err = x509.MarshalPKCS8PrivateKey(priv)

	block = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: b,
	}

	err = ioutil.WriteFile(path, pem.EncodeToMemory(block), 0600)
	if err != nil {
		return err
	}
	return nil
}

func FromPem(file string) (ed25519.PrivateKey, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(b)
	if block.Type != "PRIVATE KEY" || block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}
	sk, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return sk.(ed25519.PrivateKey), nil
}

func New(anonymous bool, pkBytes []byte) *Identity {
	if anonymous == true {
		return &Identity{
			Anonymous: anonymous,
		}
	}
	privKey := ed25519.NewKeyFromSeed(pkBytes)
	//fmt.Println(privKey)
	pubKey := privKey.Public()

	//fmt.Println(pubKey)

	return &Identity{
		anonymous,
		privKey,
		pubKey,
	}
}

type Identity struct {
	Anonymous bool
	PriKey    ed25519.PrivateKey
	PubKey    crypto.PublicKey
}

func (identity *Identity) Sign(m []byte) ([]byte, error) {

	if identity.Anonymous == true {
		return []byte{}, nil
	}

	sign := ed25519.Sign(identity.PriKey, m[:])

	return sign, nil
}

func (identity *Identity) PubKeyBytes() []byte {
	var senderPubKey []byte
	if identity.Anonymous == false {
		pkBytes, _ := MarshalEd25519PublicKey(identity.PubKey)
		return pkBytes
	}
	return senderPubKey
}
