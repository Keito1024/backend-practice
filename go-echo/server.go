package main

import (
	"go-echo/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	users = map[int]*domain.User{} //{1:{ID: 1, Name: "sample"}}
	seq   = 1
)

// handler

// ユーザー一覧取得
func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

// ユーザー作成
func createUser(c echo.Context) error {
	u := &domain.User{
		ID: seq,
	}
	// jsonパラメタを取得するには受け取るための構造体が必要(user)
	// 構造体をBindすることで構造体のメンバとして取得できる
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

// ユーザー取得（id）
func getUser(c echo.Context) error {
	// Atoiのみ二番目の返り値としてerrを返す
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

// ユーザー更新
func updateUser(c echo.Context) error {
	u := new(domain.User) // var u *User
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

// ユーザー削除
func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi("id")
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: nil, // 許可するドメインを追加
		AllowMethods: nil, // 許可するHTTP Methodを追加
		AllowHeaders: nil, // 許可するheaderを追加
	}))

	e.POST("/users", createUser)
	e.GET("/users", getUsers)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
