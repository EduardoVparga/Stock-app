<template>
    <section aria-labelledby="top-ranked-heading" v-if="stocks && stocks.length > 0">
        <h2 id="top-ranked-heading" class="text-2xl font-bold text-foreground mb-4">
        Top 5 Ranked Stocks
        </h2>
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-5 gap-4">
        <RankedStockCard
            v-for="stock in stocks"
            :key="stock.id"
            :stock="stock"
            @cardClick="$emit('stockSelect', $event)"
        />
        </div>
    </section>
</template>

<script setup lang="ts">
    import type { StockAction } from '@/types/stock-action';
    import RankedStockCard from './RankedStockCard.vue';

    interface Props {
    stocks: StockAction[];
    }

    defineProps<Props>();
    defineEmits(['stockSelect']);
</script>