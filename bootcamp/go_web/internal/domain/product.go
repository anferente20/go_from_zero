package domain

type Product struct {
	Id           int     `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	Quantity     int     `json:"quantity,omitempty"`
	Code_value   string  `json:"code___value,omitempty"`
	Is_published bool    `json:"is___published,omitempty"`
	Expiration   string  `json:"expiration,omitempty"`
	Price        float64 `json:"price,omitempty"`
}
