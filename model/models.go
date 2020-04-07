package model

// Person is your comment
type Tenant struct {
	SocialSecurity string `json:"ss,omitempty" bson:"ss,omitempty"`
	FirstName      string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName       string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

// Property is yourcomment.
type Property struct {
	ID       string `json:"id,omitempty" bson:"id,omitempty"`
	Address  string `json:"address,omitempty" bson:"address,omitempty"`
	Zipcode  string `json:"zipcode,omitempty" bson:"zipcode,omitempty"`
	Price    string `json:"price,omitempty" bson:"price,omitempty"`
	Category string `json:"category,omitempty" bson:"category,omitempty"`
}
