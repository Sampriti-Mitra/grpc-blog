package repos

import (
	"errors"

	"github.com/Sampriti-Mitra/blog-post/protogen/golang/blogs"
	"github.com/google/uuid"
)

type BlogRepo struct {
	mp map[string]*blogs.Blog
}

func NewBlogRepo() *BlogRepo {
	mp := make(map[string]*blogs.Blog)
	return &BlogRepo{mp: mp}
}

func (br *BlogRepo) FetchById(postId string) (*blogs.Blog, error) {
	if _, ok := br.mp[postId]; !ok {
		return nil, errors.New("blog not found")
	}

	return br.mp[postId], nil

}

func (br *BlogRepo) Create(blog *blogs.Blog) (*blogs.Blog, error) {

	blog.PostId = uuid.New().String()
	br.mp[blog.PostId] = blog

	return blog, nil

}

func (br *BlogRepo) UpdatePost(blogUpdateReq *blogs.BlogUpdateRequest) (*blogs.Blog, error) {

	postId := blogUpdateReq.GetPostId()
	var existingBlog *blogs.Blog
	var ok bool
	if existingBlog, ok = br.mp[postId]; !ok {
		return nil, errors.New("blog not found")
	}

	existingBlog.Tags = blogUpdateReq.GetTags()
	existingBlog.Title = blogUpdateReq.GetTitle()
	existingBlog.Content = blogUpdateReq.GetContent()

	return existingBlog, nil

}

func (br *BlogRepo) DeletePost(postId string) error {
	if _, ok := br.mp[postId]; !ok {
		return errors.New("blog not found")
	}

	delete(br.mp, postId)

	return nil

}
