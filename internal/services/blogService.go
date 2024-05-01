package services

import (
	"context"

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
	return bs.repo.Create(req)
}
func (bs *BlogService) ReadPost(_ context.Context, req *blogs.BlogFetchRequest) (*blogs.Blog, error) {
	return bs.repo.FetchById(req.PostId)
}
func (bs *BlogService) UpdatePost(_ context.Context, req *blogs.BlogUpdateRequest) (*blogs.Blog, error) {
	return bs.repo.UpdatePost(req)
}
func (bs *BlogService) DeletePost(_ context.Context, req *blogs.BlogFetchRequest) (*blogs.Result, error) {
	err := bs.repo.DeletePost(req.PostId)
	if err != nil {
		return &blogs.Result{Success: false, Message: "unable to delete"}, nil
	}
	return &blogs.Result{Success: true, Message: "deleted successfully"}, nil
}
