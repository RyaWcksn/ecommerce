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
	CreateProductEndpoint   = "/api/v1/seller/create"
	ListProductEndpoint     = "/api/v1/seller/products"
	ListOrderSellerEndpoint = "/api/v1/seller/orders"

	// Buyer
	CreateOrderEndpoint = "/api/v1/order"
)

// Header
const (
	HeaderContentType   = "Content-Type"
	HeaderAccept        = "Accept"
	MIMEApplicationJson = "application/json"
)

// Status
const (
	PendingMessage  = "Waiting Seller to accept order"
	Pending         = "PENDING"
	AcceptedMessage = "Seller is accepted the order"
	Accepted        = "ACCEPTED"
)
