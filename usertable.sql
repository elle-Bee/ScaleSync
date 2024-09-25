-- Create users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL
);

-- Insert dummy data
INSERT INTO users (name, email, password) 
VALUES 
('Alice Johnson', 'alice.johnson@example.com', 'hashed_password_1'),
('Bob Smith', 'bob.smith@example.com', 'hashed_password_2'),
('Charlie Brown', 'charlie.brown@example.com', 'hashed_password_3'),
('Diana Prince', 'diana.prince@example.com', 'hashed_password_4'),
('Eve Adams', 'eve.adams@example.com', 'hashed_password_5');


--psql -U postgres -d scalesync -f ./usertable.sql