<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { ArrowUp } from 'lucide-vue-next'

const isVisible = ref(false)
const handleScroll = () => { isVisible.value = window.scrollY > 400 }
const scrollToTop = () => { window.scrollTo({ top: 0, behavior: 'smooth' }) }
onMounted(() => window.addEventListener('scroll', handleScroll))
onUnmounted(() => window.removeEventListener('scroll', handleScroll))
</script>

<template>
  <transition name="pop">
    <button v-if="isVisible" @click="scrollToTop"
      class="seal-button fixed bottom-8 right-8 z-40 w-12 h-12 flex flex-col items-center justify-center group"
      aria-label="Back to Top">
      <span class="text-base font-kai font-bold leading-none group-hover:scale-110 transition-transform duration-300">上</span>
    </button>
  </transition>
</template>

<style scoped>
.seal-button {
  background: linear-gradient(135deg, #c0392b, #a93226);
  color: rgba(255, 255, 255, 0.9);
  border: 2px solid rgba(180, 130, 80, 0.3);
  border-radius: 6px;
  box-shadow:
    0 2px 12px rgba(192, 57, 43, 0.2),
    inset 0 0 0 1px rgba(255, 255, 255, 0.1);
  transition: all 0.3s cubic-bezier(0.22, 1, 0.36, 1);
  font-family: 'KaiTi', 'STKaiti', serif;
}

.seal-button::before {
  content: '';
  position: absolute;
  inset: 3px;
  border: 1px solid rgba(255, 255, 255, 0.15);
  border-radius: 3px;
  pointer-events: none;
}

.seal-button:hover {
  transform: translateY(-2px);
  box-shadow:
    0 4px 20px rgba(192, 57, 43, 0.3),
    0 0 15px rgba(180, 130, 80, 0.15),
    inset 0 0 0 1px rgba(255, 255, 255, 0.15);
  border-color: rgba(180, 130, 80, 0.5);
  background: linear-gradient(135deg, #d44637, #b83227);
}

.pop-enter-active, .pop-leave-active { transition: opacity 0.3s ease, transform 0.3s cubic-bezier(0.22,1,0.36,1); }
.pop-enter-from, .pop-leave-to { opacity: 0; transform: translateY(12px) scale(0.9); }
</style>
