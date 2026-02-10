package model

import "time"

// Order is a submitted purchase (checkout result).
type Order struct {
	ID        string     `json:"id"`
	Items     []CartItem `json:"items"`
	Total     float64    `json:"total"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
}
