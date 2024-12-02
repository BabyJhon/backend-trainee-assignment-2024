CREATE TABLE houses (
    id SERIAL PRIMARY KEY,
    address varchar(255) NOT NULL,
    year int NOT NULL CHECK (year >= 1),
    developer varchar(255),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

CREATE TYPE moderation_status AS  ENUM('created', 'approved', 'declined', 'on moderation');

CREATE TABLE flats (
    id int NOT NULL CHECK (id >= 1) ,
    house_id int NOT NULL CHECK (house_id >= 1),
    UNIQUE (id, house_id),
    FOREIGN KEY (house_id) REFERENCES houses(id) ON DELETE CASCADE,
    price int NOT NULL CHECK (price >= 1),
    rooms int NOT NULL CHECK (rooms >= 1),
    status moderation_status NOT NULL
);

CREATE TYPE user_status as ENUM('client', 'moderator');

CREATE TABLE users (
    id uuid DEFAULT gen_random_uuid(), 
    email TEXT NOT NULL UNIQUE,
    CHECK (email ~ '[a-zA-Z0-9_\-]+@([a-zA-Z0-9_\-]+\.)+(com|org|edu|nz|au)'),
    password_hash varchar NOT NULL,
    user_type user_status NOT NULL
);
