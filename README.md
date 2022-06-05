# Sensor

## Criação do projeto
Para gerenciar os pacotes com mais flexibilidade o golang 1.13> possibilita criar  você criará seu próprio módulo Go público e adicionará um pacote ao seu novo módulo.Para informações detalhadas acesso o [link](https://www.digitalocean.com/community/tutorials/how-to-use-go-modules). Para este projeto nosso módulo chamará sensor.

`go mod init sensor`

## Pasta DB
Inicialmente este projeto seguiu o padrão data first. Esse padrão foi escolhido por ser um reescritura de código, sem nenhuma alteração nas entidades. No entanto com evolução deste projeto o qual será criado diversos micro-serviços o padrão class first será adotado.

### Schema Migrations
Para este projeto utilizamos migration para versionamento da base de dados, criadas. O repositório do migrate para golang está disponível no [link](https://github.com/golang-migrate/migrate). A seguir está descrito o passo a passo para estrurar migrations dentro da sua solução.

1. Instale o pacote:

    `go get github.com/golang-migrate/migrate`

2. Para criar o sql dos seus serviços execute o comando a seguir. SERVICE representa o nome do serviço, isso criará uma pasta onde os arquivos sql serão do serviço serão inseridos, NAME será o nome do arquivo sql, será gerado dois, são eles:up e down.
|

    ` make migrate-create SERVICE=nome-servico NAME=nome-do-arquivo-sql`

3. Todas as automações referentes ao postgres serão inseridas dentro da pasta postgres.

4. Rode o docker utilizando o comando a seguir:

    `make run-docker-postgres`
    
Caso na execução ocorra o erro abaixo, siga as instruções abaixo:

        /usr/local/bin/docker-entrypoint.sh: /docker-entrypoint-initdb.d/init-database.sh: /bin/bash: bad interpreter: Permission denied

* Pegue o ID do seu container e de stop no container. Adicione permissão de execução para os arquivos .sh que você deseja.  E execute novamente o container.

    `docker container ls`

    `docker stop a53547da805e `
    
    `cd db/migrations/postgres`

    `chmod +x init-database.sh`

    `docker-compose up postgres`



Exemplo da resposta do comando ls:    
| CONTAINER ID | IMAGE | COMMAND | CREATED | STATUS | PORTS | NAMES|
|--- |--- |--- |--- |--- |--- |--- |
| a53547da805e | postgres:13.4-alpine| "docker-entrypoint.s…"| 37 minutes ago | Up 4 minutes | 0.0.0.0:5432->5432/tcp  3 | ensor_postgres_1 |


