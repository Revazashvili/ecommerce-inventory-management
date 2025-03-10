create table products.stocks(
    id uuid primary key ,
    product_id uuid,
    quantity int,
    reserved_quantity int,
    version int,
    create_date timestamp,
    last_update_date timestamp
);

create table products.stock_reservations(
    id uuid primary key ,
    product_id uuid,
    order_number uuid,
    quantity int,
    create_date timestamp,
    cancel_date timestamp
);
