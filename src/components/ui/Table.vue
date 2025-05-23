<template>
    <div class="relative w-full overflow-auto">
      <table :class="cn('w-full caption-bottom text-sm', props.class)" v_bind="$attrs">
        <slot />
      </table>
    </div>
  </template>
  <script setup lang="ts">
  import { cn } from '@/lib/utils';
  const props = defineProps<{ class?: string }>();
  </script>