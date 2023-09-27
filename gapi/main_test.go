package gapi

import (
	"testing"
	"time"

	db "github.com/jasonliuhe/simplebank/db/sqlc"
	"github.com/jasonliuhe/simplebank/util"
	"github.com/jasonliuhe/simplebank/worker"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}