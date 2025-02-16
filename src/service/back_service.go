package service

import (
	"blog1/src/model"
	"blog1/src/model/gorm_model"
	"blog1/src/utils"
	"fmt"
	"github.com/gosimple/slug"
	"mime/multipart"
)

func (s *Service) BackAllDateService(categoryName string) (allBackDate model.BackAllDataModel, err error) {
	var posts []gorm_model.Post
	if categoryName == "" {
		posts, err = s.dao.GetPosts()
		if err != nil {
			fmt.Println("get posts err:", err)
			return
		}
	} else {
		posts, err = s.dao.GetPostsByCategory(categoryName)
		if err != nil {
			fmt.Println("get GetPostsByCategory err:", err)
			return
		}
	}

	base, err := s.dao.GetBase()
	if err != nil {
		fmt.Println("get base err:", err)
		return
	}

	categories, err := s.dao.GetCategories()
	if err != nil {
		fmt.Println("get categories err:", err)
		return
	}

	comments, err := s.dao.GetComments()
	if err != nil {
		fmt.Println("get comments err:", err)
		return
	}

	users, err := s.dao.GetUsers()
	if err != nil {
		fmt.Println("get users err:", err)
		return
	}

	allBackDate = model.BackAllDataModel{
		Posts:      posts,
		Base:       base,
		Categories: categories,
		Comments:   comments,
		Users:      users,
	}
	return
}

func (s *Service) NewPostService(postModel gorm_model.Post, file multipart.File, fileHeader *multipart.FileHeader) (err error) {
	filePath, err := utils.SaveLocalFile(file, fileHeader)
	if err != nil {
		fmt.Println("utils.SaveLocalFile() err: ", err)
		return err
	}

	slug1 := slug.Make(postModel.Title)

	postModel.Slug = slug1
	postModel.CoverImg = filePath

	err = s.dao.InsertPost(postModel)

	fmt.Printf("postModel:%#v\n", postModel)

	return

}

func (s *Service) GetBackPostService(slug string) (post gorm_model.Post, err error) {
	post, err = s.dao.GetPostByslug(slug)
	return
}
