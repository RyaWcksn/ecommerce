package order

const (
	InsertOrder = "INSERT INTO orders (buyer, seller, delivery_source_address, delivery_destination_address, items, quantity, price, total_price, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

	GetBySellerId = "SELECT * FROM ecommerce.orders WHERE seller = ?"
)
