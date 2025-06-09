CREATE TABLE IF NOT EXISTS Symbols (
  ticker     TEXT        PRIMARY KEY,
  isActive   INT         NOT NULL DEFAULT 1,
  loaded_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS SpotPrices (
  ticker     TEXT        PRIMARY KEY,
  price  	 NUMERIC(10,2),
  loaded_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS RatingScale (
    value INT,          
    rating TEXT,         
    category TEXT        
);

create table if not exists StockInfo (
  action TEXT,
  brokerage TEXT,
  company TEXT,
  rating_from TEXT,
  rating_to TEXT,
  target_from numeric(10,2),
  target_to numeric(10,2),
  ticker TEXT,
  time timestamptz,
  loaded_at timestamptz not null
);

CREATE TABLE IF NOT EXISTS StockInfoHistory (
  action       TEXT,
  brokerage    TEXT,
  company      TEXT,
  rating_from  TEXT,
  rating_to    TEXT,
  target_from  NUMERIC(10,2),
  target_to    NUMERIC(10,2),
  ticker       TEXT,
  time         TIMESTAMPTZ,
  archived_at  TIMESTAMPTZ DEFAULT now()
);