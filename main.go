package main

import (
    "bytes"
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

func extrairCampos(aviso *Aviso) error {
    // Parse a descrição como HTML
    descHTML, err := html.Parse(strings.NewReader(aviso.Description))
    if err != nil {
        return err
    }

    // Encontrar todas as células da tabela
    cells := findTableCells(descHTML)

    // Atribuir valores aos campos correspondentes
    if len(cells) >= 8 {
        aviso.Status = cells[0]
        aviso.Evento = cells[1]
        aviso.Severidade = cells[2]
        aviso.Início = cells[3]
        aviso.Fim = cells[4]
        aviso.Descrição = cells[5]
        aviso.Área = cells[6]
        aviso.LinkGráfico = cells[7]
    }

    return nil
}

func findTableCells(n *html.Node) []string {
    var cells []string

    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "td" {
            cells = append(cells, textContent(n))
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }

    f(n)

    return cells
}

func textContent(n *html.Node) string {
    var buffer bytes.Buffer
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.TextNode {
            buffer.WriteString(n.Data)
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(n)
    return buffer.String()
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

	// Adicionar a linha abaixo para imprimir o XML completo
	// fmt.Printf("Resposta da API:\n%s\n", conteúdo)

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
