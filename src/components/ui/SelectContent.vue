<template>
    <Transition name="fade-scale">
      <div v-if="isOpenRef?.value" :class="cn('absolute z-50 min-w-[8rem] overflow-hidden mt-1 w-full rounded-md border bg-popover text-popover-foreground shadow-md data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2', props.class)">
        <ul class="max-h-60 overflow-auto p-1">
          <slot />
        </ul>
      </div>
    </Transition>
  </template>
  <script setup lang="ts">
  import { inject, type Ref } from 'vue';
  import { cn } from '@/lib/utils';
  const props = defineProps<{ class?: string }>();
  const isOpenRef = inject<Ref<boolean>>('selectIsOpen'); // Renamed to avoid conflict
  </script>
  <style scoped>
  .fade-scale-enter-active,
  .fade-scale-leave-active {
    transition: opacity 0.1s ease-out, transform 0.1s ease-out;
  }
  .fade-scale-enter-from,
  .fade-scale-leave-to {
    opacity: 0;
    transform: scale(0.95);
  }
  </style>