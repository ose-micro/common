package common

type Address struct {
	Line1       string   `json:"line1" bson:"line1"`
	Line2       string   `json:"line2,omitempty" bson:"line2,omitempty"`
	City        string   `json:"city" bson:"city"`
	State       string   `json:"state,omitempty" bson:"state,omitempty"`
	PostalCode  string   `json:"postal_code,omitempty" bson:"postal_code,omitempty"`
	CountryCode string   `json:"country_code" bson:"country_code"`
	Latitude    *float64 `json:"latitude,omitempty" bson:"latitude,omitempty"`
	Longitude   *float64 `json:"longitude,omitempty" bson:"longitude,omitempty"`
}
