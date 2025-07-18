-- up
CREATE TABLE IF NOT EXISTS feed_follows (
    id uuid PRIMARY KEY NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    user_id uuid NOT NULL,
    feed_id uuid NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_feed FOREIGN KEY (feed_id) REFERENCES feeds (id) ON DELETE CASCADE,
    CONSTRAINT user_feed UNIQUE (user_id, feed_id)
);

-- down
DROP TABLE feed_follows;

