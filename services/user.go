package services

import (
	"fmt"
	"qdjr-api/initializers"
	"qdjr-api/models"
	"reflect"
)

type UserService struct{}

func (service UserService) Search(params interface{}, page int, perPage int, _ interface{}) (users *[]models.User, totalItems int64) {
	users = &[]models.User{}
	totalItems = 0
	paramValues := reflect.ValueOf(params)
	query := initializers.DB
	if keyword := paramValues.FieldByName("keyword").String(); keyword != "" {
		query.Where("username LIKE ?", "%"+keyword+"%")
	}
	//query.Count(&totalItems)
	query.Offset((page - 1) * perPage).Limit(perPage).Find(users).Count(&totalItems)
	fmt.Println(users)
	fmt.Println(totalItems)
	return users, totalItems
}

func (service UserService) Detail(id int) *models.User {
	user := &models.User{}
	initializers.DB.Where("id = ?", id).First(&user)
	return user
}

func (service UserService) DetailByParam(param string) *models.User {
	user := &models.User{}
	initializers.DB.Where("username = ?", param).Or("email = ?", param).First(user)
	return user
}
