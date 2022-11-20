CREATE TABLE IF NOT EXISTS work_schedule (
	id serial PRIMARY KEY,
	weekday int NOT NULL,
	beggin_time time,
	end_time time,
    is_day_off boolean NOT NULL
);

CREATE TABLE IF NOT EXISTS contacts (
	id serial PRIMARY KEY,
	city varchar(255) NOT NULL,
	street varchar(255) NOT NULL,
    work_schedule_id bigint REFERENCES work_schedule (id) NOT NULL
);

CREATE TABLE IF NOT EXISTS categories (
	id serial PRIMARY KEY,
	name varchar(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS places (
	id serial PRIMARY KEY,
	name varchar(255) NOT NULL,
	describe text,
    url varchar(255) NOT NULL,
	photo varchar(255),
	class varchar(50),
    contact_id bigint REFERENCES contacts (id) NOT NULL
);

CREATE TABLE IF NOT EXISTS items (
	id serial PRIMARY KEY,
	name varchar(255) NOT NULL,
	describe text,
	price numeric(15,2) NOT NULL,
	weight float,
	photo varchar(255),
    category_id bigint REFERENCES categories (id) NOT NULL,
	place_id bigint REFERENCES places (id) NOT NULL
);
