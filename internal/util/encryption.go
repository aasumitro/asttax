package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
)

func NormalizeKey(key string) []byte {
	hash := sha256.Sum256([]byte(key)) // Produces a 32-byte key
	return hash[:]
}

func Encrypt(data, passphrase string) (string, error) {
	key := sha256.Sum256([]byte(passphrase))
	block, err := aes.NewCipher(key[:]) // Use normalized key
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nonce, nonce, []byte(data), nil)
	return hex.EncodeToString(ciphertext), nil
}

func Decrypt(encrypted, passphrase string) (string, error) {
	key := sha256.Sum256([]byte(passphrase))
	data, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key[:]) // Use normalized key
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("ciphertext too short")
	}
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	return string(plaintext), err
}

// get account from secret
// secretKey := string(util.NormalizeKey(srv.secretKey))
//		privateKeyDecrypted, err := util.Decrypt(user.PrivateKey, secretKey)
//		if err != nil {
//			fmt.Println("error decrypting private key")
//			return nil, err
//		}
//		privateKeyBytes, err := base58.Decode(privateKeyDecrypted)
//		if err != nil {
//			fmt.Println("error decode private key")
//		}
//		account, err := types.AccountFromBytes(privateKeyBytes)
//		if err != nil {
//			fmt.Println("error get account from base58")
//		}
//		fmt.Println(user.WalletAddress)
//		fmt.Println(account.PublicKey.ToBase58())
