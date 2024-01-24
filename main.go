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
	Status      string `xml:"Status"`  // Corrigido para Status com a letra inicial em maiúsculo
	Evento      string `xml:"Evento"`
	Severidade  string `xml:"Severidade"`
	Início      string `xml:"Início"`
	Fim         string `xml:"Fim"`
	Descrição   string `xml:"Descrição"`  // Adicionado o campo Descrição
	Área        string `xml:"Área"`
	LinkGráfico string `xml:"Link Gráfico"`  // Alterado para Link Gráfico
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
		fmt.Printf("Título: %s\n", aviso.Title)
		fmt.Printf("Link: %s\n", aviso.Link)
		fmt.Printf("Data de Publicação: %s\n", aviso.Published)

		// Parse HTML dentro do campo de descrição
		descricaoHTML := aviso.Description
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
				break
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

		// Exibir detalhes do aviso
		for chave, valor := range detalhes {
			fmt.Printf("%-15s: %s\n", chave, valor)
		}

		fmt.Println("-----")
	}
}
