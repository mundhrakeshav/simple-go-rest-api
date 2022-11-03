package models

type Address struct {
	State   string `json:"state,omitempty" bson:"state"`
	City    string `json:"city,omitempty" bson:"city"`
	Pincode uint32 `json:"pincode,omitempty" bson:"pincode"`
}

type User struct {
	Name    string  `json:"name,omitempty" bson:"user_name"`
	Age     uint8   `json:"age,omitempty" bson:"user_age"`
	Address Address `json:"address,omitempty" bson:"user_address"`
}
