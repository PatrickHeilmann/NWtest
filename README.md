# NWtest
Projeto NEOWAY - Prova de conhecimento

Concebido 
Data: 29/03/2020
Autor: Patrick Heilmann
Contato: psheilmann@hotmail.com

Instalação e configuração

Para este projeto será utilizado os seguintes softwares:
* Sistema Operacional Windows 10 64bits;
* Linguagem de programação google (golang) versão "go version go1.14.1 windows/amd64";
* Controle de versão git versão "2.26.0 64 bit";
* Github desktop versão "2.4.0";
* Banco de dados PostgreSQL versão "12.2 windows x64";

Neste projeto utilizaremos a linguagem golang para realizar a importação de um arquivo txt e o procedimento de validação do cpf e cnpj.
Instalação:
Realizar instalação padrão para todos os aplicativos citados.

Configuração:
GOLANG - Criar pasta no path do usuário local "go\src\github.com\pgrpatrick\nwfiletext"
Ex:  
C:\Users\"usuário local"\go\src\github.com\pgrpatrick\nwfiletext
Incluir o manipulação dados do banco:
Executar o CMD.exe do winows;
Na pasta criada "C:\Users\"usuário local"\go\src\github.com\pgrpatrick\nwfiletext" executar o comando:
"go get github.com/lib/pq"

Baixando o projeto do GITHUB


Criação da dase de dados e tabelas
No banco postgres(pgadmin), abrir o script "script_postgresql.sql" e executar.


