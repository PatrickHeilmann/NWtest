# NWtest
Projeto NEOWAY - Prova de conhecimento

Projeto criado para importação de um arquivo txt e o procedimento de validação do cpf e cnpj.
Utilizando lingagem de programação GO e banco de dados PostgreSQL.

Data: 29/03/2020
Autor: Patrick Heilmann
Contato: psheilmann@hotmail.com

Instalação e configuração

Será utilizado os seguintes softwares:
* Sistema Operacional Windows 10 64bits;
* Linguagem de programação google (golang) versão "go version go1.14.1 windows/amd64";
* Controle de versão git versão "2.26.0 64 bit";
* Github desktop versão "2.4.0";
* Banco de dados PostgreSQL versão "12.2 windows x64";

Instalação:
Realizar instalação padrão para todos os aplicativos citados.

Configuração:
GOLANG
Criar pasta no path do usuário local "go\src\github.com\pgrpatrick\nwfiletext"
Ex:  
"C:\Users\"usuário local"\go\src\github.com\pgrpatrick\nwfiletext"

Incluir o manipulação dados do banco:
Executar o CMD.exe do windows;
Na pasta criada "C:\Users\"usuário local"\go\src\github.com\pgrpatrick\nwfiletext" executar o comando:
"go get github.com/lib/pq"
Será criado a pasta 
"C:\Users\"usuário local"\go\src\github.com\lib\pq\..."


GITHUB
Realizar o clone do projeto no endereço:
$ git clone https://github.com/PatrickHeilmann/NWtest


POSTGRESQL
Criação da dase de dados, tabelas e funtions.
Obs.: Usuário "postgres" e senha "123456"; O usuário deve ter permisão para criação de base de dados, tabelas, functions etc, 
No banco postgres, executar os seguintes scritps em ordem:

Primeiro: 
Criação da base de dados "nw_db".
script_createdb_postgresql.sql

Segundo: 
Criação de tabelas e funcions.
Obs.: Dentro do banco "nw_db", executar o seguinte script:
script_tables_function_postgresql.sql

Execução do aplicativo:
Dentro do diretório do projeto "NWtest", executar o arquivo "main.exe"
