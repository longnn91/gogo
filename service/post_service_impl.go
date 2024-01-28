package service

import (
	"context"
	"gogo/data/request"
	"gogo/data/response"
	"gogo/helper"
	"gogo/model"
	"gogo/repository"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
}

func NewPostServiceImpl(postRepository repository.PostRepository) PostService {
	return &PostServiceImpl{PostRepository: postRepository}
}

// Create implements PostService.
func (p *PostServiceImpl) Create(cxt context.Context, request request.PostCreateRequest) {
	postData := model.Post{
		Title:       request.Title,
		Published:   request.Published,
		Description: request.Description,
	}
	p.PostRepository.Save(cxt, postData)
}

// Delete implements PostService.
func (p *PostServiceImpl) Delete(cxt context.Context, postId string) {
	post, err := p.PostRepository.FindById(cxt, postId)
	helper.ErrorPanic(err)
	p.PostRepository.Delete(cxt, post.Id)
}

// FindAll implements PostService.
func (p *PostServiceImpl) FindAll(cxt context.Context) []response.PostResponse {
	post := p.PostRepository.FindAll(cxt)
	var postResp []response.PostResponse
	for _, post := range post {
		postResp = append(postResp, response.PostResponse{
			Id:          post.Id,
			Title:       post.Title,
			Published:   post.Published,
			Description: post.Description,
		})
	}
	return postResp
}

// FindById implements PostService.
func (p *PostServiceImpl) FindById(cxt context.Context, postId string) response.PostResponse {
	post, err := p.PostRepository.FindById(cxt, postId)
	helper.ErrorPanic(err)
	return response.PostResponse{
		Id:          post.Id,
		Title:       post.Title,
		Published:   post.Published,
		Description: post.Description,
	}
}

// Update implements PostService.
func (p *PostServiceImpl) Update(cxt context.Context, request request.PostUpdateRequest) {
	postData := model.Post{
		Id:          request.Id,
		Title:       request.Title,
		Published:   request.Published,
		Description: request.Description,
	}
	p.PostRepository.Update(cxt, postData)
}
