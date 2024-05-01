package services

import (
	"context"
	"fmt"

	"github.com/Sampriti-Mitra/blog-post/internal/repos"
	"github.com/Sampriti-Mitra/blog-post/protogen/golang/blogs"
)

type BlogService struct {
	repo *repos.BlogRepo
	blogs.UnimplementedBlogsServer
}

func NewBlogService(repo *repos.BlogRepo) *BlogService {
	return &BlogService{repo: repo}
}

func (bs *BlogService) AddPost(_ context.Context, req *blogs.Blog) (*blogs.Blog, error) {
	err := validatePost(req)
	if err != nil {
		return nil, err
	}
	return bs.repo.Create(req)
}
func (bs *BlogService) ReadPost(_ context.Context, req *blogs.BlogFetchRequest) (*blogs.Blog, error) {
	blog, err := bs.repo.FetchById(req.PostId)
	if err != nil {
		return nil, err
	}
	return blog, nil
}
func (bs *BlogService) UpdatePost(_ context.Context, req *blogs.BlogUpdateRequest) (*blogs.Blog, error) {
	err := validatePostUpdate(req)
	if err != nil {
		return nil, err
	}
	return bs.repo.UpdatePost(req)
}
func (bs *BlogService) DeletePost(_ context.Context, req *blogs.BlogFetchRequest) (*blogs.Result, error) {
	err := bs.repo.DeletePost(req.PostId)
	if err != nil {
		return &blogs.Result{Success: false, Message: "unable to delete"}, nil
	}
	return &blogs.Result{Success: true, Message: "deleted successfully"}, nil
}

func validatePost(req *blogs.Blog) error {
	if req.Title == "" {
		return fmt.Errorf("title cannot be empty")
	}
	if req.Content == "" {
		return fmt.Errorf("content cannot be empty")
	}
	if req.Author == "" {
		return fmt.Errorf("author cannot be empty")
	}
	if req.PublicationDate.Day == 0 {
		return fmt.Errorf("publication date cannot be empty")
	}
	if len(req.Tags) == 0 {
		return fmt.Errorf("at least one tag is required")
	}
	return nil
}

func validatePostUpdate(req *blogs.BlogUpdateRequest) error {
	if req.Title == "" && req.Content == "" && len(req.Author) == 0 && len(req.Tags) == 0 {
		return fmt.Errorf("no fields provided for update")
	}
	return nil
}
