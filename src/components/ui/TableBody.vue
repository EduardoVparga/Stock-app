<template>
    <tbody :class="cn('[&_tr:last-child]:border-0', props.class)" v_bind="$attrs">
      <slot />
    </tbody>
  </template>
  <script setup lang="ts">
  import { cn } from '@/lib/utils';
  const props = defineProps<{ class?: string }>();
  </script>