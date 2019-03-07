CREATE TABLE students (
  id SERIAL,
  fio varchar(300) NOT NULL,
  info text NULL,
  score int NOT NULL
);

INSERT INTO students (fio, info, score)
VALUES ('Vasily Romanov', 'company: Mail.ru Group', '10'),
       ('Oleg Nesterov', 'company: SoftServe', '20');

DROP TABLE students

select * from students