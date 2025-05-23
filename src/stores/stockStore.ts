// src/stores/stockStore.ts
import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { StockAction } from '@/types/stock-action'
import { generateId } from '@/lib/utils'  

export const useStockStore = defineStore('stock', () => {
  const stockActions = ref<StockAction[]>([])
  const isLoading = ref(false)
  const error = ref<Error | null>(null)

  async function fetchStockActions() {
    isLoading.value = true
    error.value = null
    try {
      const res = await fetch("http://localhost:8080/stocks", {
        cache: "no-store", // Standard fetch cache option
      });
      if (!res.ok) {
        throw new Error(`Error fetching /stocks: ${res.status}`);
      }
      const rawData: Omit<StockAction, "id">[] = await res.json();
      stockActions.value = rawData.map(item => ({
        ...item,
        id: generateId(),
      }));
    } catch (e) {
      error.value = e as Error
      console.error("Failed to fetch stock actions:", e)
    } finally {
      isLoading.value = false
    }
  }

  return {
    stockActions,
    isLoading,
    error,
    fetchStockActions,
  }
})