package services

import (
	"fmt"
	"qdjr-api/initializers"
	"qdjr-api/models"
	"reflect"
	"time"
)

type ArticleService struct{}

func (service ArticleService) Search(params interface{}, page int, perPage int, _ interface{}) (articles *[]models.Article, totalItems int64) {
	articles = &[]models.Article{}
	totalItems = 0
	paramValues := reflect.ValueOf(params)
	query := initializers.DB
	if keyword := paramValues.FieldByName("Keyword").String(); keyword != "" {
		query.Where("username LIKE ?", "%"+keyword+"%")
	}
	query.Offset((page - 1) * perPage).Limit(perPage).Find(articles).Count(&totalItems)
	return articles, totalItems
}

func (service ArticleService) Detail(id int) *models.Article {
	article := &models.Article{}
	initializers.DB.Where("id = ?", id).First(&article)
	return article
}

func (service ArticleService) Create(data interface{}) (*models.Article, error) {
	dataValues := reflect.ValueOf(data)
	publishedAt, _ := time.Parse("2006-01-02 15:04:05", dataValues.FieldByName("PublishedAt").String())
	article := &models.Article{
		Title:       dataValues.FieldByName("Title").String(),
		Slug:        dataValues.FieldByName("Slug").String(),
		Description: dataValues.FieldByName("Description").String(),
		Content:     dataValues.FieldByName("Content").String(),
		Status:      dataValues.FieldByName("Status").Interface().(uint8),
		Author:      dataValues.FieldByName("Author").String(),
		PublishedAt: publishedAt,
	}
	result := initializers.DB.Create(&article)
	if result.Error != nil {
		return article, result.Error
	} else {
		return article, nil
	}
}

func (service ArticleService) Update(id int, data interface{}) (*models.Article, error) {
	dataValues := reflect.ValueOf(data)
	var publishedAt time.Time
	if strPublishedAt := dataValues.FieldByName("PublishedAt").String(); strPublishedAt != "" {
		publishedAt, _ = time.Parse("2006-01-02 15:04:05", strPublishedAt)
	}
	dataUpdate := map[string]any{
		"Title":       dataValues.FieldByName("Title").String(),
		"Slug":        dataValues.FieldByName("Slug").String(),
		"Description": dataValues.FieldByName("Description").String(),
		"Content":     dataValues.FieldByName("Content").String(),
		"Status":      dataValues.FieldByName("Status").Interface().(uint8),
		"Author":      dataValues.FieldByName("Author").String(),
		"PublishedAt": publishedAt,
	}

	/**
	 * 1. convert all type of v to boolean
	 * 2. delete key if it is zero value
	 */
	for k, v := range dataUpdate {
		if v == reflect.Zero(reflect.TypeOf(v)).Interface() {
			delete(dataUpdate, k)
		}
	}
	fmt.Println(dataUpdate)
	article := &models.Article{Id: uint64(id)}
	result := initializers.DB.Model(article).UpdateColumns(dataUpdate).First(article)
	if result.Error != nil {
		return article, result.Error
	} else {
		return article, nil
	}
}
