CREATE KEYSPACE IF NOT EXISTS twitter
    WITH REPLICATION = {
        'class' : 'NetworkTopologyStrategy',
        'datacenter1' : 2
        };

CONSISTENCY ONE;

CREATE TABLE IF NOT EXISTS twitter.tweet_by_user (
    user_id uuid,
    tweet_id uuid,
    content text,
    created_at timestamp,
    PRIMARY KEY ((user_id), tweet_id)
) WITH
    compaction = { 'class' :  'LeveledCompactionStrategy'  };