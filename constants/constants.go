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
	AcceptOrderEndpoint     = "/api/v1/seller/order/accept"

	// Buyer
	CreateOrderEndpoint = "/api/v1/order"
	OrderListEndpoint   = "/api/v1/orders"
	ProductListEndpoint = "/api/v1/products"
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
