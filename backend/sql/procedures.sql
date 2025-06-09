CREATE OR REPLACE PROCEDURE sync_symbols()
LANGUAGE PLpgSQL
AS $$
BEGIN
  INSERT INTO Symbols (ticker, isActive, loaded_at)
    SELECT DISTINCT
      ticker,
      1 AS isActive,
      now() AS loaded_at
    FROM StockInfo
  ON CONFLICT (ticker) DO UPDATE
    SET
      isActive   = EXCLUDED.isActive,
      loaded_at  = EXCLUDED.loaded_at;

  UPDATE Symbols
  SET
    isActive   = 0,
    loaded_at  = now()
  WHERE
    isActive != 0
    AND ticker NOT IN (
      SELECT DISTINCT ticker FROM StockInfo
    );
END;
$$;

CREATE OR REPLACE PROCEDURE ArchiveAndTruncate()
LANGUAGE PLpgSQL
AS $$
BEGIN
  INSERT INTO StockInfoHistory (
    action, brokerage, company,
    rating_from, rating_to,
    target_from, target_to,
    ticker, time, archived_at
  )
  SELECT
    action, brokerage, company,
    rating_from, rating_to,
    target_from, target_to,
    ticker, time, loaded_at
  FROM StockInfo;

  DELETE FROM StockInfo;
END;
$$;
