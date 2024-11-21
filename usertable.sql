-- users
DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL
);

INSERT INTO users (name, email, password) 
VALUES 
('aa', 'aa', 'aa'),
('Bob Smith', 'bob.smith@example.com', 'hashed_password_2'),
('Charlie Brown', 'charlie.brown@example.com', 'hashed_password_3'),
('Diana Prince', 'diana.prince@example.com', 'hashed_password_4'),
('Eve Adams', 'eve.adams@example.com', 'hashed_password_5');





-- items
DROP TABLE IF EXISTS items CASCADE;

CREATE TABLE items (
    item_id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    category VARCHAR(50),
    description TEXT,
    quantity INT,
    unit_price NUMERIC(10, 2),
    total_price NUMERIC(10, 2)
);

INSERT INTO items (item_id, name, category, description, quantity, unit_price, total_price) VALUES
(1, 'Steel Beam', 'Construction', 'High-strength steel beam for structural support.', 10, 150.00, 1500.00),
(2, 'LED Light Bulb', 'Electronics', 'Energy-efficient LED light bulb, 60W equivalent.', 20, 3.25, 65.00),
(3, 'Plywood Sheet', 'Construction', 'Standard 4x8 ft plywood sheet for framing.', 15, 25.00, 375.00),
(4, 'Copper Wire', 'Electrical', 'High-quality copper wire for electrical wiring.', 5, 100.00, 500.00),
(5, 'Laptop - Model X', 'Electronics', '15-inch laptop with Intel i5 processor, 8GB RAM.', 30, 550.00, 16500.00),
(6, 'PVC Pipe', 'Plumbing', '10 ft PVC pipe for plumbing installations.', 25, 12.50, 312.50),
(7, 'Office Chair', 'Furniture', 'Ergonomic office chair with lumbar support.', 8, 120.00, 960.00),
(8, 'Air Conditioner', 'Appliances', '1.5 ton split air conditioner with inverter technology.', 50, 299.99, 14999.50),
(9, 'Bluetooth Speaker', 'Electronics', 'Portable Bluetooth speaker with 10-hour battery life.', 40, 50.00, 2000.00),
(10, 'Steel Nails', 'Hardware', 'Box of 500 steel nails for construction use.', 12, 9.00, 108.00),
(11, 'Desk Lamp', 'Lighting', 'Adjustable desk lamp with LED light.', 14, 15.75, 220.50),
(12, 'Electric Drill', 'Tools', 'Cordless electric drill with battery and charger.', 35, 85.00, 2975.00),
(13, 'Concrete Mix', 'Construction', 'Ready-mix concrete for foundations and walls.', 22, 11.00, 242.00),
(14, 'Ceramic Tiles', 'Flooring', '12x12 inch ceramic tiles for flooring.', 19, 14.00, 266.00),
(15, 'Ceiling Fan', 'Appliances', '52-inch ceiling fan with remote control.', 17, 120.00, 2040.00),
(16, 'Gardening Soil', 'Gardening', 'Organic gardening soil, 40 lb bag.', 10, 5.25, 52.50),
(17, 'Shovel', 'Gardening', 'Heavy-duty steel shovel for gardening.', 9, 15.00, 135.00),
(18, 'USB Cable', 'Accessories', '6 ft USB-C cable for charging and data transfer.', 15, 3.00, 45.00),
(19, 'Hammer', 'Tools', '16 oz steel hammer with shock-absorbing handle.', 12, 7.50, 90.00),
(20, 'Portable Heater', 'Appliances', '1500W portable heater with thermostat.', 20, 29.99, 599.80),
(21, 'Fire Extinguisher', 'Safety', '5 lb ABC dry chemical fire extinguisher.', 25, 35.00, 875.00),
(22, 'LED Floodlight', 'Lighting', '100W LED floodlight for outdoor use.', 18, 19.99, 359.82),
(23, 'Wall Paint - White', 'Paint', '1 gallon of high-quality white wall paint.', 7, 30.00, 210.00),
(24, 'Cordless Screwdriver', 'Tools', 'Compact cordless screwdriver with bits set.', 13, 40.00, 520.00),
(25, 'Solar Panel', 'Energy', '100W solar panel for renewable energy systems.', 11, 99.99, 1099.89),
(26, 'Security Camera', 'Electronics', '1080p outdoor security camera with night vision.', 6, 75.00, 450.00),
(27, 'Water Bottle', 'Accessories', '32 oz stainless steel water bottle.', 14, 18.75, 262.50),
(28, 'Router', 'Networking', 'Dual-band wireless router with parental control.', 9, 50.00, 450.00),
(29, 'Hand Sanitizer', 'Hygiene', '8 oz bottle of hand sanitizer.', 4, 4.50, 18.00),
(30, 'Printer Ink Cartridge', 'Office Supplies', 'Black ink cartridge for office printers.', 3, 25.00, 75.00),
(31, 'Cement Bag', 'Construction', '50 lb bag of cement for concrete work.', 28, 10.00, 280.00),
(32, 'Network Cable', 'Networking', 'Cat6 network cable, 50 ft.', 15, 22.00, 330.00),
(33, 'Fire Alarm', 'Safety', 'Smoke and CO detector with 10-year battery.', 5, 30.00, 150.00),
(34, 'Dining Table', 'Furniture', 'Wooden dining table for 6 people.', 33, 200.00, 6600.00),
(35, 'Microwave Oven', 'Appliances', '1200W microwave oven with defrost function.', 8, 80.00, 640.00),
(36, 'Vacuum Cleaner', 'Cleaning', 'Bagless vacuum cleaner for home use.', 20, 120.00, 2400.00),
(37, 'Screwdriver Set', 'Tools', 'Precision screwdriver set with 30 pieces.', 11, 20.00, 220.00),
(38, 'Dishwasher', 'Appliances', 'Stainless steel dishwasher with energy efficiency.', 19, 300.00, 5700.00),
(39, 'Garden Hose', 'Gardening', '100 ft expandable garden hose.', 7, 25.00, 175.00),
(40, 'Power Bank', 'Electronics', '10000mAh power bank for mobile devices.', 14, 12.00, 168.00),
(41, 'Pressure Washer', 'Tools', '2000 PSI pressure washer for outdoor cleaning.', 21, 150.00, 3150.00),
(42, 'Extension Cord', 'Electrical', '25 ft heavy-duty extension cord.', 9, 15.00, 135.00),
(43, 'Lawn Mower', 'Gardening', 'Electric lawn mower with adjustable height.', 13, 200.00, 2600.00),
(44, 'Paint Roller', 'Painting', '9-inch paint roller with replacement covers.', 17, 8.75, 148.75),
(45, 'Storage Box', 'Furniture', 'Plastic storage box with lid, 20 gallons.', 23, 11.20, 257.60),
(46, 'Power Drill', 'Tools', 'Electric power drill with adjustable speed.', 30, 45.00, 1350.00),
(47, 'Surge Protector', 'Electrical', '6-outlet surge protector with 3 ft cord.', 22, 10.00, 220.00),
(48, 'Electric Kettle', 'Appliances', '1.7L electric kettle with auto shut-off.', 10, 25.00, 250.00),
(49, 'Wi-Fi Range Extender', 'Networking', 'Wi-Fi range extender with dual-band support.', 5, 35.00, 175.00),
(50, 'Hand Saw', 'Tools', '20-inch hand saw for woodworking.', 6, 15.00, 90.00);





