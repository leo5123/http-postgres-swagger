## Informação geral
API que serve como o back-end de uma aplicação TO DO list.

	
## Setup

Inicializar o projeto:

```
$ cd ToDoListAPI-go-postgres
$ docker-compose up 
```

#### A api será iniciada em:   
> http://localhost:5050/

#### Serviço do swagger será iniciado em: 
> http://localhost:81
>Testes localizados em main_test.go



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
