package models

import (
	"Devitallo/db"
)

type Pendencia struct {
	Id          int
	Titulo      string
	Descricao   string
	Evento      string
	Protocolo   int
	Responsavel string
}

func BuscaPendencias() []Pendencia {
	db := db.ConnectDB()
	selectDasPendencias, err := db.Query("select *from Pendencias order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Pendencia{}
	pendencias := []Pendencia{}
	for selectDasPendencias.Next() {
		var id, protocolo int
		var titulo, descricao, evento, responsavel string

		err = selectDasPendencias.Scan(&id, &titulo, &descricao, &evento, &protocolo, &responsavel)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Titulo = titulo
		p.Descricao = descricao
		p.Evento = evento
		p.Protocolo = protocolo
		p.Responsavel = responsavel

		pendencias = append(pendencias, p)
	}
	defer db.Close()
	return pendencias
}

func CriarPendencia(titulo, descricao, evento string, protocolo int, responsavel string) {
	db.ConnectDB()

	InsertDb, err := db.ConnectDB().Prepare("insert into pendencias(titulo, descricao, evento, protocolo, responsavel) values($1, $2, $3, $4, $5)")
	if err != nil {
		panic(err.Error())
	}
	InsertDb.Exec(titulo, descricao, evento, protocolo, responsavel)
	defer db.ConnectDB().Close()
}

func DeletaPendencia(id string) {
	db := db.ConnectDB()

	deletarAPendencia, err := db.Prepare("delete from pendencias where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarAPendencia.Exec(id)
	defer db.Close()
}

func EditaPendencia(id string) Pendencia {
	db := db.ConnectDB()

	pendenciaDoBanco, err := db.Query("select * from pendencias where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	pendenciaParaAtualizar := Pendencia{}
	for pendenciaDoBanco.Next() {
		var id int
		var titulo, descricao, evento string
		var protocolo int
		var responsavel string
		err = pendenciaDoBanco.Scan(&id, &titulo, &descricao, &evento, &protocolo, &responsavel)
		if err != nil {
			panic(err.Error())
		}
		pendenciaParaAtualizar.Id = id
		pendenciaParaAtualizar.Titulo = titulo
		pendenciaParaAtualizar.Descricao = descricao
		pendenciaParaAtualizar.Evento = evento
		pendenciaParaAtualizar.Protocolo = protocolo
		pendenciaParaAtualizar.Responsavel = responsavel
	}
	defer db.Close()
	return pendenciaParaAtualizar
}
func AtualizarPendencia(id int, titulo, descricao, evento string, protocolo int, responsavel string) {
	db := db.ConnectDB()

	AtualizarPendencia, err := db.Prepare("update pendencias set titulo=$1, descricao=$2, evento=$3, protocolo=$4, responsavel=$5 where id=$6")
	if err != nil {
		panic(err.Error())
	}
	AtualizarPendencia.Exec(titulo, descricao, evento, protocolo, responsavel, id)
	defer db.Close()
}
