-- +goose Up
-- +goose StatementBegin
CREATE TABLE urls (
    id          SERIAL PRIMARY KEY,
    longhand    TEXT NOT NULL,
    shorthand   TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE urls;
-- +goose StatementEnd
