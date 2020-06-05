package db

import (
	"database/sql"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

// Epoch is the epoch for all Mantle ULIDs
var Epoch, _ = time.Parse("Jan 2 2006", "Jan 1 2020")

func newUUID() string {
	t := time.Unix(0, time.Now().UnixNano()-Epoch.UnixNano())
	var entropy = ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}

func IsUID(s string) bool {
	_, err := ulid.Parse(s)
	return err == nil
}

func now() Time {
	s := time.Now().UTC().String()[0:19]
	t, _ := time.Parse(timeFormat, s)
	return Time(t)
}

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func uHighLow(a, b int) (int, int) {
	if a >= b {
		return a, b
	}
	return b, a
}

func queryCount(rows *sql.Rows) int64 {
	s := int64(0)
	defer rows.Close()
	for rows.Next() {
		s++
	}
	return s
}
