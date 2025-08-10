CREATE TABLE IF NOT EXISTS articles (
    article_id       TEXT PRIMARY KEY,
    title            TEXT ,
    link             TEXT ,
    keywords         TEXT[] ,
    creator          TEXT[], -- nullable
    description      TEXT ,
    content          TEXT ,
    pub_date         TEXT ,
    pub_date_tz      TEXT ,
    image_url        TEXT, -- nullable
    video_url        TEXT, -- nullable
    source_id        TEXT ,
    source_name      TEXT ,
    source_priority  INTEGER ,
    source_url       TEXT ,
    source_icon      TEXT ,
    language         TEXT ,
    country          TEXT[] ,
    category         TEXT[] ,
    sentiment        TEXT ,
    sentiment_stats  TEXT ,
    ai_tag           TEXT ,
    ai_region        TEXT ,
    ai_org           TEXT ,
    ai_summary       TEXT ,
    ai_content       TEXT ,
    duplicate        BOOLEAN 
);
