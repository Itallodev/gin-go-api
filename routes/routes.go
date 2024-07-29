package routes

import (
	"github.com/devitallo/gin-go-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets/", "./assets/")
	r.GET("/alunos", controllers.ExibeAlunos)
	r.GET("/:nome", controllers.Salve)
	r.POST("/alunos", controllers.CriaAlunos)
	r.GET("/alunos/:id", controllers.BuscaId)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.AtualizaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaCPF)
	r.GET("/index", controllers.ExibeIndex)
	r.NoRoute(controllers.RoutesNotFound)
	r.Run()
}
