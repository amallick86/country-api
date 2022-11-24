-- +goose Up
-- +goose StatementBegin
CREATE TABLE "country" (
 "id" serial PRIMARY KEY,
 "name" varchar NOT NULL UNIQUE,
 "country_short_name" varchar ,
 "created_at" Date NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE country;
-- +goose StatementEnd
