<script setup>
import { ref, onMounted } from 'vue'
import { ChevronDown } from 'lucide-vue-next'

const text = ref('')
const fullText = "Summer day in the Summer's Town ......"
const isBlinking = ref(true)

const typeText = async () => {
  for (let i = 0; i <= fullText.length; i++) {
    text.value = fullText.slice(0, i)
    await new Promise(r => setTimeout(r, 150))
  }
}

const scrollDown = () => {
  window.scrollTo({
    top: window.innerHeight,
    behavior: 'smooth'
  })
}

onMounted(() => {
  typeText()
})
</script>

<template>
  <div class="relative h-screen w-full overflow-hidden">
    <!-- Background Image -->
    <div class="absolute inset-0 bg-cover bg-center bg-no-repeat transition-transform duration-[10s] hover:scale-105"
         style="background-image: url('https://images.unsplash.com/photo-1451187580459-43490279c0fa?q=80&w=2072&auto=format&fit=crop');">
      <!-- Overlay -->
      <div class="absolute inset-0 bg-black/40"></div>
    </div>

    <!-- Content -->
    <div class="relative z-10 h-full flex flex-col items-center justify-center text-white px-4 text-center">
      <h1 class="text-4xl md:text-6xl font-bold mb-6 animate-title-scale drop-shadow-lg">
        GoNow Blog
      </h1>
      <div class="text-xl md:text-2xl font-light text-white/90 min-h-[2rem] drop-shadow-md">
        {{ text }}<span class="typed-cursor">|</span>
      </div>
    </div>

    <!-- Scroll Down -->
    <div class="absolute bottom-10 left-1/2 -translate-x-1/2 cursor-pointer animate-bounce z-20" @click="scrollDown">
      <ChevronDown size="32" class="text-white drop-shadow" />
    </div>
  </div>
</template>
