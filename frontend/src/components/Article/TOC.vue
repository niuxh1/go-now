<script setup>
import { ref, onMounted, onUnmounted, watch } from 'vue'

const props = defineProps({
  contentSelector: {
    type: String,
    required: true
  }
})

const headings = ref([])
const activeId = ref('')

const parseHeadings = () => {
  const content = document.querySelector(props.contentSelector)
  if (!content) return

  const elements = content.querySelectorAll('h1, h2, h3')
  const items = []

  elements.forEach((el, index) => {
    // Generate ID if missing
    if (!el.id) {
      el.id = `heading-${index}`
    }
    
    items.push({
      id: el.id,
      text: el.innerText,
      level: parseInt(el.tagName.substring(1)),
      el: el // keep reference for observer
    })
  })

  headings.value = items
}

const scrollToHeading = (id) => {
  const el = document.getElementById(id)
  if (el) {
    const top = el.getBoundingClientRect().top + window.scrollY - 100 // offset for fixed header
    window.scrollTo({ top, behavior: 'smooth' })
  }
}

// Scroll Spy
let observer = null
const setupObserver = () => {
  if (observer) observer.disconnect()

  const options = {
    rootMargin: '-100px 0px -60% 0px',
    threshold: 0
  }

  observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        activeId.value = entry.target.id
      }
    })
  }, options)

  headings.value.forEach(h => {
    if (h.el) observer.observe(h.el)
  })
}

onMounted(() => {
  // Give DOM time to render
  setTimeout(() => {
    parseHeadings()
    setupObserver()
  }, 500)
})

onUnmounted(() => {
  if (observer) observer.disconnect()
})

// Re-parse if content likely changed (e.g. route change trigger)
// But ArticleView handles full reload.
</script>

<template>
  <div class="fish-card p-6 sticky top-24 max-h-[calc(100vh-8rem)] overflow-y-auto custom-scrollbar">
    <h3 class="font-bold text-zinc-800 dark:text-white mb-4 flex items-center gap-2">
      <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-list-collapse"><path d="m3 10 2.5-2.5L3 5"/><path d="m3 19 2.5-2.5L3 14"/><path d="M10 6h11"/><path d="M10 12h11"/><path d="M10 18h11"/></svg>
      目录
    </h3>
    
    <div v-if="headings.length === 0" class="text-sm text-zinc-400 italic">
      暂无目录
    </div>

    <ul class="space-y-2 text-sm">
      <li v-for="h in headings" :key="h.id" 
        class="transition-colors duration-200 border-l-2 pl-3 py-1"
        :class="[
          activeId === h.id 
            ? 'border-fish-blue text-fish-blue font-medium' 
            : 'border-transparent text-zinc-500 hover:text-zinc-800 dark:text-zinc-400 dark:hover:text-zinc-200 hover:border-zinc-300',
          h.level === 3 ? 'ml-4' : ''
        ]"
      >
        <a href="#" @click.prevent="scrollToHeading(h.id)" class="block truncate">
          {{ h.text }}
        </a>
      </li>
    </ul>
  </div>
</template>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
</style>
