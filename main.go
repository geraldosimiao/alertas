package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Aviso representa a estrutura de dados para um aviso meteorológico
type Aviso struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Published   time.Time `xml:"published"`
	Status      string    `xml:"status"`
	Evento      string    `xml:"evento"`
	Severidade  string    `xml:"severidade"`
	Início      string    `xml:"início"`
	Fim         string    `xml:"fim"`
	Área        string    `xml:"área"`
}

// Feed representa a estrutura de dados para um feed RSS
type Feed struct {
	Channel struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Items       []Aviso `xml:"item"`
	} `xml:"channel"`
}

func obterAvisos(apiURL string) ([]Aviso, error) {
    fmt.Printf("Fazendo solicitação para: %s\n", apiURL)

    resposta, err := http.Get(apiURL)
    if err != nil {
        return nil, err
    }
    defer resposta.Body.Close()

    conteúdo, err := ioutil.ReadAll(resposta.Body)
    if err != nil {
        return nil, err
    }

    fmt.Printf("Resposta da API:\n%s\n", conteúdo)

    var feed Feed
    err = xml.Unmarshal(conteúdo, &feed)
    if err != nil {
        return nil, err
    }

    return feed.Channel.Items, nil
}

func main() {
	apiURL := "https://apiprevmet3.inmet.gov.br/avisos/rss"

	avisos, err := obterAvisos(apiURL)
	if err != nil {
		fmt.Println("Erro ao obter avisos:", err)
		os.Exit(1)
	}

	// Para imprimir informações dos avisos
	fmt.Printf("Avisos retornados: %d\n", len(avisos))
	for i, aviso := range avisos {
		fmt.Printf("Aviso #%d\n", i+1)
		fmt.Printf("Título: %s\n", aviso.Title)
		fmt.Printf("Link: %s\n", aviso.Link)
		fmt.Printf("Data de Publicação: %s\n", aviso.Published)
		fmt.Printf("Status: %s\n", aviso.Status)
		fmt.Printf("Evento: %s\n", aviso.Evento)
		fmt.Printf("Severidade: %s\n", aviso.Severidade)
		fmt.Printf("Início: %s\n", aviso.Início)
		fmt.Printf("Fim: %s\n", aviso.Fim)
		fmt.Printf("Área: %s\n", aviso.Área)
		fmt.Println("-----")
	}
}
