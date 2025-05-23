// src/types/stock-action.ts
export interface StockAction {
  id: string; // Add an ID for unique key prop
  action: string;
  brokerage: string;
  company: string;
  rating_from: string | null;
  rating_to: string;
  target_from: string | null;
  target_to: string | null;
  ticker: string;
  time: string; // ISO date string
}