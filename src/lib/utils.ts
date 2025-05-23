// src/lib/utils.ts
import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function parsePrice(priceString: string | null | undefined): number | null {
  if (!priceString) return null;
  const numberString = priceString.replace('$', '').replace(',', '');
  const price = parseFloat(numberString);
  return isNaN(price) ? null : price;
}

export function formatPrice(price: number | null | undefined): string {
  if (price === null || price === undefined) return "N/A";
  return `$${price.toFixed(2)}`;
}

// Helper to generate unique IDs, useful if data doesn't have them
export function generateId(): string {
  return Math.random().toString(36).substr(2, 9);
}