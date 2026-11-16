package config

import "testing"

func TestLoadUsesDefaults(t *testing.T) {
	cfg, err := Load(func(string) string { return "" })
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg.Address != "0.0.0.0" || cfg.Port != 8080 {
		t.Fatalf("Load() = %+v, want default address and port", cfg)
	}
}

func TestLoadParsesEnvironment(t *testing.T) {
	values := map[string]string{
		"ALVORA_ADDRESS": "127.0.0.1",
		"ALVORA_PORT":    "9090",
	}
	cfg, err := Load(func(key string) string { return values[key] })
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg.Address != "127.0.0.1" || cfg.Port != 9090 {
		t.Fatalf("Load() = %+v, want parsed configuration", cfg)
	}
}

func TestLoadRejectsInvalidPort(t *testing.T) {
	_, err := Load(func(key string) string {
		if key == "ALVORA_PORT" {
			return "70000"
		}
		return ""
	})
	if err == nil {
		t.Fatal("Load() error = nil, want invalid port error")
	}
}
