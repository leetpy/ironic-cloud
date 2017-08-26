package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"github.com/leetpy/common"
)

func routeHtml(templateDir string, e *echo.Echo) {
	files, err := common.WalkTemplates(templateDir, "html")
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, value := range files {
		content, err := ioutil.ReadFile(value)
		if err != nil {
			fmt.Println(err.Error)
		} else {
			e.GET(value[len(templateDir):], func(c echo.Context) error {
				return c.HTML(http.StatusOK, string(content))
			})
		}
	}
}

func main() {
	e := echo.New()

	// static
	e.Static("/static", "assets")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, world")
	})

	e.GET("/index", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<strong>Hello, World!</strong>")
	})

	//
	routeHtml("templates", e)

	e.Logger.Fatal(e.Start(":8080"))
}
