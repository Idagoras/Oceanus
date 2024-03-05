package token

import (
	"bluesell/src/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPasetoMakerCreateToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32), util.RandomString(32), util.RandomString(10))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.Equal(t, payload.Username, username)
	require.WithinDuration(t, payload.IssuedTime, issuedAt, time.Second)
	require.WithinDuration(t, payload.Expiration, expiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32), util.RandomString(32), util.RandomString(10))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
