package seller

// Query for buyer
const (
	GetPasswordByEmailQuery = "SELECT id,email,password from ecommerce.seller where email = ?"
)
