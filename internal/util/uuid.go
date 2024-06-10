package util

import (
	nanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateUUID(figures int) (string, error) {
	uuid, err := nanoid.Generate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", figures)
	if err != nil {
		return "", err
	}
	return uuid, nil
}
