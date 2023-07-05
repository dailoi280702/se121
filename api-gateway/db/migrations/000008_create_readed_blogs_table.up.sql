CREATE TABLE readed_blogs(
    id SERIAL NOT NULL,
    user_id UUID NOT NULL,
    blog_id BIGINT NOT NULL,
    at TIMESTAMP WITHOUT TIME ZONE NULL DEFAULT NOW()
);
ALTER TABLE
    readed_blogs ADD PRIMARY KEY(id);
ALTER TABLE
    readed_blogs ADD CONSTRAINT readed_blogs_user_id_foreign FOREIGN KEY(user_id) REFERENCES users(id);
