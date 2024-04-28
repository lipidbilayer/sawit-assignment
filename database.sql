-- This is the SQL script that will be used to initialize the database schema.
-- We will evaluate you based on how well you design your database.
-- 1. How you design the tables.
-- 2. How you choose the data types and keys.
-- 3. How you name the fields.
-- In this assignment we will use PostgreSQL as the database.

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE estates (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	width INT NOT NULL,
	length INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);


CREATE TYPE estate_object_type AS ENUM('tree','water');

create table estate_objects(
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    estate_id uuid NOT NULL,
	x_location INT NOT NULL,
	y_location INT NOT NULL,
	height INT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP,
    CONSTRAINT tree_duplicate UNIQUE (estate_id, x_location, y_location)
)