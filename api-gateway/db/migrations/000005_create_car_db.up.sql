CREATE TABLE car_transmissions(
    id SERIAL NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    car_transmissions ADD PRIMARY KEY(id);

CREATE TABLE car_series(
    id SERIAL NOT NULL,
    brand_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NULL
);
ALTER TABLE
    car_series ADD PRIMARY KEY(id);

CREATE TABLE fuel_types(
    id SERIAL NOT NULL,
    name TEXT NOT NULL,
    description TEXT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NULL
);
ALTER TABLE
    fuel_types ADD PRIMARY KEY(id);

CREATE TABLE car_models(
    id SERIAL NOT NULL,
    series_id INTEGER NULL,
    brand_id INTEGER NULL,
    name TEXT NOT NULL,
    year INTEGER NULL,
    horsepower INTEGER NULL,
    torque INTEGER NULL,
    transmission INT NULL,
    fuel_type INTEGER NULL,
    create_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NULL,
    review TEXT NULL,
    image_url TEXT NULL
);
ALTER TABLE
    car_models ADD PRIMARY KEY(id);

CREATE TABLE car_brands(
    id SERIAL NOT NULL,
    name TEXT NOT NULL,
    country_of_origin TEXT NULL,
    founded_year INTEGER NULL,
    website_url TEXT NULL,
    logo_url TEXT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NULL
);
ALTER TABLE
    car_brands ADD PRIMARY KEY(id);

ALTER TABLE
    car_models ADD CONSTRAINT car_models_brand_id_foreign FOREIGN KEY(brand_id) REFERENCES car_brands(id);
ALTER TABLE
    car_models ADD CONSTRAINT car_models_series_id_foreign FOREIGN KEY(series_id) REFERENCES car_series(id);
ALTER TABLE
    car_models ADD CONSTRAINT car_models_fuel_type_foreign FOREIGN KEY(fuel_type) REFERENCES fuel_types(id);
ALTER TABLE
    car_models ADD CONSTRAINT car_models_transmission_foreign FOREIGN KEY(transmission) REFERENCES car_transmissions(id);
ALTER TABLE
    car_series ADD CONSTRAINT car_series_brand_id_foreign FOREIGN KEY(brand_id) REFERENCES car_brands(id);
