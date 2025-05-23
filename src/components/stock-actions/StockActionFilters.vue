<template>
    <div class="p-4 md:p-6 space-y-4 rounded-lg shadow-sm border bg-card">
      <div class="flex flex-col md:flex-row gap-4 items-center">
        <div class="relative w-full md:flex-grow">
          <SearchIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-muted-foreground" />
          <Input
            type="text"
            placeholder="Search company, ticker, or brokerage..."
            :modelValue="searchTerm"
            @update:modelValue="$emit('update:searchTerm', $event)"
            class="pl-10 pr-4 py-2 w-full"
          />
        </div>
        <Button variant="outline" @click="showFilters = !showFilters" class="w-full md:w-auto">
          <FilterIcon class="h-4 w-4 mr-2" />
          {{ showFilters ? "Hide Filters" : "Show Filters" }}
        </Button>
        <Button v-if="hasActiveFilters" variant="ghost" @click="clearFilters" class="w-full md:w-auto text-sm">
          <XIcon class="h-4 w-4 mr-2" />
          Clear Filters
        </Button>
      </div>
  
      <div v-if="showFilters" class="grid grid-cols-1 md:grid-cols-3 gap-4 pt-4 border-t">
        <!-- Using basic select for brevity, replace with custom Select.vue or Headless UI -->
        
        
        
        <select
          :value="selectedAction || ALL_FILTER_VALUE"
          @change="$emit('update:selectedAction', ($event.target as HTMLSelectElement).value === ALL_FILTER_VALUE ? '' : ($event.target as HTMLSelectElement).value)"
          class="h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background focus:outline-none focus:ring-2 focus:ring-ring"
        >
          <option :value="ALL_FILTER_VALUE">All Action Types</option>
          <option v-for="action in actionTypes" :key="action" :value="action" class="capitalize">
            {{ action }}
          </option>
        </select>
        


        <select
          :value="selectedBrokerage || ALL_FILTER_VALUE"
          @change="$emit('update:selectedBrokerage', ($event.target as HTMLSelectElement).value === ALL_FILTER_VALUE ? '' : ($event.target as HTMLSelectElement).value)"
          class="h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background focus:outline-none focus:ring-2 focus:ring-ring"
        >
          <option :value="ALL_FILTER_VALUE">All Brokerages</option>
          <option v-for="brokerage in brokerageFirms" :key="brokerage" :value="brokerage">
            {{ brokerage }}
          </option>
        </select>
  
        <DateRangePicker :date="dateRange" @update:date="$emit('update:dateRange', $event)" class="w-full" />
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref, computed } from 'vue'
  import type { AppDateRange } from './DateRangePicker.vue' // Assuming AppDateRange is exported or define it here
  import Input from "@/components/ui/Input.vue"
  // If using the custom SelectProvider and parts:
  // import SelectProvider from "@/components/ui/SelectProvider.vue"
  // import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/SelectParts.vue"
  import DateRangePicker from "./DateRangePicker.vue"
  import { XIcon, SearchIcon, FilterIcon } from "lucide-vue-next"
  import Button from "@/components/ui/Button.vue"
  
  interface DateRange { // Define AppDateRange if not imported
    from?: Date
    to?: Date
  }
  
  interface StockActionFiltersProps {
    searchTerm: string
    selectedAction: string
    selectedBrokerage: string
    dateRange?: DateRange
    actionTypes: string[]
    brokerageFirms: string[]
  }
  
  const props = defineProps<StockActionFiltersProps>()
  
  const emit = defineEmits([
    'update:searchTerm',
    'update:selectedAction',
    'update:selectedBrokerage',
    'update:dateRange',
    'clearFilters'
  ])
  
  const showFilters = ref(false)
  const ALL_FILTER_VALUE = "__ALL_FILTERS__";
  
  const hasActiveFilters = computed(() => 
    props.searchTerm || props.selectedAction || props.selectedBrokerage || props.dateRange?.from
  )
  
  const clearFilters = () => {
    emit('clearFilters')
  }
  </script>