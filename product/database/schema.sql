create schema products;

create table products.products(
    id uuid not null primary key ,
    name text not null
);