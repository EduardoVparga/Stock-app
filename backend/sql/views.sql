CREATE OR REPLACE VIEW stock_scores AS
WITH inf AS (
  SELECT 
    spt.ticker,
    sto.action,
    sto.brokerage,
    sto.company,
    sto.rating_from,
    sto.rating_to,
    sto.target_from,
    sto.target_to,
    sto.time,
    rtg.value                                               AS sentiment,
    (sto.target_to - spt.price) / NULLIF(spt.price, 0)      AS upside,
    spt.price                                               AS spot,
    spt.loaded_at                                           AS last_updated
  FROM stockinfo AS sto
    JOIN spotprices AS spt    ON sto.ticker = spt.ticker
    JOIN ratingscale AS rtg   ON sto.rating_to = rtg.rating
),
stats AS (
  SELECT
    AVG(upside)      AS mean_upside,
    STDDEV(upside)   AS std_upside,
    AVG(sentiment)   AS mean_sentiment,
    STDDEV(sentiment) AS std_sentiment
  FROM inf
),
scored AS (
  SELECT
    inf.*,
    1 * ((inf.sentiment - stats.mean_sentiment) / NULLIF(stats.std_sentiment, 0)) + 10 * ((inf.upside    - stats.mean_upside)    / NULLIF(stats.std_upside,    0)) AS score
  FROM inf
  CROSS JOIN stats
)
SELECT
  ticker,
  action,
  brokerage,
  company,
  rating_from,
  rating_to,
  target_from,
  target_to,
  time,
  sentiment,
  upside,
  spot,
  last_updated,
  score,
  RANK() OVER (ORDER BY score DESC) AS rank
FROM scored;
