# Alertas
Alertas é um programa que reúne e filtra avisos meteorológicos do INMET [Instituto Nacional de Meteorologia](https://alertas2.inmet.gov.br/)  

## Dependências
Este programa usa a biblioteca padrão para fazer requisições HTTP (net/http) e analisar XML (encoding/xml). Certifique-se de ter uma versão mais recente do Go que inclua essas bibliotecas.

## Observações
Para compilar e executar um programa em Go, geralmente, você não precisa de muitos arquivos adicionais. No entanto, aqui estão algumas considerações:

    Ambiente Go configurado: Verifique se o Go está instalado no seu sistema e se o seu ambiente Go está configurado corretamente. Certifique-se de ter o $GOPATH e $GOBIN definidos.
    Bibliotecas externas: Se você estiver usando bibliotecas externas, o Go baixará automaticamente as dependências necessárias durante o processo de compilação.

## Como compilar o programa
1- Abra um terminal numa pasta de sua preferência e rode o comando:

    git clone https://github.com/geraldosimiao/alertas.git

2- Em seguida entre na pasta criada:
    
    cd alertas
    
3- Compile para gerar o binário:

    go build

4- Execute o binário:

    ./alertas <número de horas>
Ex. ```./alertas 12``` vai retornar os avisos do INMET nas últimas 12 horas.

![Screenshot_20240124_224503](https://github.com/geraldosimiao/alertas/assets/72357821/880e85b0-41af-4f2e-86ac-e012909704a3)


# Instalação em Fedora (RPM)
Se tiver Fedora Linux e desejar instalar via repositorio copr, pode rodar os comandos como descrito no [link do repositório](https://copr.fedorainfracloud.org/coprs/geraldosimiao/alertas/). Se preferir somente baixar um rpm e instalar manualmente, pode usar o pacote que está em "releases".

Autor: [Geraldo Simião](https://fedoraproject.org/wiki/User:Geraldosimiao)

