package config

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateDeviceID(userAgent string, ipAddress string) (string, error) {
	rawDeviceID := userAgent + ipAddress

	// Hash the combined string using SHA-256
	hasher := sha256.New()
	_, err := hasher.Write([]byte(rawDeviceID))
	if err != nil {
		return "", err
	}

	// Get the hex string representation of the hash
	deviceID := hex.EncodeToString(hasher.Sum(nil))
	return deviceID, nil
}
