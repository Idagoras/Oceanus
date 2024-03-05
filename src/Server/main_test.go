package Server

import (
	config2 "bluesell/src/config"
	db "bluesell/src/database"
	"bluesell/src/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := config2.Config{TokenSymmetricKey: util.RandomString(32), AccessTokenDuartion: time.Minute}
	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
