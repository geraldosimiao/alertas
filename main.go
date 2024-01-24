package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"golang.org/x/net/html"
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

	//Descomente a linha abaixo para mostrar o conteúdo completo do xml
	//fmt.Printf("Resposta da API:\n%s\n", conteúdo)

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
		fmt.Printf("Título: %s\n", aviso.Título)
		fmt.Printf("Link: %s\n", aviso.Link)
		fmt.Printf("Data de Publicação: %s\n", aviso.DataPublicação)
		
		fmt.Println("Detalhes:")
		fmt.Printf("%-15s: %s\n", "Status", aviso.Status)
		fmt.Printf("%-15s: %s\n", "Evento", aviso.Evento)
		fmt.Printf("%-15s: %s\n", "Severidade", aviso.Severidade)
		fmt.Printf("%-15s: %s\n", "Início", aviso.Início)
		fmt.Printf("%-15s: %s\n", "Fim", aviso.Fim)
		fmt.Printf("%-15s: %s\n", "Descrição", aviso.Descrição)
		fmt.Printf("%-15s: %s\n", "Área", aviso.Área)
		fmt.Printf("%-15s: %s\n", "Link Gráfico", aviso.LinkGráfico)
		
		fmt.Println("-----")


		// Agora, você precisa analisar o conteúdo HTML da descrição
		// Aqui, estamos extraindo informações usando strings e HTML parsing
		reader := strings.NewReader(aviso.Description)
		tokenizer := html.NewTokenizer(reader)

		for {
			tokenType := tokenizer.Next()
			switch tokenType {
			case html.ErrorToken:
				return
			case html.StartTagToken, html.SelfClosingTagToken:
				token := tokenizer.Token()
				if token.Data == "th" {
					tokenType = tokenizer.Next()
					if tokenType == html.TextToken {
						fmt.Printf("%s: ", token.Data)
						fmt.Println(tokenizer.Token().Data)
					}
				} else if token.Data == "td" {
					tokenType = tokenizer.Next()
					if tokenType == html.TextToken {
						fmt.Printf("%s\n", tokenizer.Token().Data)
					}
				}
			}
		}
		fmt.Println("-----")
	}
}
