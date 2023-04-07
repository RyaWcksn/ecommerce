package constants

// Role const
const (
	BUYER  = "buyer"
	SELLER = "seller"
)

// Endpoints
const (
	LoginEndpoint = "/api/v1/login"

	// Seller
	CreateProductEndpoint = "/api/v1/seller/create"
	ListProductEndpoint   = "/api/v1/seller/products"
)

// Header
const (
	HeaderContentType   = "Content-Type"
	HeaderAccept        = "Accept"
	MIMEApplicationJson = "application/json"
)
