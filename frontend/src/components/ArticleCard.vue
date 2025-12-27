<script setup>
import { Calendar, Folder } from 'lucide-vue-next'

defineProps({
  article: Object
})
</script>

<template>
  <div 
    class="fish-card flex flex-col overflow-hidden h-auto group p-6 md:p-8 hover:shadow-lg transition-all duration-300"
    @click="$emit('click')"
  >
    <!-- Content -->
    <div class="flex flex-col justify-center text-left w-full">
      <h2 class="text-xl md:text-2xl font-bold text-zinc-900 dark:text-zinc-50 group-hover:text-fish-blue transition-colors cursor-pointer">
        {{ article.title }}
      </h2>
      
      <div class="flex items-center gap-4 text-xs text-zinc-600 dark:text-zinc-300 mb-4">
        <span class="flex items-center gap-1">
          <Calendar size="14" /> {{ article.date }}
        </span>
        <span class="flex items-center gap-1 text-fish-blue bg-blue-50 dark:bg-blue-900/20 px-2 py-0.5 rounded">
          <Folder size="14" /> {{ article.category || '默认分类' }}
        </span>
      </div>

      <p class="text-zinc-600 dark:text-zinc-300 text-sm line-clamp-3 leading-relaxed mb-4">
        {{ article.summary || article.content?.substring(0, 100) + '...' }}
      </p>

      <div class="mt-auto">
        <span class="inline-block text-fish-blue font-bold text-sm group-hover:translate-x-2 transition-transform cursor-pointer">
          阅读全文 →
        </span>
      </div>
    </div>
  </div>
</template>