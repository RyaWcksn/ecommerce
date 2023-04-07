package product

const (
	InsertProduct     = "INSERT INTO ecommerce.product (product_name, description, price, seller) VALUES (?, ?, ?, ?)"
	GetProductsSeller = "SELECT * FROM ecommerce.product where seller = ?"
	GetProductById    = "SELECT id, product_name, description, price, seller FROM ecommerce.seller WHERE id = ?"
)
