CREATE TABLE IF NOT EXISTS users (
  user_id UUID PRIMARY KEY,
  first_name VARCHAR(30),
  last_name VARCHAR(30),
  username VARCHAR(15) UNIQUE,
  password TEXT,
  phone VARCHAR(20),
  email TEXT,
  gender BOOLEAN,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS user_photo (
  image_id UUID,
  user_id UUID REFERENCES users (user_id),
  type TEXT,
  baseCode TEXT
)