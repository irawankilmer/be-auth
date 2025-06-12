package idgen

import (
	"crypto/rand"
	"github.com/oklog/ulid/v2"
	"time"
)

func NewULID() string {
	t := time.Now().UTC()
	return ulid.MustNew(ulid.Timestamp(t), ulid.Monotonic(rand.Reader, 0)).String()
}
