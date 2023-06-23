package obfuscationprocessor

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/cyrildever/feistel"
	"github.com/cyrildever/feistel/common/utils/hash"
	"github.com/spf13/cast"
	"testing"
)

func TestFormatPreservingEncryption(t *testing.T) {
	source := 99 // 9 digits
	cipher := feistel.NewFPECipher(hash.SHA_256, "some-32-byte-long-key-to-be-safe", 128)

	obfuscated, err := cipher.EncryptNumber(cast.ToUint64(source))
	assert.NoError(t, err)
	assert.Equal(t, obfuscated.Uint64(), uint64(22375))
	fmt.Println(obfuscated.Uint64())
}
