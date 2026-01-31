package main

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("rigid constructor")
	badExample()

	fmt.Println("builder pattern")
	builderExample()

	fmt.Println("functional options")
  functionalOptionsExample()
}
// rigid constructor 

type Server struct {
	addr string
	port int
}

func NewServerBad(addr string, port int) (*Server, error) {
	if port < 0 {
		return nil, errors.New("port must be positive")
	}
	if port == 0 {
		port = randomPort()
	}
	return &Server{addr: addr, port:port}, nil
}

func badExample() {
	s, _ := NewServerBad("localhost", 8080)
	fmt.Printf("server: ", s)

	fmt.Println("Problem: add more params breaks callers")
}

// Builder pattern


type Config struct {
	port int
}

type ConfigBuilder struct {
	port *int
}

func(b *ConfigBuilder) Port(p int) *ConfigBuilder {
	b.port = &p
	return b 
}

func (b *ConfigBuilder) Build() (Config, error) {
	var cfg Config

	if b.port == nil {
		cfg.port = defaultPort
	} else if *b.port < 0 {
		return Config{}, errors.New("port must be positive")
	} else if *b.port == 0 {
		cfg.port = randomPort()
	} else {
		cfg.port = *b.port
	}

	return cfg, nil
}

func builderExample() {
	builder := &ConfigBuilder{}
	builder.Port(8080)

	cfg, _ := builder.Build()
	fmt.Println("builder config:", cfg)

	fmt.Println("Works, but verbose + delayed validation")
}

// Functional Options 

type options struct {
	port *int
}

type Option func(*options) error 

func WithPort(port int) Option {
	return func(o *options) error {
		if port < 0 {
			return errors.New("port must be positive")
		}
		o.port = &port 
		return nil
	}
}

func NewServer(addr string, opts ...Option) (*Server, error) {
	var o options

	for _, opt := range opts {
		if err := opt(&o); err != nil {
			return nil, err
		}
	}

	port := defaultPort
	if o.port != nil {
		if *o.port == 0 {
			port = randomPort()
		} else {
			port = *o.port
		}
	}

	return &Server{addr: addr, port: port}, nil
}

func functionalOptionsExample() {
	s1, _ := NewServer("localhost")
	s2, _ := NewServer("localhost", WithPort(8080))
	s3, _ := NewServer("localhost", WithPort(0))

	fmt.Printf("default: %+v\n", s1)
	fmt.Printf("custom : %+v\n", s2)
	fmt.Printf("random : %+v\n", s3)

	fmt.Println("Clean API, extensible, no breaking changes")
}

// helpers


const defaultPort = 80

func randomPort() int {
	rand.Seed(time.Now().UnixNano())
	return 10000 + rand.Intn(50000)
}
