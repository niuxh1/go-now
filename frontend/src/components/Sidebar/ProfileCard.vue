<script setup>
import { ref, onMounted, computed } from 'vue'
import { fetchSettings, fetchArticles, fetchTags } from '../../services/api'
import { Github, Mail, MapPin } from 'lucide-vue-next'

const profile = ref({ name: 'GoNow', bio: '...', avatar: '/avatar.jpg', location: '', social: {} })
const avatarUrl = computed(() => {
  const p = profile.value.avatar || '/avatar.jpg'
  return (p.startsWith('http') || p.startsWith('/')) ? p : '/' + p
})
const stats = ref({ articles: 0, categories: 0, tags: 0 })

onMounted(async () => {
  try { const { data } = await fetchSettings(); if (data.profile) profile.value = { ...profile.value, ...data.profile } } catch (e) {}
  try {
    const [ar, tr] = await Promise.all([fetchArticles(), fetchTags()])
    stats.value.articles = ar.data.length
    stats.value.categories = new Set(ar.data.map(a => a.category).filter(Boolean)).size
    stats.value.tags = tr.data.length
  } catch (e) {}
})
</script>

<template>
  <div class="profile-card sticky top-24 text-center overflow-hidden">
    <!-- Background calligraphy watermark -->
    <div class="absolute inset-0 flex items-center justify-center pointer-events-none select-none z-0">
      <span class="text-[160px] font-kai text-sumi-800/[0.02] leading-none">雅</span>
    </div>

    <!-- Decorative header band with pattern -->
    <div class="profile-header-band h-20 relative overflow-hidden">
      <div class="absolute inset-0 opacity-20">
        <svg class="w-full h-full" viewBox="0 0 400 80" preserveAspectRatio="none">
          <pattern id="wave" x="0" y="0" width="40" height="40" patternUnits="userSpaceOnUse">
            <path d="M0 20 Q10 10 20 20 Q30 30 40 20" fill="none" stroke="currentColor" stroke-width="0.5" class="text-white"/>
          </pattern>
          <rect width="400" height="80" fill="url(#wave)"/>
        </svg>
      </div>
    </div>

    <!-- Avatar with ornamental double ring -->
    <div class="relative -mt-12 z-10 inline-block mb-4 group">
      <div class="absolute -inset-4 rounded-full border border-dashed border-sakura/20 group-hover:border-sakura/40 transition-all duration-700 animate-spin-slow"></div>
      <div class="absolute -inset-2 rounded-full border-2 border-amber-600/15 group-hover:border-amber-600/30 transition-all duration-700"></div>
      <div class="absolute -inset-2 rounded-full bg-gradient-to-br from-sakura/10 to-mizu/10 opacity-0 group-hover:opacity-100 transition-opacity duration-700 blur-sm"></div>
      <img :src="avatarUrl" class="relative w-20 h-20 rounded-full border-2 border-white object-cover transition-all duration-700 group-hover:scale-105 shadow-lg ring-2 ring-amber-600/10" alt="Avatar" />
    </div>

    <div class="relative z-10 px-6 pb-6">
      <h2 class="text-lg font-kai font-bold text-sumi-800 mb-1">{{ profile.name }}</h2>
      <p v-if="profile.location" class="flex items-center justify-center gap-1 text-sumi-400 text-xs mb-2 font-kai">
        <MapPin size="10" /> {{ profile.location }}
      </p>
      <p class="text-sm text-sumi-500 leading-relaxed whitespace-pre-line mb-5 px-2 font-song">{{ profile.bio }}</p>

      <!-- Stats with decorative borders -->
      <div class="stats-grid grid grid-cols-3 gap-0 mb-5 rounded-lg overflow-hidden">
        <div class="stat-item py-3 px-1">
          <div class="text-lg font-kai font-bold text-sumi-800">{{ stats.articles }}</div>
          <div class="text-[9px] text-sumi-400 tracking-widest font-kai mt-0.5">文章</div>
        </div>
        <div class="stat-item py-3 px-1 border-x border-amber-600/10">
          <div class="text-lg font-kai font-bold text-sumi-800">{{ stats.categories }}</div>
          <div class="text-[9px] text-sumi-400 tracking-widest font-kai mt-0.5">分類</div>
        </div>
        <div class="stat-item py-3 px-1">
          <div class="text-lg font-kai font-bold text-sumi-800">{{ stats.tags }}</div>
          <div class="text-[9px] text-sumi-400 tracking-widest font-kai mt-0.5">標籤</div>
        </div>
      </div>

      <!-- Social links with seal-style hover -->
      <div class="flex justify-center gap-3">
        <a v-if="profile.social?.github" :href="`https://github.com/${profile.social.github}`" target="_blank"
          class="social-seal w-9 h-9 flex items-center justify-center rounded text-sumi-400 hover:text-white hover:bg-gradient-to-br hover:from-rose-700 hover:to-rose-800 transition-all duration-300">
          <Github size="14" />
        </a>
        <a v-if="profile.social?.email" :href="`mailto:${profile.social.email}`"
          class="social-seal w-9 h-9 flex items-center justify-center rounded text-sumi-400 hover:text-white hover:bg-gradient-to-br hover:from-rose-700 hover:to-rose-800 transition-all duration-300">
          <Mail size="14" />
        </a>
      </div>

      <!-- Bottom ornamental line -->
      <div class="flex justify-center mt-5">
        <div class="flex items-center gap-2">
          <div class="w-8 h-px bg-gradient-to-r from-transparent to-sakura/20"></div>
          <div class="w-1.5 h-1.5 rounded-full border border-amber-600/20"></div>
          <div class="w-8 h-px bg-gradient-to-l from-transparent to-sakura/20"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.profile-card {
  position: relative;
  background: linear-gradient(180deg, rgba(255,255,255,0.95), rgba(252,248,244,0.92));
  border: 2px solid rgba(180, 130, 80, 0.1);
  box-shadow: 0 0 0 1px rgba(180, 130, 80, 0.05), 0 4px 20px rgba(139, 92, 75, 0.06);
  border-radius: 12px;
  outline: 1px solid rgba(180, 130, 80, 0.04);
  outline-offset: 3px;
}

.profile-header-band {
  background: linear-gradient(135deg, rgba(232, 160, 176, 0.15), rgba(139, 175, 196, 0.1));
  border-bottom: 1px solid rgba(180, 130, 80, 0.08);
}

.stats-grid {
  background: rgba(252, 248, 244, 0.6);
  border: 1px solid rgba(180, 130, 80, 0.08);
}

.stat-item {
  text-align: center;
  transition: background 0.3s ease;
}

.stat-item:hover {
  background: rgba(232, 160, 176, 0.05);
}

.social-seal {
  border: 1px solid rgba(180, 130, 80, 0.12);
  background: rgba(252, 248, 244, 0.5);
}

.social-seal:hover {
  box-shadow: 0 2px 8px rgba(192, 57, 43, 0.2);
  border-color: transparent;
}

@keyframes spin-slow {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.animate-spin-slow {
  animation: spin-slow 30s linear infinite;
}
</style>
