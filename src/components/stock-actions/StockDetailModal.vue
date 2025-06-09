<template>
    <Sheet :open="isOpen" @update:open="$emit('update:isOpen', $event)" dialogClass="sm:max-w-lg w-[90vw]">
      <SheetContent v-if="stockAction">
        <SheetHeader class="mb-6">
          <SheetTitle class="text-2xl font-bold">
            {{ stockAction.company }} ({{ stockAction.ticker }})
          </SheetTitle>
          <SheetDescription>
            Action by {{ stockAction.brokerage }} on
            {{ format(new Date(stockAction.time), "MMMM dd, yyyy 'at' HH:mm") }}
          </SheetDescription>
        </SheetHeader>
  
        <div class="space-y-4 text-sm">
          <div>
            <h4 class="font-semibold text-muted-foreground">Action Type</h4>
            <p class="capitalize">{{ stockAction.action }}</p>
          </div>
          <Separator />
          <div>
            <h4 class="font-semibold text-muted-foreground">Rating Change</h4>
            <p>
              {{ stockAction.rating_from || "N/A" }} → {{ stockAction.rating_to }}
            </p>
          </div>
          <Separator />
          <div>
            <h4 class="font-semibold text-muted-foreground">Price Target Change</h4>
            <p :class="priceTargetColorClass">
              {{ formatPrice(targetFrom) }} → {{ formatPrice(targetTo) }}
            </p>
          </div>
          <Separator />
          <div>
            <h4 class="font-semibold text-muted-foreground">External Links</h4>
            <div class="space-y-1 mt-1">
              <Button variant="link" class="p-0 h-auto text-primary hover:underline" as-child>
                <a :href="`https://finance.yahoo.com/quote/${stockAction.ticker}`" target="_blank" rel="noopener noreferrer">
                  Yahoo Finance <ExternalLink class="ml-1 h-3 w-3 inline" />
                </a>
              </Button>
              <br />
              <Button variant="link" class="p-0 h-auto text-primary hover:underline" as-child>
                <a :href="`https://seekingalpha.com/symbol/${stockAction.ticker}`" target="_blank" rel="noopener noreferrer">
                  Seeking Alpha <ExternalLink class="ml-1 h-3 w-3 inline" />
                </a>
              </Button>
            </div>
          </div>
        </div>
        
        <SheetFooter class="mt-8">
          <Button @click="$emit('update:isOpen', false)" variant="outline">
            Close
          </Button>
        </SheetFooter>
      </SheetContent>
    </Sheet>
  </template>
  
  <script setup lang="ts">
  import { computed } from 'vue'
  import type { StockAction } from "@/types/stock-action"

  import Sheet from "@/components/ui/Sheet.vue"
  import SheetContent from "@/components/ui/SheetContent.vue"
  import SheetHeader from "@/components/ui/SheetHeader.vue"
  import SheetTitle from "@/components/ui/SheetTitle.vue"
  import SheetDescription from "@/components/ui/SheetDescription.vue"
  import SheetFooter from "@/components/ui/SheetFooter.vue"
  import Button from "@/components/ui/Button.vue"
  import Separator from "@/components/ui/Separator.vue"
  import { format } from "date-fns"
  import { ExternalLink } from "lucide-vue-next"
  import { formatPrice } from "@/lib/utils"
  
  interface StockDetailModalProps {
    stockAction: StockAction | null
    isOpen: boolean
  }
  
  const props = defineProps<StockDetailModalProps>()
  const emit = defineEmits(['update:isOpen'])
  
  const targetFrom = computed(() => props.stockAction ? props.stockAction.target_from : null)
  const targetTo = computed(() => props.stockAction ? props.stockAction.target_to : null)
  
  const priceTargetColorClass = computed(() => {
    if (!props.stockAction) return "text-foreground";
    
    const tf = targetFrom.value;
    const tt = targetTo.value;
  
    let colorClass = "text-foreground";
    if (tt !== null && tf !== null) {
      if (tt > tf) colorClass = "text-green-600 dark:text-green-500";
      else if (tt < tf) colorClass = "text-red-600 dark:text-red-500";
    } else if (tt !== null && tf === null && props.stockAction.action.toLowerCase().includes("initiated")) {
       colorClass = "text-green-600 dark:text-green-500";
    }
    return colorClass;
  })
  
  </script>