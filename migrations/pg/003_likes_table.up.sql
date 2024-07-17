CREATE TABLE IF NOT EXISTS likes (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  tweet_id uuid NOT NULL,
  username VARCHAR(30) NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),

  CONSTRAINT fk_user FOREIGN KEY(username) REFERENCES users(username),
  CONSTRAINT fk_tweet FOREIGN KEY(tweet_id) REFERENCES tweets(id)
);

CREATE INDEX IF NOT EXISTS "idx-likes-username" ON public.likes USING btree (username);
CREATE INDEX IF NOT EXISTS "idx-likes-tweet_id" ON public.likes USING btree (tweet_id);