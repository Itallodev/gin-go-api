package controllers

import (
	"net/http"

	"github.com/devitallo/gin-go-api/database"
	"github.com/devitallo/gin-go-api/models"
	"github.com/gin-gonic/gin"
)

func ExibeAlunos(c *gin.Context) {
	var alunos []models.Aluno

	// Fetch all alunos from the database
	database.DB.Find(&alunos)

	c.JSON(200, gin.H{
		"API diz:": alunos,
	})
}

func Salve(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "Saalve! " + nome + ", Tudo certinho?"})
}

func CriaAlunos(c *gin.Context) {
	var aluno models.Aluno

	// Bind JSON request body to struct
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := models.ValidarDados(&aluno); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	database.DB.Create(&aluno)

	c.JSON(200, gin.H{
		"API diz:": "Aluno criado com sucesso!",
	})
}

func BuscaId(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno

	// Find aluno by ID
	if err := database.DB.Where("id =?", id).First(&aluno).Error; err != nil {
		c.JSON(404, gin.H{"error": "Aluno não encontrado"})
		return
	}

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Não encontrado",
		})
	}

	c.JSON(200, gin.H{"aluno": aluno})
}

func DeletaAluno(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno

	// Find aluno by ID
	if err := database.DB.Where("id =?", id).First(&aluno).Error; err != nil {
		c.JSON(404, gin.H{"error": "Aluno não encontrado"})
		return
	}

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Não encontrado",
		})
	}

	// Delete aluno from database
	database.DB.Delete(&aluno)

	c.JSON(200, gin.H{"success": "Aluno deletado com sucesso"})
}

func AtualizaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := c.ShouldBindBodyWithJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado"})
		return
	}
	if err := models.ValidarDados(&aluno); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaCPF(c *gin.Context) {
	cpf := c.Params.ByName("cpf")
	var aluno models.Aluno

	// Find aluno by CPF
	if err := database.DB.Where("cpf =?", cpf).First(&aluno).Error; err != nil {
		c.JSON(404, gin.H{"error": "Aluno não encontrado"})
		return
	}

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Não encontrado",
		})
	}

	c.JSON(200, gin.H{"aluno": aluno})
}

func ExibeIndex(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RoutesNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
