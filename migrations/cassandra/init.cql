CREATE KEYSPACE IF NOT EXISTS twitter
    WITH REPLICATION = {
        'class' : 'NetworkTopologyStrategy',
        'DC1' : 1
        };

CONSISTENCY ONE;

CREATE TABLE IF NOT EXISTS twitter.tweet_by_user (
    user_id UUID,
    tweet_id UUID,
    content TEXT,
    created_at TIMESTAMP,
    PRIMARY KEY ((user_id), created_at, tweet_id)
) WITH CLUSTERING ORDER BY (created_at DESC, tweet_id DESC)
    AND COMPACTION = { 'class' :  'LeveledCompactionStrategy'  };