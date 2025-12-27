<script setup>
import { ref, onMounted, computed } from 'vue'
import { fetchSettings, fetchArticles, fetchTags } from '../../services/api'
import { Github, Mail } from 'lucide-vue-next'

const profile = ref({
  name: 'GoNow',
  bio: 'Loading...',
  avatar: '/avatar.jpg',
  social: {}
})

// Ensure avatar path is always absolute
const avatarUrl = computed(() => {
  const path = profile.value.avatar || '/avatar.jpg'
  if (path.startsWith('http') || path.startsWith('/')) {
    return path
  }
  return '/' + path
})

const stats = ref({
  articles: 0,
  categories: 0,
  tags: 0
})

onMounted(async () => {
  // Fetch Profile
  try {
    const { data } = await fetchSettings()
    if (data.profile) {
      profile.value = { ...profile.value, ...data.profile }
    }
  } catch (e) {
    console.error("Failed to load profile", e)
  }

  // Fetch Stats
  try {
    const [articlesRes, tagsRes] = await Promise.all([
      fetchArticles(),
      fetchTags()
    ])
    
    const articles = articlesRes.data
    stats.value.articles = articles.length
    
    // Count unique categories
    const categories = new Set(articles.map(a => a.category).filter(Boolean))
    stats.value.categories = categories.size
    
    stats.value.tags = tagsRes.data.length
  } catch (e) {
    console.error("Failed to load stats", e)
  }
})
</script>

<template>
  <div class="fish-card p-6 text-center sticky top-24">
    <div class="relative inline-block mb-4 group">
      <img :src="avatarUrl" 
           class="w-28 h-28 rounded-full border-2 border-zinc-100 shadow-lg object-cover transition-transform duration-500 group-hover:rotate-[360deg]" 
           alt="Avatar" />
    </div>
    <h2 class="text-xl font-bold text-zinc-900 dark:text-zinc-50 mb-2">{{ profile.name }}</h2>
    <p class="text-sm text-zinc-600 dark:text-zinc-300 mb-6 px-4 leading-relaxed whitespace-pre-line">{{ profile.bio }}</p>

    <!-- Stats -->
    <div class="flex justify-center gap-6 mb-6 text-sm">
      <div class="flex flex-col">
        <span class="font-bold text-lg text-zinc-900 dark:text-zinc-50">{{ stats.articles }}</span>
        <span class="text-zinc-600 dark:text-zinc-400 text-xs font-medium">文章</span>
      </div>
      <div class="flex flex-col">
        <span class="font-bold text-lg text-zinc-900 dark:text-zinc-50">{{ stats.categories }}</span>
        <span class="text-zinc-600 dark:text-zinc-400 text-xs font-medium">分类</span>
      </div>
      <div class="flex flex-col">
        <span class="font-bold text-lg text-zinc-900 dark:text-zinc-50">{{ stats.tags }}</span>
        <span class="text-zinc-600 dark:text-zinc-400 text-xs font-medium">标签</span>
      </div>
    </div>

    <!-- Socials -->
    <div class="flex justify-center gap-4">
       <a v-if="profile.social.github" :href="`https://github.com/${profile.social.github}`" target="_blank" class="text-zinc-400 hover:text-fish-blue transition-colors"><Github size="20" /></a>
       <a v-if="profile.social.email" :href="`mailto:${profile.social.email}`" class="text-zinc-400 hover:text-fish-blue transition-colors"><Mail size="20" /></a>
    </div>
  </div>
</template>
