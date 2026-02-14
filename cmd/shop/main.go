package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"shop.gyeongho.dev/internal/apiclient"
	"shop.gyeongho.dev/internal/sshsrv"
	"shop.gyeongho.dev/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
)

const (
	sshPortDefault  = "2222"
	sshHostKeyPath  = ".ssh/id_ed25519"
	shutdownTimeout = 30 * time.Second
)

func main() {
	runSSH := flag.Bool("ssh", false, "run as SSH server instead of local TUI")
	flag.Parse()
	if !*runSSH && os.Getenv("SHOP_SSH") != "" {
		if v, _ := strconv.ParseBool(os.Getenv("SHOP_SSH")); v {
			*runSSH = true
		}
	}

	if *runSSH {
		runSSHServer()
		return
	}
	runLocalTUI()
}

func runSSHServer() {
	addr := net.JoinHostPort("", sshPortDefault)
	s, err := sshsrv.NewServer(addr, sshHostKeyPath, sshsrv.ShopHandler)
	if err != nil {
		fmt.Fprintf(os.Stderr, "SSH server: %v\n", err)
		os.Exit(1)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			fmt.Fprintf(os.Stderr, "SSH server: %v\n", err)
			done <- nil
		}
	}()

	<-done
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		fmt.Fprintf(os.Stderr, "SSH shutdown: %v\n", err)
	}
}

func runLocalTUI() {
	client := apiclient.MockClient{}
	m := tui.NewModel(client)
	ctx := context.Background()
	if prods, err := client.GetProducts(ctx); err == nil {
		m.Products = prods
	} else {
		m.Error = err.Error()
	}
	if about, err := client.GetAbout(ctx); err == nil {
		m.About = about
	} else if m.Error == "" {
		m.Error = err.Error()
	}
	if faq, err := client.GetFAQ(ctx); err == nil {
		m.FAQ = faq
	} else if m.Error == "" {
		m.Error = err.Error()
	}
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
