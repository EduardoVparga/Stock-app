<template>
    <span :class="cn(props.placeholder && !hasValue ? 'text-muted-foreground' : '', props.class)">
      <slot v-if="hasValue" />
      <template v-else>{{ props.placeholder }}</template>
    </span>
  </template>
  <script setup lang="ts">
  import { inject, computed, type Ref} from 'vue';
  import { cn } from '@/lib/utils';
  
  const props = defineProps<{ placeholder?: string, class?: string }>()
  const selectedValue = inject<Ref<string | number | undefined>>('selectSelectedValue');
  const hasValue = computed(() => selectedValue?.value !== undefined && selectedValue.value !== '' && selectedValue.value !== null)
  </script>