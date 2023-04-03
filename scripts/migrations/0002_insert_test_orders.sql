-- +goose Up
-- +goose StatementBegin
insert into kitchen_order (id, shop_order_id, status)
VALUES (1, 1, 'new'),
       (2, 2, 'new'),
       (3, 3, 'new'),
       (4, 4, 'new'),
       (5, 5, 'new'),
       (6, 6, 'new');

insert into item (kitchen_order_id, item_name, item_comment)
VALUES (1, 'pizza1', 'comment1'),
       (2, 'pizza2', 'comment2'),
       (2, 'pizza2', 'comment2'),
       (3, 'pizza3', 'comment3'),
       (3, 'pizza3', 'comment3'),
       (4, 'pizza1', 'comment1'),
       (4, 'pizza2', 'comment2'),
       (4, 'pizza3', 'comment3'),
       (5, 'pizza1', 'comment1'),
       (5, 'pizza2', 'comment2'),
       (5, 'pizza3', 'comment3'),
       (6, 'pizza3', 'comment3'),
       (6, 'pizza1', 'comment1'),
       (6, 'pizza2', 'comment2'),
       (6, 'pizza3', 'comment3');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from item where true;
delete from kitchen_order where true;
-- +goose StatementEnd
