package gapi

import (
	"fmt"

	db "github.com/jasonliuhe/simplebank/db/sqlc"
	"github.com/jasonliuhe/simplebank/pb"
	"github.com/jasonliuhe/simplebank/token"
	"github.com/jasonliuhe/simplebank/util"
	"github.com/jasonliuhe/simplebank/worker"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	pb.UnimplementedSimplebankServer
	config util.Config
	store db.Store
	tokenMaker token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config: config,
		store: store,
		tokenMaker: tokenMaker,
		taskDistributor: taskDistributor,
	}
	
	return server, nil
}