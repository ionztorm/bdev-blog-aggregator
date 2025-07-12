-- up
CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    name text NOT NULL UNIQUE
);

-- down
DROP TABLE users;

