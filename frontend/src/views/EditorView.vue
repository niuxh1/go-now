<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { createArticle } from '../services/api'
import { Save, FileText, ArrowLeft, AlertCircle, Eye } from 'lucide-vue-next'

const router = useRouter()
const title = ref('')
const content = ref('')
const loading = ref(false)
const errorMsg = ref('')

const handlePublish = async () => {
  if (!title.value || !content.value) return
  
  loading.value = true
  errorMsg.value = ''
  try {
    const { data } = await createArticle({ 
      title: title.value, 
      content: content.value,
      date: new Date().toISOString().split('T')[0]
    })
    router.push(`/articles/${data.id}`)
  } catch (error) {
    console.error("Publish Error:", error)
    errorMsg.value = error.response?.data?.error || '发布失败：' + error.message
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="max-w-5xl mx-auto py-12 animate-fade-in">
    <!-- Editor Top Bar -->
    <div class="flex flex-col md:flex-row items-center justify-between gap-6 mb-12">
      <div class="flex items-center gap-6">
        <button @click="router.back()" class="p-3 rounded-2xl bg-white border border-zinc-200 hover:bg-zinc-50 dark:bg-zinc-900/50 dark:border-white/5 text-zinc-600 hover:text-zinc-900 dark:text-zinc-400 dark:hover:text-white dark:hover:bg-zinc-800 transition-all">
          <ArrowLeft size="20" />
        </button>
        <div>
          <h2 class="text-3xl font-bold text-zinc-900 dark:text-white tracking-tight">撰写文章</h2>
          <p class="text-zinc-600 dark:text-zinc-400 text-sm mt-1">捕捉您的灵感与洞察</p>
        </div>
      </div>
      
      <div class="flex items-center gap-4">
        <button 
          @click="handlePublish" 
          class="flex items-center gap-2 bg-zinc-900 text-white dark:bg-white dark:text-black px-8 py-3 rounded-2xl font-bold hover:opacity-90 transition-all shadow-xl shadow-zinc-200 dark:shadow-white/5"
          :disabled="loading"
        >
          <Save size="18" />
          {{ loading ? '发布中...' : '发布文章' }}
        </button>
      </div>
    </div>

    <!-- Error State -->
    <div v-if="errorMsg" class="mb-8 p-4 bg-red-50 dark:bg-red-900/10 border border-red-100 dark:border-red-900/20 rounded-2xl flex items-center gap-3 text-red-600 dark:text-red-400">
      <AlertCircle size="18" /> {{ errorMsg }}
    </div>

    <!-- Editor UI -->
    <div class="grid gap-8">
      <div class="fish-card p-2 rounded-[2rem]">
        <input 
          v-model="title" 
          type="text" 
          placeholder="在此输入引人注目的标题..." 
          class="w-full bg-transparent border-none text-zinc-900 dark:text-white text-4xl font-extrabold p-8 focus:ring-0 placeholder-zinc-400 dark:placeholder-zinc-600"
        />
      </div>

      <div class="fish-card rounded-[2.5rem] overflow-hidden flex flex-col min-h-[600px]">
        <div class="flex items-center justify-between px-8 py-4 border-b border-gray-100 dark:border-zinc-700 bg-gray-50 dark:bg-zinc-900/20">
           <div class="flex items-center gap-2 text-zinc-600 dark:text-zinc-400 text-xs font-bold uppercase tracking-widest">
             <FileText size="14" /> Markdown 编辑器
           </div>
           <div class="flex items-center gap-4 text-zinc-600 dark:text-zinc-400 text-xs">
             支持语法高亮 & 实时预览
           </div>
        </div>
        <textarea 
          v-model="content" 
          class="flex-grow w-full bg-transparent border-none text-zinc-800 dark:text-zinc-200 text-lg leading-relaxed p-8 focus:ring-0 resize-none font-mono placeholder-zinc-400" 
          placeholder="开始您的创作之旅..."
        ></textarea>
      </div>
    </div>

    <div class="mt-8 flex justify-center">
       <div class="px-4 py-2 rounded-full bg-white border border-zinc-200 dark:bg-zinc-900/50 dark:border-white/5 text-xs text-zinc-600 dark:text-zinc-400 flex items-center gap-2">
         <Eye size="14" /> 自动保存已启用 (本地)
       </div>
    </div>
  </div>
</template>