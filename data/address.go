package data

type Address struct {
	ID           int     `json:"id"`
	StreetName   string  `json:"street_name"`
	StreetNumber string  `json:"street_number"`
	Apartment    *string `json:"apartment"`
	City         string  `json:"city"`
	State        string  `json:"state"`
	ZipCode      string  `json:"zip_code"`
}
