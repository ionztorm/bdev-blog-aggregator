-- +goose Up
CREATE TABLE posts (
    id uuid PRIMARY KEY NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    title text NOT NULL,
    url text UNIQUE NOT NULL,
    description text,
    published_at timestamp,
    feed_id uuid NOT NULL,
    CONSTRAINT fk_feed_id FOREIGN KEY (feed_id) REFERENCES feeds (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;

