package Server

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	config2 "oceanus/src/config"
	db "oceanus/src/database"
	"oceanus/src/util"
	"os"
	"testing"
	"time"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := config2.Config{TokenSymmetricKey: util.RandomString(32), AccessTokenDuration: time.Minute}
	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
