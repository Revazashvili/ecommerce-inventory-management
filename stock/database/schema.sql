create schema if not exists products;

create table products.stocks(
    id uuid not null primary key ,
    product_id uuid not null ,
    quantity int not null ,
    reserved_quantity int not null ,
    version int not null ,
    create_date timestamp not null ,
    last_update_date timestamp not null
);

create table products.stock_reservations(
    id uuid not null primary key ,
    product_id uuid not null ,
    order_number uuid not null ,
    quantity int not null ,
    create_date timestamp not null ,
    cancel_date timestamp
);