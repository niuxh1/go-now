<script setup>
import { ref, onMounted } from 'vue'
import { fetchArticles, fetchTags } from '../services/api'
import HeroBanner from '../components/Home/HeroBanner.vue'
import ArticleCard from '../components/ArticleCard.vue'
import ProfileCard from '../components/Sidebar/ProfileCard.vue'

const articles = ref([])
const tags = ref([])
const loading = ref(true)

onMounted(async () => {
  try {
    const [ar, tr] = await Promise.all([fetchArticles(), fetchTags()])
    articles.value = ar.data
    tags.value = tr.data
  } catch (e) { console.error(e) }
  finally { loading.value = false }
})
</script>

<template>
  <div class="min-h-screen">
    <HeroBanner />

    <!-- Ornamental transition section -->
    <div class="relative py-16 overflow-hidden">
      <!-- Subtle background pattern -->
      <div class="absolute inset-0 opacity-[0.015]" style="background-image: url('data:image/svg+xml,<svg xmlns=%22http://www.w3.org/2000/svg%22 width=%2260%22 height=%2260%22><circle cx=%2230%22 cy=%2230%22 r=%221%22 fill=%22%23654321%22/></svg>');"></div>

      <div class="flex flex-col items-center justify-center mb-0">
        <!-- Top decorative line -->
        <div class="flex items-center gap-4 mb-6">
          <div class="w-24 h-px bg-gradient-to-r from-transparent via-sakura/30 to-sakura/5"></div>
          <svg class="w-3 h-3 text-kin/40" viewBox="0 0 10 10"><rect x="2" y="2" width="6" height="6" transform="rotate(45 5 5)" fill="currentColor"/></svg>
          <div class="w-12 h-px bg-gradient-to-r from-sakura/20 via-kin/15 to-mizu/20"></div>
          <svg class="w-3 h-3 text-kin/40" viewBox="0 0 10 10"><rect x="2" y="2" width="6" height="6" transform="rotate(45 5 5)" fill="currentColor"/></svg>
          <div class="w-24 h-px bg-gradient-to-l from-transparent via-mizu/30 to-mizu/5"></div>
        </div>

        <!-- Center emblem and title -->
        <div class="relative flex items-center gap-5">
          <svg class="w-6 h-6 text-kin/25" viewBox="0 0 100 100">
            <circle cx="50" cy="50" r="24" fill="none" stroke="currentColor" stroke-width="1.5"/>
            <circle cx="50" cy="50" r="14" fill="none" stroke="currentColor" stroke-width="0.8"/>
            <circle cx="50" cy="50" r="3" fill="currentColor"/>
            <line x1="50" y1="20" x2="50" y2="26" stroke="currentColor" stroke-width="0.8"/>
            <line x1="50" y1="74" x2="50" y2="80" stroke="currentColor" stroke-width="0.8"/>
            <line x1="20" y1="50" x2="26" y2="50" stroke="currentColor" stroke-width="0.8"/>
            <line x1="74" y1="50" x2="80" y2="50" stroke="currentColor" stroke-width="0.8"/>
          </svg>
          <div class="flex flex-col items-center">
            <span class="text-sumi-500/70 text-xs font-kai tracking-[0.6em] mb-1">最新文章</span>
            <span class="text-sumi-300/40 text-[9px] tracking-[0.3em] font-song">LATEST ARTICLES</span>
          </div>
          <svg class="w-6 h-6 text-kin/25" viewBox="0 0 100 100">
            <circle cx="50" cy="50" r="24" fill="none" stroke="currentColor" stroke-width="1.5"/>
            <circle cx="50" cy="50" r="14" fill="none" stroke="currentColor" stroke-width="0.8"/>
            <circle cx="50" cy="50" r="3" fill="currentColor"/>
            <line x1="50" y1="20" x2="50" y2="26" stroke="currentColor" stroke-width="0.8"/>
            <line x1="50" y1="74" x2="50" y2="80" stroke="currentColor" stroke-width="0.8"/>
            <line x1="20" y1="50" x2="26" y2="50" stroke="currentColor" stroke-width="0.8"/>
            <line x1="74" y1="50" x2="80" y2="50" stroke="currentColor" stroke-width="0.8"/>
          </svg>
        </div>

        <!-- Bottom decorative line -->
        <div class="flex items-center gap-4 mt-6">
          <div class="w-32 h-px bg-gradient-to-r from-transparent via-sakura/20 to-transparent"></div>
          <svg class="w-2 h-2 text-sakura/30" viewBox="0 0 10 10"><rect x="2" y="2" width="6" height="6" transform="rotate(45 5 5)" fill="currentColor"/></svg>
          <div class="w-32 h-px bg-gradient-to-l from-transparent via-mizu/20 to-transparent"></div>
        </div>
      </div>
    </div>

    <!-- Main content area -->
    <div class="relative max-w-5xl mx-auto px-4 md:px-6 pb-16 md:pb-24">
      <!-- Subtle dot pattern behind articles -->
      <div class="absolute inset-0 opacity-[0.008] pointer-events-none" style="background-image: url('data:image/svg+xml,<svg xmlns=%22http://www.w3.org/2000/svg%22 width=%2240%22 height=%2240%22><circle cx=%2220%22 cy=%2220%22 r=%220.8%22 fill=%22%23654321%22/></svg>');"></div>

      <div class="relative grid grid-cols-1 lg:grid-cols-4 gap-8">
        <!-- Article list -->
        <div class="lg:col-span-3 space-y-6">
          <!-- Loading skeletons -->
          <div v-if="loading" class="space-y-6">
            <div v-for="i in 3" :key="i" class="washi-card p-8 animate-pulse relative overflow-hidden">
              <div class="absolute inset-0 opacity-[0.02]" style="background-image: url('data:image/svg+xml,<svg xmlns=%22http://www.w3.org/2000/svg%22 width=%2220%22 height=%2220%22><circle cx=%2210%22 cy=%2210%22 r=%220.5%22 fill=%22%23654321%22/></svg>');"></div>
              <div class="relative">
                <div class="h-3 bg-sumi-100 rounded w-20 mb-5"></div>
                <div class="h-6 bg-sumi-100 rounded w-3/4 mb-4"></div>
                <div class="h-4 bg-sumi-100 rounded w-full mb-2"></div>
                <div class="h-4 bg-sumi-100 rounded w-2/3"></div>
              </div>
            </div>
          </div>

          <!-- Articles -->
          <template v-else>
            <div v-for="(article, index) in articles" :key="article.id"
              class="animate-fade-up" :style="{ animationDelay: `${index * 120}ms` }">
              <ArticleCard :article="article" @click="$router.push(`/articles/${article.id}`)" />
            </div>

            <!-- Empty state -->
            <div v-if="articles.length === 0" class="text-center py-24">
              <div class="relative inline-block mb-6">
                <div class="font-brush text-7xl text-sumi-100 select-none leading-none">空</div>
                <div class="absolute inset-0 flex items-center justify-center">
                  <svg class="w-20 h-20 text-sakura/10" viewBox="0 0 100 100">
                    <circle cx="50" cy="50" r="40" fill="none" stroke="currentColor" stroke-width="0.5"/>
                    <circle cx="50" cy="50" r="30" fill="none" stroke="currentColor" stroke-width="0.3"/>
                  </svg>
                </div>
              </div>
              <div class="flex items-center justify-center gap-3 mb-3">
                <div class="w-10 h-px bg-gradient-to-r from-transparent to-sakura/20"></div>
                <svg class="w-2 h-2 text-kin/30" viewBox="0 0 10 10"><rect x="2" y="2" width="6" height="6" transform="rotate(45 5 5)" fill="currentColor"/></svg>
                <div class="w-10 h-px bg-gradient-to-l from-transparent to-mizu/20"></div>
              </div>
              <p class="text-sumi-300 font-kai italic text-sm">春風拂過空庭，靜待花開</p>
            </div>
          </template>
        </div>

        <!-- Sidebar -->
        <div class="hidden lg:block lg:col-span-1 space-y-6">
          <ProfileCard />

          <!-- Announcement card with ornamental frame -->
          <div class="relative group">
            <!-- Decorative corner marks -->
            <div class="absolute -top-1 -left-1 w-4 h-4 border-t border-l border-kin/15 rounded-tl-sm"></div>
            <div class="absolute -top-1 -right-1 w-4 h-4 border-t border-r border-kin/15 rounded-tr-sm"></div>
            <div class="absolute -bottom-1 -left-1 w-4 h-4 border-b border-l border-kin/15 rounded-bl-sm"></div>
            <div class="absolute -bottom-1 -right-1 w-4 h-4 border-b border-r border-kin/15 rounded-br-sm"></div>

            <div class="washi-card p-5 relative overflow-hidden">
              <!-- Pattern background -->
              <div class="absolute inset-0 opacity-[0.015]" style="background-image: url('data:image/svg+xml,<svg xmlns=%22http://www.w3.org/2000/svg%22 width=%2230%22 height=%2230%22><path d=%22M15 0L30 15L15 30L0 15Z%22 fill=%22none%22 stroke=%22%23B8860B%22 stroke-width=%220.3%22/></svg>');"></div>

              <div class="relative">
                <h3 class="text-xs font-kai text-sumi-400 mb-3 tracking-widest flex items-center gap-2">
                  <span class="inline-flex items-center justify-center w-4 h-4 rounded-sm bg-gradient-to-br from-sakura/10 to-mizu/10 text-mizu/60 text-[8px]">&#9670;</span>
                  公告
                  <span class="flex-1 h-px bg-gradient-to-r from-kin/10 to-transparent"></span>
                </h3>
                <div class="border-l-2 border-sakura/20 pl-3 py-1">
                  <p class="text-sm text-sumi-500 leading-relaxed font-song">
                    歡迎造訪！此處記錄著技術的點滴與生活的感悟。
                  </p>
                </div>
                <!-- Small decorative seal -->
                <div class="flex justify-end mt-3">
                  <div class="w-6 h-6 rounded-sm border border-shu/10 flex items-center justify-center">
                    <span class="text-[7px] text-shu/25 font-kai">記</span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Tags card with pill-shaped gradient borders -->
          <div class="washi-card p-5">
            <h3 class="text-xs font-kai text-sumi-400 mb-4 tracking-widest flex items-center gap-2">
              <span class="inline-flex items-center justify-center w-4 h-4 rounded-sm bg-gradient-to-br from-take/10 to-mizu/10 text-take/60 text-[8px]">&#9670;</span>
              標籤
              <span class="flex-1 h-px bg-gradient-to-r from-take/10 to-transparent"></span>
            </h3>
            <div class="flex flex-wrap gap-2">
              <span v-for="tag in (tags.length ? tags.map(t=>t.name) : ['Go','Vue','Docker','K8s','Rust'])" :key="tag"
                class="relative px-3 py-1.5 text-[11px] font-kai text-sumi-500 rounded-full cursor-pointer transition-all duration-300 hover:text-sakura-deep hover:shadow-[0_0_12px_rgba(219,112,147,0.15)] group/tag bg-gradient-to-r from-white to-white border border-transparent"
                style="background-clip: padding-box;">
                <!-- Gradient border effect via box-shadow -->
                <span class="absolute inset-0 rounded-full border border-sakura/12 group-hover/tag:border-sakura/35 transition-all duration-300"></span>
                <span class="relative">{{ tag }}</span>
              </span>
            </div>
          </div>

          <!-- Site info card -->
          <div class="washi-card p-5">
            <h3 class="text-xs font-kai text-sumi-400 mb-3 tracking-widest flex items-center gap-2">
              <span class="inline-flex items-center justify-center w-4 h-4 rounded-sm bg-gradient-to-br from-kin/10 to-sakura/10 text-kin/60 text-[8px]">&#9670;</span>
              站點
              <span class="flex-1 h-px bg-gradient-to-r from-kin/10 to-transparent"></span>
            </h3>
            <div class="space-y-2.5 text-xs font-song text-sumi-400">
              <div class="flex justify-between items-center">
                <span>文章數</span>
                <span class="text-sumi-600 font-kai">{{ articles.length }}</span>
              </div>
              <div class="h-px bg-sumi-100/50"></div>
              <div class="flex justify-between items-center">
                <span>標籤數</span>
                <span class="text-sumi-600 font-kai">{{ tags.length || 5 }}</span>
              </div>
              <div class="h-px bg-sumi-100/50"></div>
              <div class="flex justify-between items-center">
                <span>運行天數</span>
                <span class="text-sumi-600 font-kai">{{ Math.floor((Date.now() - new Date('2025-01-01').getTime()) / 86400000) }}</span>
              </div>
            </div>
          </div>

          <!-- Brush calligraphy -->
          <div class="text-center py-4">
            <div class="font-brush text-3xl text-sumi-100 tracking-widest select-none mb-2">清風明月</div>
            <div class="flex items-center justify-center gap-2">
              <div class="w-8 h-px bg-gradient-to-r from-transparent to-sakura/15"></div>
              <svg class="w-1.5 h-1.5 text-kin/20" viewBox="0 0 10 10"><rect x="2" y="2" width="6" height="6" transform="rotate(45 5 5)" fill="currentColor"/></svg>
              <div class="w-8 h-px bg-gradient-to-l from-transparent to-mizu/15"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
