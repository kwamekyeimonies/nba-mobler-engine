-- +goose Up
CREATE TABLE "user"
(
    id          UUID PRIMARY KEY,
    firstName   TEXT        NOT NULL,
    lastName    TEXT        NOT NULL,
    password    TEXT        NOT NULL DEFAULT '',
    createdAt   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updateAt    TIMESTAMPTZ,
    deletedAt   TIMESTAMPTZ,
    activatedAt TIMESTAMPTZ,
    email       TEXT        NOT NULL UNIQUE,
    phoneNumber TEXT        NOT NULL UNIQUE,
    dateOfBirth TEXT        NOT NULL DEFAULT '',
    isDeleted   BOOLEAN              DEFAULT FALSE
);

-- +goose Down
DROP TABLE "user";