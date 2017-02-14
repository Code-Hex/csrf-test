package likesns

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Run() error {
	e := echo.New()
	e.GET("/", Home)
	e.POST("/login", Login)

	// /users 以下は jwt による認証が必要
	secret := e.Group("/users")
	secret.Use(middleware.JWT([]byte("secret")))
	secret.POST("/post", Post)

	return e.Start(":8080")
}

// 簡易ログインボタン
func Home(c echo.Context) error {
	return c.HTML(200, `<form action="/login" method="post"><input type="submit" value="send"></form>`)
}

// title=value でポスト可能
func Post(c echo.Context) error {
	title := c.FormValue("title")
	return c.String(200, fmt.Sprintf("post: %s", title))
}

// ログイン後 jwt を Cookie に仕込む
func Login(c echo.Context) error {

	// jwt 作成
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	tok, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.String(500, err.Error())
	}

	// Cookie のセット
	cookie := new(http.Cookie)
	cookie.HttpOnly = true
	cookie.Name = "token"
	cookie.Value = tok
	c.SetCookie(cookie)

	// ログイン中に誰かが悪意のあるサイトへ誘導するためのリンクを設置していると仮定
	return c.HTML(200, "login success: <a href='http://localhost:8081/'>CSRF!!</a>")
}
