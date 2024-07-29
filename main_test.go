package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/devitallo/gin-go-api/controllers"
	"github.com/devitallo/gin-go-api/database"
	"github.com/devitallo/gin-go-api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func TestsSetup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func TestStatusCode(t *testing.T) {
	r := TestsSetup()
	r.GET("/:nome", controllers.Salve)
	req, _ := http.NewRequest("GET", "/gui", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, resposta.Code, http.StatusOK)
	mockResposta := `{"API diz:":"Saalve! gui, Tudo certinho?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, string(respostaBody), mockResposta)
}

func CreateAlunoMock(t *testing.T) {
	aluno := models.Aluno{Nome: "Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
}

func DeleteAlunoMock(t *testing.T) {
	aluno := models.Aluno{Nome: "Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Where("nome =?", aluno.Nome).Delete(&aluno)
}

func TestaCPF(t *testing.T) {
	database.ConectaDB()
	CreateAlunoMock(t)
	defer DeleteAlunoMock(t)
	r := TestsSetup()
	r.GET("/aluno/:cpf", controllers.BuscaCPF)
	req, _ := http.NewRequest("GET", "/aluno/12345678901", nil)
	assert.Equal(t, http.StatusOK, req)
}

func TestBuscaID(t *testing.T) {
	database.ConectaDB()
	CreateAlunoMock(t)
	defer DeleteAlunoMock(t)
	r := TestsSetup()
	r.GET("/aluno/:id", controllers.BuscaId)
	path := "/aluno/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	assert.Equal(t, alunoMock.Nome, 3)

}

func TestAlunosHandler(t *testing.T) {
	database.ConectaDB()
	CreateAlunoMock(t)
	defer DeleteAlunoMock(t)
	r := TestsSetup()
	r.GET("/alunos", controllers.ExibeAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, resposta.Code, http.StatusOK)
}
