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

func obterAvisos(apiURL string, filtro map[string]string) ([]Aviso, error) {
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

	var avisosFiltrados []Aviso
	for _, aviso := range feed.Channel.Items {
		if atendeCritérios(aviso, filtro) {
			avisosFiltrados = append(avisosFiltrados, aviso)
		}
	}

	return avisosFiltrados, nil
}

func atendeCritérios(aviso Aviso, filtro map[string]string) bool {
	for chave, valor := range filtro {
		switch chave {
		case "status":
			if aviso.Status != valor {
				return false
			}
		case "evento":
			if aviso.Evento != valor {
				return false
			}
		case "severidade":
			if aviso.Severidade != valor {
				return false
			}
		case "início":
			if aviso.Início != valor {
				return false
			}
		case "fim":
			if aviso.Fim != valor {
				return false
			}
		case "área":
			if aviso.Área != valor {
				return false
			}
		// Adicione mais casos conforme necessário
		}
	}

	return true
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Por favor, forneça um critério de filtro. Exemplo: go run main.go Crítica")
		os.Exit(1)
	}

	critérioFiltro := os.Args[1]

	apiURL := "https://apiprevmet3.inmet.gov.br/avisos/rss"

	// Exemplo de filtro baseado no argumento de linha de comando
	filtro := map[string]string{
		"severidade": critérioFiltro,
	}

	avisos, err := obterAvisos(apiURL, filtro)
	if err != nil {
		fmt.Println("Erro ao obter avisos:", err)
		os.Exit(1)
	}

	// Linhas para imprimir informações dos avisos
	fmt.Printf("Avisos retornados: %d\n", len(avisos))
	for _, aviso := range avisos {
    		fmt.Printf("Severidade: %s\n", aviso.Severidade)
    	// Adicionar mais informações conforme necessário
	}

	fmt.Printf("Avisos encontrados para o critério de filtro '%s': %d\n", critérioFiltro, len(avisos))

	if len(avisos) == 0 {
		fmt.Println("Nenhum aviso encontrado para o critério de filtro:", critérioFiltro)
	} else {
		for _, aviso := range avisos {
			fmt.Println("Título:", aviso.Title)
			fmt.Println("Link:", aviso.Link)
			fmt.Println("Data de Publicação:", aviso.Published)
			fmt.Println("Status:", aviso.Status)
			fmt.Println("Evento:", aviso.Evento)
			fmt.Println("Severidade:", aviso.Severidade)
			fmt.Println("Início:", aviso.Início)
			fmt.Println("Fim:", aviso.Fim)
			fmt.Println("Área:", aviso.Área)
			fmt.Println()
		}
	}
}
