-- +goose Up
-- +goose StatementBegin
-- Создаём тип ENUM для статуса запроса в друзья
DO
$$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'friend_request_status') THEN
            CREATE TYPE friend_request_status AS ENUM (
                'PENDING',
                'ACCEPTED',
                'DECLINED'
                );
        END IF;
    END
$$;

-- Создаём таблицу friend_requests
CREATE TABLE IF NOT EXISTS friend_requests
(
    id           UUID                     NOT NULL,                   -- Уникальный идентификатор заявки
    from_user_id UUID                     NOT NULL,                   -- Кто отправил заявку
    to_user_id   UUID                     NOT NULL,                   -- Кому отправлена заявка
    status       friend_request_status    NOT NULL DEFAULT 'PENDING', -- Статус заявки
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),     -- Когда создана заявка
    PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS uniq_pending_friend_request
    ON public.friend_requests (from_user_id, to_user_id)
    WHERE status = 'PENDING';

COMMENT ON TABLE friend_requests IS 'Таблица заявок в друзья между пользователями';

COMMENT ON COLUMN friend_requests.id IS 'Уникальный идентификатор заявки';
COMMENT ON COLUMN friend_requests.from_user_id IS 'Пользователь, отправивший заявку';
COMMENT ON COLUMN friend_requests.to_user_id IS 'Пользователь, которому отправлена заявка';
COMMENT ON COLUMN friend_requests.status IS 'Статус заявки: PENDING, ACCEPTED, DECLINED';
COMMENT ON COLUMN friend_requests.created_at IS 'Когда создана заявка';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS friend_requests;
DROP TYPE IF EXISTS friend_request_status;
-- +goose StatementEnd
