## Informação geral
API que serve como o back-end de uma aplicação TO DO list.

	
## Setup

Iniciar a imagem docker:
```
$ docker-compose up
```

##### A api será iniciada em:   
> http://localhost:5050/

##### Serviço do swagger será iniciado em: 
> http://localhost:81


Instalar dependencias do projeto e iniciar aplicação na porta 5050
```
$ cd ToDoListAPI-go-postgres
$ go run main.go 
```

### Exemplo de método disponivel:
> http://localhost:5050/api/tarefas

#### Resposta:
  ```json
    {
        "id": 1,
        "nome": "Tarefa 1",
        "filtro": "Javascript",
        "pronto": true
    },
    {
        "id": 2,
        "nome": "Tarefa 2",
        "filtro": "React",
        "pronto": false
    }
  ```
