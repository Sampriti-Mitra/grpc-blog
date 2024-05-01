package repos_test

import (
	"errors"
	"testing"

	"github.com/Sampriti-Mitra/blog-post/internal/repos"
	"github.com/Sampriti-Mitra/blog-post/protogen/golang/blogs"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/type/date"
)

func TestBlogCreate(t *testing.T) {
	blog := blogs.Blog{
		Title:   "How to make a blog",
		Content: "Just start blogging by adding text",
		Author:  "Sam",
		PublicationDate: &date.Date{
			Year:  2024,
			Month: 05,
			Day:   1,
		},
		Tags: []string{"starter", "newbie", "blog"},
	}
	testCases := []struct {
		name          string
		expectedError error
		expectedBlog  *blogs.Blog
	}{
		{
			"Create Blog",
			nil,
			&blog,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := repos.NewBlogRepo()
			blog, err := repo.Create(&blog)
			assert.Nil(t, err)
			assert.Equal(t, blog.Author, tc.expectedBlog.Author)
			assert.Equal(t, blog.Content, tc.expectedBlog.Content)
			assert.Equal(t, blog.Tags, tc.expectedBlog.Tags)
			assert.Equal(t, blog.PublicationDate, tc.expectedBlog.PublicationDate)
			assert.Equal(t, blog.Title, tc.expectedBlog.Title)
			assert.NotNil(t, blog.PostId)
		})
	}
}

func TestBlogRead(t *testing.T) {
	blog := blogs.Blog{
		Title:   "How to make a blog",
		Content: "Just start blogging by adding text",
		Author:  "Sam",
		PublicationDate: &date.Date{
			Year:  2024,
			Month: 05,
			Day:   1,
		},
		Tags: []string{"starter", "newbie", "blog"},
	}

	repo := repos.NewBlogRepo()
	createdBlog, _ := repo.Create(&blog)

	testCases := []struct {
		name          string
		postId        string
		expectedError error
		expectedBlog  *blogs.Blog
	}{
		{
			"Read Blog",
			createdBlog.PostId,
			nil,
			createdBlog,
		},
		{
			"Blog Not Found",
			uuid.New().String(),
			errors.New("blog not found"),
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fetchedBlog, err := repo.FetchById(tc.postId)

			if tc.expectedError != nil {
				assert.Equal(t, err, tc.expectedError)
				return
			}

			assert.Equal(t, fetchedBlog.Author, tc.expectedBlog.Author)
			assert.Equal(t, fetchedBlog.Content, tc.expectedBlog.Content)
			assert.Equal(t, fetchedBlog.Tags, tc.expectedBlog.Tags)
			assert.Equal(t, fetchedBlog.PublicationDate, tc.expectedBlog.PublicationDate)
			assert.Equal(t, fetchedBlog.Title, tc.expectedBlog.Title)
			assert.NotNil(t, fetchedBlog.PostId)
		})
	}
}

func TestBlogDelete(t *testing.T) {
	blog := blogs.Blog{
		Title:   "How to make a blog",
		Content: "Just start blogging by adding text",
		Author:  "Sam",
		PublicationDate: &date.Date{
			Year:  2024,
			Month: 05,
			Day:   1,
		},
		Tags: []string{"starter", "newbie", "blog"},
	}

	repo := repos.NewBlogRepo()
	createdBlog, _ := repo.Create(&blog)

	testCases := []struct {
		name          string
		postId        string
		expectedError error
	}{
		{
			"Delete Blog",
			createdBlog.PostId,
			nil,
		},
		{
			"Delete Not Found",
			uuid.New().String(),
			errors.New("blog not found"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.DeletePost(tc.postId)
			assert.Equal(t, err, tc.expectedError)

		})
	}
}

func TestBlogUpdate(t *testing.T) {
	blog := blogs.Blog{
		Title:   "How to make a blog",
		Content: "Just start blogging by adding text",
		Author:  "Sam",
		PublicationDate: &date.Date{
			Year:  2024,
			Month: 05,
			Day:   1,
		},
		Tags: []string{"starter", "newbie", "blog"},
	}

	repo := repos.NewBlogRepo()
	createdBlog, _ := repo.Create(&blog)

	testCases := []struct {
		name          string
		updateReq     blogs.BlogUpdateRequest
		expectedError error
		updatedBlog   blogs.Blog
	}{
		{
			"Update Blog",
			blogs.BlogUpdateRequest{
				PostId:  createdBlog.PostId,
				Title:   "How to update a blog",
				Content: "Update blog",
				Author:  "Sampriti",
				Tags:    []string{"blog"},
			},
			nil,
			blogs.Blog{
				PostId:  createdBlog.PostId,
				Title:   "How to update a blog",
				Content: "Update blog",
				Author:  "Sampriti",
				PublicationDate: &date.Date{
					Year:  2024,
					Month: 05,
					Day:   1,
				},
				Tags: []string{"blog"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			updatedBlog, err := repo.UpdatePost(&tc.updateReq)
			assert.Equal(t, err, tc.expectedError)
			assert.Equal(t, updatedBlog.Author, tc.updatedBlog.Author)
			assert.Equal(t, updatedBlog.PostId, tc.updatedBlog.PostId)
			assert.Equal(t, updatedBlog.Content, tc.updatedBlog.Content)
			assert.Equal(t, updatedBlog.Title, tc.updatedBlog.Title)
			assert.Equal(t, updatedBlog.PublicationDate, tc.updatedBlog.PublicationDate)
			assert.Equal(t, updatedBlog.Tags, tc.updatedBlog.Tags)

		})
	}
}
