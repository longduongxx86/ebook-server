create database if not exists books character set utf8mb4 collate utf8mb4_unicode_ci;
use books;

create table categories (
    id bigint unsigned auto_increment primary key,
    name varchar(255) not null,
    slug varchar(255) not null,
    description text null,
    created_at bigint null,
    updated_at bigint null,
    constraint name unique (name),
    constraint slug unique (slug)
);

create table users (
    id bigint unsigned auto_increment primary key,
    username varchar(191) not null,
    email varchar(191) not null,
    password_hash longtext not null,
    full_name longtext null,
    phone longtext null,
    address longtext null,
    avatar_url longtext null,
    role varchar(191) default 'customer' null,
    shop_name longtext null,
    provider varchar(191) default 'local' null,
    google_id varchar(191) null,
    email_verified tinyint(1) default 0 null,
    verification_token varchar(191) null,
    created_at bigint null,
    updated_at bigint null,
    deleted_at datetime(3) null,
    constraint email unique (email),
    constraint email_2 unique (email),
    constraint username unique (username),
    constraint username_2 unique (username)
);

create table books (
    id bigint unsigned auto_increment primary key,
    title varchar(191) not null,
    author longtext null,
    description text null,
    price bigint not null,
    cost bigint default 0 null,
    stock bigint default 0 null,
    slug varchar(191) null,
    image_url longtext null,
    seller_id bigint unsigned not null,
    isbn varchar(191) null,
    average_rating double default 0 null,
    review_count bigint default 0 null,
    created_at bigint null,
    updated_at bigint null,
    deleted_at datetime(3) null,
    category_id bigint unsigned null,
    constraint isbn unique (isbn),
    constraint isbn_2 unique (isbn),
    constraint slug unique (slug),
    constraint slug_2 unique (slug),
    constraint books_ibfk_1 foreign key (seller_id) references users (id),
    constraint books_ibfk_2 foreign key (category_id) references categories (id)
);

create index category_id on books (category_id);

create index idx_books_created_at on books (created_at);

create index idx_books_deleted_at on books (deleted_at);

create index idx_books_isbn on books (isbn);

create index idx_books_seller_id on books (seller_id);

create index idx_books_slug on books (slug);

create index idx_books_title on books (title);

create index idx_books_updated_at on books (updated_at);

create table carts (
    id bigint unsigned auto_increment primary key,
    user_id bigint unsigned not null,
    created_at bigint null,
    updated_at bigint null,
    deleted_at datetime(3) null,
    constraint carts_ibfk_1 foreign key (user_id) references users (id)
);

create table cart_items (
    id bigint unsigned auto_increment primary key,
    cart_id bigint unsigned not null,
    book_id bigint unsigned not null,
    quantity bigint not null,
    created_at bigint null,
    updated_at bigint null,
    deleted_at datetime(3) null,
    constraint cart_items_ibfk_1 foreign key (cart_id) references carts (id),
    constraint cart_items_ibfk_2 foreign key (book_id) references books (id),
    constraint fk_cart_items_book foreign key (book_id) references books (id),
    constraint fk_carts_items foreign key (cart_id) references carts (id)
);

create index idx_cart_items_book_id on cart_items (book_id);

create index idx_cart_items_cart_id on cart_items (cart_id);

create index idx_cart_items_created_at on cart_items (created_at);

create index idx_cart_items_deleted_at on cart_items (deleted_at);

create index idx_cart_items_updated_at on cart_items (updated_at);

create index idx_carts_created_at on carts (created_at);

create index idx_carts_deleted_at on carts (deleted_at);

create index idx_carts_updated_at on carts (updated_at);

create index idx_carts_user_id on carts (user_id);

create table orders (
    id bigint unsigned auto_increment primary key,
    order_number varchar(191) not null,
    buyer_id bigint unsigned not null,
    total_amount bigint not null,
    status varchar(191) default 'pending' null,
    shipping_address longtext null,
    notes longtext null,
    created_at bigint null,
    updated_at bigint null,
    deleted_at datetime(3) null,
    constraint order_number unique (order_number),
    constraint order_number_2 unique (order_number),
    constraint fk_orders_buyer foreign key (buyer_id) references users (id),
    constraint orders_ibfk_1 foreign key (buyer_id) references users (id)
);

create table order_items (
    id bigint unsigned auto_increment primary key,
    order_id bigint unsigned not null,
    book_id bigint unsigned not null,
    quantity bigint not null,
    price bigint not null,
    cost bigint default 0 null,
    created_at bigint null,
    updated_at bigint null,
    deleted_at datetime(3) null,
    constraint fk_order_items_book foreign key (book_id) references books (id),
    constraint fk_orders_items foreign key (order_id) references orders (id),
    constraint order_items_ibfk_1 foreign key (order_id) references orders (id),
    constraint order_items_ibfk_2 foreign key (book_id) references books (id)
);

