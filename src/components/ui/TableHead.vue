<template>
    <th :class="cn('h-12 px-4 text-left align-middle font-medium text-muted-foreground [&:has([role=checkbox])]:pr-0', props.class)" v_bind="$attrs">
      <slot />
    </th>
  </template>
  <script setup lang="ts">
  import { cn } from '@/lib/utils';
  const props = defineProps<{ class?: string }>();
  </script>