package posts

import (
	"context"
	"math/rand"
	"strconv"
	"time"
)

type PostService struct {
	storage map[string][]*Post
}

func NewService() *PostService {
	rand.Seed(time.Now().UnixNano())
	return &PostService{storage: map[string][]*Post{}}
}

func (ps PostService) Get(ctx context.Context, r *GetRequest) (*GetResponse, error) {
	p := ps.storage[r.UserId]

	return &GetResponse{Posts: p}, nil
}

func (ps PostService) Create(ctx context.Context, r *CreateRequest) (*CreateResponse, error) {
	id := strconv.Itoa(rand.Int())
	posts := ps.storage[r.UserId]
	posts = append(posts, &Post{
		Id:          id,
		Name:        r.Name,
		Description: r.Description,
		Type:        r.Type,
	})
	ps.storage[r.UserId] = posts
	return &CreateResponse{Id: id}, nil
}
