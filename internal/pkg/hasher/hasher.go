package hasher

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Hash(obj string, cost int) (string, error) {

	if cost == 0 {
		cost = bcrypt.DefaultCost
	}

	objHash, err := bcrypt.GenerateFromPassword([]byte(obj), cost)
	if err != nil {
		return "", fmt.Errorf("failed to hash object - %v", err)
	}

	return string(objHash), nil

}
