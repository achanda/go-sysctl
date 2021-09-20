package sysctl

import (
	"runtime"
	"testing"
)

func TestGet(t *testing.T) {
	var got string
	var err error
	switch runtime.GOOS {
	case "freebsd":
		got, err = Get("net.inet.tcp.sack.enable")
	default:
		got, err = Get("net.ipv4.tcp_sack")
	}
	ex := "1"
	if err != nil {
		t.Fatalf("Could not read key")
	}
	if got != ex {
		t.Errorf("Expected: %s, got %s", ex, got)
	}
}

func TestGetFailPattern(t *testing.T) {
	param := "invalid_kernel_param"
	got, err := Get(param)

	if got != "" {
		t.Fatalf("Expected: \"\", got %s", got)
	}
	if err == nil {
		t.Fatal("Expected: returns some error but got nil")
	}
}

func TestSet(t *testing.T) {
	var param, ex string
	switch runtime.GOOS {
	case "freebsd":
		param = "compat.linux.debug"
		ex = "2"
	default:
		param = "net.ipv4.ip_forward"
		ex = "0"
	}
	err := Set(param, ex)
	if err != nil {
		t.Fatalf("Failed to call Set(%s, %s)", param, ex)
	}
	got, err := Get(param)
	if err != nil {
		t.Fatalf("Could not read key from %s", param)
	}
	if got != ex {
		t.Errorf("Expected: %s, got %s from %s", ex, got, param)
	}
}

func TestSetFailPattern(t *testing.T) {
	param := "invalid_kernel_param"
	ex := "0"
	err := Set(param, ex)

	if err == nil {
		t.Fatalf("Expected: Set(%s, %s) returns error, returned nil", param, ex)
	}
}
