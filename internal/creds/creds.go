package creds

import (
	"fmt"
	"strings"

	"github.com/germandv/rememberpass/internal/argon"
)

type Credential struct {
	ID     string
	Secret string
}

func (c Credential) String() string {
	return fmt.Sprintf("%s:%s", c.ID, c.Secret)
}

func (c Credential) Compare(candidate string) bool {
	return argon.Compare(candidate, c.Secret)
}

func New(id string, secret string) (Credential, error) {
	if id == "" || secret == "" {
		return Credential{}, fmt.Errorf("ID and secret are required")
	}

	if strings.Contains(id, ":") {
		return Credential{}, fmt.Errorf("ID cannot contain ':'")
	}

	secret, err := argon.Hash(secret)
	if err != nil {
		return Credential{}, err
	}

	return Credential{ID: id, Secret: secret}, nil
}

func Parse(lines []string) []Credential {
	credentials := []Credential{}

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			panic("Corrupted file: " + line)
		}

		c := Credential{ID: parts[0], Secret: strings.TrimSpace(parts[1])}
		credentials = append(credentials, c)
	}

	return credentials
}
