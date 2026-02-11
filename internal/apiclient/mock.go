package apiclient

import (
	"context"

	"mushroom.gyeongho.dev/internal/model"
)

// MockClient is an in-memory client for tests and development (no server).
type MockClient struct{}

// GetProducts returns static mushroom products (Mushroom Department Store lineup).
func (MockClient) GetProducts(ctx context.Context) ([]model.Product, error) {
	return []model.Product{
		{
			ID:          "1",
			Name:        "Oyster Mushroom",
			Attributes:  []string{"fresh"},
			Price:       800,
			Description: "Oyster mushroom (Neutari), per 100g.",
			Quantity:    100,
		},
		{
			ID:          "2",
			Name:        "Enoki Mushroom",
			Attributes:  []string{"fresh"},
			Price:       350,
			Description: "Enoki mushroom (Paeng-i), per 100g.",
			Quantity:    100,
		},
		{
			ID:          "3",
			Name:        "Button Mushroom",
			Attributes:  []string{"fresh"},
			Price:       2000,
			Description: "Button mushroom (Yang song-i), per 100g.",
			Quantity:    100,
		},
		{
			ID:          "4",
			Name:        "King Oyster Mushroom",
			Attributes:  []string{"fresh"},
			Price:       300,
			Description: "King oyster mushroom (Sae song-i), per 100g.",
			Quantity:    100,
		},
		{
			ID:          "5",
			Name:        "Shiitake Mushroom",
			Attributes:  []string{"fresh"},
			Price:       700,
			Description: "Shiitake mushroom (Pyogo), per 100g.",
			Quantity:    100,
		},
	}, nil
}

// GetAbout returns static store info (Mushroom Department Store).
func (MockClient) GetAbout(ctx context.Context) (model.StoreInfo, error) {
	return model.StoreInfo{
		ID:    "about",
		Title: "Mushroom Department Store",
		Body:  "Welcome to Mushroom Department Store (Buseot Baekhwajeom). We sell fresh mushrooms by weight. SSH store at mushroom.gyeongho.dev.",
	}, nil
}

// GetFAQ returns static FAQ entries.
func (MockClient) GetFAQ(ctx context.Context) ([]model.FAQEntry, error) {
	return []model.FAQEntry{
		{ID: "faq-1", Question: "How do I order?", Answer: "Use key (a) for Shop, add to cart, then (c) for cart and checkout."},
		{ID: "faq-2", Question: "What payment do you accept?", Answer: "Payment options are configured at checkout."},
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
