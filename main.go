package main

import (
    "fmt"
    "database/sql"
    "regexp"
    "bufio"
    "os"
    "log"
)

const (
  host = "localhost"
  port = 5432
  user = "nw_patrick"
  password = "123456"
  dbname = "nw_db"
  localFilePath = "C:\\users\\Gizeli.zancanaro\\go\\src\\github.com\\pgrpatrick\\NWtest\\base_testepatrick.txt"
  localFile = "base_testepatrick.txt"
)   

func main() {
    var dadosGravacao []string
    dadosGravacao, err := lerArquivo(localFilePath)
    if err != nil {
        log.Fatalf("Erro de leitura:", err)
    }
    inseriBancoDados(dadosGravacao)
    if err != nil {
        log.Fatalf("Erro de gravacao:", err)
    }
}

// Funcao leitura do arquivo, retorna a string apenas com dados necessarios
func lerArquivo(localFilePath string) ([]string, error) {

    // abrir o arquivo
    dadosArquivoTexto, err := os.Open(localFilePath)
    if err != nil {
        return nil, err
    }
   
    // fecha o arquivo
    defer dadosArquivoTexto.Close()
   
    //faz a leitura das linhas
    scanLinhas := bufio.NewScanner(dadosArquivoTexto)
    
    var linhaArray []string
    //scanear linha a linha
    for scanLinhas.Scan(){
        
        linhaScanner := scanLinhas.Text()
        //compila com expressao regular o delimitador dos espacos em branco
        re := regexp.MustCompile(`\s+.*?`)
        

        //separacao das colunas conforme formula re
        coluna := re.Split(linhaScanner, -1)

        //teste para saber se está tudo ok
        for i := range(coluna) {

            fmt.Println(coluna[i])
        }
    }
    // Retorna as linhas lidas e um erro se ocorrer algum erro no scanner
    return linhaArray, scanLinhas.Err()
}

// Funcao que escreve um texto no arquivo e retorna um erro caso tenha algum problema
func inseriBancoDados(linhas []string) error {
    
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
    panic(err)
    }
    defer db.Close()
 
    sqlquerie := `
    INSERT INTO tb_ticket(nu_cpf_cnpj,st_private,st_incompleto,dt_ultima_compra,vl_ticket_medio,vl_ticket_ultima_compra,nu_cnpj_loja_frequente character,nu_cnpj_loja_ultima_compra) 
    VALUES($1, $2, $3, $4, $5, $6, 7$, 8$) 
    returning id_ticket.Scan(&lastInsertId)`

    var lastInsertId int
    id_ticket := 0
    err = db.QueryRow(sqlquerie).Scan(&id_ticket)
    //checkErr(err) -- undefined checkErr
    
    fmt.Println("Ultimo id_ticket inserido=", lastInsertId)
    return err
}
