CREATE TABLE IF NOT EXISTS follows (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  follower UUID NOT NULL,
  following UUID NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),

  CONSTRAINT fk_follower FOREIGN KEY(follower) REFERENCES users(id),
  CONSTRAINT fk_following FOREIGN KEY(following) REFERENCES users(id)
);


CREATE INDEX IF NOT EXISTS "idx-follows-follower" ON public.follows USING btree (follower);
CREATE INDEX IF NOT EXISTS "idx-follows-following" ON public.follows USING btree (following);