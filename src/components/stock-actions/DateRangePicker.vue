<template>
    <div :class="cn('grid gap-2', $attrs.class as string)">
      <VDatePicker
        v-model.range="selectedDateRange"
        mode="date"
        :masks="{ modelValue: 'YYYY-MM-DD' }"
        :popover="{ visibility: 'click' }"
        is-range
        :locale="enUS"
      >
        <template #default="{ togglePopover, inputValue }">
          <Button
            id="date"
            variant="outline"
            :class="cn(
              'w-full justify-start text-left font-normal',
              !selectedDateRange?.start && 'text-muted-foreground'
            )"
            @click="() => togglePopover()"
          >
            <CalendarIcon class="mr-2 h-4 w-4" />
            <span v-if="inputValue.start && inputValue.end">
              {{ format(new Date(inputValue.start), "LLL dd, y") }} - {{ format(new Date(inputValue.end), "LLL dd, y") }}
            </span>
            <span v-else-if="inputValue.start">
              {{ format(new Date(inputValue.start), "LLL dd, y") }}
            </span>
            <span v-else>Pick a date range</span>
          </Button>
        </template>
      </VDatePicker>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref, watch, computed } from 'vue'
  import { format } from 'date-fns'
  import { enUS } from 'date-fns/locale'
  import { Calendar as CalendarIcon } from 'lucide-vue-next'
  import { DatePicker as VDatePicker } from 'v-calendar'
  import 'v-calendar/style.css' // Import v-calendar styles
  
  import { cn } from '@/lib/utils'
  import Button from '@/components/ui/Button.vue'
  
  // For v-calendar, DateRange is typically { start: Date | null, end: Date | null }
  interface VCalendarDateRange {
    start: Date | null
    end: Date | null
  }
  // Your app uses DateRange { from?: Date, to?: Date }
  interface AppDateRange {
    from?: Date
    to?: Date
  }
  
  interface Props {
    date?: AppDateRange
    className?: string // $attrs.class will handle this
  }
  const props = defineProps<Props>()
  const emit = defineEmits(['update:date']) // Emitting 'update:date' for v-model like behavior
  
  const internalDate = ref<VCalendarDateRange>({
    start: props.date?.from || null,
    end: props.date?.to || null,
  });
  
  // Computed property to bridge v-calendar's range format and your app's format
  const selectedDateRange = computed<VCalendarDateRange>({
    get: () => ({
      start: props.date?.from || null,
      end: props.date?.to || null,
    }),
    set: (value) => {
      emit('update:date', {
        from: value.start || undefined,
        to: value.end || undefined,
      });
    },
  });
  
  // Watch for external changes to props.date
  watch(() => props.date, (newDate) => {
    internalDate.value = {
      start: newDate?.from || null,
      end: newDate?.to || null,
    };
  }, { deep: true });
  
  </script>
  
  <style>
  /* You might need to style v-calendar to match your theme */
  /* Example: .vc-container { border-radius: var(--radius); } */
  </style>