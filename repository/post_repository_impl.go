package repository

import (
	"context"
	"fmt"
	"gogo/helper"
	"gogo/model"
	"gogo/prisma/db"
)

type PostRepositoryImpl struct {
	Db *db.PrismaClient
}

// Delete implements PostRepository.
func (p *PostRepositoryImpl) Delete(ctx context.Context, postId string) {
	result, err := p.Db.Post.FindUnique(db.Post.ID.Equals(postId)).Delete().Exec(ctx)
	helper.ErrorPanic(err)
	fmt.Println("Rows affected: ", result)
}

// FindAll implements PostRepository.
func (p *PostRepositoryImpl) FindAll(ctx context.Context) []model.Post {
	allPost, err := p.Db.Post.FindMany().Exec(ctx)
	helper.ErrorPanic(err)
	var posts []model.Post
	for _, post := range allPost {
		published, _ := post.Published()
		description, _ := post.Description()

		postData := model.Post{
			Id:          post.ID,
			Title:       post.Title,
			Published:   published,
			Description: description,
		}

		posts = append(posts, postData)
	}
	return posts
}

// FindById implements PostRepository.
func (p *PostRepositoryImpl) FindById(ctx context.Context, postId string) (model.Post, error) {
	post, err := p.Db.Post.FindUnique(db.Post.ID.Equals(postId)).Exec(ctx)
	helper.ErrorPanic(err)
	published, _ := post.Published()
	description, _ := post.Description()

	postData := model.Post{
		Id:          post.ID,
		Title:       post.Title,
		Published:   published,
		Description: description,
	}
	if post != nil {
		return postData, nil
	} else {
		return postData, err
	}

}

// Save implements PostRepository.
func (p *PostRepositoryImpl) Save(ctx context.Context, post model.Post) {
	result, err := p.Db.Post.CreateOne(
		db.Post.Title.Set(post.Title),
		db.Post.Published.Set(post.Published),
		db.Post.Description.Set(post.Description),
	).Exec(ctx)
	helper.ErrorPanic(err)
	fmt.Println("Rows affected: ", result)
}

// Update implements PostRepository.
func (p *PostRepositoryImpl) Update(ctx context.Context, post model.Post) {
	result, err := p.Db.Post.FindUnique(db.Post.ID.Equals(post.Id)).Update(
		db.Post.Title.Set(post.Title),
		db.Post.Published.Set(post.Published),
		db.Post.Description.Set(post.Description),
	).Exec(ctx)
	helper.ErrorPanic(err)
	fmt.Println("Rows affected: ", result)
}

func NewPostRepository(db *db.PrismaClient) PostRepository {
	return &PostRepositoryImpl{Db: db}
}
