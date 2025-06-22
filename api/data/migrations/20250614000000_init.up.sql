-- UP

CREATE TABLE provinces (
    id BIGINT PRIMARY KEY,  -- int64 id do code gen
    name VARCHAR(255) NOT NULL UNIQUE,
    slug VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE places (
    id BIGINT PRIMARY KEY,  -- int64 id do code gen
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT NOT NULL DEFAULT '',
    province_id BIGINT NOT NULL REFERENCES provinces(id),
    slug VARCHAR(255) NOT NULL UNIQUE,
    average_vote FLOAT NOT NULL DEFAULT 0.0,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

-- Create table to store types/categories of places
CREATE TABLE place_types (
     id BIGINT PRIMARY KEY,  -- ID generated in code (int64)
     name VARCHAR(255) NOT NULL UNIQUE,  -- Unique name of the place type (e.g., Cultural, Playground, Food court)
     created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),  -- Record creation timestamp
     updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()   -- Record update timestamp
);

-- Create junction table to link places with their types (many-to-many relationship)
CREATE TABLE place_place_types (
    place_id BIGINT NOT NULL REFERENCES places(id),  -- Reference to place
    place_type_id BIGINT NOT NULL REFERENCES place_types(id),  -- Reference to place type
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),  -- Record creation timestamp
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),  -- Record update timestamp
    PRIMARY KEY (place_id, place_type_id)  -- Composite primary key to avoid duplicate mappings
);


CREATE TABLE images (
    id BIGINT PRIMARY KEY,
    place_id BIGINT NOT NULL REFERENCES places(id),
    url TEXT NOT NULL DEFAULT '',
    type VARCHAR(255) NOT NULL DEFAULT '',  -- image, thumbnail
    caption TEXT NOT NULL DEFAULT '',
    order_num INT NOT NULL DEFAULT 0,
    is_header_image BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE comments (
    id BIGINT PRIMARY KEY,
    place_id BIGINT NOT NULL REFERENCES places(id),
    content TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE admins (
    id BIGINT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
