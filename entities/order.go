package entities

type CreateOrder struct {
	Buyer               int
	Seller              int
	DeliverySource      string
	DeliveryDestination string
	Items               string
	Quantity            int
	Price               float64
	TotalPrice          float64
	Status              int
}

type OrderStatus struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type Order struct {
	Id                  int
	Buyer               int
	Seller              int
	DeliverySource      string
	DeliveryDestination string
	Items               string
	Quantity            int
	Price               float64
	TotalPrice          float64
	Status              OrderStatus
}
