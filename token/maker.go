package token

import "time"

type Maker interface {
	createToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
