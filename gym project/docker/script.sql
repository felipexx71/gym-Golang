\connect exemplo_db

create table if not exists person  (
    id serial not null
        constraint pk_person primary key,
    name varchar(64) not null,
    weight decimal(4, 2) not null,
    height decimal(4, 2) not null
);

insert into
    person (name, weight, height)
values ('Jorginho',130 , 150);

select id,
       name,
       weight,
       height
    from person;

-- update product set count = 13 where id = 7;