create index idx_order_items_book_id on order_items (book_id);

create index idx_order_items_created_at on order_items (created_at);

create index idx_order_items_deleted_at on order_items (deleted_at);

create index idx_order_items_order_id on order_items (order_id);

create index idx_order_items_updated_at on order_items (updated_at);

create index idx_orders_buyer_id on orders (buyer_id);

create index idx_orders_created_at on orders (created_at);

create index idx_orders_deleted_at on orders (deleted_at);

create index idx_orders_order_number on orders (order_number);

create index idx_orders_updated_at on orders (updated_at);

create table payments (
    id bigint unsigned auto_increment primary key,
    order_id bigint unsigned not null,
    amount bigint not null,
    method longtext null,
    status varchar(191) default 'pending' null,
    transaction_id longtext null,
    qr_code text null,
    bank_info text null,
    created_at bigint null,
    updated_at bigint null,
    deleted_at datetime(3) null,
    constraint fk_payments_order foreign key (order_id) references orders (id),
    constraint payments_ibfk_1 foreign key (order_id) references orders (id)
);

create index idx_payments_created_at on payments (created_at);

create index idx_payments_deleted_at on payments (deleted_at);

create index idx_payments_order_id on payments (order_id);

create index idx_payments_updated_at on payments (updated_at);

create table reviews (
    id bigint unsigned auto_increment primary key,
    book_id bigint unsigned not null,
    user_id bigint unsigned not null,
    rating bigint not null,
    comment text null,
    approved tinyint(1) default 1 null,
    created_at bigint null,
    updated_at bigint null,
    deleted_at datetime(3) null,
    constraint fk_reviews_user foreign key (user_id) references users (id),
    constraint reviews_ibfk_1 foreign key (book_id) references books (id),
    constraint reviews_ibfk_2 foreign key (user_id) references users (id),
    constraint chk_reviews_rating check (
        (`rating` >= 1)
        and (`rating` <= 5)
    ),
    check (
        (`rating` >= 1)
        and (`rating` <= 5)
    )
);

create index idx_reviews_book_id on reviews (book_id);

create index idx_reviews_created_at on reviews (created_at);

create index idx_reviews_deleted_at on reviews (deleted_at);

create index idx_reviews_updated_at on reviews (updated_at);

create index idx_reviews_user_id on reviews (user_id);

create index idx_users_created_at on users (created_at);

create index idx_users_deleted_at on users (deleted_at);

create index idx_users_email on users (email);

create index idx_users_google_id on users (google_id);

create index idx_users_provider on users (provider);

create index idx_users_updated_at on users (updated_at);

create index idx_users_username on users (username);

create index idx_users_verification_token on users (verification_token);

ALTER TABLE
    books DROP FOREIGN KEY books_ibfk_1;

ALTER TABLE
    books DROP FOREIGN KEY books_ibfk_2;

ALTER TABLE
    carts DROP FOREIGN KEY carts_ibfk_1;

ALTER TABLE
    cart_items DROP FOREIGN KEY cart_items_ibfk_1;

ALTER TABLE
    cart_items DROP FOREIGN KEY cart_items_ibfk_2;

ALTER TABLE
    cart_items DROP FOREIGN KEY fk_cart_items_book;

ALTER TABLE
    cart_items DROP FOREIGN KEY fk_carts_items;

ALTER TABLE
    orders DROP FOREIGN KEY fk_orders_buyer;

ALTER TABLE
    orders DROP FOREIGN KEY orders_ibfk_1;

ALTER TABLE
    order_items DROP FOREIGN KEY fk_order_items_book;

ALTER TABLE
    order_items DROP FOREIGN KEY fk_orders_items;

ALTER TABLE
    order_items DROP FOREIGN KEY order_items_ibfk_1;

ALTER TABLE
    order_items DROP FOREIGN KEY order_items_ibfk_2;

ALTER TABLE
    payments DROP FOREIGN KEY fk_payments_order;

ALTER TABLE
    payments DROP FOREIGN KEY payments_ibfk_1;

ALTER TABLE
    reviews DROP FOREIGN KEY fk_reviews_user;

ALTER TABLE
    reviews DROP FOREIGN KEY reviews_ibfk_1;

ALTER TABLE
    reviews DROP FOREIGN KEY reviews_ibfk_2;

CREATE TABLE notifications (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    type VARCHAR(50) NOT NULL, -- "order", "stock", "system"
    is_read BOOLEAN DEFAULT FALSE,
    reference_id BIGINT UNSIGNED,
    created_at BIGINT NULL,
    updated_at BIGINT NULL,
    INDEX idx_type (type),
    INDEX idx_is_read (is_read),
    INDEX idx_created_at (created_at)
);