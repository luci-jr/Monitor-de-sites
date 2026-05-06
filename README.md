# <img src="https://www.svgrepo.com/show/373635/go-gopher.svg" width="32" alt="Go Gopher"> Monitor de Sites

![Go](https://img.shields.io/badge/Go-1.26.2-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Terminal](https://img.shields.io/badge/Interface-Terminal-2E3440?style=for-the-badge)
![Status](https://img.shields.io/badge/Status-Em%20desenvolvimento-2EA44F?style=for-the-badge)


## 📌 Sobre o projeto

O **Monitor de Sites** é uma aplicação de terminal escrita em **Go** para cadastrar, listar, remover e monitorar sites de forma simples.

O programa verifica o status HTTP dos sites cadastrados, registra logs com data e hora e oferece um menu interativo com aparência de painel de monitoramento.

A primeira versão deste projeto foi executada durante um curso introdutório à linguagem Go. Depois, o projeto foi melhorado com novas funcionalidades, ajustes visuais no terminal, persistência simples em arquivos e uma documentação mais completa para publicação no GitHub.

Durante esse processo, foram praticados conceitos como entrada de dados, leitura e escrita de arquivos, requisições HTTP, estruturas condicionais, laços de repetição e organização de funções.

## ✨ Funcionalidades

- ▣ Menu interativo em formato de painel no terminal
- ▣ Boas-vindas personalizada para o usuário
- ▣ Nome do usuário salvo após a primeira execução
- ▣ Cadastro de sites para monitoramento
- ▣ Preenchimento automático de `https://` quando a URL for digitada sem protocolo
- ▣ Listagem dos sites cadastrados
- ▣ Remoção de sites por número
- ▣ Confirmação antes de remover um site
- ▣ Reorganização automática da lista após remoção
- ▣ Monitoramento HTTP dos sites cadastrados
- ▣ Registro de logs com data, hora, site e status

## 🖥️ Prévia do terminal

```text
+------------------------------------------------------------+
| SEJA BEM-VINDO(A), Lucivaldo                               |
+------------------------------------------------------------+
| Este será seu monitoramento pessoal de sites.              |
| Aplicação de terminal escrita em Go.                       |
| Versão do app: 1.2 | Status HTTP dos sites.                |
| Versão do Go: go1.26.2                                     |
+------------------------------------------------------------+

+------------------------------------------------------------+
| MONITOR DE SITES                                           |
+------------------------------------------------------------+
| Status: pronto | Sites cadastrados: 3                      |
| Ciclos: 3 | Intervalo: 5 segundos                          |
+------------------------------------------------------------+
| [1] Iniciar monitoramento                                  |
| [2] Exibir logs                                            |
| [3] Adicionar sites para monitoramento                     |
| [4] Remover site do monitoramento                          |
| [5] Exibir lista de monitoramento                          |
| [0] Sair do programa                                       |
+------------------------------------------------------------+
Escolha uma opção:
```

## 📦 Pacotes utilizados

O projeto foi desenvolvido em **Go 1.26.2** e utiliza apenas pacotes da biblioteca padrão da linguagem.

| Pacote | Função no projeto |
| --- | --- |
| `bufio` | Cria um leitor com buffer para capturar entradas do usuário pelo terminal e ler o arquivo `sites.txt` linha por linha. |
| `fmt` | Exibe mensagens no terminal, monta textos formatados e lê opções digitadas pelo usuário. |
| `net/http` | Faz as requisições HTTP para testar se os sites cadastrados estão respondendo. |
| `os` | Manipula arquivos e interage com o sistema operacional, criando, lendo, escrevendo e abrindo arquivos como `sites.txt`, `log.txt` e `usuario.txt`. |
| `runtime` | Obtém informações do ambiente de execução, como a versão do Go exibida no painel inicial. |
| `strconv` | Converte valores booleanos para texto ao registrar no log se o site estava online ou não. |
| `strings` | Trata textos, removendo espaços, comparando prefixos de URL e normalizando respostas digitadas pelo usuário. |
| `time` | Controla o intervalo entre os ciclos de monitoramento e registra data e hora nos logs. |

## 📋 Requisitos

- Go instalado na máquina
- Versão usada no desenvolvimento: `go1.26.2 linux/amd64`

## 📁 Estrutura do projeto

```text
Monitor de Sites/
├── monitoramento_sites.go
├── sites.txt
├── log.txt
├── usuario.txt
└── README.md
```

## 🚀 Como executar

### 1. Clone o repositório

No terminal, escolha a pasta onde deseja salvar o projeto e execute:

```bash
git clone https://github.com/luci-jr/monitor-de-sites.git
```

### 2. Acesse a pasta do repositório

```bash
cd monitor-de-sites
```

### 3. Acesse a pasta do projeto

Como o projeto está dentro da pasta `Monitor de Sites`, entre nela:

```bash
cd "Monitor de Sites"
```

### 4. Execute diretamente com Go

```bash
go run monitoramento_sites.go
```

### 5. Ou compile o programa

```bash
go build -o monitoramento_sites monitoramento_sites.go
```

Depois execute o binário gerado:

```bash
./monitoramento_sites
```

## 📖 Como usar

Ao iniciar o programa pela primeira vez, ele solicita seu nome e salva essa informação em `usuario.txt`.

Depois disso, o menu principal será exibido com as opções:

```text
[1] Iniciar monitoramento
[2] Exibir logs
[3] Adicionar sites para monitoramento
[4] Remover site do monitoramento
[5] Exibir lista de monitoramento
[0] Sair do programa
```

### ➕ Adicionar site

Digite somente o domínio, se quiser:

```text
comunidade.tech.lucivaldo.cloud/inicio
```

O programa salva automaticamente como:

```text
https://comunidade.tech.lucivaldo.cloud/inicio
```

### 🗑️ Remover site

O programa exibe a lista numerada, pede o número do site e solicita confirmação antes de remover.

Após a remoção, a lista é salva novamente já reorganizada.

### 📊 Monitorar sites

Ao iniciar o monitoramento, o programa percorre os sites cadastrados, faz requisições HTTP e informa se cada site respondeu com sucesso.

Os resultados são gravados no arquivo `log.txt`. Caso esse arquivo ainda não exista, o programa cria automaticamente no primeiro registro de monitoramento.

## 📝 Arquivos de dados

| Arquivo | Função |
| --- | --- |
| `sites.txt` | Armazena os sites cadastrados para monitoramento |
| `log.txt` | Armazena os registros de monitoramento e é criado automaticamente caso não exista |
| `usuario.txt` | Armazena o nome do usuário após a primeira execução |

## 💡 Melhorias possíveis

Como este é um projeto simples, feito como resultado de um curso introdutório à linguagem Go, as melhorias abaixo podem ser implementadas apenas se fizerem sentido para a necessidade de uso:

- Adicionar opção para editar um site cadastrado
- Permitir configurar quantidade de ciclos e intervalo pelo menu
- Evitar cadastro duplicado de URLs
- Exportar logs em formato CSV
- Criar testes automatizados
- Separar o projeto em pacotes Go

## <img src="https://go.dev/doc/gopher/pencil/gopherswrench.jpg" width="32" alt="Go Gopher"> Desenvolvido por

### Lucivaldo

Projeto criado inicialmente durante um curso introdutório de Go e posteriormente aprimorado como prática pessoal de estudo, organização de código e documentação.

[![GitHub](https://img.shields.io/badge/GitHub-luci--jr-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/luci-jr)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-lucivaldojr-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/lucivaldojr/)
[![Email](https://img.shields.io/badge/Email-lucivaldojr25%40gmail.com-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:lucivaldojr25@gmail.com)

<sub>Go Gopher desenhado por Renee French. Imagem disponível na documentação oficial do Go.</sub>
