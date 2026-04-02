<script setup>
import { Calendar, ArrowRight, Eye } from 'lucide-vue-next'
defineProps({ article: Object })
</script>

<template>
  <article class="article-card group relative overflow-hidden rounded-xl cursor-pointer"
    @click="$emit('click')">
    <!-- Background kanji watermark -->
    <div class="absolute top-4 right-4 text-[120px] leading-none font-kai text-sumi-800/[0.05] select-none pointer-events-none z-0">
      文
    </div>

    <!-- Decorative corner ornament (top-right) -->
    <svg class="absolute top-0 right-0 w-16 h-16 text-sakura/15 pointer-events-none z-0" viewBox="0 0 64 64">
      <path d="M64 0 L64 64 L0 0 Z" fill="none" stroke="currentColor" stroke-width="0.5"/>
      <path d="M64 0 Q48 16 64 32" fill="none" stroke="currentColor" stroke-width="0.8"/>
      <path d="M64 0 Q40 8 32 0" fill="none" stroke="currentColor" stroke-width="0.8"/>
      <circle cx="56" cy="8" r="1.5" fill="currentColor" opacity="0.4"/>
    </svg>

    <!-- Decorative corner ornament (bottom-left) -->
    <svg class="absolute bottom-0 left-0 w-12 h-12 text-mizu/15 pointer-events-none z-0" viewBox="0 0 48 48">
      <path d="M0 48 Q16 32 0 16" fill="none" stroke="currentColor" stroke-width="0.8"/>
      <circle cx="4" cy="40" r="1.2" fill="currentColor" opacity="0.3"/>
    </svg>

    <div class="relative z-10 flex">
      <!-- Gradient left border -->
      <div class="w-1 min-h-full rounded-full bg-gradient-to-b from-sakura via-sakura/60 to-mizu/40 group-hover:from-sakura-deep group-hover:to-mizu transition-all duration-700 shrink-0"></div>

      <div class="flex flex-col gap-3 p-7 md:p-9 flex-1">
        <div class="flex items-center gap-3 mb-1">
          <!-- Seal/stamp style category -->
          <div class="category-seal">
            <span class="relative z-10">{{ article.category || '技術' }}</span>
          </div>
          <span class="flex items-center gap-1.5 text-sumi-300 text-xs">
            <Calendar size="11" /> {{ article.date }}
          </span>
          <span v-if="article.views" class="flex items-center gap-1 text-sumi-300 text-xs ml-auto">
            <Eye size="11" /> {{ article.views }}
          </span>
        </div>

        <h2 class="article-title text-xl md:text-2xl font-kai font-bold text-sumi-800 leading-snug">
          {{ article.title }}
        </h2>

        <div class="w-8 h-0.5 rounded-full bg-gradient-to-r from-sakura/30 to-transparent group-hover:w-16 transition-all duration-500"></div>

        <p class="text-sumi-500 text-sm leading-relaxed line-clamp-2 font-song">
          {{ article.summary || article.content?.substring(0, 120) + '...' }}
        </p>

        <div class="flex items-center gap-2 text-mizu-deep/50 text-xs font-kai tracking-wider group-hover:text-mizu-deep transition-all duration-500 mt-1">
          <span>閱讀全文</span>
          <ArrowRight size="12" class="group-hover:translate-x-2 transition-transform duration-500" />
        </div>
      </div>
    </div>
  </article>
</template>

<style scoped>
.article-card {
  background: linear-gradient(135deg, rgba(255,255,255,0.97), rgba(252,248,244,0.95));
  border: 1px solid rgba(232, 160, 176, 0.2);
  box-shadow: 0 2px 16px rgba(139, 92, 75, 0.08), 0 0 0 1px rgba(201,169,89,0.06);
  transition: all 0.5s cubic-bezier(0.22, 1, 0.36, 1);
  background-image: url("data:image/svg+xml,%3Csvg width='20' height='20' viewBox='0 0 20 20' xmlns='http://www.w3.org/2000/svg'%3E%3Ccircle cx='1' cy='1' r='0.4' fill='%23d4a574' fill-opacity='0.04'/%3E%3C/svg%3E");
}

.article-card:hover {
  transform: translateY(-4px);
  box-shadow:
    0 8px 30px rgba(232, 160, 176, 0.12),
    inset 0 0 30px rgba(232, 160, 176, 0.03);
  border-color: rgba(232, 160, 176, 0.25);
}

.category-seal {
  position: relative;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 2px 10px;
  background: linear-gradient(135deg, #c0392b, #a93226);
  color: white;
  font-size: 10px;
  font-weight: 600;
  letter-spacing: 0.1em;
  border-radius: 2px;
  box-shadow: 1px 1px 3px rgba(192, 57, 43, 0.25);
  font-family: 'KaiTi', 'STKaiti', serif;
}

.category-seal::before,
.category-seal::after {
  content: '';
  position: absolute;
  border: 1px solid rgba(255, 255, 255, 0.3);
  border-radius: 1px;
}

.category-seal::before {
  inset: 2px;
}

.article-title {
  transition: all 0.5s ease;
}

.group:hover .article-title {
  color: #c0392b;
  text-shadow: 0 0 20px rgba(232, 160, 176, 0.15);
}
</style>
