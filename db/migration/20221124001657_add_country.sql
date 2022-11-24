-- +goose Up
-- +goose StatementBegin
CREATE TABLE "country" (
 "id" serial PRIMARY KEY,
 "country_name" varchar NOT NULL UNIQUE,
 "created_at" Date NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE country;
-- +goose StatementEnd
