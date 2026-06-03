# openssl gem: Ruby bindings to OpenSSL
require 'openssl'

module OpenSSLHelper
  def self.sha256(data)
    OpenSSL::Digest::SHA256.digest(data)
  end

  def self.rsa_key(bits = 2048)
    OpenSSL::PKey::RSA.new(bits)
  end

  def self.aes_encrypt(key, data)
    cipher = OpenSSL::Cipher.new('AES-256-CBC')
    cipher.encrypt
    cipher.key = key
    cipher.update(data) + cipher.final
  end
end
