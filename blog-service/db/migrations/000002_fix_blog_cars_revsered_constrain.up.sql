ALTER TABLE blogs DROP CONSTRAINT blogs_id_foreign;
ALTER TABLE blog_cars ADD CONSTRAINT blog_cars_blogs_id_foreign FOREIGN KEY (id) REFERENCES blogs (id);
