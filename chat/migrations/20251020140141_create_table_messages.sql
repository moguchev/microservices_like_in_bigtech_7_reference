-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS messages
(
    id         UUID                     NOT NULL,               -- Уникальный идентификатор сообщения
    chat_id    UUID                     NOT NULL,               -- Чат, к которому относится сообщение
    sender_id  UUID                     NOT NULL,               -- Отправитель сообщения
    text       TEXT                     NOT NULL,               -- Текст сообщения
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), -- Когда сообщение отправлено
    PRIMARY KEY (id)
);

-- Индекс для быстрого получения сообщений по чату и времени
CREATE INDEX IF NOT EXISTS idx_messages_chat_id_created_at ON messages (chat_id, created_at);

COMMENT ON TABLE messages IS 'Сообщения, отправленные пользователями в чатах';

COMMENT ON COLUMN messages.id IS 'Уникальный идентификатор сообщения';
COMMENT ON COLUMN messages.chat_id IS 'Чат, к которому относится сообщение';
COMMENT ON COLUMN messages.sender_id IS 'Пользователь, отправивший сообщение';
COMMENT ON COLUMN messages.text IS 'Текст сообщения';
COMMENT ON COLUMN messages.created_at IS 'Когда сообщение было отправлено';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_messages_chat_id_created_at;
DROP TABLE IF EXISTS messages;
-- +goose StatementEnd
