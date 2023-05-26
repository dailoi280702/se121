INSERT INTO car_brands (id, name, country_of_origin, founded_year, website_url, logo_url)
VALUES
    (1, 'Toyota', 'Japan', 1937, 'https://www.toyota.com/', 'https://www.carlogos.org/car-logos/toyota-logo-2019-1350x1500.png'),
    (2, 'Ford', 'United States', 1903, 'https://www.ford.com/', 'https://www.carlogos.org/car-logos/ford-logo-2017.png'),
    (3, 'Honda', 'Japan', 1948, 'https://www.honda.com/', 'https://www.carlogos.org/car-logos/honda-logo-1700x1150.png'),
    (4, 'Chevrolet', 'United States', 1911, 'https://www.chevrolet.com/', 'https://www.carlogos.org/car-brands/chevrolet-logo.html'),
    (5, 'Nissan', 'Japan', 1933, 'https://www.nissanusa.com/', 'https://www.carlogos.org/car-logos/nissan-logo-2020-black.png'),
    (6, 'Hyundai', 'South Korea', 1967, 'https://www.hyundaiusa.com/', 'https://www.carlogos.org/logo/Hyundai-logo-silver-2560x1440.png'),
    (7, 'Kia', 'South Korea', 1944, 'https://www.kia.com/', 'https://www.carlogos.org/logo/Kia-logo-2560x1440.png'),
    (8, 'Jeep', 'United States', 1941, 'https://www.jeep.com/', 'https://www.carlogos.org/car-logos/jeep-logo-1993.png'),
    (9, 'Subaru', 'Japan', 1953, 'https://www.subaru.com/', 'https://www.carlogos.org/car-logos/subaru-logo-2003.png'),
    (10, 'Mercedes-Benz', 'Germany', 1926, 'https://www.mercedes-benz.com/', 'https://www.carlogos.org/logo/Mercedes-Benz-logo-2011-1920x1080.png'),
    (11, 'BMW', 'Germany', 1916, 'https://www.bmw.com/', 'https://www.carlogos.org/car-logos/bmw-logo-2020-blue-white.png'),
    (12, 'Audi', 'Germany', 1909, 'https://www.audiusa.com/', 'https://www.carlogos.org/car-logos/audi-logo-2016.png'),
    (13, 'Mazda', 'Japan', 1920, 'https://www.mazdausa.com/', 'https://www.carlogos.org/logo/Mazda-logo-1997-1920x1080.png'),
    (14, 'Lexus', 'Japan', 1983,'https://www.lexus.com/', 'https://www.carlogos.org/logo/Lexus-logo-1988-1920x1080.png'),
    (15, 'Tesla', 'United States', 2003, 'https://www.tesla.com/', 'https://www.carlogos.org/car-logos/tesla-logo-2200x2800.png'),
    (16, 'GMC', 'United States', 1911, 'https://www.gmc.com/', 'https://www.carlogos.org/logo/GMC-logo-2200x600.png'),
    (17, 'Dodge', 'United States', 1900, 'https://www.dodge.com/', 'https://www.carlogos.org/logo/Dodge-logo-1990-2100x2100.png'),
    (18, 'Volkswagen', 'Germany', 1937, 'https://www.vw.com/', 'https://www.carlogos.org/logo/Volkswagen-logo-2019-1500x1500.png'),
    (19, 'Ram', 'United States', 2010 ,'https://www.ramtrucks.com/', 'https://www.carlogos.org/logo/RAM-logo-2009-1920x1080.png'),
    (20, 'Porsche', 'Germany', 1948, 'https://www.porsche.com/', 'https://www.carlogos.org/car-logos/porsche-logo-2100x1100.png');


INSERT INTO fuel_types (name, description, created_at) 
VALUES 
    ('Gasoline', 'A liquid fuel used in spark-ignited internal combustion engines.', NOW()),
    ('Diesel', 'A liquid fuel used in diesel engines.', NOW()),
    ('Electricity', 'A form of energy resulting from the flow of electric charge.', NOW()),
    ('Hybrid', 'A vehicle that uses two or more distinct types of power, such as internal combustion engine and electric motor.', NOW());


INSERT INTO car_transmissions (name, description, updated_at)
VALUES
    ('Manual', 'Manual transmission', NOW()),
    ('Automatic', 'Automatic transmission', NOW()),
    ('CVT', 'Continuously Variable Transmission', NOW()),
    ('DCT', 'Dual-Clutch Transmission', NOW()),
    ('AMT', 'Automated Manual Transmission', NOW());
