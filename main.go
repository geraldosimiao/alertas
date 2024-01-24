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
	XMLName xml.Name `xml:"rss"`
	Channel struct {
		Items []Aviso `xml:"channel>item"`
	} `xml:"channel"`
}

func obterAvisos(apiURL string) ([]Aviso, error) {
	resposta, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resposta.Body.Close()

	conteúdo, err := ioutil.ReadAll(resposta.Body)
	if err != nil {
		return nil, err
	}

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

	// Adicione estas linhas para imprimir informações dos avisos
	fmt.Printf("Avisos retornados: %d\n", len(avisos))
	for _, aviso := range avisos {
		fmt.Printf("Severidade: %s\n", aviso.Severidade)
		// Adicione mais informações conforme necessário
	}
}
