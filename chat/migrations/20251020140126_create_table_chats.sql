-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS chats
(
    id         UUID                     NOT NULL,               -- Уникальный идентификатор чата
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), -- Когда создан чат
    PRIMARY KEY (id)
);

COMMENT ON TABLE chats IS 'Таблица чатов между пользователями';

COMMENT ON COLUMN chats.id IS 'Уникальный идентификатор чата';
COMMENT ON COLUMN chats.created_at IS 'Когда создан чат';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS chats;
-- +goose StatementEnd
