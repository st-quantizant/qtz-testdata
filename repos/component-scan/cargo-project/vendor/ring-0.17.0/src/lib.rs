// ring: safe, fast, small crypto
use ring::digest;
use ring::hmac;
use ring::rand::SystemRandom;
use ring::signature::{self, EcdsaKeyPair};

pub fn sha256_hash(data: &[u8]) -> digest::Digest {
    digest::digest(&digest::SHA256, data)
}

pub fn hmac_sign(key: &hmac::Key, msg: &[u8]) -> hmac::Tag {
    hmac::sign(key, msg)
}
