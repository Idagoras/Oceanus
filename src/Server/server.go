package Server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"oceanus/src/config"
	accountController "oceanus/src/controller/account"
	userController "oceanus/src/controller/user"
	db "oceanus/src/database"
	"oceanus/src/middleware"
	"oceanus/src/token"
	valid "oceanus/src/validator/currency"
)

type Server struct {
	config     config.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(config config.Config, store db.Store) (*Server, error) {
	server := &Server{store: store, config: config}
	router := gin.Default()
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey, "", "app")
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server.tokenMaker = tokenMaker
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", valid.CurrencyValid)
	}
	server.router = router
	server.initRouter()
	return server, nil
}
func (server *Server) initRouter() *Server {
	r := server.router

	r.POST("/users/signup", server.createUser)
	r.POST("/users/login", server.loginUser)

	authRoutes := r.Group("/").Use(server.authorization)
	authRoutes.POST("/account/add", server.createAccount)
	authRoutes.GET("/account/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)
	authRoutes.POST("/transfers", server.createAccounts)

	return server
}

func (server *Server) Router() *gin.Engine {
	return server.router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) createAccount(ctx *gin.Context) {
	accountController.CreateAccount(server.store, ctx)
}

func (server *Server) getAccount(ctx *gin.Context) {
	accountController.GetAccount(server.store, ctx)
}

func (server *Server) listAccounts(ctx *gin.Context) {
	accountController.ListAccounts(server.store, ctx)
}

func (server *Server) createAccounts(ctx *gin.Context) {
	accountController.CreateAccount(server.store, ctx)
}
func (server *Server) createUser(ctx *gin.Context) {
	userController.CreateUser(server.store, ctx)
}

func (server *Server) loginUser(ctx *gin.Context) {
	userController.LoginUser(server.store, ctx, server.tokenMaker, server.config)
}

func (server *Server) authorization(ctx *gin.Context) {
	middleware.Authorization(server.tokenMaker)(ctx)
}
