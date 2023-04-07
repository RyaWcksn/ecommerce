package entities

type CreateOrder struct {
	Buyer               int
	Seller              int
	DeliverySource      string
	DeliveryDestination string
	Items               string
	Quantity            int
	Price               string
	TotalPrice          string
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
	Price               string
	TotalPrice          string
	Status              OrderStatus
}
