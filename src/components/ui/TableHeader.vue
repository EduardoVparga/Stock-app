<template>
    <thead :class="cn('[&_tr]:border-b', props.class)" v_bind="$attrs">
      <slot />
    </thead>
  </template>
  <script setup lang="ts">
  import { cn } from '@/lib/utils';
  const props = defineProps<{ class?: string }>();
  </script>