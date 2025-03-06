create database products;

create schema products;

create table products.products(
    id uuid primary key ,
    name text
);