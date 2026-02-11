package e2e

import (
	"context"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"mushroom.gyeongho.dev/internal/apiclient"
	"mushroom.gyeongho.dev/internal/model"
	"mushroom.gyeongho.dev/internal/testserver"
)

func fixtureDir(t *testing.T) string {
	t.Helper()
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Join(filepath.Dir(file), "fixtures")
	return dir
}

func TestE2E_GetProducts(t *testing.T) {
	dir := fixtureDir(t)
	srv := httptest.NewServer(testserver.NewFixtureHandler(dir))
	defer srv.Close()

	client := apiclient.NewHTTPClient(srv.URL)
	ctx := context.Background()
	products, err := client.GetProducts(ctx)
	if err != nil {
		t.Fatalf("GetProducts: %v", err)
	}
	if len(products) < 1 {
		t.Fatalf("expected at least one product, got %d", len(products))
	}
	if products[0].Name == "" || products[0].Price <= 0 {
		t.Errorf("first product should have Name and Price, got %+v", products[0])
	}
}

func TestE2E_GetAbout(t *testing.T) {
	dir := fixtureDir(t)
	srv := httptest.NewServer(testserver.NewFixtureHandler(dir))
	defer srv.Close()

	client := apiclient.NewHTTPClient(srv.URL)
	ctx := context.Background()
	about, err := client.GetAbout(ctx)
	if err != nil {
		t.Fatalf("GetAbout: %v", err)
	}
	if about.Title == "" {
		t.Errorf("about should have Title, got %+v", about)
	}
}

func TestE2E_GetFAQ(t *testing.T) {
	dir := fixtureDir(t)
	srv := httptest.NewServer(testserver.NewFixtureHandler(dir))
	defer srv.Close()

	client := apiclient.NewHTTPClient(srv.URL)
	ctx := context.Background()
	faq, err := client.GetFAQ(ctx)
	if err != nil {
		t.Fatalf("GetFAQ: %v", err)
	}
	if len(faq) < 1 {
		t.Fatalf("expected at least one FAQ entry, got %d", len(faq))
	}
}

func TestE2E_GetCart(t *testing.T) {
	dir := fixtureDir(t)
	srv := httptest.NewServer(testserver.NewFixtureHandler(dir))
	defer srv.Close()

	client := apiclient.NewHTTPClient(srv.URL)
	ctx := context.Background()
	cart, err := client.GetCart(ctx)
	if err != nil {
		t.Fatalf("GetCart: %v", err)
	}
	if cart == nil {
		t.Fatal("cart should not be nil")
	}
}

func TestE2E_AddToCart(t *testing.T) {
	dir := fixtureDir(t)
	srv := httptest.NewServer(testserver.NewFixtureHandler(dir))
	defer srv.Close()

	client := apiclient.NewHTTPClient(srv.URL)
	ctx := context.Background()
	cart, err := client.AddToCart(ctx, "1", 2)
	if err != nil {
		t.Fatalf("AddToCart: %v", err)
	}
	if cart == nil {
		t.Fatal("cart should not be nil")
	}
}

func TestE2E_SubmitOrder(t *testing.T) {
	dir := fixtureDir(t)
	srv := httptest.NewServer(testserver.NewFixtureHandler(dir))
	defer srv.Close()

	client := apiclient.NewHTTPClient(srv.URL)
	ctx := context.Background()
	cart, err := client.GetCart(ctx)
	if err != nil {
		t.Fatalf("GetCart: %v", err)
	}
	if cart == nil {
		cart = &model.Cart{}
	}
	if len(cart.Items) == 0 {
		cart.Items = []model.CartItem{
			{ProductID: "1", Name: "Oyster Mushroom", UnitPrice: 800, Quantity: 1},
		}
	}
	order, err := client.SubmitOrder(ctx, cart)
	if err != nil {
		t.Fatalf("SubmitOrder: %v", err)
	}
	if order.ID == "" || order.Status == "" {
		t.Errorf("order should have ID and Status, got %+v", order)
	}
}

func TestE2E_XUserIdSentAndEchoed(t *testing.T) {
	dir := fixtureDir(t)
	var lastUserID string
	base := testserver.NewFixtureHandler(dir)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lastUserID = r.Header.Get("X-User-Id")
		base.ServeHTTP(w, r)
	})
	srv := httptest.NewServer(handler)
	defer srv.Close()

	client := apiclient.NewHTTPClient(srv.URL)
	client.UserID = "e2e-test-user"
	ctx := context.Background()
	_, err := client.GetProducts(ctx)
	if err != nil {
		t.Fatalf("GetProducts: %v", err)
	}
	if lastUserID != "e2e-test-user" {
		t.Errorf("expected X-User-Id e2e-test-user, got %q", lastUserID)
	}
}
