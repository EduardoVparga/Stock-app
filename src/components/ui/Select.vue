<template>
  <div class="relative" ref="selectRef">
    <slot />
  </div>
</template>
<script setup lang="ts">
import { ref, provide, onMounted, onUnmounted, watch } from 'vue';
const props = defineProps<{ modelValue: string | number | undefined }>();
const emit = defineEmits(['update:modelValue']);

const isOpen = ref(false);
const selectedValue = ref(props.modelValue);
const selectRef = ref<HTMLElement | null>(null);

provide('selectIsOpen', isOpen);
provide('selectSelectedValue', selectedValue);
provide('selectToggleOpen', () => isOpen.value = !isOpen.value);
provide('selectClose', () => isOpen.value = false);
provide('selectUpdateSelectedValue', (val: string | number) => {
  selectedValue.value = val;
  emit('update:modelValue', val);
  isOpen.value = false; // Close after selection
});

const handleClickOutside = (event: MouseEvent) => {
  if (selectRef.value && !selectRef.value.contains(event.target as Node)) {
    isOpen.value = false;
  }
};

onMounted(() => {
  if(typeof document !== 'undefined') document.addEventListener('click', handleClickOutside);
});
onUnmounted(() => {
  if(typeof document !== 'undefined') document.removeEventListener('click', handleClickOutside);
});

watch(() => props.modelValue, (newVal) => {
  selectedValue.value = newVal;
});
</script>