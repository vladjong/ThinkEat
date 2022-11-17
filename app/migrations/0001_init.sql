-- +goose Up
-- +goose StatementBegin
CREATE TABLE places (
	id serial PRIMARY KEY,
	name varchar(255) NOT NULL,
	describe text,
    url varchar(255) NOT NULL,
	photo varchar(255),
	class varchar(255)
);
--
CREATE TABLE items (
	id serial PRIMARY KEY,
	name varchar(255) NOT NULL,
	describe text,
	price numeric(15,2) NOT NULL,
	weight float,
	photo varchar(255),
	place_id bigint REFERENCES places (id) NOT NULL
);
--
CREATE TABLE categories (
	id serial PRIMARY KEY,
	name varchar(255) NOT NULL,
	item_id bigint REFERENCES items (id) NOT NULL
);
--
CREATE TABLE contacts (
	id serial PRIMARY KEY,
	city varchar(255) NOT NULL,
	street varchar(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE places CASCADE;
--
DROP TABLE items CASCADE;
--
DROP TABLE categories CASCADE;
--
DROP TABLE contacts CASCADE;
-- +goose StatementEnd
