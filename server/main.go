package main

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/iskorotkov/minimail/handlers"
	"github.com/iskorotkov/minimail/models"
	"github.com/iskorotkov/minimail/services"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Renderer struct {
	templates *template.Template
}

func (r Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Renderer = Renderer{
		templates: template.Must(template.ParseGlob("static/templates/*.html")),
	}

	if err := godotenv.Load(); err != nil {
		e.Logger.Warn(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("CONN_STRING")))
	if err != nil {
		e.Logger.Fatal(err)
	}

	if err := db.AutoMigrate(models.Message{}); err != nil {
		e.Logger.Fatal(err)
	}

	service, err := services.NewMessages(db)
	if err != nil {
		e.Logger.Fatal(err)
	}

	c, err := handlers.NewContainer(service)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Middleware
	e.Use(
		middleware.NonWWWRedirect(),
		middleware.Logger(),
		middleware.Recover(),
		middleware.Gzip(),
	)

	apiGroup := e.Group("/api",
		middleware.CORS(),
	)

	// AddMessage - Создаёт новое сообщение
	apiGroup.POST("/messages", c.AddMessage)

	// ClapMessage - Увеличивает количество хлопков сообщения на 1
	apiGroup.POST("/messages/:messageId/claps", c.ClapMessage)

	// GetMessage - Возвращает сообщение по ID
	apiGroup.GET("/messages/:messageId", c.GetMessage)

	// GetMessages - Возвращает список всех сообщений
	apiGroup.GET("/messages", c.GetMessages)

	// Swagger
	swagger := e.Group("/swagger")
	swagger.Static("/", "static/swagger-ui")
	swagger.File("/openapi.yml", "openapi.yml")

	// Assets
	e.Static("/css", "static/common/css")

	// Frontend based on templates
	templates := e.Group("/templates")
	templates.POST("/", c.AddMessagePage)
	templates.POST("/messages/:messageId/claps", c.ClapMessagePage)
	templates.GET("/", c.GetMessagesPage)
	templates.GET("/messages/:messageId", c.GetMessagePage)

	// Frontend based on AJAX
	ajax := e.Group("/ajax")
	ajaxMessageHTML := regexp.MustCompile("/messages/[0-9]*$")
	ajax.GET("/*", func(c echo.Context) error {
		requestPath := c.Request().URL.Path
		if ajaxMessageHTML.MatchString(requestPath) {
			return c.File("static/ajax/messages/index.html")
		}

		filePath := filepath.Join("static/ajax", strings.TrimPrefix(requestPath, "/ajax"))
		return c.File(filePath)
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
