// Package config loads env-driven configuration for musig-stream.
package config

import "os"

type Config struct {
	Addr    string
	Product string
}

func FromEnv() Config {
	c := Config{Addr: ":8080", Product: "musig"}
	if v := os.Getenv("ADDR"); v != "" {
		c.Addr = v
	}
	return c
}
