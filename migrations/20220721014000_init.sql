-- +goose Up
-- +goose StatementBegin
CREATE TABLE message_from_chepush (
                      id int NOT NULL,
                      title text,
                      message text,
                      PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE message_from_chepush;
-- +goose StatementEnd
