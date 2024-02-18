-- name: CreateRefreshToken :one
INSERT INTO
    "refresh_token" (
        "id",
        "token",
        "user_id",
        "expiry"
    )
VALUES
    ($1, $2, $3, $4) RETURNING *;