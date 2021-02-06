CREATE DATABASE todolist;

create type task_state as enum('TODO', 'COMPLETED', 'DISCARTED');

CREATE TABLE tasks (
  id VARCHAR(138),
  title VARCHAR(200) NOT NULL,
  description TEXT NOT NULL,
  due_date DATETIME NOT NULL,
  state task_state NOT NULL,
  CONSTRAINT tasks_pk PRIMARY KEY(id)
);