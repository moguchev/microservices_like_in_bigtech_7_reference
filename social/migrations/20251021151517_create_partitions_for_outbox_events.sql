-- +goose Up
-- +goose StatementBegin

-- 1) Месячная партиция (range) OCT-2025 → сама partitioned BY LIST (aggregate_type)
CREATE TABLE IF NOT EXISTS public.outbox_events_2025_11
    PARTITION OF public.outbox_events
        FOR VALUES FROM ('2025-11-01 00:00:00+00') TO ('2025-12-01 00:00:00+00')
    PARTITION BY LIST (aggregate_type);

COMMENT ON TABLE public.outbox_events_2025_11 IS 'Партиция outbox за ноябрь 2025';

-- 2) Партиция по aggregate_type='FriendRequest' (и внутри — LIST по event_type)
CREATE TABLE IF NOT EXISTS public.outbox_events_2025_11__agg_friend_req
    PARTITION OF public.outbox_events_2025_11
        FOR VALUES IN ('FriendRequest')
    PARTITION BY LIST (event_type);

COMMENT ON TABLE public.outbox_events_2025_11__agg_friend_req IS 'Партиция для aggregate_type=FriendRequest (внутри — по event_type)';

-- 3) Листья по event_type внутри aggregate_type='FriendRequest'
CREATE TABLE IF NOT EXISTS public.outbox_events_2025_11__agg_friend_req__evt_friend_req_created
    PARTITION OF public.outbox_events_2025_11__agg_friend_req
        FOR VALUES IN ('FriendRequestCreated');

COMMENT ON TABLE public.outbox_events_2025_11__agg_friend_req__evt_friend_req_created IS 'Лист: FriendRequest / FriendRequestCreated';

CREATE TABLE IF NOT EXISTS public.outbox_events_2025_11__agg_friend_req__evt_friend_req_updated
    PARTITION OF public.outbox_events_2025_11__agg_friend_req
        FOR VALUES IN ('FriendRequestUpdated');

COMMENT ON TABLE public.outbox_events_2025_11__agg_friend_req__evt_friend_req_updated IS 'Лист: FriendRequest / FriendRequestUpdated';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
