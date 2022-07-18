package shortener

import (
	"crypto/sha256"
	"encoding/base64"
	"math/big"
	"strconv"
)

func ShortenURL(original string, reqId string) string {
	// using sha256 to encrypt the URL
	alg := sha256.New()
	alg.Write([]byte(original + reqId))

	hashedUrl := alg.Sum(nil)
	generatedNumber := new(big.Int).SetBytes(hashedUrl).Uint64()

	// base64 encoding the URL from a generated number to store in db
	encodedUrl := base64.StdEncoding.EncodeToString([]byte(strconv.FormatUint(generatedNumber, 10)))

	// shortening to 10 characters and returning
	return encodedUrl[:10]
}
