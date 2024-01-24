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
