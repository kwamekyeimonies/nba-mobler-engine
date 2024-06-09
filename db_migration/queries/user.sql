-- name: CreateUserAccount :one
INSERT INTO "user"(
    id,
    firstName,
    lastName,
    createdAt,
    updateAt,
    deletedAt,
    activatedAt,
    email,
    phoneNumber,
    dateOfBirth,
    password,
    isDeleted
)VALUES(
           $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12
       )RETURNING *;


-- name: GetUserByEmail :one
SELECT * FROM "user" WHERE email=$1 AND isDeleted=FALSE;

-- name: GetUserByPhone :one
SELECT * FROM "user" WHERE phoneNumber=$1 AND isDeleted=FALSE;