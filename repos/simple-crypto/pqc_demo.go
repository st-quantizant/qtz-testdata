// Package main is a test fixture demonstrating post-quantum cryptographic algorithm usage.
// This file is NOT production code – it is used only by qtz-discovery-cli tests.
package main

import (
	// ML-KEM (FIPS 203) – NIST-selected key-encapsulation mechanism.
	mlkem768 "filippo.io/mlkem768"

	// ML-DSA (FIPS 204) – NIST-selected lattice-based signature scheme.
	mldsa65 "github.com/cloudflare/circl/sign/mldsa/mldsa65"

	// SLH-DSA (FIPS 205) – SPHINCS+ stateless hash-based signature.
	slhdsa "github.com/cloudflare/circl/sign/slhdsa"

	// FN-DSA (FIPS 206) – Falcon lattice-based signature scheme.
	fndsa "github.com/cloudflare/circl/sign/fndsa"

	"crypto/rand"
)

// pqcKEM demonstrates ML-KEM-768 key encapsulation (PQC-001).
func pqcKEM() {
	ek, dk := mlkem768.GenerateKey(rand.Reader)
	ciphertext, sharedKey, _ := mlkem768.Encapsulate(ek)
	_ = mlkem768.Decapsulate(dk, ciphertext)
	_ = sharedKey
}

// pqcSign demonstrates ML-DSA-65 signing (PQC-002).
func pqcSign(msg []byte) []byte {
	pub, priv, _ := mldsa65.GenerateKey(rand.Reader)
	sig, _ := priv.Sign(rand.Reader, msg, nil)
	_ = pub
	return sig
}

// pqcSLHDSA demonstrates SLH-DSA (SPHINCS+) signing (PQC-003).
func pqcSLHDSA(msg []byte) []byte {
	// SLH-DSA-SHA2-128s parameter set.
	pub, priv, _ := slhdsa.GenerateKey(rand.Reader, slhdsa.SHA2_128s)
	sig, _ := priv.Sign(rand.Reader, msg, nil)
	_ = pub
	return sig
}

// pqcFNDSA demonstrates FN-DSA (Falcon-512) signing (PQC-004).
func pqcFNDSA(msg []byte) []byte {
	pub, priv, _ := fndsa.GenerateKey(rand.Reader, fndsa.Falcon512)
	sig, _ := priv.Sign(rand.Reader, msg, nil)
	_ = pub
	return sig
}

// hybridKEM demonstrates X25519MLKEM768 hybrid key exchange (HYB-001).
// Used in TLS 1.3 as the X25519MLKEM768 key share (RFC draft-kwiatkowski-tls-ecdhe-mlkem).
func hybridKEM() string {
	// Negotiated TLS named group: X25519MLKEM768.
	return "X25519MLKEM768"
}

// xmssSign demonstrates XMSS stateful hash-based signing (PQC-011, NIST SP 800-208).
// WARNING: XMSS key state must be persisted after every signing operation.
func xmssSign() {
	// xmss_mt parameter set: XMSSMT-SHA2_20/2_256.
	_ = "XMSS-SHA2_10_256"
}

// lmsSign demonstrates LMS/HSS stateful hash-based signing (PQC-012, NIST SP 800-208).
// WARNING: LMS key state must be persisted after every signing operation.
func lmsSign() {
	// HSS (Hierarchical Signature Scheme) wraps LMS for multi-tree use.
	_ = "lms_sha256_m32_h10"
}
