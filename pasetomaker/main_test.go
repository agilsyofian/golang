package pasetomaker

import (
	"log"
	"testing"
	"time"

	"github.com/agilsyofian/golang/util"
	"github.com/stretchr/testify/require"
)

func TestCretaToken(t *testing.T) {
	tokenSymmetricKey := "12345678901234567890123456789012"
	tokenMaker, err := NewPasetoMaker(tokenSymmetricKey)
	if err != nil {
		log.Fatal("cannot create token maker: %w", err)
	}

	var username string = util.RandomString(10)
	var duration time.Duration = 15 * time.Minute
	token, payload, err := tokenMaker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)
}
