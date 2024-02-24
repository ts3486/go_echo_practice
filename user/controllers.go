package user

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	database "myapp/database"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

var (
	users = map[int]*User{}
	lock  = sync.Mutex{}
)

func Create(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
        db := database.DB;
        result := new(User)
        if err := c.Bind(result); err != nil {
            return err
        }
        db.Create(&result)    
        return c.JSON(http.StatusOK, result)
}

func Get(c echo.Context) error {
	db := database.DB
	user := User{}
	user_id := c.Param("id")
	result := db.Table("users").Select([]string{
		"id",
		"first_name",
		"family_name",
		"email",
		"created_at",
		"updated_at",
		"deleted_at"}).Find(&user,"id = ?", user_id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound){
		fmt.Println("no record")
	}
	return c.JSON(http.StatusOK, user)
}

func Update(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
		db := database.DB
			newPost := new(User)
			if err := c.Bind(newPost); err != nil {
				return err
			}
			user_id := c.Param("id")
			if user_id != "" {
				post := User{}
				db.Model(&post).Where("id = ?", user_id).Updates(newPost)
				fmt.Println(post)
				return c.JSON(http.StatusOK, post)
			} else {
	return c.JSON(http.StatusNotFound, nil)
 }
}

func Delete(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	db := database.DB
	newDelete := new(User)
	user_id := c.Param("id")
	db.Delete(&newDelete, "id = ?", user_id)
	return c.JSON(http.StatusOK, newDelete)
}

func GetAll(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	return c.JSON(http.StatusOK, users)
}
