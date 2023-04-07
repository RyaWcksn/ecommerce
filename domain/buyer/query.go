package buyer

// Query for buyer
const (
	GetPasswordByEmailQuery = "SELECT id,email,password from ecommerce.buyer where email = ?"
)
