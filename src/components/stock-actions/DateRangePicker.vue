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

interface AppDateRange {
  from?: Date
  to?: Date
}

interface VCalendarDateRange {
  start: Date | null
  end: Date | null
}

interface Props {
  date?: AppDateRange
}
const props = defineProps<Props>()
const emit = defineEmits(['update:date'])

const clearDate = () => {
  emit('update:date', undefined);
};

const selectedDateRange = computed<VCalendarDateRange | undefined>({
  get: () => {
    if (!props.date || !props.date.from) {
      return undefined
    }
    return {
      start: props.date.from,
      end: props.date.to || null,
    }
  },
  set: (value) => {
    if (!value || !value.start) {
      emit('update:date', undefined);
      return;
    }

    emit('update:date', {
      from: value.start,
      to: value.end || undefined,
    });
  },
});
</script>