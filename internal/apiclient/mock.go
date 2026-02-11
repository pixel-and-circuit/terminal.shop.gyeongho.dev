package apiclient

import (
	"context"

	"mushroom.gyeongho.dev/internal/model"
)

// MockClient is an in-memory client for tests and development (no server).
type MockClient struct{}

// GetProducts returns static mushroom products.
func (MockClient) GetProducts(ctx context.Context) ([]model.Product, error) {
	return []model.Product{
		{
			ID:          "1",
			Name:        "Shiitake",
			Attributes:  []string{"fresh", "organic"},
			Price:       12.50,
			Description: "Rich, savory shiitake mushrooms.",
			Quantity:    100,
		},
		{
			ID:          "2",
			Name:        "Oyster",
			Attributes:  []string{"whole", "dried"},
			Price:       8.00,
			Description: "Mild oyster mushrooms, great for cooking.",
			Quantity:    50,
		},
	}, nil
}

// GetAbout returns static store info.
func (MockClient) GetAbout(ctx context.Context) (model.StoreInfo, error) {
	return model.StoreInfo{
		ID:    "about",
		Title: "About the store",
		Body:  "Welcome to the mushroom farm. We sell fresh and dried mushrooms. SSH store at mushroom.gyeongho.dev.",
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
