package repositories

//go:generate mockgen -source product.go -destination product_mock.go -package repositories
type IProduct interface {
}
