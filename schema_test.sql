drop database if exists devgo_test cascade;
create database if not exists devgo_test;
set database = devgo_test;

create table if not exists users(
	   id serial not null primary key,
	   username varchar not null unique,
	   email varchar not null unique,
	   password varchar not null,
	   name varchar not null
);
