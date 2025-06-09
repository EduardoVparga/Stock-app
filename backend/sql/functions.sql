CREATE OR REPLACE FUNCTION LastLoadTime()
RETURNS timestamptz
LANGUAGE SQL
AS $$
  SELECT COALESCE(MAX(loaded_at), '1970-01-01T00:00:00Z'::timestamptz)
    FROM StockInfo;
$$;

CREATE OR REPLACE FUNCTION get_active_tickers()
  RETURNS TABLE (ticker STRING)
  LANGUAGE SQL
AS $$
  SELECT ticker
  FROM symbols
  WHERE isactive = 1;
$$;