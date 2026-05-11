"""Sample Python file demonstrating weak cryptographic usage detected by tree-sitter AST."""

import hashlib
from Crypto.Cipher import DES, ARC4
from Crypto.PublicKey import RSA


def weak_hash_md5(data: bytes) -> bytes:
    # Weak: MD5 — AST should flag this
    h = hashlib.md5()
    h.update(data)
    return h.digest()


def weak_hash_sha1_new(data: bytes) -> bytes:
    # Weak: SHA-1 via hashlib.new — AST should flag this
    h = hashlib.new("sha1")
    h.update(data)
    return h.digest()


def strong_hash_sha256(data: bytes) -> bytes:
    # Safe: SHA-256
    h = hashlib.sha256()
    h.update(data)
    return h.digest()


def weak_cipher_des(key: bytes, plaintext: bytes) -> bytes:
    # Weak: DES — AST should flag this
    cipher = DES.new(key, DES.MODE_ECB)
    return cipher.encrypt(plaintext)


def weak_cipher_rc4(key: bytes, plaintext: bytes) -> bytes:
    # Weak: RC4 — AST should flag this
    cipher = ARC4.new(key)
    return cipher.encrypt(plaintext)


def weak_rsa_keygen() -> object:
    # Weak: RSA key generation — AST should flag this
    key = RSA.generate(2048)
    return key


# ── Post-Quantum Cryptography ──────────────────────────────────────────────────
# These demonstrate PQC algorithm usage detected by qtz-discovery-cli (PQC-001–012).

import oqs  # liboqs-python — open-quantum-safe


def pqc_ml_kem() -> bytes:
    # PQC-001: ML-KEM-768 (FIPS 203) – NIST-selected KEM.
    kem = oqs.KeyEncapsulation('ML-KEM-768')
    public_key = kem.generate_keypair()
    ciphertext, shared_secret = kem.encap_secret(public_key)
    return shared_secret


def pqc_ml_dsa(msg: bytes) -> bytes:
    # PQC-002: ML-DSA-65 (FIPS 204) – NIST-selected lattice signature.
    signer = oqs.Signature('ML-DSA-65')
    public_key = signer.generate_keypair()
    return signer.sign(msg)


def pqc_slh_dsa(msg: bytes) -> bytes:
    # PQC-003: SLH-DSA / SPHINCS+-SHA2-128f (FIPS 205) – stateless hash-based signature.
    signer = oqs.Signature('SPHINCS+-SHA2-128f-simple')
    signer.generate_keypair()
    return signer.sign(msg)


def pqc_fn_dsa(msg: bytes) -> bytes:
    # PQC-004: FN-DSA / Falcon-512 (FIPS 206) – lattice-based signature.
    signer = oqs.Signature('Falcon-512')
    signer.generate_keypair()
    return signer.sign(msg)


def pqc_xmss(msg: bytes) -> bytes:
    # PQC-011: XMSS stateful hash-based signature (NIST SP 800-208).
    # WARNING: persist key state after every sign to avoid reuse.
    signer = oqs.Signature('XMSS-SHA2_10_256')
    signer.generate_keypair()
    return signer.sign(msg)


def pqc_lms(msg: bytes) -> bytes:
    # PQC-012: LMS/HSS stateful hash-based signature (NIST SP 800-208).
    # HSS = Hierarchical Signature Scheme built on LMS.
    signer = oqs.Signature('lms_sha256_m32_h10')
    signer.generate_keypair()
    return signer.sign(msg)


def hybrid_kem_tls() -> str:
    # HYB-001: X25519MLKEM768 hybrid key exchange for TLS 1.3.
    # Combines X25519 classical security with ML-KEM-768 post-quantum security.
    return 'X25519MLKEM768'
