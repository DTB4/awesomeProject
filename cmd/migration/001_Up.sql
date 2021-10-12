create table if not exists suppliers
(
    id      int auto_increment
        primary key,
    name    varchar(45)             null,
    created datetime                null,
    updated datetime                null,
    deleted tinyint(1)  default 0   null,
    img_url varchar(500)            null,
    type    varchar(32) default ' ' null,
    opening time                    null,
    closing time                    null
);

create table if not exists products
(
    id          int auto_increment
        primary key,
    name        varchar(45)          null,
    type        varchar(45)          null,
    price       float                not null,
    created     datetime             null,
    updated     datetime             null,
    deleted     tinyint(1) default 0 null,
    id_supplier int                  null,
    img_url     varchar(500)         null,
    ingredients varchar(200)         null,
    constraint supplier_has_products
        foreign key (id_supplier) references suppliers (id)
            on update cascade on delete cascade
);

create index suplier_has_products_idx
    on products (id_supplier);

create table if not exists users
(
    id            int auto_increment,
    first_name    varchar(45)          not null,
    last_name     varchar(45)          null,
    email         varchar(45)          not null,
    password_hash varchar(90)          not null,
    created       datetime             null,
    updated       datetime             null,
    deleted       tinyint(1) default 0 null,
    constraint id_UNIQUE
        unique (id),
    constraint users_email_uindex
        unique (email)
);

alter table users
    add primary key (id);

create table if not exists orders
(
    id      int auto_increment,
    id_user int                                                                                     null,
    status  enum ('created', 'preparing', 'prepared', 'on road', 'delivered', 'canceled', 'closed') null,
    created datetime                                                                                null,
    updated datetime                                                                                null,
    deleted tinyint(1) default 0                                                                    null,
    constraint id_UNIQUE
        unique (id),
    constraint user_have_orders
        foreign key (id_user) references users (id)
            on update cascade on delete cascade
);

create index user_have_orders_idx
    on orders (id_user);

alter table orders
    add primary key (id);

create table if not exists order_product
(
    order_id   int not null,
    product_id int not null,
    quantity   int null,
    primary key (order_id, product_id),
    constraint fk_order_from_orders
        foreign key (order_id) references orders (id)
            on update cascade on delete cascade,
    constraint fk_order_has_product
        foreign key (product_id) references products (id)
            on update cascade on delete cascade
);

create index fk_orders_has_Products_Products1_idx
    on order_product (product_id);

create index fk_orders_has_Products_orders1_idx
    on order_product (order_id);

create table if not exists uids
(
    user_id int          not null,
    uid     varchar(128) null,
    constraint user_have_tokens
        foreign key (user_id) references users (id)
            on update cascade on delete cascade
);