-- warehouses
DROP TABLE IF EXISTS warehouses;

CREATE TABLE warehouses (
    warehouse_id SERIAL PRIMARY KEY,
    location VARCHAR(100),
    current_capacity INT,
    total_capacity INT,
    admin_id INT
);

INSERT INTO warehouses (warehouse_id, location, current_capacity, total_capacity, admin_id) VALUES
(1, 'Okhla Industrial Area, New Delhi', 40, 100, 1),
(2, 'Udyog Vihar, Gurgaon', 20, 150, 2),
(3, 'Greater Noida Industrial Area, Greater Noida', 60, 200, 3),
(4, 'Sohna Road, Gurgaon', 80, 250, 4),
(5, 'Sector 62, Noida', 100, 300, 5),
(6, 'Kirti Nagar, New Delhi', 90, 350, 1),
(7, 'Manesar Industrial Area, Gurgaon', 70, 400, 2),
(8, 'Narela Industrial Area, New Delhi', 30, 250, 3),
(9, 'Sector 59, Faridabad', 150, 450, 4),
(10, 'Bhiwadi Industrial Area, Bhiwadi', 200, 500, 5);





-- warehouseItems
DROP TABLE IF EXISTS warehouseItems;

CREATE TABLE warehouseItems (
    warehouse_id INT NOT NULL,
    item_id INT NOT NULL,
    quantity INT,  -- quantity of this item in this specific warehouse
    PRIMARY KEY (warehouse_id, item_id),
    FOREIGN KEY (warehouse_id) REFERENCES Warehouses(warehouse_id) ON DELETE CASCADE,
    FOREIGN KEY (item_id) REFERENCES Items(item_id) ON DELETE CASCADE
);

INSERT INTO warehouseItems (warehouse_id, item_id, quantity) VALUES
(1, 1, 5),
(1, 2, 10),
(1, 3, 7),
(2, 4, 20),
(2, 5, 15),
(2, 6, 8),
(3, 7, 12),
(3, 8, 30),
(3, 9, 18),
(4, 10, 25),
(4, 11, 14),
(4, 12, 9),
(5, 13, 20),
(5, 14, 22),
(5, 15, 10),
(6, 16, 8),
(6, 17, 15),
(6, 18, 5),
(7, 19, 12),
(7, 20, 16),
(8, 21, 30),
(8, 22, 18),
(9, 23, 22),
(9, 24, 25),
(10, 25, 28),
(10, 26, 10);

--psql -U postgres -d scalesync -f ./usertable.sql