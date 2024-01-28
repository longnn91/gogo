package service

import (
	"context"
	"gogo/data/request"
	"gogo/data/response"
)

type PostService interface {
	Create(cxt context.Context, request request.PostCreateRequest)
	Update(cxt context.Context, request request.PostUpdateRequest)
	Delete(cxt context.Context, postId string)
	FindById(cxt context.Context, postId string) response.PostResponse
	FindAll(cxt context.Context) []response.PostResponse
}
