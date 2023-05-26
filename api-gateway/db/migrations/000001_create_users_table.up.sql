CREATE EXTENSION "uuid-ossp";

CREATE TABLE users (
    id          UUID        PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        TEXT        NOT NULL,
    email       TEXT        NULL,
    image_url   TEXT        NULL,
    create_at   TIMESTAMP   WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    is_admin    BOOLEAN     NOT NULL DEFAULT true,
    password    TEXT        NOT NULL
);
