-- +goose Up
-- +goose StatementBegin
CREATE TABLE "sessions" (
 "id" uuid PRIMARY KEY,
 "user_id" int NOT NULL,
 "refresh_token" varchar NOT NULL,
 "user_agent" varchar NOT NULL,
 "client_ip" varchar NOT NULL,
 "is_blocked" boolean NOT NULL DEFAULT false,
 "expires_at" timestamp NOT NULL,
 "created_at" Date NOT NULL DEFAULT now()
);
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sessions;
-- +goose StatementEnd
