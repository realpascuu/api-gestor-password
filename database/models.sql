CREATE TABLE IF NOT EXISTS users (
    id serial NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    password varchar(256) NOT NULL,
    salt varchar(150) NOT NULL,
    CONSTRAINT pk_users PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS passwords (
    id serial NOT NULL,
    user_id int NOT NULL,
    content varchar(256) NOT NULL,
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_passwords PRIMARY KEY(id),
    CONSTRAINT fk_posts_users FOREIGN KEY(user_id) REFERENCES users(id)
);