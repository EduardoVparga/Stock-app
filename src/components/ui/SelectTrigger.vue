<template>
    <Button variant="outline" :class="cn('w-full justify-between', props.class)" @click="toggleOpen" aria-haspopup="listbox" :aria-expanded="isOpenRef ? isOpenRef.value : false">
      <slot /> <!-- This slot will contain SelectValue -->
      <ChevronDownIcon class="ml-2 h-4 w-4 shrink-0 opacity-50" />
    </Button>
  </template>
  <script setup lang="ts">
  import { inject, type Ref } from 'vue';
  import Button from './Button.vue';
  import { cn } from '@/lib/utils';
  import { ChevronDownIcon } from 'lucide-vue-next';
  const props = defineProps<{ class?: string }>();
  const isOpenRef = inject<Ref<boolean>>('selectIsOpen'); // Renamed to avoid conflict
  const toggleOpen = inject<() => void>('selectToggleOpen');
  </script>