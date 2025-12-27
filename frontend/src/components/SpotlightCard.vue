<script setup>
import { ref, onMounted, onUnmounted, computed } from 'vue'
import { useMouseInElement } from '@vueuse/core'

const target = ref(null)
const { elementX, elementY, isOutside } = useMouseInElement(target)

// Configuration props
const props = defineProps({
  gradientColor: {
    type: String,
    default: '#6366f1' // Indigo-500
  },
  slotClass: {
    type: String,
    default: ''
  }
})

// Dynamic style for the spotlight gradient
const spotlightStyle = computed(() => {
  if (isOutside.value) return { opacity: 0 }
  
  return {
    background: `radial-gradient(600px circle at ${elementX.value}px ${elementY.value}px, ${props.gradientColor}15, transparent 40%)`,
    opacity: 1
  }
})

// Dynamic style for the border spotlight
const borderStyle = computed(() => {
   if (isOutside.value) return { opacity: 0 }
   return {
     background: `radial-gradient(600px circle at ${elementX.value}px ${elementY.value}px, ${props.gradientColor}40, transparent 40%)`,
     opacity: 1
   }
})
</script>

<template>
  <div 
    ref="target"
    class="relative group rounded-3xl bg-zinc-900/40 border border-white/5 overflow-hidden transition-all duration-300 hover:shadow-2xl hover:shadow-black/50"
    :class="slotClass"
  >
    <!-- Border Glow Effect (Overlay on border) -->
    <div 
      class="absolute inset-0 z-0 pointer-events-none transition-opacity duration-500"
      :style="borderStyle"
    ></div>

    <!-- Inner Content Glow -->
    <div 
      class="absolute inset-[1px] rounded-[calc(1.5rem-1px)] z-0 pointer-events-none transition-opacity duration-500"
      :style="spotlightStyle"
    ></div>

    <!-- Content Background (To cover the full gradient but keep transparency) -->
    <div class="absolute inset-[1px] bg-zinc-900/80 rounded-[calc(1.5rem-1px)] z-0 backdrop-blur-sm"></div>

    <!-- Actual Content -->
    <div class="relative z-10 h-full">
      <slot />
    </div>
  </div>
</template>