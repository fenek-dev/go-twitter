CREATE TABLE IF NOT EXISTS tweets (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  content VARCHAR(255) NOT NULL,
  username VARCHAR(30) NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),

  CONSTRAINT fk_user FOREIGN KEY(username) REFERENCES users(username)
);

CREATE INDEX IF NOT EXISTS "idx-tweets-username" ON PUBLIC.tweets USING btree (username);

CREATE TRIGGER update_tweets_updated_at
BEFORE UPDATE ON public.tweets
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();