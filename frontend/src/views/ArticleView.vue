<script setup>
import { onMounted, ref, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import markdownit from 'markdown-it'
import { fetchArticle, fetchComments, postComment, deleteArticle } from '../services/api'
import { useAuthStore } from '../stores/auth'
import { Calendar, MessageCircle, Send, Clock, Trash2, Copy, Check } from 'lucide-vue-next'
import ProfileCard from '../components/Sidebar/ProfileCard.vue'
import TOC from '../components/Article/TOC.vue'

const route = useRoute()
const router = useRouter()
const article = ref(null)
const comments = ref([])
const loading = ref(true)
const newComment = ref('')
const authStore = useAuthStore()
const md = markdownit({
  html: true,
  linkify: true,
  typographer: true
})

// Copy Button Logic
const addCopyButtons = () => {
  const preTags = document.querySelectorAll('.prose pre')
  preTags.forEach(pre => {
    // Check if button already exists to prevent duplicates
    if (pre.querySelector('.copy-btn')) return

    const wrapper = document.createElement('div')
    wrapper.className = 'relative group'
    
    // Wrap pre
    pre.parentNode.insertBefore(wrapper, pre)
    wrapper.appendChild(pre)

    // Create button
    const btn = document.createElement('button')
    btn.className = 'copy-btn absolute top-2 right-2 p-1.5 rounded bg-zinc-700/50 text-zinc-300 hover:bg-zinc-600 hover:text-white transition-all opacity-0 group-hover:opacity-100'
    btn.innerHTML = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/></svg>` // Copy Icon
    
    btn.addEventListener('click', async () => {
      const code = pre.querySelector('code')?.innerText || pre.innerText
      try {
        await navigator.clipboard.writeText(code)
        btn.innerHTML = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-green-400"><path d="M20 6 9 17l-5-5"/></svg>` // Check Icon
        setTimeout(() => {
          btn.innerHTML = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/></svg>`
        }, 2000)
      } catch (e) {
        console.error('Copy failed', e)
      }
    })

    wrapper.appendChild(btn)
  })
}

const loadArticle = async () => {
  const id = route.params.id
  if (!id || id === 'undefined') {
    loading.value = false
    return
  }
  
  try {
    const { data } = await fetchArticle(id)
    article.value = {
      ...data,
      html: md.render(data.content || '')
    }
    
    // Fetch comments
    const commentsData = await fetchComments(id)
    comments.value = commentsData.data

    // Setup Code Copy buttons after render
    nextTick(() => {
      addCopyButtons()
    })

  } catch (error) {
    console.error('Failed to load article', error)
  } finally {
    loading.value = false
  }
}

const submitComment = async () => {
  if (!newComment.value.trim()) return
  try {
    await postComment(route.params.id, newComment.value)
    newComment.value = ''
    const commentsData = await fetchComments(route.params.id)
    comments.value = commentsData.data
  } catch (error) {
    alert('发送评论失败')
  }
}

const handleDelete = async () => {
  if (!confirm('确定要删除这篇文章吗？此操作不可逆。')) return
  try {
    await deleteArticle(route.params.id)
    router.push('/')
  } catch (error) {
    alert('删除失败: ' + (error.response?.data?.error || error.message))
  }
}

onMounted(loadArticle)
</script>

<template>
  <div class="max-w-7xl mx-auto px-4 md:px-6 py-12 md:py-28 animate-fade-in">
    <div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
      
      <!-- Main Content (Col-3) -->
      <div class="lg:col-span-3">
        <div v-if="!loading && article" class="fish-card p-8 md:p-12">
          
          <!-- Header -->
          <header class="mb-10 text-center border-b border-gray-100 dark:border-zinc-700 pb-10">
            <div class="flex justify-center mb-6">
              <span class="px-3 py-1 bg-fish-blue/10 text-fish-blue text-xs font-bold rounded-full uppercase tracking-wider">
                 {{ article.category || 'Tech' }}
              </span>
            </div>
            
            <h1 class="text-3xl md:text-5xl font-bold text-zinc-900 dark:text-zinc-50 mb-6 leading-tight">
              {{ article.title }}
            </h1>

            <div class="flex items-center justify-center gap-6 text-sm text-zinc-600 dark:text-zinc-300">
               <span class="flex items-center gap-2"><Calendar size="14" /> {{ article.date }}</span>
               <span class="flex items-center gap-2" v-if="article.readingTime"><Clock size="14" /> {{ article.readingTime }}m</span>
               <button v-if="authStore.isAdmin" @click="handleDelete" class="text-red-600 hover:underline flex items-center gap-1 font-medium">
                 <Trash2 size="14" /> 删除
               </button>
            </div>
          </header>

          <!-- Markdown Content -->
          <article id="article-content" class="prose prose-lg max-w-none dark:prose-invert
             prose-headings:font-bold prose-headings:text-zinc-900 dark:prose-headings:text-zinc-50
             prose-a:text-fish-blue hover:prose-a:text-fish-orange prose-a:no-underline
             prose-img:rounded-xl prose-img:shadow-lg prose-pre:bg-zinc-900 prose-pre:shadow-inner
             mb-12 text-zinc-800 dark:text-zinc-200">
             <div v-html="article.html"></div>
          </article>

          <!-- Comments -->
          <div class="mt-16 pt-10 border-t border-gray-100 dark:border-zinc-700">
             <h3 class="flex items-center gap-2 text-xl font-bold text-zinc-900 dark:text-zinc-50 mb-8">
               <MessageCircle class="text-fish-blue" /> 评论
             </h3>

             <!-- Input -->
             <div v-if="authStore.isAuthenticated" class="mb-10">
               <div class="relative">
                 <textarea 
                   v-model="newComment" 
                   rows="3" 
                   class="w-full bg-gray-50 dark:bg-zinc-900 border border-gray-200 dark:border-zinc-700 rounded-xl p-4 focus:ring-2 focus:ring-fish-blue focus:outline-none transition-all resize-none text-zinc-800 dark:text-zinc-200"
                   placeholder="写下你的评论..."
                 ></textarea>
                 <button 
                   @click="submitComment" 
                   class="absolute bottom-3 right-3 bg-fish-blue text-white p-2 rounded-lg hover:bg-fish-orange transition-colors shadow-lg shadow-fish-blue/30"
                   :disabled="!newComment.trim()"
                 >
                   <Send size="16" />
                 </button>
               </div>
             </div>
             <!-- Guest Comment -->
             <div v-else class="mb-10">
                <div class="bg-blue-50 dark:bg-blue-900/10 p-4 rounded-xl flex items-center justify-between mb-4">
                  <span class="text-sm text-fish-blue font-bold">访客模式</span>
                  <RouterLink to="/login" class="text-xs text-zinc-600 dark:text-zinc-400 hover:underline">我是管理员?</RouterLink>
                </div>
                <div class="relative">
                 <textarea 
                   v-model="newComment" 
                   rows="3" 
                   class="w-full bg-gray-50 dark:bg-zinc-900 border border-gray-200 dark:border-zinc-700 rounded-xl p-4 focus:ring-2 focus:ring-fish-blue focus:outline-none transition-all resize-none text-zinc-800 dark:text-zinc-200"
                   placeholder="既然来了，不妨留个言..."
                 ></textarea>
                 <button 
                   @click="submitComment" 
                   class="absolute bottom-3 right-3 bg-fish-blue text-white p-2 rounded-lg hover:bg-fish-orange transition-colors shadow-lg shadow-fish-blue/30"
                   :disabled="!newComment.trim()"
                 >
                   <Send size="16" />
                 </button>
               </div>
             </div>

             <!-- List -->
             <div class="space-y-6">
                <div v-for="comment in comments" :key="comment.id" class="flex gap-4">
                  <div class="w-10 h-10 rounded-full bg-gray-200 dark:bg-zinc-700 flex items-center justify-center font-bold text-zinc-600 dark:text-zinc-300 shrink-0">
                    {{ (comment.user?.username || 'U')[0].toUpperCase() }}
                  </div>
                  <div class="flex-1 bg-gray-50 dark:bg-zinc-900/50 p-4 rounded-xl rounded-tl-none border border-gray-100 dark:border-zinc-800">
                     <div class="flex justify-between items-start mb-2">
                       <span class="font-bold text-zinc-900 dark:text-zinc-50 text-sm">
                         {{ comment.user?.username === 'Guest' ? '访客' : comment.user?.username }}
                         <span v-if="comment.user?.role === 'admin'" class="ml-1 text-[10px] px-1.5 py-0.5 rounded bg-fish-blue text-white">ADMIN</span>
                       </span>
                       <span class="text-xs text-zinc-500 dark:text-zinc-400">{{ new Date(comment.created_at).toLocaleDateString() }}</span>
                     </div>
                     <p class="text-zinc-700 dark:text-zinc-200 text-sm leading-relaxed">{{ comment.content }}</p>
                  </div>
                </div>
             </div>
          </div>

        </div>
        <div v-else class="fish-card p-12 text-center animate-pulse">加载中...</div>
      </div>

      <!-- Sidebar (Col-1) -->
      <div class="hidden lg:block lg:col-span-1 space-y-8">
         <ProfileCard />
         
         <!-- TOC Component -->
         <TOC content-selector="#article-content" v-if="!loading && article" />
      </div>

    </div>
  </div>
</template>
