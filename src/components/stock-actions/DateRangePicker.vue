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
      <template #default="{ togglePopover }">
        <Button
          id="date"
          variant="outline"
          :class="cn(
            'w-full justify-start text-left font-normal',
            !date?.from && 'text-muted-foreground'
          )"
          @click="togglePopover"
        >
          <div class="flex items-center justify-between w-full">
            <div class="flex items-center">
              <CalendarIcon class="mr-2 h-4 w-4" />
              <span v-if="date?.from && date.to">
                {{ format(date.from, "LLL dd, y") }} - {{ format(date.to, "LLL dd, y") }}
              </span>
              <span v-else-if="date?.from">
                {{ format(date.from, "LLL dd, y") }}
              </span>
              <span v-else>Pick a date range</span>
            </div>

            <button
              v-if="date?.from"
              class="p-1 rounded-full hover:bg-secondary"
              @click.stop="clearDate"
            >
              <XIcon class="h-4 w-4 text-muted-foreground" />
            </button>
          </div>
        </Button>
      </template>
    </VDatePicker>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { format } from 'date-fns'
import { enUS } from 'date-fns/locale'
import { Calendar as CalendarIcon, XIcon } from 'lucide-vue-next'
import { DatePicker as VDatePicker } from 'v-calendar'
import 'v-calendar/style.css'

import { cn } from '@/lib/utils'
import Button from '@/components/ui/Button.vue'
import type { DatePickerModel } from 'v-calendar/dist/types/src/use/datePicker.js'

interface AppDateRange {
  from?: Date
  to?: Date
}

interface Props {
  date?: AppDateRange
}

const props = defineProps<Props>()
const emit = defineEmits<{
  (e: 'update:date', value: AppDateRange | undefined): void
}>()

const clearDate = () => {
  emit('update:date', undefined)
}

const selectedDateRange = computed<DatePickerModel | undefined>({
  get(): DatePickerModel | undefined {
    if (!props.date?.from) return undefined
    return {
      start: props.date.from,
      ...(props.date.to ? { end: props.date.to } : {}),
    } as DatePickerModel
  },
  set(value: DatePickerModel | undefined) {
    if (!value || typeof value !== 'object' || !('start' in value) || !value.start) {
      emit('update:date', undefined)
      return
    }
    const { start, end } = value as { start: Date; end?: Date }
    emit('update:date', {
      from: start,
      to: end,
    })
  },
})
</script>
