<script setup>
import { ref, onMounted, watch } from 'vue'
import { fetchArticles } from '../services/api'
import { Calendar, History } from 'lucide-vue-next'

const archives = ref([])
const loading = ref(true)
const selectedCategory = ref('All')
const categories = ['All', 'Tech', 'Life', 'Reading']

const fetchArchives = async () => {
  loading.value = true
  archives.value = [] // Clear previous results immediately
  try {
    const cat = selectedCategory.value === 'All' ? '' : selectedCategory.value
    console.log("Fetching archives for category:", cat)
    const { data } = await fetchArticles(cat)
    console.log("Fetched articles count:", data.length)
    
    // Group by year
    const groups = data.reduce((acc, article) => {
      const year = new Date(article.date).getFullYear() || '未知'
      if (!acc[year]) acc[year] = []
      acc[year].push(article)
      return acc
    }, {})
    
    // Sort years descending
    archives.value = Object.keys(groups)
      .sort((a, b) => b - a)
      .map(year => ({
        year,
        articles: groups[year].sort((a, b) => new Date(b.date) - new Date(a.date))
      }))
      
  } catch (error) {
    console.error(error)
  } finally {
    loading.value = false
  }
}

const formatDate = (dateStr) => {
  const date = new Date(dateStr)
  return `${(date.getMonth() + 1).toString().padStart(2, '0')}-${date.getDate().toString().padStart(2, '0')}`
}

watch(selectedCategory, fetchArchives)

onMounted(fetchArchives)
</script>

<template>
  <div class="max-w-3xl mx-auto py-20 animate-fade-in">
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6 mb-16">
      <div class="flex items-center gap-4">
        <div class="w-12 h-12 rounded-2xl bg-indigo-500/10 flex items-center justify-center text-indigo-400">
          <History size="24" />
        </div>
        <div>
          <h1 class="text-4xl font-bold text-zinc-900 dark:text-zinc-50 tracking-tight">归档</h1>
          <p class="text-zinc-600 dark:text-zinc-400 text-sm mt-1">记录成长的足迹与思绪的流转</p>
        </div>
      </div>
      
      <!-- Category Tabs -->
      <div class="flex p-1 bg-zinc-900 rounded-xl border border-white/5">
        <button 
          v-for="cat in categories" 
          :key="cat"
          @click="selectedCategory = cat"
          class="px-4 py-2 rounded-lg text-sm font-medium transition-all duration-300"
          :class="selectedCategory === cat ? 'bg-indigo-600 text-white shadow-lg' : 'text-zinc-400 hover:text-zinc-200'"
        >
          {{ cat }}
        </button>
      </div>
    </div>
    
    <div v-if="loading" class="flex justify-center py-20 text-zinc-600 dark:text-zinc-400 animate-pulse">
      正在载入回忆...
    </div>
    
    <div v-else class="space-y-16">
      <div v-for="group in archives" :key="group.year" class="relative pl-8 border-l border-zinc-800">
        <div class="absolute -left-2 top-0 w-4 h-4 rounded-full bg-zinc-900 border-2 border-indigo-500"></div>
        
        <h2 class="text-3xl font-bold text-zinc-900 dark:text-zinc-50 mb-8">{{ group.year }}</h2>
        
        <div class="space-y-4">
          <RouterLink 
            v-for="article in group.articles" 
            :key="article.id" 
            :to="`/articles/${article.id}`"
            class="group flex items-center justify-between p-4 rounded-2xl bg-zinc-900/40 border border-white/5 hover:bg-zinc-800/60 hover:border-indigo-500/30 transition-all duration-300"
          >
            <div class="flex items-center gap-4">
              <span class="text-zinc-500 dark:text-zinc-400 font-mono text-sm group-hover:text-indigo-400 w-12">{{ formatDate(article.date) }}</span>
              <span class="text-zinc-800 dark:text-zinc-200 font-medium group-hover:text-fish-blue transition-colors">{{ article.title }}</span>
            </div>
            <div class="flex items-center gap-3">
               <span v-if="article.category" class="text-xs px-2 py-0.5 rounded bg-white/5 text-zinc-500 dark:text-zinc-400 border border-white/5 uppercase tracking-wider group-hover:border-indigo-500/30 group-hover:text-indigo-400 transition-colors">
                 {{ article.category }}
               </span>
               <Calendar size="16" class="text-zinc-600 group-hover:text-indigo-400 opacity-0 group-hover:opacity-100 transition-all" />
            </div>
          </RouterLink>
        </div>
      </div>
      
      <div v-if="archives.length === 0" class="text-center py-20 text-zinc-600 dark:text-zinc-400 italic">
        暂无文章归档
      </div>
    </div>
  </div>
</template>