<template>
  <div class="container mx-auto py-10 px-4 md:px-0">
    <header class="mb-8">
      <h1 class="text-4xl font-bold text-foreground">Equity Pulse</h1>
    </header>

    <div class="space-y-8">
      <div>
        <h2 class="text-2xl font-semibold mb-4">Top Movers</h2>
        <TopRankedStocks 
          v-if="topRankedStocks.length > 0"
          :stocks="topRankedStocks"
          @stockSelect="handleRowClick"
        />
      </div>

      <div>
        <h2 class="text-2xl font-semibold mb-4">Quick Filters</h2>
        <StockActionFilters
          :search-term="searchTerm"
          @update:searchTerm="searchTerm = $event"
          :selected-action="selectedAction"
          @update:selectedAction="selectedAction = $event"
          :selected-brokerage="selectedBrokerage"
          @update:selectedBrokerage="selectedBrokerage = $event"
          :date-range="dateRange"
          @update:dateRange="dateRange = $event"
          :action-types="actionTypes"
          :brokerage-firms="brokerageFirms"
          @clearFilters="clearFilters"
        />
      </div>

      <div>
        <h2 class="text-2xl font-semibold mb-4">Ratings Grid</h2>
        <StockActionsTable
          :columns="columns"
          :data="filteredData"
          :on-row-click="handleRowClick"
        />
      </div>
    </div>

    <StockDetailModal
      :stock-action="selectedStockAction"
      :is-open="isModalOpen"
      @update:isOpen="isModalOpen = $event"
    />
  </div>
</template>

<script setup lang="ts">
  import { ref, computed } from 'vue'
  import type { StockAction } from "@/types/stock-action"
  import StockActionsTable from "./StockActionsTable.vue"
  import { columns } from "./columns"
  import StockActionFilters from "./StockActionFilters.vue"
  import StockDetailModal from "./StockDetailModal.vue"
  import TopRankedStocks from './TopRankedStocks.vue'

  interface DateRange {
    from?: Date
    to?: Date
  }
    
  interface StockActionsClientProps {
    initialData: StockAction[]
  }
  
  const props = defineProps<StockActionsClientProps>()
  
  const searchTerm = ref("")
  const selectedAction = ref("")
  const selectedBrokerage = ref("")
  const dateRange = ref<DateRange | undefined>(undefined)
    
  const selectedStockAction = ref<StockAction | null>(null)
  const isModalOpen = ref(false)

  const topRankedStocks = computed(() => {
    return [...props.initialData]
      .sort((a, b) => a.rank - b.rank)
      .slice(0, 5);
  });
  
  const actionTypes = computed(() => {
    const uniqueActions = new Set(props.initialData.map(item => item.action.toLowerCase()));
    return Array.from(uniqueActions).map(action => 
        action.split(' ').map(word => word.charAt(0).toUpperCase() + word.slice(1)).join(' ')
    );
  });
  
  const brokerageFirms = computed(() => 
    Array.from(new Set(props.initialData.map(item => item.brokerage))).sort()
  );
  
  const filteredData = computed(() => {
    let data = [...props.initialData]
  
    if (searchTerm.value) {
      const lowerSearchTerm = searchTerm.value.toLowerCase()
      data = data.filter(
        (item) =>
          item.company.toLowerCase().includes(lowerSearchTerm) ||
          item.ticker.toLowerCase().includes(lowerSearchTerm) ||
          item.brokerage.toLowerCase().includes(lowerSearchTerm)
      )
    }
  
    if (selectedAction.value) {
      data = data.filter(
        (item) => item.action.toLowerCase() === selectedAction.value.toLowerCase()
      )
    }
  
    if (selectedBrokerage.value) {
      data = data.filter(
        (item) => item.brokerage === selectedBrokerage.value
      )
    }
  
    if (dateRange.value?.from) {
      data = data.filter((item) => {
        const itemDate = new Date(item.time)
        itemDate.setHours(0,0,0,0); 
        
        let fromDate = new Date(dateRange.value!.from!);
        fromDate.setHours(0,0,0,0); 
  
        if (dateRange.value!.to) {
          let toDate = new Date(dateRange.value!.to!);
          toDate.setHours(23,59,59,999); 
          return itemDate >= fromDate && itemDate <= toDate;
        }
        return itemDate >= fromDate;
      })
    }
    return data
  })
  
  const handleRowClick = (stockActionItem: StockAction) => {
    selectedStockAction.value = stockActionItem
    isModalOpen.value = true
  }
  
  const clearFilters = () => {
    searchTerm.value = "";
    selectedAction.value = "";
    selectedBrokerage.value = "";
    dateRange.value = undefined;
  };
</script>