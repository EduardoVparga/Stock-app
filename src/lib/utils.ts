// src/lib/utils.ts
import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function formatPrice(price: number | null | undefined): string {
  if (price === null || price === undefined) return "N/A";
  // Puedes decidir si usar toFixed(2) si los precios pueden tener decimales
  return `$${price.toFixed(2)}`;
}

// Helper to generate unique IDs, useful if data doesn't have them
export function generateId(): string {
  return Math.random().toString(36).substr(2, 9);
}