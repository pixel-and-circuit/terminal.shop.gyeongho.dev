package apiclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"mushroom.gyeongho.dev/internal/model"
)

// HTTPClient calls mushroom.gyeongho.dev/api (or base URL from env).
type HTTPClient struct {
	BaseURL string
	Client  *http.Client
}

// NewHTTPClient returns a client for the given base URL (e.g. https://mushroom.gyeongho.dev/api).
func NewHTTPClient(baseURL string) *HTTPClient {
	if baseURL == "" {
		baseURL = os.Getenv("MUSHROOM_API_BASE")
		if baseURL == "" {
			baseURL = "https://mushroom.gyeongho.dev/api"
		}
	}
	return &HTTPClient{BaseURL: baseURL, Client: http.DefaultClient}
}

// GetProducts implements Client.
func (c *HTTPClient) GetProducts(ctx context.Context) ([]model.Product, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseURL+"/products", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("products: %s", resp.Status)
	}
	var out []model.Product
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetAbout implements Client.
func (c *HTTPClient) GetAbout(ctx context.Context) (model.StoreInfo, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseURL+"/about", nil)
	if err != nil {
		return model.StoreInfo{}, err
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return model.StoreInfo{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return model.StoreInfo{}, fmt.Errorf("about: %s", resp.Status)
	}
	var out model.StoreInfo
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return model.StoreInfo{}, err
	}
	return out, nil
}

// GetFAQ implements Client.
func (c *HTTPClient) GetFAQ(ctx context.Context) ([]model.FAQEntry, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseURL+"/faq", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("faq: %s", resp.Status)
	}
	var out []model.FAQEntry
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetCart implements Client.
func (c *HTTPClient) GetCart(ctx context.Context) (*model.Cart, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseURL+"/cart", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return &model.Cart{}, nil
	}
	var out model.Cart
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AddToCart implements Client.
func (c *HTTPClient) AddToCart(ctx context.Context, productID string, quantity int) (*model.Cart, error) {
	body := struct {
		ProductID string `json:"productId"`
		Quantity  int    `json:"quantity"`
	}{productID, quantity}
	// Simplified: no request body sent; real impl would POST JSON
	_ = body
	return c.GetCart(ctx)
}

// SubmitOrder implements Client.
func (c *HTTPClient) SubmitOrder(ctx context.Context, cart *model.Cart) (*model.Order, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BaseURL+"/orders", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("orders: %s", resp.Status)
	}
	var out model.Order
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}
