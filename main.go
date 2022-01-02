package main

import (
	"html/template"
	"io"
	"os"

	"github.com/iskorotkov/minimail/handlers"
	"github.com/iskorotkov/minimail/models"
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

	c, err := handlers.NewContainer(db)
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
	swagger.File("/openapi.yml", ".docs/api/openapi.yaml")

	// Assets
	e.Static("/css", "static/common/css")

	// Frontend based on templates
	simple := e.Group("/simple")
	simple.GET("/", c.GetMessagesPage)
	simple.GET("/:messageId", c.GetMessagePage)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
