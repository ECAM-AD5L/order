package schema

type Customer struct {
	Email     string `json:"email" bson:"email"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
}

type ShippingAddress struct {
	Name    string `json:"name" bson:"name"`
	Street  string `json:"street" bson:"street"`
	City    string `json:"city" bson:"city"`
	Country string `json:"country" bson:"country"`
	ZipCode string `json:"zip_code" bson:"zip_code"`
}

type BillingAddress struct {
	Name    string `json:"name" bson:"name"`
	Street  string `json:"street" bson:"street"`
	City    string `json:"city" bson:"city"`
	Country string `json:"country" bson:"country"`
	ZipCode string `json:"zip_code" bson:"zip_code"`
}
