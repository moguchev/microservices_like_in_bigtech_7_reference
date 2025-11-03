-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_profiles
(
    id         UUID                     NOT NULL,               -- Уникальный идентификатор пользователя
    email      TEXT                     NOT NULL,               -- email пользователя
    nickname   TEXT                     NOT NULL,               -- Имя пользователя
    bio        TEXT,                                            -- Био пользователя
    avatar_url TEXT,                                            -- url аватарки пользователя
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), -- Когда создан
    PRIMARY KEY (id),
    UNIQUE (email),
    UNIQUE (nickname)
);

COMMENT ON TABLE user_profiles IS 'Таблица пользователей';

COMMENT ON COLUMN user_profiles.id IS 'Уникальный идентификатор пользователя';
COMMENT ON COLUMN user_profiles.email IS 'email пользователя';
COMMENT ON COLUMN user_profiles.nickname IS 'Имя пользователя';
COMMENT ON COLUMN user_profiles.bio IS 'Био пользователя';
COMMENT ON COLUMN user_profiles.avatar_url IS 'url аватарки пользователя';
COMMENT ON COLUMN user_profiles.created_at IS 'Когда создан';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_profiles;
-- +goose StatementEnd
