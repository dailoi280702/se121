CREATE TABLE blog_cars(
    id SERIAL NOT NULL,
    blog_id INTEGER NOT NULL,
    car_id INTEGER NOT NULL
);
ALTER TABLE
    blog_cars ADD PRIMARY KEY(id);

CREATE TABLE blogs(
    id SERIAL NOT NULL,
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    image_url TEXT NULL,
    author TEXT NULL,
    tldr TEXT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NULL
);
ALTER TABLE
    blogs ADD PRIMARY KEY(id);

CREATE TABLE blog_comments(
    id SERIAL NOT NULL,
    blog_id INTEGER NOT NULL,
    user_id TEXT NOT NULL,
    comment TEXT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NULL
);
ALTER TABLE
    blog_comments ADD PRIMARY KEY(id);

CREATE TABLE tags(
    id SERIAL NOT NULL,
    name TEXT NOT NULL,
    description TEXT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE NULL DEFAULT NOW()
);
ALTER TABLE
    tags ADD PRIMARY KEY(id);

CREATE TABLE blog_tags(
    id SERIAL NOT NULL,
    tag_id INTEGER NOT NULL,
    blog_id INTEGER NOT NULL
);
ALTER TABLE
    blog_tags ADD PRIMARY KEY(id);

ALTER TABLE
    blog_tags ADD CONSTRAINT blog_tags_blog_id_foreign FOREIGN KEY(blog_id) REFERENCES blogs(id);
ALTER TABLE
    tags ADD CONSTRAINT tags_id_foreign FOREIGN KEY(id) REFERENCES blog_tags(id);
ALTER TABLE
    blog_comments ADD CONSTRAINT blog_comments_blog_id_foreign FOREIGN KEY(blog_id) REFERENCES blogs(id);
ALTER TABLE
    blogs ADD CONSTRAINT blogs_id_foreign FOREIGN KEY(id) REFERENCES blog_cars(id);
