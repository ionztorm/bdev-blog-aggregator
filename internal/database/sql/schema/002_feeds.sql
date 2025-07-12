-- up
CREATE TABLE IF NOT EXISTS feeds (
    id uuid PRIMARY KEY NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    name text NOT NULL,
    url text UNIQUE NOT NULL,
    user_id uuid NOT NULL,
    last_fetched_at timestamp,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- down
DROP TABLE feeds;

