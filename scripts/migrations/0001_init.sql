-- +goose Up
-- +goose StatementBegin
CREATE TABLE kitchen_order
(
    id            BIGSERIAL PRIMARY KEY NOT NULL,
    shop_order_id BIGINT UNIQUE         NOT NULL,
    status        TEXT                  NOT NULL,
    CONSTRAINT shop_order_id UNIQUE (shop_order_id)
);

CREATE TABLE item
(
    kitchen_order_id BIGINT REFERENCES kitchen_order (id),
    item_name        TEXT NOT NULL,
    item_comment     TEXT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE item;
DROP TABLE kitchen_order;
-- +goose StatementEnd