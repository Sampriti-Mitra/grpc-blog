package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Sampriti-Mitra/blog-post/protogen/golang/blogs"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	blogServerAddr := "localhost:5501"
	conn, err := grpc.Dial(blogServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to order service: %v", err)
	}
	defer conn.Close()

	client := blogs.NewBlogsClient(conn)
	postId := createPost(client)
	getPost(client, postId)
	updatePost(client, postId)
	deletePost(client, postId)

}

func createPost(client blogs.BlogsClient) string {
	blog, err := client.AddPost(context.Background(), &blogs.Blog{
		Title:   "How to make a blog",
		Content: "Just start blogging by adding text",
		Author:  "Sam",
		PublicationDate: &date.Date{
			Year:  2024,
			Month: 05,
			Day:   1,
		},
		Tags: []string{"starter", "newbie", "blog"},
	})
	if err != nil {
		fmt.Println("unexpected error in create", err)
	}
	fmt.Println(blog)
	return blog.PostId
}

func getPost(client blogs.BlogsClient, postId string) {
	blog, err := client.ReadPost(context.Background(), &blogs.BlogFetchRequest{PostId: postId})
	if err != nil {
		fmt.Println("unexpected error in create", err)
	}
	fmt.Println(blog)
}

func updatePost(client blogs.BlogsClient, postId string) {
	blog, err := client.UpdatePost(context.Background(), &blogs.BlogUpdateRequest{
		PostId:  postId,
		Title:   "How to update a blog post",
		Content: "Updating a blog post is easy",
		Author:  "Sampriti",
		Tags:    []string{"update", "blog"},
	})
	if err != nil {
		fmt.Println("unexpected error in post", err)
	}
	fmt.Println(blog)
}

func deletePost(client blogs.BlogsClient, postId string) {
	res, err := client.DeletePost(context.Background(), &blogs.BlogFetchRequest{
		PostId: postId,
	})
	if err != nil {
		fmt.Println("unexpected error in post", err)
	}
	fmt.Println(res)
}
