-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS urls (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    longhand    TEXT NOT NULL,
    shorthand   TEXT UNIQUE NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS shorthand_idx ON urls(shorthand);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE urls;
-- +goose StatementEnd
