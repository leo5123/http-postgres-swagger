
## Informação geral
API que serve como o back-end de uma lista de tarefas.


	
## Setup
Para rodar esse projeto, instale-o localmente usando o npm:

```
$ cd ToDoListAPI-go-postgres
$ go run main.go 
```
Para preparar o swagger para o uso

```
$ cd ToDoListAPI-go-postgres
$ cd swagger

$ docker pull swaggerapi/swagger-editor
$ docker run -p 80:8080 -v ${pwd}:/tmp -e SWAGGER_FILE=/tmp/openapi.yaml swaggerapi/swagger-editor
```


