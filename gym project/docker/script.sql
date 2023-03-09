create database postgres

\connect postgres

create table if not exists person  (
    id serial not null
        constraint pk_person primary key,
    name varchar(64) not null,
    weight decimal(4, 2) not null,
    height decimal(3, 2) not null
);


select id,
       name,
       weight,
       height
    from person;
