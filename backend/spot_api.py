from fastapi import FastAPI
from starlette.concurrency import run_in_threadpool
import yfinance as yf

app = FastAPI()

def _fetch_price(symbol: str) -> float:
    try:
        ticker = yf.Ticker(symbol)
        return ticker.info.get("regularMarketPrice", 0.0) or 0.0
    except Exception:
        # Si ocurre cualquier error con yfinance, devolvemos 0.0
        return 0.0

@app.get("/{symbol}")
async def get_spot_price(symbol: str):
    price = await run_in_threadpool(_fetch_price, symbol)
    return {"spot": price}

if __name__ == "__main__":
    import uvicorn
    uvicorn.run("main:app", host="0.0.0.0", port=8000, reload=True, workers=4)
