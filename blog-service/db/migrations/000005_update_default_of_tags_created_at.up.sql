-- Drop the default constraint on updated_at column
ALTER TABLE tags ALTER COLUMN updated_at DROP DEFAULT;

-- Modify the column definitions
ALTER TABLE tags ALTER COLUMN created_at SET DEFAULT NOW();
ALTER TABLE tags ALTER COLUMN updated_at DROP DEFAULT;
