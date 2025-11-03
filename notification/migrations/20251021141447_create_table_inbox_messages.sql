-- +goose Up
-- +goose StatementBegin
DO
$$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status') THEN
            CREATE TYPE status AS ENUM (
                'received',
                'processing',
                'processed',
                'failed'
                );
        END IF;
    END
$$;

CREATE TABLE IF NOT EXISTS public.inbox_messages
(
    id           UUID        NOT NULL PRIMARY KEY,
    topic        TEXT        NOT NULL,
    partition    INT         NOT NULL,
    ofset        BIGINT      NOT NULL,
    payload      JSONB       NOT NULL,
    status       status      NOT NULL,
    attempts     INT         NOT NULL DEFAULT 0,
    last_error   TEXT,
    received_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    processed_at TIMESTAMPTZ
);

COMMENT ON TABLE public.inbox_messages IS 'Таблица входящих сообщений для паттерна Transactional Inbox';
COMMENT ON COLUMN public.inbox_messages.id IS 'Уникальный идентификатор сообщения';
COMMENT ON COLUMN public.inbox_messages.topic IS 'Kafka-топик, из которого получено сообщение';
COMMENT ON COLUMN public.inbox_messages.partition IS 'Номер раздела (partition) Kafka';
COMMENT ON COLUMN public.inbox_messages.ofset IS 'Позиция (offset) сообщения в разделе';
COMMENT ON COLUMN public.inbox_messages.payload IS 'Тело входящего сообщения (JSONB)';
COMMENT ON COLUMN public.inbox_messages.status IS 'Статус обработки сообщения: received, processing, processed, failed';
COMMENT ON COLUMN public.inbox_messages.attempts IS 'Количество попыток обработки сообщения';
COMMENT ON COLUMN public.inbox_messages.last_error IS 'Описание последней ошибки обработки';
COMMENT ON COLUMN public.inbox_messages.received_at IS 'Время получения сообщения';
COMMENT ON COLUMN public.inbox_messages.processed_at IS 'Время успешной обработки сообщения';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.inbox_messages;
-- +goose StatementEnd
