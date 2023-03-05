package shortener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const uuid = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestGenerateShortUrl(t *testing.T) {
	var generated1, err1 = GenerateShortUrl("https://google.com", uuid)

	assert.Nil(t, err1)
	assert.Equal(t, "SrqwF4nH", generated1)

	var generated2, err2 = GenerateShortUrl("https://facebook.com", uuid)

	assert.Nil(t, err2)
	assert.Equal(t, "eNZPCSio", generated2)

	var generated3, err3 = GenerateShortUrl("https://github.com", uuid)

	assert.Nil(t, err3)
	assert.Equal(t, "cCd1qpQg", generated3)
}
