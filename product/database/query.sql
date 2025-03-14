-- name: Get :many
select id, name from products.products 
where name ilike concat('%', @name::text ,'%');

-- name: GetCount :one
select count(*) from products.products 
where name ilike concat('%', @name::text ,'%');

-- name: Insert :exec
insert into products.products (id, name) values ($1, $2);

-- name: Update :exec
update products.products set name=$1 where id=$2;

-- name: Delete :exec
delete from products.products where id=$1;