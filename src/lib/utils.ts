import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function formatPrice(price: number | null | undefined): string {
  if (price === null || price === undefined) return "N/A";
  return `$${price.toFixed(2)}`;
}

export function generateId(): string {
  return Math.random().toString(36).substr(2, 9);
}