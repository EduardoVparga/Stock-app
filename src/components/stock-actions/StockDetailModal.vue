<template>
    <Sheet :open="isOpen" @update:open="$emit('update:isOpen', $event)" dialogClass="sm:max-w-lg w-[90vw]">
      <!-- SheetContent is just a div wrapper in my basic example, so structure inside Sheet -->
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
  
  // The asChild prop for Button is a pattern from Radix.
  // For simple cases, you can wrap the <a> tag in the <Button> slot,
  // or style the <a> tag to look like a button if `asChild` implies rendering the child as the button.
  // My basic Button.vue doesn't support `asChild`. You'd typically just put an <a> inside a <button> or style the <a>.
  // Here, I'll assume the Button with variant="link" is styled like a link and can contain an <a>.
  // For true `asChild` behavior (where the Button component renders an `<a>` tag), you'd need a more complex Button component.
  // A common Vue pattern is to pass `tag="a"` to the Button component if it supports dynamic root elements.
  // Given the usage `variant="link"`, it's likely an `<a>` styled as a button or a button that behaves like a link.
  // I'll adjust the Button component to accept `as="a"` or similar for this behavior. Or simply:
  // <a :href="..." class="p-0 h-auto text-primary hover:underline ... (button link styles)">...</a>
  // For simplicity, I'll keep the Button component and rely on its `link` variant styling.
  </script>