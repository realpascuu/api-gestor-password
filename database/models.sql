CREATE TABLE IF NOT EXISTS users (
    id serial NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    password varchar(256) NOT NULL,
    salt varchar(150) NOT NULL,
    CONSTRAINT pk_users PRIMARY KEY(id)
);
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS passwords (
    id uuid DEFAULT uuid_generate_v4(),
    user_id int NOT NULL,
    content varchar(256) NOT NULL,
    updated_at timestamp,
    CONSTRAINT pk_passwords PRIMARY KEY(id),
    CONSTRAINT fk_posts_users FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);