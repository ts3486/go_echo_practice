package user

import (
	"net/http"
	"strconv"
	"sync"

	database "myapp/db"

	"github.com/labstack/echo/v4"
)

var (
	users = map[int]*User{}
	seq   = 1
	lock  = sync.Mutex{}
)

func Create(c echo.Context) error {
        db := database.Connect();
        result := new(User)
        if err := c.Bind(result); err != nil {
            return err
        }
        db.Create(&result)    
        return c.JSON(http.StatusOK, result)
}

func CreateUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := &User{
		Id: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.Id] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func GetUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func UpdateUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].First_name = u.First_name
	return c.JSON(http.StatusOK, users[id])
}

func DeleteUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

func GetAllUsers(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	return c.JSON(http.StatusOK, users)
}
