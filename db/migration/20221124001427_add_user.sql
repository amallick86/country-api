-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
 "id" serial PRIMARY KEY,
 "username" varchar NOT NULL UNIQUE,
 "password" varchar NOT NULL,
 "created_at" Date NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
