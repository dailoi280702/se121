ALTER TABLE blog_cars DROP CONSTRAINT blog_cars_blogs_id_foreign;
ALTER TABLE blog_cars ADD CONSTRAINT blog_cars_blogs_id_foreign FOREIGN KEY (blog_id) REFERENCES blogs (id);

ALTER TABLE blog_tags DROP CONSTRAINT blog_tags_tag_id_foreign;
ALTER TABLE blog_tags ADD CONSTRAINT blog_tags_tag_id_foreign FOREIGN KEY (tag_id) REFERENCES tags (id);
