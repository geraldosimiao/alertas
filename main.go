package main

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Aviso representa a estrutura de dados para um aviso meteorológico
type Aviso struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Published   string `xml:"pubDate"`
}

// Feed representa a estrutura de dados para um feed RSS
type Feed struct {
	Channel struct {
		Title       string  `xml:"title"`
		Link        string  `xml:"link"`
		Description string  `xml:"description"`
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

	var feed Feed
	err = xml.Unmarshal(conteúdo, &feed)
	if err != nil {
		return nil, err
	}

	return feed.Channel.Items, nil
}

func extrairDetalhesHTML(descricaoHTML string) map[string]string {
	reader := strings.NewReader(descricaoHTML)
	tokenizer := html.NewTokenizer(reader)

	// Mapa para armazenar os detalhes do aviso
	detalhes := make(map[string]string)

	// Variável para rastrear a chave atual enquanto percorremos o HTML
	var chaveAtual string

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return detalhes
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "th" {
				tokenType = tokenizer.Next()
				if tokenType == html.TextToken {
					// Remover espaços extras e definir como chave atual
					chaveAtual = strings.TrimSpace(tokenizer.Token().Data)
				}
			} else if token.Data == "td" {
				tokenType = tokenizer.Next()
				if tokenType == html.TextToken {
					// Adicionar valor associado à chave atual no mapa
					detalhes[chaveAtual] = strings.TrimSpace(tokenizer.Token().Data)
				}
			}
		}
	}
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

		// Extrair detalhes do HTML e exibir
		detalhes := extrairDetalhesHTML(aviso.Description)
		for chave, valor := range detalhes {
			fmt.Printf("%-15s: %s\n", chave, valor)
		}

		fmt.Println("-----")
	}
}
