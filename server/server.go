package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/RyaWcksn/ecommerce/apis/v1/handlers"
	"github.com/RyaWcksn/ecommerce/apis/v1/services"
	"github.com/RyaWcksn/ecommerce/configs"
	"github.com/RyaWcksn/ecommerce/constants"
	"github.com/RyaWcksn/ecommerce/domain/buyer"
	"github.com/RyaWcksn/ecommerce/domain/order"
	"github.com/RyaWcksn/ecommerce/domain/product"
	"github.com/RyaWcksn/ecommerce/domain/seller"
	"github.com/RyaWcksn/ecommerce/pkgs/database"
	"github.com/RyaWcksn/ecommerce/pkgs/logger"
	"github.com/RyaWcksn/ecommerce/server/middleware"
)

type Server struct {
	cfg     *configs.Config
	log     logger.ILogger
	service services.IService
	handler handlers.IHandler
}

var addr string
var SVR *Server
var db *sql.DB
var signalChan chan (os.Signal) = make(chan os.Signal, 1)

func (s *Server) initServer() {
	addr = ":9000"
	cfg := s.cfg
	if len(cfg.Server.HTTPAddress) > 0 {
		if _, err := strconv.Atoi(cfg.Server.HTTPAddress); err == nil {
			addr = fmt.Sprintf(":%v", cfg.Server.HTTPAddress)
		} else {
			addr = cfg.Server.HTTPAddress
		}
	}
}

func (s *Server) Register() {

	s.initServer()

	// MYSQL
	dbConn := database.NewDatabaseConnection(*s.cfg, s.log)
	if dbConn == nil {
		s.log.Fatal("Expecting DB connection but received nil")
	}

	db = dbConn.DBConnect()
	if db == nil {
		s.log.Fatal("Expecting DB connection but received nil")
	}

	buyerImpl := buyer.NewBuyerImpl(db, s.log)
	sellerImpl := seller.NewSellerImpl(db, s.log)
	productImpl := product.NewProductImpl(db, s.log)
	orderImpl := order.NewOrderImpl(db, s.log)

	// Register service
	s.service = services.NewServiceImpl().
		WithBuyer(buyerImpl).
		WithSeller(sellerImpl).
		WithProduct(productImpl).
		WithOrder(orderImpl).
		WithConfig(*s.cfg).
		WithLog(s.log)

	// Register handler
	s.handler = handlers.NewHandlerImpl(s.service, s.log)
}

func NewService(cfg *configs.Config, logger logger.ILogger) *Server {
	if SVR != nil {
		return SVR
	}
	SVR = &Server{
		cfg: cfg,
		log: logger,
	}

	SVR.Register()

	return SVR
}

func (s Server) Start() {

	http.Handle(constants.LoginEndpoint, middleware.ErrHandler(s.handler.LoginHandler))

	// Seller
	http.Handle(constants.CreateProductEndpoint, middleware.AuthrorizationMiddleware(constants.SELLER, *s.cfg, s.handler.CreateProductHandler))
	http.Handle(constants.ListProductEndpoint, middleware.AuthrorizationMiddleware(constants.SELLER, *s.cfg, s.handler.GetProductListsHandler))
	http.Handle(constants.ListOrderSellerEndpoint, middleware.AuthrorizationMiddleware(constants.SELLER, *s.cfg, s.handler.GetSellerOrdersHandler))
	http.Handle(constants.AcceptOrderEndpoint, middleware.AuthrorizationMiddleware(constants.SELLER, *s.cfg, s.handler.AcceptOrderHandler))

	// Buyer
	http.Handle(constants.CreateOrderEndpoint, middleware.AuthrorizationMiddleware(constants.BUYER, *s.cfg, s.handler.CreateOrderHandler))
	http.Handle(constants.OrderListEndpoint, middleware.AuthrorizationMiddleware(constants.BUYER, *s.cfg, s.handler.GetBuyerOrdersHandler))

	go func() {
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			s.log.Fatalf("error listening to address %v, err=%v", addr, err)
		}
		s.log.Infof("HTTP server started %v", addr)
	}()

	sig := <-signalChan
	s.log.Infof("%s signal caught", sig)

	// Doing cleanup if received signal from Operating System.
	err := db.Close()
	if err != nil {
		s.log.Errorf("Error in closing DB connection. Err : %+v", err.Error())
	}
}
