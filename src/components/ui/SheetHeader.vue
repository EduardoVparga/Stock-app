<template>
    <div :class="cn('flex flex-col space-y-2 text-center sm:text-left mb-4', props.class)"> 
      <slot />
    </div>
  </template>
  <script setup lang="ts">
  import { cn } from '@/lib/utils';
  const props = defineProps<{ class?: string }>();
  </script>