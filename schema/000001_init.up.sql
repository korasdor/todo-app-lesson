CREATE TABLE
    users (
        id serial NOT NULL,
        name VARCHAR(255) NOT NULL,
        username VARCHAR(255) NOT NULL UNIQUE,
        password_hash VARCHAR(255) NOT NULL
    );

CREATE TABLE
    todo_lists (
        id serial NOT NULL UNIQUE,
        title VARCHAR(255) NOT NULL,
        description VARCHAR(255)
    );

CREATE TABLE
    users_lists (
        id serial NOT NULL UNIQUE,
        user_id BIGINT UNSIGNED NOT NULL,
        list_id BIGINT UNSIGNED NOT NULL,

        FOREIGN KEY (user_id) REFERENCES users (id) on delete CASCADE,
        FOREIGN KEY (list_id) REFERENCES todo_lists (id) on delete CASCADE
    );

CREATE TABLE
    todo_items (
        id serial NOT NULL UNIQUE,
        title VARCHAR(255) NOT NULL,
        description VARCHAR(255),
        done BOOLEAN NOT NULL DEFAULT false
    );

CREATE TABLE
    lists_items (
        id serial NOT NULL UNIQUE,
        item_id BIGINT UNSIGNED NOT NULL,
        list_id BIGINT UNSIGNED NOT NULL,

        FOREIGN KEY (item_id) REFERENCES todo_items (id) on delete CASCADE,
        FOREIGN KEY (list_id) REFERENCES todo_lists (id) on delete CASCADE
    );