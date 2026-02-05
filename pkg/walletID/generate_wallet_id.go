package walletid

import (
	"crypto/rand"
	"math/big"
)

func GenerateWalletID(prefix string) (string, error) {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 16

	bytes := make([]byte, length)
	for i := range bytes {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		bytes[i] = chars[num.Int64()]
	}

	return prefix + "_" + string(bytes), nil
}
