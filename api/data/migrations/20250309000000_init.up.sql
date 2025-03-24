-- Up

-- Create categories table
CREATE TABLE categories (
                            id BIGINT CONSTRAINT categories_pkey PRIMARY KEY
                                CONSTRAINT categories_id_check CHECK ((id)::NUMERIC > (0)::NUMERIC),
                            uuid UUID NOT NULL DEFAULT gen_random_uuid(),
                            name VARCHAR(255) NOT NULL,
                            description TEXT,
                            created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
                            updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
                            deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create users table
CREATE TABLE users (
                       id BIGINT CONSTRAINT users_pkey PRIMARY KEY
                           CONSTRAINT users_id_check CHECK ((id)::NUMERIC > (0)::NUMERIC),
                       uuid UUID NOT NULL DEFAULT gen_random_uuid(),
                       username VARCHAR(255) UNIQUE NOT NULL,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password_hash TEXT NOT NULL,
                       created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
                       updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
                       deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create reviews table
CREATE TABLE reviews (
                         id BIGINT CONSTRAINT reviews_pkey PRIMARY KEY
                             CONSTRAINT reviews_id_check CHECK ((id)::NUMERIC > (0)::NUMERIC),
                         uuid UUID NOT NULL DEFAULT gen_random_uuid(),
                         user_id BIGINT NOT NULL CONSTRAINT reviews_user_id_fkey REFERENCES users(id) ON DELETE CASCADE,
                         category_id BIGINT NOT NULL CONSTRAINT reviews_category_id_fkey REFERENCES categories(id) ON DELETE CASCADE,
                         rating INT CHECK (rating >= 1 AND rating <= 5),
                         comment TEXT,
                         created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
                         updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
                         deleted_at TIMESTAMP WITH TIME ZONE
);
