<template>
    <Teleport to="body">
      <Transition name="fade">
        <div v-if="open" @click.self="handleOverlayClick" class="fixed inset-0 z-50 bg-black/50 flex items-center justify-center backdrop-blur-sm">
          <div :class="cn('bg-card shadow-lg rounded-lg m-4 relative', props.dialogClass)" v_bind="$attrs" role="dialog" aria-modal="true">
            <slot />
          </div>
        </div>
      </Transition>
    </Teleport>
  </template>
  <script setup lang="ts">
  import { cn } from '@/lib/utils';
  import { watch, onUnmounted } from 'vue';
  
  interface Props {
    open: boolean;
    dialogClass?: string; 
    persistent?: boolean; 
  }
  const props = defineProps<Props>();
  const emit = defineEmits(['update:open']);
  
  const handleOverlayClick = () => {
    if (!props.persistent) {
      emit('update:open', false);
    }
  }
  
  watch(() => props.open, (isOpen) => {
    if (typeof document !== 'undefined') {
      if (isOpen) {
        document.body.style.overflow = 'hidden';
      } else {
        document.body.style.overflow = '';
      }
    }
  }, { immediate: true });
  
  onUnmounted(() => {
    if (typeof document !== 'undefined') {
      document.body.style.overflow = '';
    }
  });
  </script>
  <style scoped>
  .fade-enter-active, .fade-leave-active {
    transition: opacity 0.2s ease-in-out;
  }
  .fade-enter-active .bg-card, .fade-leave-active .bg-card { /* Target inner dialog for scale/translate */
    transition: opacity 0.2s ease-in-out, transform 0.2s ease-in-out;
  }
  
  .fade-enter-from, .fade-leave-to {
    opacity: 0;
  }
  .fade-enter-from .bg-card, .fade-leave-to .bg-card {
    opacity: 0;
    transform: scale(0.95) translateY(10px);
  }
  </style>