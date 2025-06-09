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
  
  <template>
    <thead :class="cn('[&_tr]:border-b', props.class)" v_bind="$attrs">
      <slot />
    </thead>
  </template>
  <script setup lang="ts">
  import { cn } from '@/lib/utils';
  const props = defineProps<{ class?: string }>();
  </script>
  
  <template>
    <tbody :class="cn('[&_tr:last-child]:border-0', props.class)" v_bind="$attrs">
      <slot />
    </tbody>
  </template>
  <script setup lang="ts">
  import { cn } from '@/lib/utils';
  const props = defineProps<{ class?: string }>();
  </script>
  
  <template>
    <tr :class="cn('border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted', props.class)" v_bind="$attrs">
      <slot />
    </tr>
  </template>
  <script setup lang="ts">
  import { cn } from '@/lib/utils';
  const props = defineProps<{ class?: string }>();
  </script>
  
  <template>
    <th :class="cn('h-12 px-4 text-left align-middle font-medium text-muted-foreground [&:has([role=checkbox])]:pr-0', props.class)" v_bind="$attrs">
      <slot />
    </th>
  </template>
  <script setup lang="ts">
  import { cn } from '@/lib/utils';
  const props = defineProps<{ class?: string }>();
  </script>
  
  <template>
    <td :class="cn('p-4 align-middle [&:has([role=checkbox])]:pr-0', props.class)" v_bind="$attrs">
      <slot />
    </td>
  </template>
  <script setup lang="ts">
  import { cn } from '@/lib/utils';
  const props = defineProps<{ class?: string }>();
  </script>
  
  <template>
    <div :class="badgeVariants({ variant, class: props.class })" v_bind="$attrs">
      <slot />
    </div>
  </template>
  <script setup lang="ts">
  import { cva, type VariantProps } from 'class-variance-authority'
  import { cn } from '@/lib/utils'
  
  const badgeVariants = cva(
    "inline-flex items-center rounded-full border px-2.5 py-0.5 text-xs font-semibold transition-colors focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2",
    {
      variants: {
        variant: {
          default: "border-transparent bg-primary text-primary-foreground hover:bg-primary/80",
          secondary: "border-transparent bg-secondary text-secondary-foreground hover:bg-secondary/80",
          destructive: "border-transparent bg-destructive text-destructive-foreground hover:bg-destructive/80",
          outline: "text-foreground",
        },
      },
      defaultVariants: {
        variant: "default",
      },
    }
  )
  type BadgeProps = VariantProps<typeof badgeVariants>
  interface Props {
    variant?: BadgeProps["variant"]
    class?: string
  }
  const props = withDefaults(defineProps<Props>(), {
    variant: 'default'
  })
  </script>
  
  <template>
    <hr :class="cn('shrink-0 bg-border', orientation === 'horizontal' ? 'h-[1px] w-full' : 'h-full w-[1px]', props.class)" v_bind="$attrs" />
  </template>
  <script setup lang="ts">
  import { cn } from '@/lib/utils';
  interface Props {
    orientation?: 'horizontal' | 'vertical'
    class?: string
  }
  const props = withDefaults(defineProps<Props>(), {
    orientation: 'horizontal'
  })
  </script>
  
  <template>
    <Teleport to="body">
      <Transition name="fade">
        <div v-if="open" @click.self="handleOverlayClick" class="fixed inset-0 z-50 bg-black/50 flex items-center justify-center">
          <div :class="cn('bg-card p-6 shadow-lg rounded-lg sm:max-w-lg w-[90vw] m-4', props.class)" v_bind="$attrs" role="dialog" aria-modal="true">
            <slot />
          </div>
        </div>
      </Transition>
    </Teleport>
  </template>
  <script setup lang="ts">
  import { cn } from '@/lib/utils';
  import { watch } from 'vue';
  
  interface Props {
    open: boolean;
    class?: string;
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
    if (isOpen) {
      document.body.style.overflow = 'hidden';
    } else {
      document.body.style.overflow = '';
    }
  });
  import { onUnmounted } from 'vue';
  onUnmounted(() => {
    document.body.style.overflow = '';
  });
  </script>
  <style scoped>
  .fade-enter-active, .fade-leave-active {
    transition: opacity 0.3s ease;
  }
  .fade-enter-from, .fade-leave-to {
    opacity: 0;
  }
  </style>
  
  <template><div :class="props.class"><slot /></div></template>
  <script setup lang="ts">const props = defineProps<{ class?: string }>();</script>
  
  <template><div :class="cn('mb-6', props.class)"><slot /></div></template>
  <script setup lang="ts">import { cn } from '@/lib/utils'; const props = defineProps<{ class?: string }>();</script>
  
  <template><h2 :class="cn('text-2xl font-bold', props.class)"><slot /></h2></template>
  <script setup lang="ts">import { cn } from '@/lib/utils'; const props = defineProps<{ class?: string }>();</script>
  
  <template><p :class="cn('text-muted-foreground', props.class)"><slot /></p></template>
  <script setup lang="ts">import { cn } from '@/lib/utils'; const props = defineProps<{ class?: string }>();</script>
  
  <template><div :class="cn('mt-8 flex flex-col-reverse sm:flex-row sm:justify-end sm:space-x-2', props.class)"><slot /></div></template>
  <script setup lang="ts">import { cn } from '@/lib/utils'; const props = defineProps<{ class?: string }>();</script>