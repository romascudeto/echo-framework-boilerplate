package services

import (
	"echo-framework/article/repositories"
	"echo-framework/models"
)

func ListArticle() (dataRet []models.Article, err error) {
	err = repositories.Fetch(&dataRet)
	return dataRet, err
}

func DetailArticle(id int32) (dataRet models.Article, err error) {
	err = repositories.Detail(&dataRet, id)
	return dataRet, err
}

func CreateArticle(article models.Article) (dataRet models.Article, err error) {
	err = repositories.Create(&article)
	err = repositories.Detail(&article, int32(article.ID))

	return article, err
}

func UpdateArticle(article models.Article) (dataRet models.Article, err error) {
	err = repositories.Update(&article)
	return article, err
}

func DeleteArticle(article models.Article) (dataRet models.Article, err error) {
	err = repositories.Delete(&article)
	return article, err
}
