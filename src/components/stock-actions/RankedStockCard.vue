<template>
    <div class="bg-card border rounded-lg p-4 shadow-sm transition-all duration-300 hover:shadow-md hover:-translate-y-1 cursor-pointer"
         @click="$emit('cardClick', stock)">
      <div class="flex items-center gap-4">
        <div class="flex flex-col items-center justify-center text-primary w-12">
          <TrophyIcon class="h-6 w-6 mb-1" />
          <span class="text-2xl font-bold leading-none">{{ stock.rank }}</span>
        </div>
  
        <Separator orientation="vertical" class="h-10" />
  
        <div class="flex-grow">
          <p class="font-semibold text-foreground truncate" :title="stock.company">
            {{ stock.company }}
          </p>
          <Badge variant="secondary" class="mt-1">{{ stock.ticker }}</Badge>
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import type { StockAction } from '@/types/stock-action';
  import Badge from '@/components/ui/Badge.vue';
  import Separator from '@/components/ui/Separator.vue';
  import { TrophyIcon } from 'lucide-vue-next';
  
  interface Props {
    stock: StockAction;
  }
  
  defineProps<Props>();
  defineEmits(['cardClick']);
  </script>