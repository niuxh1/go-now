<script setup>
import { ref, onMounted } from 'vue'
import { fetchArticles } from '../services/api'
import HeroBanner from '../components/Home/HeroBanner.vue'
import ArticleCard from '../components/ArticleCard.vue'
import ProfileCard from '../components/Sidebar/ProfileCard.vue'

const articles = ref([])
const loading = ref(true)

onMounted(async () => {
  try {
    const res = await fetchArticles()
    articles.value = res.data
  } catch (e) {
    console.error(e)
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div class="min-h-screen">
    <!-- Hero Section -->
    <HeroBanner />

    <!-- Content Section -->
    <div class="max-w-7xl mx-auto px-4 md:px-6 py-12 md:py-16">
      <div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
        
        <!-- Left: Articles (75%) -->
        <div class="lg:col-span-3 space-y-8">
           <!-- Loading State -->
           <div v-if="loading" class="space-y-8">
              <div v-for="i in 3" :key="i" class="fish-card p-6 animate-pulse flex flex-col">
                <div class="h-6 bg-gray-200 dark:bg-zinc-800 rounded w-3/4 mb-4"></div>
                <div class="h-4 bg-gray-200 dark:bg-zinc-800 rounded w-1/4 mb-4"></div>
                <div class="h-20 bg-gray-200 dark:bg-zinc-800 rounded w-full"></div>
              </div>
           </div>

           <!-- Article List -->
           <template v-else>
             <ArticleCard 
               v-for="article in articles" 
               :key="article.id" 
               :article="article"
               @click="$router.push(`/articles/${article.id}`)"
             />
             
             <!-- Pagination Placeholder -->
             <div class="flex justify-center mt-12">
               <button class="px-6 py-2 bg-white dark:bg-zinc-800 border border-gray-200 dark:border-zinc-700 rounded-full text-zinc-600 hover:text-fish-blue hover:shadow-md transition-all">
                 Load More
               </button>
             </div>
           </template>
        </div>

        <!-- Right: Sidebar (25%) -->
        <div class="hidden lg:block lg:col-span-1 space-y-8">
          <!-- Profile -->
          <ProfileCard />

          <!-- Notice Board -->
          <div class="fish-card p-6">
            <h3 class="flex items-center gap-2 font-bold text-zinc-800 dark:text-white mb-4 border-b border-gray-100 dark:border-zinc-700 pb-2">
              公告
            </h3>
            <div class="space-y-3 text-sm text-zinc-600 dark:text-zinc-400">
               <p class="bg-blue-50 dark:bg-blue-900/10 p-3 rounded-lg border-l-4 border-fish-blue">
                 欢迎访问 GoNow Blog! 这里记录着技术的点滴。
               </p>
               <p>
                 GitHub Repo 已更新，欢迎 Star!
               </p>
            </div>
          </div>

          <!-- Tags -->
          <div class="fish-card p-6">
            <h3 class="font-bold text-zinc-800 dark:text-white mb-4">标签</h3>
            <div class="flex flex-wrap gap-2">
              <span v-for="tag in ['Go', 'Vue', 'Docker', 'K8s', 'Rust']" :key="tag" 
                class="px-3 py-1 bg-gray-100 dark:bg-zinc-700 text-xs rounded-md text-zinc-600 dark:text-zinc-300 hover:bg-fish-blue hover:text-white transition-colors cursor-pointer">
                {{ tag }}
              </span>
            </div>
          </div>

        </div>
      </div>
    </div>
  </div>
</template>
