-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS friends
(
    user_id        UUID                     NOT NULL,               -- Пользователь
    friend_user_id UUID                     NOT NULL,               -- Его друг
    created_at     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), -- Когда они стали друзьями
    PRIMARY KEY (user_id, friend_user_id)
);

-- Индекс для быстрого получения друзей по пользователю и времени
CREATE INDEX IF NOT EXISTS idx_friends_user_id_created_at ON friends (user_id, created_at);

COMMENT ON TABLE friends IS 'Таблица дружбы между пользователями';

COMMENT ON COLUMN friends.user_id IS 'Пользователь';
COMMENT ON COLUMN friends.friend_user_id IS 'Друг пользователя';
COMMENT ON COLUMN friends.created_at IS 'Когда они стали друзьями';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS friends;
-- +goose StatementEnd
