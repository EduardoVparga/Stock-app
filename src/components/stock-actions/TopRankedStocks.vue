<template>
  <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-5 gap-4">
    <RankedStockCard 
      v-for="stock in stocks" 
      :key="stock.id"
      :stock="stock"
      @card-click="$emit('stockSelect', stock)"
    />
  </div>
</template>

<script setup lang="ts">
import type { StockAction } from '@/types/stock-action';
import RankedStockCard from './RankedStockCard.vue'; 

defineProps<{
  stocks: StockAction[]
}>();

defineEmits(['stockSelect']);
</script>