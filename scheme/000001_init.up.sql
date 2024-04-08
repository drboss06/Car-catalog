CREATE TABLE people (
                        id serial not null unique,
                        name varchar(255) not null,
                        surname varchar(255) not null,
                        patronymic varchar(255)
);

CREATE TABLE car (
                     id serial not null unique,
                     regNum varchar(255) not null,
                     mark varchar(255) not null,
                     model varchar(255) not null,
                     year int,
                     owner int references people(id)
);
