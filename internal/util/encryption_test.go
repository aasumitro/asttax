package util_test

import (
	"testing"

	"github.com/aasumitro/asttax/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestNormalizeKey(t *testing.T) {
	key := "somepassphrase"
	normalizedKey := util.NormalizeKey(key)
	assert.Len(t, normalizedKey, 32, "Normalized key should be 32 bytes long")
	normalizedKey2 := util.NormalizeKey(key)
	assert.Equal(t, normalizedKey, normalizedKey2, "Normalization of the same key should return the same result")
	anotherKey := "anotherpassphrase"
	normalizedKey3 := util.NormalizeKey(anotherKey)
	assert.NotEqual(t, normalizedKey, normalizedKey3, "Different keys should return different normalized keys")
}

func TestEncryptDecrypt(t *testing.T) {
	// Test case 1: Normal encryption and decryption (success)
	passphrase := "supersecret"
	plaintext := "This is a test message."

	// Encrypt the plaintext
	encrypted, err := util.Encrypt(plaintext, passphrase)
	assert.NoError(t, err, "Encryption should succeed")
	assert.NotEmpty(t, encrypted, "Encrypted text should not be empty")

	// Decrypt the encrypted text
	decrypted, err := util.Decrypt(encrypted, passphrase)
	assert.NoError(t, err, "Decryption should succeed")
	assert.Equal(t, plaintext, decrypted, "Decrypted text should match original")

	// Test case 2: Decrypt with wrong passphrase (should fail)
	invalidPassphrase := "wrongpassphrase"
	_, err = util.Decrypt(encrypted, invalidPassphrase)
	assert.Error(t, err, "Decryption with an invalid passphrase should fail")

	// Test case 3: Decrypt with invalid encrypted data (should fail)
	invalidEncrypted := "invalid-encrypted-data"
	_, err = util.Decrypt(invalidEncrypted, passphrase)
	assert.Error(t, err, "Decryption with invalid encrypted data should fail")

	// Test case 4: Encrypt with empty string
	emptyText := ""
	encryptedEmpty, err := util.Encrypt(emptyText, passphrase)
	assert.NoError(t, err, "Encryption of an empty string should succeed")
	assert.NotEmpty(t, encryptedEmpty, "Encrypted text for empty string should not be empty")

	// Test case 5: Decrypt with empty encrypted string (should fail)
	_, err = util.Decrypt("", passphrase)
	assert.Error(t, err, "Decryption with empty encrypted text should fail")
}
