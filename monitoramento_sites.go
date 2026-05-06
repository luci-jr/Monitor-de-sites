package main

import (
	"bufio"    // Cria leitores com buffer; aqui le o arquivo linha por linha.
	"fmt"      // Formata textos e faz entrada/saida no terminal; aqui usa Println e Scan.
	"net/http" // Trabalha com HTTP; aqui usa Get para acessar sites.
	"os"       // Interage com o sistema operacional; aqui usa Exit e Open.
	"runtime"  // Informa dados da execução; aqui mostra a versão do Go.
	"strconv"  // Converte strings para numeros
	"strings"  // Manipula textos; aqui usa TrimSpace para remover espacos e quebras de linha.
	"time"     // Trabalha com tempo, datas, duracoes e pausas; aqui usa Sleep e Second.
)

const monitoramentos = 3
const delay = 5
const nomeDoPrograma = "Monitor de Sites"
const larguraLinhaMenu = 58
const arquivoUsuario = "usuario.txt"

var leitor = bufio.NewReader(os.Stdin)

func main() {

	exibirIntro()

	for {

		exibirMenu()
		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			exibeLogs()
		case 3:
			adicionarSitesParaMonitoramento()
		case 4:
			removerSiteDoMonitoramento()
		case 5:
			exibirListaDeMonitoramento()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esta opção.")
			os.Exit(-1)
		}
	}

}

func exibirIntro() {
	nome := lerNomeUsuario()
	versao := 1.2

	fmt.Println("")
	exibirBordaMenu()
	exibirLinhaMenu(fmt.Sprintf("SEJA BEM-VINDO(A), %s", nome))
	exibirBordaMenu()
	exibirLinhaMenu("Este será seu monitoramento pessoal de sites.")
	exibirLinhaMenu("Aplicação de terminal escrita em Go.")
	exibirLinhaMenu(fmt.Sprintf("Versão do app: %.1f | Status HTTP dos sites.", versao))
	exibirLinhaMenu(fmt.Sprintf("Versão do Go: %s", runtime.Version()))
	exibirBordaMenu()
}

func lerNomeUsuario() string {
	nomeSalvo, err := os.ReadFile(arquivoUsuario)
	if err == nil {
		nome := strings.TrimSpace(string(nomeSalvo))
		if nome != "" {
			return nome
		}
	}

	fmt.Print("Digite seu nome: ")
	nome, _ := leitor.ReadString('\n')
	nome = strings.TrimSpace(nome)
	if nome == "" {
		nome = "visitante"
	}

	salvarNomeUsuario(nome)
	return nome
}

func salvarNomeUsuario(nome string) {
	err := os.WriteFile(arquivoUsuario, []byte(nome), 0666)
	if err != nil {
		fmt.Println("Não foi possível salvar o nome do usuário:", err)
	}
}

func exibirMenu() {
	totalSites := len(lerArqSites())

	fmt.Println("")
	exibirBordaMenu()
	exibirLinhaMenu("MONITOR DE SITES")
	exibirBordaMenu()
	exibirLinhaMenu(fmt.Sprintf("Status: pronto | Sites cadastrados: %d", totalSites))
	exibirLinhaMenu(fmt.Sprintf("Ciclos: %d | Intervalo: %d segundos", monitoramentos, delay))
	exibirBordaMenu()
	exibirOpcaoMenu(1, "Iniciar monitoramento")
	exibirOpcaoMenu(2, "Exibir logs")
	exibirOpcaoMenu(3, "Adicionar sites para monitoramento")
	exibirOpcaoMenu(4, "Remover site do monitoramento")
	exibirOpcaoMenu(5, "Exibir lista de monitoramento")
	exibirOpcaoMenu(0, "Sair do programa")
	exibirBordaMenu()
}

func exibirBordaMenu() {
	fmt.Println("+------------------------------------------------------------+")
}

func exibirLinhaMenu(texto string) {
	fmt.Printf("| %-*s |\n", larguraLinhaMenu, limitarTexto(texto, larguraLinhaMenu))
}

func exibirOpcaoMenu(numero int, descricao string) {
	exibirLinhaMenu(fmt.Sprintf("[%d] %s", numero, descricao))
}

func limitarTexto(texto string, limite int) string {
	letras := []rune(texto)
	if len(letras) <= limite {
		return texto
	}

	return string(letras[:limite-3]) + "..."
}

