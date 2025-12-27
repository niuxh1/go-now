<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { ArrowUp } from 'lucide-vue-next'

const isVisible = ref(false)

const handleScroll = () => {
  isVisible.value = window.scrollY > 300
}

const scrollToTop = () => {
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

onMounted(() => {
  window.addEventListener('scroll', handleScroll)
})

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
})
</script>

<template>
  <transition name="fade">
    <button 
      v-if="isVisible"
      @click="scrollToTop"
      class="fixed bottom-8 right-8 z-40 p-3 bg-fish-blue text-white rounded-full shadow-lg shadow-fish-blue/30 hover:bg-fish-orange hover:shadow-fish-orange/30 hover:-translate-y-1 transition-all duration-300"
      aria-label="Back to Top"
    >
      <ArrowUp size="24" />
    </button>
  </transition>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease, transform 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(20px);
}
</style>