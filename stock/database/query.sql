-- name: GetStock :one
select sqlc.embed(s) from products.stocks s where s.product_id = @productID;

-- name: GetStockReservation :one
select sqlc.embed(sr) from products.stock_reservations sr where sr.order_number = @orderNumber;

-- name: StockReservationExists :one
select exists(select 1 from products.stock_reservations where order_number = @orderNumber and cancel_date is null);

-- name: CancelStockReservation :exec
update products.stock_reservations set cancel_date = @cancelDate where id = @ID;

-- name: UpdateStockReserve :exec
update products.stocks set reserved_quantity = @reservedQuantity, version = version+1 where id = @ID and version = @version;

-- name: AddStockReservation :exec
insert into products.stock_reservations (id, product_id, order_number, quantity, create_date, cancel_date) values ($1, $2, $3, $4, $5, $6);

-- name: GetStocks :many
select * from products.stocks
where (sqlc.narg('productID')::uuid is null or product_id = sqlc.narg('productID')::uuid) 
  and (
    (sqlc.narg('from')::timestamp is null or create_date > sqlc.narg('from')::timestamp)
     and (sqlc.narg('to')::timestamp is null or create_date < sqlc.narg('to')::timestamp)
     );

-- name: Insert :exec
insert into products.stocks (id, product_id, quantity, reserved_quantity, version, create_date, last_update_date) 
values ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateStockQuantity :exec
update products.stocks set quantity = @quantity, version = version+1 where id = @ID and version = @version;