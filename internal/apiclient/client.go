package apiclient

import (
	"context"

	"mushroom.gyeongho.dev/internal/model"
)

// Client is the abstract API client for products, about, faq, cart, and orders.
// Implementations can be real HTTP or mock for tests.
type Client interface {
	GetProducts(ctx context.Context) ([]model.Product, error)
	GetAbout(ctx context.Context) (model.StoreInfo, error)
	GetFAQ(ctx context.Context) ([]model.FAQEntry, error)
	GetCart(ctx context.Context) (*model.Cart, error)
	AddToCart(ctx context.Context, productID string, quantity int) (*model.Cart, error)
	SubmitOrder(ctx context.Context, cart *model.Cart) (*model.Order, error)
}
