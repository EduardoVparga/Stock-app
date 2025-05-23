<!-- src/components/ui/Select.vue -->
<template>
    <div class="relative">
      <slot />
    </div>
  </template>
  
  <!-- src/components/ui/SelectTrigger.vue -->
  <template>
    <Button variant="outline" :class="cn('w-full justify-between', props.class)" @click="toggleOpen" aria-haspopup="listbox" :aria-expanded="isOpen">
      <slot name="valuePlaceholder" />
      <ChevronDownIcon class="h-4 w-4 opacity-50" />
    </Button>
  </template>
  <script setup lang="ts">
  import { inject, type Ref } from 'vue';
  import Button from './Button.vue';
  import { cn } from '@/lib/utils';
  import { ChevronDownIcon } from 'lucide-vue-next';
  const props = defineProps<{ class?: string }>();
  const isOpen = inject<Ref<boolean>>('selectIsOpen');
  const toggleOpen = inject<() => void>('selectToggleOpen');
  </script>
  
  <!-- src/components/ui/SelectValue.vue -->
  <template><span><slot /></span></template>
  
  <!-- src/components/ui/SelectContent.vue -->
  <template>
    <div v-if="isOpen" :class="cn('absolute z-50 mt-1 w-full rounded-md border bg-popover text-popover-foreground shadow-md animate-in fade-in-0 zoom-in-95', props.class)">
      <ul class="max-h-60 overflow-auto p-1">
        <slot />
      </ul>
    </div>
  </template>
  <script setup lang="ts">
  import { inject, type Ref } from 'vue';
  import { cn } from '@/lib/utils';
  const props = defineProps<{ class?: string }>();
  const isOpen = inject<Ref<boolean>>('selectIsOpen');
  </script>
  
  <!-- src/components/ui/SelectItem.vue -->
  <template>
    <li
      :class="cn('relative flex w-full cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50 hover:bg-accent/50', props.class)"
      @click="handleClick"
      role="option"
      :aria-selected="isSelected"
    >
      <span v-if="isSelected" class="absolute left-2 flex h-3.5 w-3.5 items-center justify-center">
        <CheckIcon class="h-4 w-4" />
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
    if (closeSelect) closeSelect();
  };
  </script>