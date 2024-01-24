package main

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"strings"
	"bytes"
	"os"
	"time"
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

// Função para fazer a solicitação à API e verificar seu recebimento
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

// Função para extrair detalhes do HTML, incluindo o link gráfico
func extrairDetalhesHTML(descricaoHTML string) map[string]string {
	// Utilizando a biblioteca "golang.org/x/net/html" para analisar HTML
	doc, err := html.Parse(strings.NewReader(descricaoHTML))
	if err != nil {
		fmt.Println("Erro ao analisar HTML:", err)
		return nil
	}

	// Mapa para armazenar os detalhes do aviso
	detalhes := make(map[string]string)

	// Função para percorrer o HTML e extrair detalhes
	var f func(*html.Node)
	var chaveAtual string

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "th" {
			// Se o nó é uma tag <th>, obtemos a chave atual
			chaveAtual = strings.TrimSpace(n.FirstChild.Data)
		} else if n.Type == html.ElementNode && n.Data == "td" {
			// Se o nó é uma tag <td>, obtemos o valor associado à chave atual
			valor := strings.TrimSpace(renderNodeTextContent(n))
			detalhes[chaveAtual] = valor
		} else if n.Type == html.ElementNode && n.Data == "a" {
			// Se o nó é uma tag <a>, procuramos o link dentro dela
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					detalhes["Link Gráfico"] = attr.Val
				}
			}
		}

		// Recursivamente chama a função para os filhos do nó
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	// Inicia o processamento recursivo do HTML
	f(doc)

	return detalhes
}

// Função para renderizar o conteúdo de texto de um nó HTML
func renderNodeTextContent(n *html.Node) string {
	var buf bytes.Buffer
	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			buf.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)
	return buf.String()
}

// Função para imprimir campos na ordem específica
func printField(label string, value string) {
	fmt.Printf("%-15s: %s\n", label, value)
}


// Função principal
func main() {
	// Aceitar um argumento de linha de comando para o intervalo de tempo em horas
	if len(os.Args) < 2 {
		fmt.Println("Por favor, forneça um intervalo de tempo em horas.")
		os.Exit(1)
	}

	intervalo, err := time.ParseDuration(os.Args[1] + "h")
	if err != nil {
		fmt.Println("Erro ao analisar o intervalo de tempo:", err)
		os.Exit(1)
	}

	// Definir a data de corte como o tempo atual subtraindo o intervalo
	dataCorte := time.Now().Add(-intervalo)

	// Obter avisos usando a lógica existente
	avisos, err := obterAvisos("https://apiprevmet3.inmet.gov.br/avisos/rss")
	if err != nil {
		fmt.Println("Erro ao obter avisos:", err)
		os.Exit(1)
	}

	// Exibir avisos que estão após a data de corte
	fmt.Printf("Avisos retornados: %d\n", len(avisos))
	for i, aviso := range avisos {
		// Converter a data de publicação para o tipo time.Time
		dataPublicacao, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", aviso.Published)
		if err == nil {
			fmt.Printf("Data de Publicação: %s\n", dataPublicacao.Format("Mon, 02 Jan 2006 15:04:05 -0700"))
		} else {
			fmt.Println("Erro ao analisar a data de publicação:", err)
		}

		if dataPublicacao.After(dataCorte) {
			fmt.Printf("Aviso #%d\n", i+1)
			fmt.Printf("Título: %s\n", aviso.Title)
			fmt.Printf("Link: %s\n", aviso.Link)
			fmt.Printf("Data de Publicação: %s\n", aviso.Published)

			// Extrair detalhes do HTML e exibir na ordem correta
			detalhes := extrairDetalhesHTML(aviso.Description)
			printField("Status", detalhes["Status"])
			printField("Evento", detalhes["Evento"])
			printField("Severidade", detalhes["Severidade"])
			printField("Início", detalhes["Início"])
			printField("Fim", detalhes["Fim"])
			printField("Descrição", detalhes["Descrição"])
			printField("Área", detalhes["Área"])
			printField("Link Gráfico", detalhes["Link Gráfico"])

			fmt.Println("-----")
		}
	}
}
