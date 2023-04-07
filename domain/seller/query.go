package seller

// Query for buyer
const (
	GetPasswordByEmailQuery = "SELECT id,email,password from ecommerce.seller where email = ?"
	GetSellerDataById       = "SELECT name, email, alamat_pickup from ecommerce.seller WHERE id = ?"
)
