-- UP

CREATE TABLE places (
                        id BIGINT PRIMARY KEY,  -- int64 id do code gen
                        name VARCHAR(255) NOT NULL UNIQUE,
                        description TEXT,
                        province VARCHAR(255),
                        slug VARCHAR(255) NOT NULL UNIQUE,
                        vote_count INT DEFAULT 0,
                        created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
                        updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE media (
                       id BIGINT PRIMARY KEY,
                       place_id BIGINT NOT NULL REFERENCES places(id),
                       url TEXT NOT NULL,
                       type VARCHAR(255) NOT NULL,  -- image, video
                       caption TEXT,
                       order_num INT,
                       is_header_image BOOLEAN DEFAULT FALSE,
                       created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
                       updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE comments (
                          id BIGINT PRIMARY KEY,
                          place_id BIGINT NOT NULL REFERENCES places(id),
                          content TEXT NOT NULL,
                          created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
                          updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE admins (
                        id BIGINT PRIMARY KEY,
                        username VARCHAR(255) NOT NULL UNIQUE,
                        password_hash TEXT NOT NULL,
                        created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
                        updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
