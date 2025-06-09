<template>
  <div class="relative w-full" ref="selectRef">
    <Button
      variant="secondary" 
      :class="cn(
        'w-full justify-between text-left font-normal border border-input', 
        !selectedLabel && 'text-muted-foreground',
        props.class
      )"
      @click="toggleDropdown"
      aria-haspopup="listbox"
      :aria-expanded="isOpen"
    >
      <span class="truncate">
        {{ selectedLabel || placeholder }}
      </span>
      <ChevronDownIcon class="ml-2 h-4 w-4 shrink-0 opacity-50" />
    </Button>

    <Transition name="fade-scale">
      <div
        v-if="isOpen"
        class="absolute z-50 mt-1 w-full rounded-md border bg-popover text-popover-foreground shadow-md"
      >
        <ul class="max-h-60 overflow-auto p-1" role="listbox">
          <li
            v-for="option in options"
            :key="option.value"
            :class="[
              'relative flex w-full cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none hover:bg-accent/50',
              { 'bg-accent': modelValue === option.value }
            ]"
            role="option"
            :aria-selected="modelValue === option.value"
            @click="selectOption(option)"
          >
            <span class="absolute left-2 flex h-3.5 w-3.5 items-center justify-center">
              <CheckIcon v-if="modelValue === option.value" class="h-4 w-4" />
            </span>
            <span class="truncate">{{ option.label }}</span>
          </li>
          <li v-if="!options.length" class="px-2 py-1.5 text-sm text-muted-foreground">
            No options available.
          </li>
        </ul>
      </div>
    </Transition>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, onUnmounted, nextTick } from 'vue';
import { cn } from '@/lib/utils';
import Button from './Button.vue';
import { ChevronDownIcon, CheckIcon } from 'lucide-vue-next';
interface SelectOption {
  value: string | number;
  label: string;
}
const props = defineProps<{
  modelValue: string | number | undefined;
  options: SelectOption[];
  placeholder?: string;
  class?: string;
}>();
const emit = defineEmits(['update:modelValue']);
const isOpen = ref(false);
const selectRef = ref<HTMLElement | null>(null);
const selectedLabel = computed(() => {
  return props.options.find(opt => opt.value === props.modelValue)?.label;
});
const toggleDropdown = () => {
  isOpen.value = !isOpen.value;
};
const selectOption = (option: SelectOption) => {
  emit('update:modelValue', option.value);
  isOpen.value = false;
};
const handleClickOutside = (event: MouseEvent) => {
  if (selectRef.value && !selectRef.value.contains(event.target as Node)) {
    isOpen.value = false;
  }
};
watch(isOpen, (isCurrentlyOpen) => {
  if (isCurrentlyOpen) {
    nextTick(() => {
      document.addEventListener('click', handleClickOutside);
    });
  } else {
    document.removeEventListener('click', handleClickOutside);
  }
});
onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside);
});
</script>
<style scoped>
.fade-scale-enter-active,
.fade-scale-leave-active {
  transition: opacity 0.1s ease-out, transform 0.1s ease-out;
}
.fade-scale-enter-from,
.fade-scale-leave-to {
  opacity: 0;
  transform: scale(0.95) translateY(-5px);
}
</style>