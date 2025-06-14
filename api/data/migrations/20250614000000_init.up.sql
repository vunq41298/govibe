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

CREATE TABLE media (
    id BIGINT PRIMARY KEY,
    place_id BIGINT NOT NULL REFERENCES places(id),
    url TEXT NOT NULL DEFAULT '',
    type VARCHAR(255) NOT NULL DEFAULT '',  -- image, video
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
