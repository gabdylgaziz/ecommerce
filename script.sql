create table users (
                       id int primary key,
                       name varchar(150),
                       surname varchar(150),
                       email varchar(150),
                       username varchar(150),
                       password varchar
);

create table roles (
                       id int primary key,
                       title varchar(25)
);

create table user_roles (
                            user_id int,
                            role_id int,
                            PRIMARY KEY (user_id, role_id),
                            FOREIGN KEY (user_id) references users(id),
                            FOREIGN KEY (role_id) references roles(id)
);

create table items (
                       id int primary key,
                       name varchar,
                       description text,
                       price real
);

create table orders (
                        id int primary key,
                        user_id int,
                        item_id int,
                        FOREIGN KEY (user_id) references users(id),
                        FOREIGN KEY (item_id) references items(id)
);

create table ratings (
                         user_id int,
                         item_id int,
                         rating int,
                         PRIMARY KEY (user_id, item_id),
                         FOREIGN KEY (item_id) references items(id),
                         FOREIGN KEY (user_id) references users(id)
);

create table comments (
                          user_id int,
                          item_id int,
                          comment text,
                          comment_date date,
                          FOREIGN KEY (user_id) references users(id),
                          FOREIGN KEY (item_id) references items(id)
);

create table cards (
                       id int,
                       user_id int,
                       cardholder varchar(150),
                       PAN varchar(16),
                       exp_date date,
                       PRIMARY KEY (user_id, PAN),
                       FOREIGN KEY (user_id) references users(id)
);