func lerComando() int {
	var comandoLido int
	fmt.Print("Escolha uma opção: ")
	fmt.Fscan(leitor, &comandoLido)
	fmt.Println("Opção escolhida:", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := lerArqSites()

	if len(sites) == 0 {
		fmt.Println("Nenhum site cadastrado para monitoramento.")
		return
	}

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			fmt.Println("Testando site:", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")

	}

	fmt.Println("")

}

func adicionarSitesParaMonitoramento() {
	fmt.Println("Adicionar sites para monitoramento")
	fmt.Println("Digite o site sem se preocupar com http:// ou https://.")
	fmt.Println("Se você não informar, vou usar https:// automaticamente.")
	fmt.Println("Digite 0 para cancelar e voltar ao menu.")

	var site string
	fmt.Print("Site: ")
	fmt.Fscan(leitor, &site)

	site = strings.TrimSpace(site)
	if site == "0" || site == "" {
		fmt.Println("Cadastro cancelado. Voltando para o menu principal.")
		return
	}

	site = prepararURL(site)
	salvarSite(site)
	fmt.Println("Site adicionado:", site)
	fmt.Println("Voltando para o menu principal.")
}

func prepararURL(site string) string {
	site = strings.TrimSpace(site)
	siteMinusculo := strings.ToLower(site)

	if strings.HasPrefix(siteMinusculo, "http://") || strings.HasPrefix(siteMinusculo, "https://") {
		return site
	}

	return "https://" + site
}

func removerSiteDoMonitoramento() {
	sites := lerArqSites()

	if len(sites) == 0 {
		fmt.Println("Nenhum site cadastrado para remover.")
		return
	}

	fmt.Println("Escolha o número do site que deseja remover:")
	exibirSites(sites)
	fmt.Println("Digite 0 para cancelar.")

	var indice int
	fmt.Print("Número: ")
	fmt.Fscan(leitor, &indice)

	if indice == 0 {
		fmt.Println("Remoção cancelada.")
		return
	}

	if indice < 1 || indice > len(sites) {
		fmt.Println("Número inválido.")
		return
	}

	siteRemovido := sites[indice-1]
	if !confirmarRemocao(siteRemovido) {
		fmt.Println("Remoção cancelada. Voltando para o menu principal.")
		return
	}

	sites = append(sites[:indice-1], sites[indice:]...)
	if !salvarListaDeSites(sites) {
		return
	}

	fmt.Println("Site removido:", siteRemovido)
	if len(sites) == 0 {
		fmt.Println("A lista de monitoramento ficou vazia.")
		return
	}

	fmt.Println("Lista atualizada:")
	exibirSites(sites)
}

func confirmarRemocao(site string) bool {
	var resposta string

	fmt.Println("Você escolheu remover este site:")
	fmt.Println(site)
	fmt.Print("É exatamente este site? Tem certeza? (s/n): ")
	fmt.Fscan(leitor, &resposta)

	resposta = strings.ToLower(strings.TrimSpace(resposta))
	return resposta == "s" || resposta == "sim"
}

func exibirListaDeMonitoramento() {
	sites := lerArqSites()

	if len(sites) == 0 {
		fmt.Println("Nenhum site cadastrado para monitoramento.")
		return
	}

	fmt.Println("Sites cadastrados para monitoramento:")
	exibirSites(sites)
}

func exibirSites(sites []string) {
	for i, site := range sites {
		fmt.Println(i+1, "-", site)
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		registroLog(site, false)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("Site:", "[", site, "]", "foi carregado com sucesso!")
		registroLog(site, true)
	} else {
		fmt.Println("Site:", "[", site, "]", "está com problemas. Status Code:", resp.StatusCode)
		registroLog(site, false)
	}
}

func lerArqSites() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Houve um erro inesperado:", err)
		return sites
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linha := strings.TrimSpace(scanner.Text())
		if linha != "" {
			sites = append(sites, linha)
		}
	}

	return sites
}

func salvarSite(site string) {
	arquivo, err := os.OpenFile("sites.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Houve um erro ao salvar o site:", err)
		return
	}
	defer arquivo.Close()

	if info, err := arquivo.Stat(); err == nil && info.Size() > 0 {
		arquivo.WriteString("\n")
	}

	arquivo.WriteString(site)
}

func salvarListaDeSites(sites []string) bool {
	conteudo := strings.Join(sites, "\n")
	err := os.WriteFile("sites.txt", []byte(conteudo), 0666)
	if err != nil {
		fmt.Println("Houve um erro ao atualizar a lista de sites:", err)
		return false
	}

	return true
}

func registroLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer arquivo.Close()

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 - ") + site + "- online: " + strconv.FormatBool(status) + "\n")
}

func exibeLogs() {
	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}
