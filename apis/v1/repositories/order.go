package repositories

//go:generate mockgen -source order.go -destination order_mock.go -package repositories
type IOrder interface{}
