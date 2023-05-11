package pasetomaker

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/agilsyofian/golang/util"
	"github.com/stretchr/testify/require"
)

type TestPayload struct {
	Username string `json:"username"`
	Tahun    int    `json:"tahun"`
}

func TestCretaToken(t *testing.T) {
	tokenSymmetricKey := "12345678901234567890123456789012"
	tokenMaker, err := NewPasetoMaker(tokenSymmetricKey)
	if err != nil {
		log.Fatal("cannot create token maker: %w", err)
	}

	pyl := TestPayload{
		Username: util.RandomString(10),
		Tahun:    2022,
	}
	var duration time.Duration = 15 * time.Minute
	token, payload, err := tokenMaker.CreateToken(pyl, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)
}

func TestVerifyToken(t *testing.T) {
	tokenSymmetricKey := "12345678901234567890123456789012"
	tokenMaker, err := NewPasetoMaker(tokenSymmetricKey)
	if err != nil {
		log.Fatal("cannot create token maker: %w", err)
	}
	pyl := TestPayload{
		Username: util.RandomString(10),
		Tahun:    2022,
	}
	var duration time.Duration = 15 * time.Minute
	token, payload, err := tokenMaker.CreateToken(pyl, duration)
	require.NoError(t, err)

	testPayload, err := tokenMaker.VerifyToken(token)
	x := testPayload.Payload.(map[string]interface{})
	fmt.Println(x["username"])
	require.NoError(t, err)
	require.Equal(t, pyl.Tahun, int(x["tahun"].(float64)))
	require.Equal(t, pyl.Username, x["username"])
	require.NotEmpty(t, payload)
}
