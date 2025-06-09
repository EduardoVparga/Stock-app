<template>
  <div class="space-y-4">
    <div class="flex flex-col md:flex-row gap-4 items-center">

      <div class="relative w-full md:[flex-grow:5] min-w-0">
        <SearchIcon class="absolute left-3 top-1/2 -translate-y-1/2 h-5 w-5 text-muted-foreground" />
        <Input
          type="text"
          placeholder="Search company or ticker..."
          :modelValue="searchTerm"
          @update:modelValue="$emit('update:searchTerm', $event)"
          class="pl-10 pr-4 py-2 w-full bg-secondary border-border"
        />
      </div>

      <Select
        :model-value="selectedAction"
        @update:modelValue="$emit('update:selectedAction', $event || '')"
        :options="actionTypeOptions"
        placeholder="Action Type"
        class="w-full md:[flex-grow:2] min-w-0"
      />
      
      <Select
        :model-value="selectedBrokerage"
        @update:modelValue="$emit('update:selectedBrokerage', $event || '')"
        :options="brokerageFirmOptions"
        placeholder="Brokerage"
        class="w-full md:[flex-grow:2] min-w-0"
      />

      <DateRangePicker 
        :date="dateRange" 
        @update:date="$emit('update:dateRange', $event)" 
        class="w-full md:[flex-grow:3] min-w-0" 
      />

      <Button
        v-if="hasActiveFilters"
        @click="$emit('clearFilters')"
        variant="ghost"
        size="sm"
        class="text-muted-foreground hover:text-foreground flex-shrink-0"
      >
        Clear Filters
      </Button>

    </div>
  </div>
</template>

<script setup lang="ts">
  import { computed } from 'vue'
  import Input from "@/components/ui/Input.vue"
  import Select from "@/components/ui/Select.vue";
  import DateRangePicker from "./DateRangePicker.vue"
  import Button from "@/components/ui/Button.vue"
  import { SearchIcon } from "lucide-vue-next"
  
  interface DateRange {
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
  defineEmits([
    'update:searchTerm',
    'update:selectedAction',
    'update:selectedBrokerage',
    'update:dateRange',
    'clearFilters'
  ])

  const actionTypeOptions = computed(() => [
    { value: '', label: 'All Action Types' },
    ...props.actionTypes.map(action => ({ value: action, label: action }))
  ]);

  const brokerageFirmOptions = computed(() => [
    { value: '', label: 'All Brokerages' },
    ...props.brokerageFirms.map(brokerage => ({ value: brokerage, label: brokerage }))
  ]);

  const hasActiveFilters = computed(() => 
    props.searchTerm || props.selectedAction || props.selectedBrokerage || props.dateRange?.from
  )
</script>