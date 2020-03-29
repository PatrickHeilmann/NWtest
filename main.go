package main

import (
    "fmt"
    "regexp"
    "database/sql"
    "bufio"
    "os"
    "log"
    "strings"
    _ "github.com/lib/pq"
)

const (
  host = "localhost"
  port = 5432
  user = "postgres"
  password = "Cancer21"
  dbname = "nw_db"
  localFile = "base_teste.txt"
)   

type ticket struct {
    nuCpfCnpj           string 
    stPrivate           string
    stIncompleto        string
    dtUltimaCompra      string
    vlTicketMedio       string
    vlTicketUltimaCompra string
    nuCnpjLojaFrequente    string
    nuCnpjLojaUltimaCompra string
}

func (t *ticket) setValues(values []string) {
    t.nuCpfCnpj = values[0]
    t.stPrivate = values[1]
    t.stIncompleto = values[2]
    t.dtUltimaCompra = values[3]
    t.vlTicketMedio = values[4]
    t.vlTicketUltimaCompra = values[5]
    t.nuCnpjLojaFrequente = values[6]
    t.nuCnpjLojaUltimaCompra = values[7]
}

func main() {
    fmt.Println("Iniciando importação do arquivo: ",localFile )
    // abrir o arquivo
    contentFile, err := os.Open(localFile)
    
    //contentFile, err := ioutil.ReadFile(localFile)
    if err != nil {
        log.Panicf("Falha ao ler o arquivo: %s", err)
    }

    // Retorna as linhas lidas e um erro se ocorrer algum erro no scanner
    //fmt.Println(contentFile)
    defer contentFile.Close()

    scanner := bufio.NewScanner(contentFile)
    scanner.Split(bufio.ScanLines)

    var lines []string
  
    for scanner.Scan() {
      lines = append(lines, scanner.Text())
    }
    //INICIALIZANDO A STRUCT

    structTicket := ticket{}
    //COMPILANDO A REGEX PARA TIRAR OS ESPAÇOS NO TXT
   // regexCompiled, _ := regexp.Compile("\\s+")
    regexCompiled := regexp.MustCompile(`\s+.*?`)

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
    panic(err)
    }
    defer db.Close()

    //tudo esta dentro de lines, lendo linha a linha
    fmt.Println("Iniciando gravação dos registros...")
    for i, line := range lines {
        if i == 0 {
            fmt.Println("Cabeçalho do arquivo texto descartado...")
            //fmt.Println(line)
        } else {
            //fmt.Println("Linha inteira:")
            //fmt.Println(line)
            //regex compilada, substitui a regex que encontrar c uma string c apenas um espaço
            replaced := regexCompiled.ReplaceAllString(line, " ")
            //separando as strings em um slice p extrair o valor
            splited := strings.Split(replaced, " ")
            //fmt.Println("Linha formatada (splited):")
            //fmt.Println(splited)

            //metodo criado para setar os valores da struct
            structTicket.setValues(splited)

            //linha a ser inserida
            //fmt.Println("Inserindo linha (structTicket): ",structTicket.nucpfcnpj)
            sqlStatement := `
            INSERT INTO tb_ticket (nu_cpf_cnpj,st_private,st_incompleto,dt_ultima_compra,vl_ticket_medio,vl_ticket_ultima_compra,nu_cnpj_loja_frequente,nu_cnpj_loja_ultima_compra)
            VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`
             _, err = db.Exec(sqlStatement,
                     structTicket.nuCpfCnpj,
                     structTicket.stPrivate,
                     structTicket.stIncompleto,
                     structTicket.dtUltimaCompra,
                     structTicket.vlTicketMedio,
                     structTicket.vlTicketUltimaCompra,
                     structTicket.nuCnpjLojaFrequente,
                     structTicket.nuCnpjLojaUltimaCompra) 
            //fmt.Println(sqlStatement)
            if err != nil {
            panic(err)
            }

            //fmt.Println(sqlStatement)
            //fmt.Println("Inserido linha:")
            //fmt.Println(structTicket)
            }
        }
        fmt.Println("Finalizado o INSERT no banco de dados.")
         

    // Update tabela tb_ticket com a validacao do CPF e CNPF, com funcoes no banco postgresql
    sqlUpdate := `
    Update tb_ticket set 
    st_valido_cpf_cnpj = 
        case when char_length(nu_cpf_cnpj)=(14) 
        then fc_valida_cpf(replace(replace(replace( nu_cpf_cnpj, '.', '' ),'/',''),'-',''), false)  
        when char_length(nu_cpf_cnpj)=(18) 
        then fc_valida_cnpj(replace(replace(replace( nu_cpf_cnpj, '.', '' ),'/',''),'-',''), false) else null end,
    nu_cnpj_loja_frequente = 
        case when char_length(nu_cnpj_loja_frequente)=(18) 
        then fc_valida_cnpj(replace(replace(replace(nu_cnpj_loja_frequente, '.', '' ),'/',''),'-',''), false) else null end,
    st_valido_cnpj_ultima_compra = 
        case when char_length(nu_cnpj_loja_ultima_compra)=(18) 
        then fc_valida_cnpj(replace(replace(replace( nu_cnpj_loja_ultima_compra, '.', '' ),'/',''),'-',''), false) else null end
    `
    fmt.Println("Validando CPF e CNPJ...")
    _, err = db.Exec(sqlUpdate)
    if err != nil {
      panic(err)
    }
    fmt.Println("Registros CPF e CNPJ verificados.")
    fmt.Println("Fim.")
return
}