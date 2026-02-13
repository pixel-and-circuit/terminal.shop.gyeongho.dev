package apiclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"shop.gyeongho.dev/internal/model"
)

// HTTPClient calls shop.gyeongho.dev/api (or base URL from env).
// If UserID is set (or MUSHROOM_USER_ID env), it sends X-User-Id on every request for user-scoped data.
type HTTPClient struct {
	BaseURL string
	Client  *http.Client
	// UserID is sent as X-User-Id header when non-empty (e.g. SSH key fingerprint from gateway).
	UserID string
}

// NewHTTPClient returns a client for the given base URL (e.g. https://shop.gyeongho.dev/api).
// UserID is set from MUSHROOM_USER_ID when empty so the gateway can inject identity via env.
func NewHTTPClient(baseURL string) *HTTPClient {
	if baseURL == "" {
		baseURL = os.Getenv("MUSHROOM_API_BASE")
		if baseURL == "" {
			baseURL = "https://shop.gyeongho.dev/api"
		}
	}
	userID := os.Getenv("MUSHROOM_USER_ID")
	return &HTTPClient{BaseURL: baseURL, Client: http.DefaultClient, UserID: userID}
}

func (c *HTTPClient) setUserIDHeader(req *http.Request) {
	if c.UserID != "" {
		req.Header.Set("X-User-Id", c.UserID)
	}
}

// GetProducts implements Client.
func (c *HTTPClient) GetProducts(ctx context.Context) ([]model.Product, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseURL+"/products", nil)
	if err != nil {
		return nil, err
	}
	c.setUserIDHeader(req)
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
	c.setUserIDHeader(req)
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
	c.setUserIDHeader(req)
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
	c.setUserIDHeader(req)
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
	bodyBytes, _ := json.Marshal(body)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BaseURL+"/cart", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	c.setUserIDHeader(req)
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("cart: %s", resp.Status)
	}
	var out model.Cart
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// SubmitOrder implements Client.
func (c *HTTPClient) SubmitOrder(ctx context.Context, cart *model.Cart) (*model.Order, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BaseURL+"/orders", nil)
	if err != nil {
		return nil, err
	}
	c.setUserIDHeader(req)
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
