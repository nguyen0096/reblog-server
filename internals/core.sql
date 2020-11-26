drop database if exists reblog with (force);
drop role if exists reblog;


create database reblog;
create user reblog encrypted password 'reblog';
grant all privileges on database reblog to reblog;
alter database reblog owner to reblog;
\c reblog

create schema rb_core; -- ISSUE: need to grant permission to access schema, or is there any way to cascade down permission from database

create table rb_core.user (
    id serial not null,
    username character varying(100) not null,
    first_name character varying(100),
    last_name character varying(100),
    address character varying(200),
    primary key (id)
);

grant all privileges on table rb_core.user to reblog;
