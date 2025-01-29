-- name: GetProduct :one
SELECT * FROM product
WHERE id = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM product
ORDER BY name;

-- name: CreateProduct :one
INSERT INTO product (
    name, price, available
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM product
WHERE id = $1;

-- name: UpdateProduct :exec
UPDATE product
SET name = $2, price = $3, available = $4
WHERE id = $1;

-- name: SumProductPrice :one
SELECT SUM(price)::FLOAT FROM product ;