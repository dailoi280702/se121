ALTER TABLE tags DROP CONSTRAINT tags_id_foreign;
ALTER TABLE blog_tags ADD CONSTRAINT blog_tags_tag_id_foreign FOREIGN KEY (id) REFERENCES blogs (id);
