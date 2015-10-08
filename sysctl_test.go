package sysctl

import (
        "testing"
)

func TestGet(t *testing.T) {
        got, err := Get("net.ipv4.tcp_sack")
        ex := "1"
        if err != nil {
                t.Fatalf("Could not read key")
        }
        if got != ex {
                t.Errorf("Expected: %s, got %s", ex, got)
        }
}

func TestSet(t *testing.T) {
        ex := "0"
        err := Set("net.ipv4.ip_forward", ex)
        if err != nil {
                t.Fatalf("err")
        }
        got, err := Get("net.ipv4.ip_forward")
        if err != nil {
                t.Fatalf("Could not read key")
        }
        if got != ex {
                t.Errorf("Expected: %s, got %s", ex, got)
        }
}
