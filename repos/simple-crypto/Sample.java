import java.security.MessageDigest;
import java.security.KeyPairGenerator;
import javax.crypto.Cipher;
import javax.crypto.SecretKeyFactory;

/**
 * Sample Java file demonstrating weak cryptographic usage detected by tree-sitter AST.
 */
public class Sample {

    // Weak: MD5 — AST should flag this
    public byte[] weakHashMd5(byte[] data) throws Exception {
        MessageDigest md = MessageDigest.getInstance("MD5");
        return md.digest(data);
    }

    // Weak: SHA-1 — AST should flag this
    public byte[] weakHashSha1(byte[] data) throws Exception {
        MessageDigest md = MessageDigest.getInstance("SHA-1");
        return md.digest(data);
    }

    // Safe: SHA-256
    public byte[] strongHashSha256(byte[] data) throws Exception {
        MessageDigest md = MessageDigest.getInstance("SHA-256");
        return md.digest(data);
    }

    // Weak: DES — AST should flag this
    public byte[] weakCipherDes(byte[] key, byte[] plaintext) throws Exception {
        Cipher cipher = Cipher.getInstance("DES/CBC/PKCS5Padding");
        return cipher.doFinal(plaintext);
    }

    // Weak: RSA key generation — AST should flag this
    public void weakRsaKeygen() throws Exception {
        KeyPairGenerator kpg = KeyPairGenerator.getInstance("RSA");
        kpg.initialize(2048);
    }

    // Weak: ECDSA key generation — AST should flag this
    public void weakEcdsaKeygen() throws Exception {
        KeyPairGenerator kpg = KeyPairGenerator.getInstance("EC");
        kpg.initialize(256);
    }

    // ── Post-Quantum Cryptography ─────────────────────────────────────────────
    // Bouncy Castle 1.78+ provides FIPS PQC implementations.

    // PQC-001: ML-KEM-768 (FIPS 203) – NIST-selected KEM.
    public void pqcMlKem() throws Exception {
        // MLKEMKeyPairGenerator uses ML-KEM-768 parameter set by default.
        KeyPairGenerator kpg = KeyPairGenerator.getInstance("ML-KEM", "BCPQC");
        kpg.initialize(new MLKEMParameterSpec(MLKEMParameterSpec.ml_kem_768));
        KeyPair keyPair = kpg.generateKeyPair();
    }

    // PQC-002: ML-DSA-65 (FIPS 204) – NIST-selected lattice signature.
    // Previously known as CRYSTALS-Dilithium (dilithium3 = ML-DSA-65).
    public void pqcMlDsa() throws Exception {
        KeyPairGenerator kpg = KeyPairGenerator.getInstance("ML-DSA", "BCPQC");
        kpg.initialize(MLDSAParameterSpec.ml_dsa_65);
    }

    // PQC-003: SLH-DSA (FIPS 205) – SPHINCS+ stateless hash-based signature.
    public void pqcSlhDsa() throws Exception {
        KeyPairGenerator kpg = KeyPairGenerator.getInstance("SLH-DSA", "BCPQC");
        kpg.initialize(SLHDSAParameterSpec.slh_dsa_sha2_128s);
    }

    // PQC-004: FN-DSA / Falcon-512 (FIPS 206) – lattice-based signature.
    public void pqcFnDsa() throws Exception {
        KeyPairGenerator kpg = KeyPairGenerator.getInstance("Falcon", "BCPQC");
        kpg.initialize(FalconParameterSpec.falcon_512);
    }

    // PQC-011: XMSS stateful hash-based signature (NIST SP 800-208).
    // WARNING: persist XMSS key state after every sign to prevent reuse.
    public void pqcXmss() throws Exception {
        KeyPairGenerator kpg = KeyPairGenerator.getInstance("XMSS", "BCPQC");
        kpg.initialize(new XMSSParameterSpec(10, "SHA256"));
    }

    // PQC-012: LMS/HSS stateful hash-based signature (NIST SP 800-208).
    // HSS (Hierarchical Signature Scheme) wraps LMS for larger signing capacity.
    public void pqcLms() throws Exception {
        KeyPairGenerator kpg = KeyPairGenerator.getInstance("LMS", "BCPQC");
        // HSS two-level tree over LMS-SHA256-M32-H10 leaves.
    }

    // HYB-001: X25519MLKEM768 hybrid key exchange for TLS 1.3.
    // Combines classical X25519 with ML-KEM-768 post-quantum security.
    public String hybridKemGroup() {
        return "X25519MLKEM768"; // TLS named group per draft-kwiatkowski-tls-ecdhe-mlkem
    }
}
