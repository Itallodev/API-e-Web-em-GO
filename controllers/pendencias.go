package controllers

import (
	"Devitallo/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todasPendencias := models.BuscaPendencias()
	temp.ExecuteTemplate(w, "Index", todasPendencias)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		titulo := r.FormValue("titulo")
		descricao := r.FormValue("descricao")
		evento := r.FormValue("evento")
		protocolo := r.FormValue("protocolo")
		responsavel := r.FormValue("responsavel")

		protocoloConv, err := strconv.Atoi(protocolo)
		if err != nil {
			log.Println("Erro de conversão", err)
		}
		models.CriarPendencia(titulo, descricao, evento, protocoloConv, responsavel)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDaPendencia := r.URL.Query().Get("id")
	models.DeletaPendencia(idDaPendencia)
	http.Redirect(w, r, "/", 301)

}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDaPendencia := r.URL.Query().Get("id")
	pendencia := models.EditaPendencia(idDaPendencia)
	temp.ExecuteTemplate(w, "Edit", pendencia)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		titulo := r.FormValue("titulo")
		descricao := r.FormValue("descricao")
		evento := r.FormValue("evento")
		protocolo := r.FormValue("protocolo")
		responsavel := r.FormValue("responsavel")

		idConv, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro de conversão id", err)
		}
		protoConv, err := strconv.Atoi(protocolo)
		if err != nil {
			log.Println("Erro de conversão proto:", err)
		}
		models.AtualizarPendencia(idConv, titulo, descricao, evento, protoConv, responsavel)
	}
	http.Redirect(w, r, "/", 301)
}
