<script setup>
import { ref, onMounted, watch } from 'vue'
import { fetchArticles } from '../services/api'
import { History } from 'lucide-vue-next'

const archives = ref([])
const loading = ref(true)
const selectedCategory = ref('All')
const categories = ['All', 'Tech', 'Life', 'Reading']

const fetchArchives = async () => {
  loading.value = true; archives.value = []
  try {
    const cat = selectedCategory.value === 'All' ? '' : selectedCategory.value
    const { data } = await fetchArticles(cat)
    const groups = data.reduce((acc, a) => {
      const y = new Date(a.date).getFullYear() || '未知'
      if (!acc[y]) acc[y] = []
      acc[y].push(a)
      return acc
    }, {})
    archives.value = Object.keys(groups).sort((a, b) => b - a)
      .map(year => ({ year, articles: groups[year].sort((a, b) => new Date(b.date) - new Date(a.date)) }))
  } catch (e) { console.error(e) }
  finally { loading.value = false }
}

const formatDate = (d) => { const x = new Date(d); return `${(x.getMonth()+1).toString().padStart(2,'0')}-${x.getDate().toString().padStart(2,'0')}` }
watch(selectedCategory, fetchArchives)
onMounted(fetchArchives)
</script>

<template>
  <div class="max-w-3xl mx-auto px-6 pt-28 pb-16 animate-fade-in relative">
    <!-- Decorative background pattern -->
    <div class="absolute inset-0 pointer-events-none opacity-[0.015]" style="background-image: url('data:image/svg+xml,%3Csvg width=&quot;60&quot; height=&quot;60&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;%3E%3Cpath d=&quot;M30 0v60M0 30h60&quot; stroke=&quot;%238B4513&quot; stroke-width=&quot;0.5&quot; fill=&quot;none&quot;/%3E%3C/svg%3E'); background-size: 60px 60px;"></div>

    <!-- Rich Header -->
    <div class="relative flex flex-col md:flex-row md:items-end justify-between gap-6 mb-16">
      <div class="relative">
        <!-- Ornamental top decoration -->
        <div class="flex items-center gap-2 mb-4">
          <span class="block w-12 h-[1px] bg-gradient-to-r from-transparent to-sumi-300/40"></span>
          <span class="text-sumi-300/60 text-xs">&#9753;</span>
          <span class="block w-8 h-[1px] bg-gradient-to-l from-transparent to-sumi-300/40"></span>
        </div>
        <div class="flex items-center gap-3 mb-3">
          <div class="w-11 h-11 rounded-full bg-gradient-to-br from-sakura/20 via-shu/10 to-kin/15 flex items-center justify-center shadow-inner border border-sakura/10">
            <History size="18" class="text-shu/70" />
          </div>
          <span class="text-[10px] font-kai text-sumi-300 tracking-[0.3em] uppercase">Archive</span>
        </div>
        <h1 class="text-4xl md:text-5xl font-kai font-bold text-sumi-800 tracking-tight">歸 檔</h1>
        <p class="text-sumi-400 text-sm mt-2 font-kai italic">歲月留痕，文墨成卷</p>
        <!-- Ornamental bottom decoration -->
        <div class="flex items-center gap-2 mt-4">
          <span class="block w-16 h-[1px] bg-gradient-to-r from-shu/30 to-transparent"></span>
          <span class="w-1.5 h-1.5 rounded-full bg-shu/20"></span>
          <span class="block w-8 h-[1px] bg-gradient-to-l from-shu/20 to-transparent"></span>
        </div>
      </div>

      <!-- Category Filter: Pill buttons with gradient borders -->
      <div class="flex p-1.5 rounded-full border border-sumi-200/30 bg-kami/80 backdrop-blur-sm shadow-sm">
        <button v-for="cat in categories" :key="cat" @click="selectedCategory = cat"
          class="px-5 py-2 rounded-full text-xs font-kai tracking-wider transition-all duration-400 relative"
          :class="selectedCategory === cat
            ? 'bg-gradient-to-r from-shu/10 via-sakura/15 to-kin/10 text-sumi-800 shadow-sm border border-shu/15'
            : 'text-sumi-400 hover:text-sumi-600 hover:bg-sumi-50/50'">
          <span class="relative z-10">{{ cat }}</span>
        </button>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex flex-col items-center py-20">
      <div class="font-brush text-4xl text-shu/10 animate-pulse">載</div>
      <div class="mt-3 flex items-center gap-1">
        <span class="w-1 h-1 rounded-full bg-shu/20 animate-bounce" style="animation-delay:0ms"></span>
        <span class="w-1 h-1 rounded-full bg-shu/20 animate-bounce" style="animation-delay:150ms"></span>
        <span class="w-1 h-1 rounded-full bg-shu/20 animate-bounce" style="animation-delay:300ms"></span>
      </div>
    </div>

    <!-- Timeline -->
    <div v-else class="space-y-20">
      <div v-for="group in archives" :key="group.year" class="relative pl-10">
        <!-- Vertical timeline line -->
        <div class="absolute left-0 top-3 bottom-0 w-[2px] bg-gradient-to-b from-shu/20 via-sakura/15 to-transparent"></div>

        <!-- Decorative timeline node -->
        <div class="absolute -left-[7px] top-2 w-4 h-4 rounded-full border-2 border-shu/30 bg-kami flex items-center justify-center shadow-sm">
          <div class="w-1.5 h-1.5 rounded-full bg-shu/40"></div>
        </div>

        <!-- Year Header -->
        <div class="mb-8 relative">
          <div class="flex items-center gap-4">
            <h2 class="text-4xl font-kai font-bold text-sumi-800 tracking-wide">{{ group.year }}</h2>
            <div class="flex items-center gap-2 flex-1">
              <span class="block h-[1px] w-12 bg-gradient-to-r from-shu/25 to-transparent"></span>
              <span class="text-shu/25 text-[10px]">&#9670;</span>
              <span class="block h-[1px] flex-1 bg-gradient-to-r from-shu/15 to-transparent"></span>
            </div>
          </div>
          <span class="text-[10px] font-kai text-sumi-300 tracking-widest ml-1">{{ group.articles.length }} 篇</span>
        </div>

        <!-- Article Items -->
        <div class="space-y-3">
          <RouterLink v-for="article in group.articles" :key="article.id" :to="`/articles/${article.id}`"
            class="group relative flex items-center justify-between p-5 rounded-lg bg-kami/60 border border-sumi-100/40 hover:border-shu/15 hover:bg-gradient-to-r hover:from-shu/[0.02] hover:to-kin/[0.02] transition-all duration-400 hover:shadow-md hover:shadow-shu/5">
            <!-- Ornamental left accent -->
            <div class="absolute left-0 top-1/2 -translate-y-1/2 w-[3px] h-0 group-hover:h-8 bg-gradient-to-b from-shu/40 to-sakura/30 rounded-full transition-all duration-400"></div>
            <div class="flex items-center gap-4 min-w-0 pl-2">
              <span class="text-sumi-300 text-xs font-kai group-hover:text-shu/60 transition-colors w-14 shrink-0">{{ formatDate(article.date) }}</span>
              <span class="text-sumi-700 font-kai font-medium group-hover:text-shu/80 transition-colors duration-300 truncate">{{ article.title }}</span>
            </div>
            <span v-if="article.category" class="text-[9px] font-kai shrink-0 ml-4 px-3 py-1 rounded-full border border-shu/10 text-shu/50 group-hover:border-shu/25 group-hover:text-shu/70 group-hover:bg-shu/[0.04] transition-all duration-300">{{ article.category }}</span>
          </RouterLink>
        </div>
      </div>

      <!-- Empty State -->
      <div v-if="archives.length === 0" class="text-center py-24">
        <div class="font-brush text-5xl text-sumi-100 mb-4">空</div>
        <p class="text-sumi-300 font-kai italic text-sm">春雨初霽，書卷待啟</p>
        <div class="flex items-center justify-center gap-2 mt-4">
          <span class="block w-8 h-[1px] bg-sumi-200/50"></span>
          <span class="text-sumi-200/60 text-xs">&#9753;</span>
          <span class="block w-8 h-[1px] bg-sumi-200/50"></span>
        </div>
      </div>
    </div>
  </div>
</template>
