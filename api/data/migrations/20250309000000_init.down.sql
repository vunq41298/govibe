-- Down

-- Delete seed data from categories
DELETE FROM categories WHERE name IN ('Food', 'Technology', 'Movies', 'Books', 'Travel');

-- Drop reviews table first since it has foreign key dependencies
DROP TABLE IF EXISTS reviews;

-- Drop users table
DROP TABLE IF EXISTS users;

-- Drop categories table
DROP TABLE IF EXISTS categories;
