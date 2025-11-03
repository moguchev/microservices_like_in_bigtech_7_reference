-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS chat_members
(
    chat_id UUID NOT NULL, -- Идентификатор чата
    user_id UUID NOT NULL, -- Участник чата
    PRIMARY KEY (chat_id, user_id)
);

COMMENT ON TABLE chat_members IS 'Связь пользователей с чатами (участники чатов)';

COMMENT ON COLUMN chat_members.chat_id IS 'Идентификатор чата';
COMMENT ON COLUMN chat_members.user_id IS 'Идентификатор пользователя, участвующего в чате';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS chat_members;
-- +goose StatementEnd
