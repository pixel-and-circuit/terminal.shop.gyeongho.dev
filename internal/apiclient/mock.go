package apiclient

import (
	"context"

	"shop.gyeongho.dev/internal/model"
)

// MockClient is an in-memory client for tests and development (no server).
type MockClient struct{}

// GetProducts returns static products (shop catalog lineup).
func (MockClient) GetProducts(ctx context.Context) ([]model.Product, error) {
	return []model.Product{
		{
			ID:          "1",
			Name:        "Oyster",
			Attributes:  []string{"fresh"},
			Price:       800,
			Description: "Oyster (Neutari), per 100g.",
			Quantity:    100,
		},
		{
			ID:          "2",
			Name:        "Enoki",
			Attributes:  []string{"fresh"},
			Price:       350,
			Description: "Enoki (Paeng-i), per 100g.",
			Quantity:    100,
		},
		{
			ID:          "3",
			Name:        "Button",
			Attributes:  []string{"fresh"},
			Price:       2000,
			Description: "Button (Yang song-i), per 100g.",
			Quantity:    100,
		},
		{
			ID:          "4",
			Name:        "King Oyster",
			Attributes:  []string{"fresh"},
			Price:       300,
			Description: "King oyster (Sae song-i), per 100g.",
			Quantity:    100,
		},
		{
			ID:          "5",
			Name:        "Shiitake",
			Attributes:  []string{"fresh"},
			Price:       700,
			Description: "Shiitake (Pyogo), per 100g.",
			Quantity:    100,
		},
	}, nil
}

// GetAbout returns static store info (shop branding).
func (MockClient) GetAbout(ctx context.Context) (model.StoreInfo, error) {
	return model.StoreInfo{
		ID:    "about",
		Title: "shop.gyeongho.dev",
		Body:  "Welcome to shop.gyeongho.dev. We sell all products gyeongho provides: produce, embedded devices, robots, and more. SSH store at shop.gyeongho.dev.",
	}, nil
}

// GetFAQ returns static FAQ entries.
func (MockClient) GetFAQ(ctx context.Context) ([]model.FAQEntry, error) {
	return []model.FAQEntry{
		{ID: "faq-1", Question: "How do I order?", Answer: "Use key (a) for Shop, add to cart, then (c) for cart and checkout."},
		{ID: "faq-2", Question: "What payment do you accept?", Answer: "Payment options are configured at checkout."},
		{ID: "faq-3", Question: "What do you sell?", Answer: "We sell all products gyeongho provides: produce, embedded devices, robots, and more."},
	}, nil
}

// GetCart returns empty cart (mock does not persist).
func (MockClient) GetCart(ctx context.Context) (*model.Cart, error) {
	return &model.Cart{Items: nil}, nil
}

// AddToCart returns a cart with the item added (in-memory only for mock).
func (MockClient) AddToCart(ctx context.Context, productID string, quantity int) (*model.Cart, error) {
	prods, _ := (MockClient{}).GetProducts(ctx)
	for _, p := range prods {
		if p.ID == productID {
			return &model.Cart{
				Items: []model.CartItem{
					{ProductID: p.ID, Name: p.Name, UnitPrice: p.Price, Quantity: quantity},
				},
			}, nil
		}
	}
	return &model.Cart{Items: nil}, nil
}

// SubmitOrder returns a fake order.
func (MockClient) SubmitOrder(ctx context.Context, cart *model.Cart) (*model.Order, error) {
	if cart == nil || len(cart.Items) == 0 {
		return nil, nil
	}
	return &model.Order{
		ID:     "mock-order-1",
		Items:  cart.Items,
		Total:  cart.Total(),
		Status: "confirmed",
	}, nil
}
