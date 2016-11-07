package main

import (
	//"fmt"
	"log"
	//"strings"
	//"unicode"
	"github.com/mattwilmott/go-foreman"
)

type Config struct {
	Username string
	Password string
	URL      string
}

type ForemanClient struct {
	foremanconn *foreman.Foreman
}

// Client() returns a new client.
func (c *Config) Client() (interface{}, error) {
	var client ForemanClient

	client.foremanconn = foreman.Client(c.URL, c.Username, c.Password)
	log.Println("[INFO] Foreman Client configured")

	return &client, nil
}
