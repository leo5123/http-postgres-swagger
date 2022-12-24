create table tarefas(
    id serial primary key,
    nome varchar,
    filtro varchar,
    pronto boolean
);

INSERT INTO tarefas(nome, filtro, pronto) VALUES
('Tarefa 1', 'Javascript', TRUE),
('Tarefa 2', 'React', FALSE)