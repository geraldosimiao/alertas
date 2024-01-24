# Alertas
Testes de códigos para um programa que reúna e filtre alertas meteorológicos do INMET

## Dependências
Este programa usa a biblioteca padrão para fazer requisições HTTP (net/http) e analisar XML (encoding/xml). Certifique-se de ter uma versão mais recente do Go que inclua essas bibliotecas.

## Observações
Para compilar e executar um programa em Go, geralmente, você não precisa de muitos arquivos adicionais. No entanto, aqui estão algumas considerações:

    Código-fonte Go: O código-fonte do programa em Go é arquivo com extensão .go.

    Módulo Go: Se estiver usando módulos Go (introduzidos no Go 1.11), você pode ter um arquivo go.mod. Isso não é obrigatório, mas é uma boa prática.

    Arquivo de configuração Go (opcional): Você pode incluir um arquivo de configuração Go chamado go.sum que contém hashes de módulos e suas versões correspondentes. Isso é usado para garantir a integridade dos módulos.

    Ambiente Go configurado: Verifique se o Go está instalado no seu sistema e se o seu ambiente Go está configurado corretamente. Certifique-se de ter o $GOPATH e $GOBIN definidos.

    Bibliotecas externas: Se você estiver usando bibliotecas externas, o Go baixará automaticamente as dependências necessárias durante o processo de compilação.

## Como rodar o programa
Abra um terminal numa pasta de sua preferência e rode o comando:

    git clone https://github.com/geraldosimiao/alertas.git

Em seguida entre na pasta criada:
    
    cd alertas
    
Para rodar o programa basta executar:

    go run main.go <número de horas>

Aqui um exemplo:

    go run main.go 15
(isso vai retornar os avisos do INMET nas últimas 15 horas)

# Em construção
Ainda vou preparar instruções para compilar ele em go, e também ainda pretendo empacotar em RPM e distribuir via repositorio copr pro Fedora.
