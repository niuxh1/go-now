<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

const props = defineProps({ contentSelector: { type: String, required: true } })
const headings = ref([])
const activeId = ref('')

const parseHeadings = () => {
  const content = document.querySelector(props.contentSelector)
  if (!content) return
  const elements = content.querySelectorAll('h1, h2, h3')
  const items = []
  elements.forEach((el, index) => {
    if (!el.id) el.id = `heading-${index}`
    items.push({ id: el.id, text: el.innerText, level: parseInt(el.tagName.substring(1)), el })
  })
  headings.value = items
}

const scrollToHeading = (id) => {
  const el = document.getElementById(id)
  if (el) window.scrollTo({ top: el.getBoundingClientRect().top + window.scrollY - 100, behavior: 'smooth' })
}

let observer = null
const setupObserver = () => {
  if (observer) observer.disconnect()
  observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => { if (entry.isIntersecting) activeId.value = entry.target.id })
  }, { rootMargin: '-100px 0px -60% 0px', threshold: 0 })
  headings.value.forEach(h => { if (h.el) observer.observe(h.el) })
}

onMounted(() => setTimeout(() => { parseHeadings(); setupObserver() }, 500))
onUnmounted(() => { if (observer) observer.disconnect() })
</script>

<template>
  <div class="toc-card sticky top-24 max-h-[calc(100vh-8rem)] overflow-y-auto">
    <!-- Decorative top border -->
    <div class="h-1 bg-gradient-to-r from-sakura/30 via-amber-600/20 to-mizu/30 rounded-t-xl"></div>

    <div class="p-5">
      <!-- Ornamental header -->
      <h3 class="flex items-center gap-2.5 text-xs font-kai text-sumi-500 mb-5 tracking-widest">
        <span class="toc-brush-char text-base font-bold text-sakura-deep/70">目</span>
        <span class="text-sumi-300">|</span>
        <span>錄</span>
        <div class="flex-1 h-px bg-gradient-to-r from-sumi-200/30 to-transparent ml-2"></div>
      </h3>

      <div v-if="headings.length === 0" class="text-xs text-sumi-300 italic font-kai">暫無目錄</div>

      <ul class="space-y-0.5 text-[13px]">
        <li v-for="h in headings" :key="h.id"
          class="toc-item transition-all duration-300 pl-4 py-1.5 rounded-r relative"
          :class="[
            activeId === h.id
              ? 'is-active text-sakura-deep bg-sakura/[0.06] font-medium'
              : 'text-sumi-400 hover:text-sumi-600 hover:bg-mizu/[0.03]',
            h.level === 3 ? 'ml-4' : ''
          ]">
          <!-- Ink-drop indicator for active state -->
          <span v-if="activeId === h.id" class="ink-drop"></span>
          <span v-else class="inactive-border"></span>
          <a href="#" @click.prevent="scrollToHeading(h.id)" class="block truncate font-kai">{{ h.text }}</a>
        </li>
      </ul>
    </div>

    <!-- Decorative bottom border -->
    <div class="h-px bg-gradient-to-r from-transparent via-amber-600/10 to-transparent mx-4 mb-3"></div>
    <div class="flex justify-center pb-3">
      <svg class="w-6 h-6 text-sakura/20" viewBox="0 0 24 24">
        <circle cx="12" cy="12" r="3" fill="none" stroke="currentColor" stroke-width="0.5"/>
        <circle cx="12" cy="12" r="1" fill="currentColor"/>
      </svg>
    </div>
  </div>
</template>

<style scoped>
.toc-card {
  background: linear-gradient(180deg, rgba(252,248,244,0.95), rgba(255,255,255,0.9));
  border: 1px solid rgba(180, 130, 80, 0.08);
  border-radius: 12px;
  box-shadow: 0 2px 16px rgba(139, 92, 75, 0.04);
  background-image: url("data:image/svg+xml,%3Csvg width='20' height='20' viewBox='0 0 20 20' xmlns='http://www.w3.org/2000/svg'%3E%3Ccircle cx='1' cy='1' r='0.3' fill='%23d4a574' fill-opacity='0.03'/%3E%3C/svg%3E");
}

.toc-brush-char {
  text-shadow: 1px 1px 2px rgba(192, 57, 43, 0.1);
}

.toc-item {
  position: relative;
}

.ink-drop {
  position: absolute;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(192, 57, 43, 0.7), rgba(192, 57, 43, 0.2));
  box-shadow: 0 0 6px rgba(192, 57, 43, 0.2);
  animation: ink-appear 0.4s ease-out;
}

.ink-drop::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 6px;
  transform: translateY(-50%);
  width: 12px;
  height: 1.5px;
  background: linear-gradient(to right, rgba(192, 57, 43, 0.3), transparent);
  border-radius: 1px;
}

.inactive-border {
  position: absolute;
  left: 2px;
  top: 50%;
  transform: translateY(-50%);
  width: 2px;
  height: 0;
  background: rgba(139, 175, 196, 0.2);
  border-radius: 1px;
  transition: height 0.3s ease;
}

.toc-item:hover .inactive-border {
  height: 60%;
}

@keyframes ink-appear {
  0% { transform: translateY(-50%) scale(0); opacity: 0; }
  60% { transform: translateY(-50%) scale(1.3); }
  100% { transform: translateY(-50%) scale(1); opacity: 1; }
}

::-webkit-scrollbar { width: 3px; }
::-webkit-scrollbar-thumb { background: rgba(232, 160, 176, 0.2); border-radius: 3px; }
</style>
