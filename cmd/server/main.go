package main

import (
	"log"
	"net"

	"github.com/Sampriti-Mitra/blog-post/internal/repos"
	"github.com/Sampriti-Mitra/blog-post/internal/services"
	"github.com/Sampriti-Mitra/blog-post/protogen/golang/blogs"
	"google.golang.org/grpc"
)

func main() {
	const addr = "0.0.0.0:5501"

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	repo := repos.NewBlogRepo()
	blogService := services.NewBlogService(repo)

	blogs.RegisterBlogsServer(server, blogService)

	log.Printf("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
