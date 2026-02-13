package model

// Product represents a sellable product.
type Product struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Attributes  []string `json:"attributes"`
	Price       float64  `json:"price"`
	Description string   `json:"description"`
	Quantity    int      `json:"quantity"`
}
