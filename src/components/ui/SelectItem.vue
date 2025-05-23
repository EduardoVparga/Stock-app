<template>
    <li
      :class="cn('relative flex w-full cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50 hover:bg-accent/50', props.class)"
      @click="handleClick"
      role="option"
      :aria-selected="isSelected"
      :data-state="isSelected ? 'checked' : 'unchecked'"
    >
      <span class="absolute left-2 flex h-3.5 w-3.5 items-center justify-center">
        <CheckIcon v-if="isSelected" class="h-4 w-4" />
      </span>
      <slot />
    </li>
  </template>
  <script setup lang="ts">
  import { inject, computed, type Ref } from 'vue';
  import { cn } from '@/lib/utils';
  import { CheckIcon } from 'lucide-vue-next';
  
  const props = defineProps<{ value: string | number, class?: string }>();
  const selectedValue = inject<Ref<string | number | undefined>>('selectSelectedValue');
  const updateSelectedValue = inject<(val: string | number) => void>('selectUpdateSelectedValue');
  const closeSelect = inject<() => void>('selectClose');
  
  const isSelected = computed(() => selectedValue?.value === props.value);
  
  const handleClick = () => {
    if (updateSelectedValue) updateSelectedValue(props.value);
    // closeSelect is now called within updateSelectedValue in Select.vue
  };
  </script>