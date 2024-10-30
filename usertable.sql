-- Create users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE Items (
    item_id INT PRIMARY KEY,
    name VARCHAR(255),
    category VARCHAR(255),
    description VARCHAR(255),
    quantity INT,
    unit_price DECIMAL(10, 2),
    total_price DECIMAL(10, 2)
);

CREATE TABLE Warehouses (
    warehouse_id INT PRIMARY KEY,
    location VARCHAR(255),
    current_capacity INT,
    total_capacity INT,
    admin_id INT,
    FOREIGN KEY (admin_id) REFERENCES Users(id)
);

-- Insert dummy data
INSERT INTO users (name, email, password) 
VALUES 
('Alice Johnson', 'alice.johnson@example.com', 'hashed_password_1'),
('Bob Smith', 'bob.smith@example.com', 'hashed_password_2'),
('Charlie Brown', 'charlie.brown@example.com', 'hashed_password_3'),
('Diana Prince', 'diana.prince@example.com', 'hashed_password_4'),
('Eve Adams', 'eve.adams@example.com', 'hashed_password_5');

-- Insert dummy data into Items
INSERT INTO Items (item_id, name, category, description, quantity, unit_price, total_price) VALUES
(1, 'Item A', 'Category 1', 'Description for Item A', 10, 5.50, 55.00),
(2, 'Item B', 'Category 2', 'Description for Item B', 20, 3.25, 65.00),
(3, 'Item C', 'Category 3', 'Description for Item C', 15, 7.10, 106.50),
(4, 'Item D', 'Category 1', 'Description for Item D', 5, 12.00, 60.00),
(5, 'Item E', 'Category 2', 'Description for Item E', 30, 1.75, 52.50),
(6, 'Item F', 'Category 3', 'Description for Item F', 25, 2.50, 62.50),
(7, 'Item G', 'Category 1', 'Description for Item G', 8, 4.00, 32.00),
(8, 'Item H', 'Category 2', 'Description for Item H', 50, 0.99, 49.50),
(9, 'Item I', 'Category 3', 'Description for Item I', 40, 6.30, 252.00),
(10, 'Item J', 'Category 1', 'Description for Item J', 12, 9.00, 108.00),
(11, 'Item K', 'Category 2', 'Description for Item K', 14, 15.75, 220.50),
(12, 'Item L', 'Category 3', 'Description for Item L', 35, 8.50, 297.50),
(13, 'Item M', 'Category 1', 'Description for Item M', 22, 11.00, 242.00),
(14, 'Item N', 'Category 2', 'Description for Item N', 19, 14.00, 266.00),
(15, 'Item O', 'Category 3', 'Description for Item O', 17, 20.00, 340.00),
(16, 'Item P', 'Category 1', 'Description for Item P', 10, 5.25, 52.50),
(17, 'Item Q', 'Category 2', 'Description for Item Q', 9, 7.00, 63.00),
(18, 'Item R', 'Category 3', 'Description for Item R', 15, 3.00, 45.00),
(19, 'Item S', 'Category 1', 'Description for Item S', 12, 4.50, 54.00),
(20, 'Item T', 'Category 2', 'Description for Item T', 20, 6.25, 125.00),
(21, 'Item U', 'Category 3', 'Description for Item U', 25, 2.10, 52.50),
(22, 'Item V', 'Category 1', 'Description for Item V', 18, 9.80, 176.40),
(23, 'Item W', 'Category 2', 'Description for Item W', 7, 13.00, 91.00),
(24, 'Item X', 'Category 3', 'Description for Item X', 13, 17.50, 227.50),
(25, 'Item Y', 'Category 1', 'Description for Item Y', 11, 25.00, 275.00),
(26, 'Item Z', 'Category 2', 'Description for Item Z', 6, 5.00, 30.00),
(27, 'Item AA', 'Category 3', 'Description for Item AA', 14, 18.75, 262.50),
(28, 'Item AB', 'Category 1', 'Description for Item AB', 9, 15.00, 135.00),
(29, 'Item AC', 'Category 2', 'Description for Item AC', 4, 8.20, 32.80),
(30, 'Item AD', 'Category 3', 'Description for Item AD', 3, 12.50, 37.50),
(31, 'Item AE', 'Category 1', 'Description for Item AE', 28, 10.00, 280.00),
(32, 'Item AF', 'Category 2', 'Description for Item AF', 15, 22.00, 330.00),
(33, 'Item AG', 'Category 3', 'Description for Item AG', 5, 9.50, 47.50),
(34, 'Item AH', 'Category 1', 'Description for Item AH', 33, 13.00, 429.00),
(35, 'Item AI', 'Category 2', 'Description for Item AI', 8, 14.50, 116.00),
(36, 'Item AJ', 'Category 3', 'Description for Item AJ', 20, 6.00, 120.00),
(37, 'Item AK', 'Category 1', 'Description for Item AK', 11, 11.50, 126.50),
(38, 'Item AL', 'Category 2', 'Description for Item AL', 19, 7.50, 142.50),
(39, 'Item AM', 'Category 3', 'Description for Item AM', 7, 18.00, 126.00),
(40, 'Item AN', 'Category 1', 'Description for Item AN', 14, 24.00, 336.00),
(41, 'Item AO', 'Category 2', 'Description for Item AO', 21, 4.50, 94.50),
(42, 'Item AP', 'Category 3', 'Description for Item AP', 9, 2.00, 18.00),
(43, 'Item AQ', 'Category 1', 'Description for Item AQ', 13, 3.30, 42.90),
(44, 'Item AR', 'Category 2', 'Description for Item AR', 17, 8.75, 148.75),
(45, 'Item AS', 'Category 3', 'Description for Item AS', 23, 11.20, 257.60),
(46, 'Item AT', 'Category 1', 'Description for Item AT', 30, 5.00, 150.00),
(47, 'Item AU', 'Category 2', 'Description for Item AU', 22, 15.00, 330.00),
(48, 'Item AV', 'Category 3', 'Description for Item AV', 10, 7.00, 70.00),
(49, 'Item AW', 'Category 1', 'Description for Item AW', 5, 9.99, 49.95),
(50, 'Item AX', 'Category 2', 'Description for Item AX', 6, 20.00, 120.00);

-- Insert dummy data into Warehouses
INSERT INTO Warehouses (warehouse_id, location, current_capacity, total_capacity, admin_id) VALUES
(1, 'Warehouse 1', 40, 100, 1),
(2, 'Warehouse 2', 20, 150, 2),
(3, 'Warehouse 3', 60, 200, 3),
(4, 'Warehouse 4', 80, 250, 4),
(5, 'Warehouse 5', 100, 300, 5),
(6, 'Warehouse 6', 90, 350, 6),
(7, 'Warehouse 7', 70, 400, 7),
(8, 'Warehouse 8', 30, 250, 8),
(9, 'Warehouse 9', 150, 450, 9),
(10, 'Warehouse 10', 200, 500, 10);

--psql -U postgres -d scalesync -f ./usertable.sql