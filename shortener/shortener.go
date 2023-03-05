package shortener

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"math/big"
)

func getSha256Hash(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encode(input []byte) (string, error) {
	encoding := base58.BitcoinEncoding

	encoded, err := encoding.Encode(input)

	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func GenerateShortUrl(input string, uuid string) (string, error) {
	sha256Hash := getSha256Hash(input + uuid)

	seedNumber := new(big.Int).SetBytes(sha256Hash).Uint64()

	encoded, err := base58Encode([]byte(fmt.Sprintf("%d", seedNumber)))

	if err != nil {
		return "", err
	}

	return encoded[:8], nil
}
