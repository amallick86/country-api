-- +goose Up
-- +goose StatementBegin
CREATE TABLE "state" (
"id" serial PRIMARY KEY,
"country_id" int NOT NULL ,
"state_name" varchar NOT NULL,
"created_at" Date NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE "state" ADD FOREIGN KEY ("country_id") REFERENCES "country" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE state;
-- +goose StatementEnd
