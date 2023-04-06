package seller

// Query for buyer
const (
	GetPasswordByEmailQuery = "SELECT email,password from ecommerce.seller where email = ?"
)
