// src/types/stock-action.ts
export interface StockAction {
  id: string; 
  action: string;
  brokerage: string;
  company: string;
  rating_from: string | null;
  rating_to: string;
  target_from: number | null;
  target_to: number | null;   
  ticker: string;
  time: string; 
  last_updated: string; 
  rank: number;         
  score: number;        
  sentiment: number;    
  spot: number;         
  upside: number;       
}