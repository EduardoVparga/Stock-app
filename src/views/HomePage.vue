<template>
  <main class="min-h-screen bg-background">
    <div v-if="stockStore.isLoading" class="container mx-auto py-6 text-center">Loading data...</div>
    <div v-else-if="stockStore.error" class="container mx-auto py-6 text-center text-red-500">
      Failed to load data: {{ stockStore.error.message }}
    </div>
    <StockActionsClient v-else :initial-data="stockStore.stockActions" />
  </main>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import StockActionsClient from "@/components/stock-actions/StockActionsClient.vue"
import { useStockStore } from '@/stores/stockStore'

const stockStore = useStockStore()

onMounted(() => {
  if (stockStore.stockActions.length === 0) { // Fetch only if not already loaded
    stockStore.fetchStockActions()
  }
})
</script>