-- Create "users" table
CREATE TABLE users (
  id int NOT NULL PRIMARY KEY,
  name text NOT NULL,
  slug text NOT NULL,
  reserved_spots int NOT NULL
);
-- Create "tokens" table
CREATE TABLE tokens (
	token text NOT NULL UNIQUE,
	expires_at int NOT NULL,
	user_id int NOT NULL,
  token_type text NOT NULL,
  PRIMARY KEY (token),
	FOREIGN KEY (user_id) REFERENCES user(id)
);
-- Create "sessions" table
CREATE TABLE sessions (
  id int NOT NULL PRIMARY KEY,
  token text NOT NULL,
  expires_at int NOT NULL,
  user_id int NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
  FOREIGN KEY (token) REFERENCES tokens(token)
);
-- Create "reservations" table
CREATE TABLE confirmations (
  id int NOT NULL PRIMARY KEY,
  user_id int NOT NULL,
  confirmed_name text NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

