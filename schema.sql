drop database if exists devgo cascade;
create database if not exists devgo;
set database = devgo;

create table if not exists users(
	   id serial not null primary key,
	   username varchar not null unique,
	   email varchar not null unique,
	   password bytes not null,
	   name varchar not null
);
