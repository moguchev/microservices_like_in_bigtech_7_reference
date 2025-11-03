-- +goose Up
-- +goose StatementBegin

-- 1) Месячная партиция (range) OCT-2025 → сама partitioned BY LIST (aggregate_type)
CREATE TABLE IF NOT EXISTS public.outbox_events_2025_11
    PARTITION OF public.outbox_events
        FOR VALUES FROM ('2025-11-01 00:00:00+00') TO ('2025-12-01 00:00:00+00')
    PARTITION BY LIST (aggregate_type);

COMMENT ON TABLE public.outbox_events_2025_11 IS 'Партиция outbox за ноябрь 2025';

-- 2) Партиция по aggregate_type='ChatMessage' (и внутри — LIST по event_type)
CREATE TABLE IF NOT EXISTS public.outbox_events_2025_11__agg_chat_msg
    PARTITION OF public.outbox_events_2025_11
        FOR VALUES IN ('ChatMessage')
    PARTITION BY LIST (event_type);

COMMENT ON TABLE public.outbox_events_2025_11__agg_chat_msg IS 'Партиция для aggregate_type=ChatMessage (внутри — по event_type)';

-- 3) Листья по event_type внутри aggregate_type='ChatMessage'
CREATE TABLE IF NOT EXISTS public.outbox_events_2025_11__agg_chat_msg__evt_chat_msg_sent
    PARTITION OF public.outbox_events_2025_11__agg_chat_msg
        FOR VALUES IN ('ChatMessageSent');

COMMENT ON TABLE public.outbox_events_2025_11__agg_chat_msg__evt_chat_msg_sent IS 'Лист: ChatMessage / ChatMessageSent';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
