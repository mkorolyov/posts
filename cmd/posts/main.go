package main

import (
	"context"
	"github.com/mkorolyov/core/config"
	"github.com/mkorolyov/core/server"
	"github.com/mkorolyov/posts"
)

func main() {
	postsService := posts.NewService()
	loader := config.Configure()
	s := server.New(loader, postsService)
	s.Serve(context.Background())
}
