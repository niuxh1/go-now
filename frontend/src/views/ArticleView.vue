<script setup>
import { onMounted, ref, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import markdownit from 'markdown-it'
import { fetchArticle, fetchComments, postComment, deleteArticle } from '../services/api'
import { useAuthStore } from '../stores/auth'
import { Calendar, MessageCircle, Send, Trash2, ArrowLeft } from 'lucide-vue-next'
import TOC from '../components/Article/TOC.vue'

const route = useRoute()
const router = useRouter()
const article = ref(null)
const comments = ref([])
const loading = ref(true)
const newComment = ref('')
const authStore = useAuthStore()
const md = markdownit({ html: true, linkify: true, typographer: true })

const addCopyButtons = () => {
  document.querySelectorAll('.prose pre').forEach(pre => {
    if (pre.querySelector('.copy-btn')) return
    const wrapper = document.createElement('div')
    wrapper.className = 'relative group'
    pre.parentNode.insertBefore(wrapper, pre)
    wrapper.appendChild(pre)
    const btn = document.createElement('button')
    btn.className = 'copy-btn absolute top-3 right-3 p-2 rounded-lg bg-white/80 border border-sakura/10 text-sumi-400 hover:text-sakura-deep hover:border-sakura/25 transition-all opacity-0 group-hover:opacity-100'
    btn.innerHTML = `<svg xmlns="http://www.w3.org/2000/svg" width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect width="14" height="14" x="8" y="8" rx="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/></svg>`
    btn.addEventListener('click', async () => {
      const code = pre.querySelector('code')?.innerText || pre.innerText
      try {
        await navigator.clipboard.writeText(code)
        btn.innerHTML = `<svg xmlns="http://www.w3.org/2000/svg" width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" class="text-take"><path d="M20 6 9 17l-5-5"/></svg>`
        setTimeout(() => { btn.innerHTML = `<svg xmlns="http://www.w3.org/2000/svg" width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect width="14" height="14" x="8" y="8" rx="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/></svg>` }, 2000)
      } catch (e) {}
    })
    wrapper.appendChild(btn)
  })
}

const loadArticle = async () => {
  const id = route.params.id
  if (!id || id === 'undefined') { loading.value = false; return }
  try {
    const { data } = await fetchArticle(id)
    article.value = { ...data, html: md.render(data.content || '') }
    comments.value = (await fetchComments(id)).data
    nextTick(() => addCopyButtons())
  } catch (e) { console.error(e) }
  finally { loading.value = false }
}

const submitComment = async () => {
  if (!newComment.value.trim()) return
  try {
    await postComment(route.params.id, newComment.value)
    newComment.value = ''
    comments.value = (await fetchComments(route.params.id)).data
  } catch (e) { alert('發送評論失敗') }
}

const handleDelete = async () => {
  if (!confirm('確定要刪除這篇文章嗎？')) return
  try { await deleteArticle(route.params.id); router.push('/') }
  catch (e) { alert('刪除失敗') }
}

onMounted(loadArticle)
</script>

<template>
  <div class="max-w-5xl mx-auto px-4 md:px-6 pt-28 pb-16 animate-fade-in">
    <div class="grid grid-cols-1 lg:grid-cols-4 gap-10">
      <div class="lg:col-span-3">
        <!-- Back button with ornamental style -->
        <button @click="router.back()" class="flex items-center gap-2.5 text-sumi-400 hover:text-sakura-deep text-sm mb-8 group transition-all duration-300 font-kai">
          <span class="inline-flex items-center justify-center w-7 h-7 rounded-full border border-sakura/15 group-hover:border-sakura/40 group-hover:bg-sakura/5 transition-all duration-300">
            <ArrowLeft size="12" class="group-hover:-translate-x-0.5 transition-transform duration-300" />
          </span>
          <span class="relative">
            返回
            <span class="absolute -bottom-1 left-0 w-0 h-px bg-sakura/30 group-hover:w-full transition-all duration-300"></span>
          </span>
        </button>

        <!-- Article loaded -->
        <div v-if="!loading && article" class="washi-card overflow-hidden">
          <!-- Ornamental header section -->
          <header class="relative p-8 md:p-12 pb-6 overflow-hidden">
            <!-- Pattern background for header -->
            <div class="absolute inset-0 opacity-[0.012]" style="background-image: url('data:image/svg+xml,<svg xmlns=%22http://www.w3.org/2000/svg%22 width=%2240%22 height=%2240%22><path d=%22M20 0L40 20L20 40L0 20Z%22 fill=%22none%22 stroke=%22%23654321%22 stroke-width=%220.3%22/></svg>');"></div>
            <!-- Warm gradient overlay -->
            <div class="absolute inset-0 bg-gradient-to-br from-sakura/[0.02] via-transparent to-mizu/[0.02]"></div>

            <!-- Large watermark kanji -->
            <div class="absolute top-2 right-6 font-brush text-[10rem] text-sumi-800/[0.018] leading-none select-none pointer-events-none">文</div>
            <!-- Secondary watermark -->
            <div class="absolute bottom-0 left-6 font-brush text-[6rem] text-sumi-800/[0.01] leading-none select-none pointer-events-none">章</div>

            <div class="relative z-10">
              <!-- Category stamp and metadata -->
              <div class="flex items-center gap-3 mb-6">
                <div class="relative">
                  <div class="shu-stamp px-2.5 py-0.5 text-xs">{{ article.category || '技術' }}</div>
                  <!-- Stamp shadow -->
                  <div class="absolute inset-0 shu-stamp px-2.5 py-0.5 text-xs opacity-20 translate-x-px translate-y-px -z-10">{{ article.category || '技術' }}</div>
                </div>
                <span class="flex items-center gap-1.5 text-sumi-300 text-xs font-song">
                  <Calendar size="11" class="text-sumi-300/70" />
                  {{ article.date }}
                </span>
                <button v-if="authStore.isAdmin" @click="handleDelete" class="ml-auto flex items-center gap-1.5 text-xs text-shu/40 hover:text-shu transition-colors duration-300 font-kai group">
                  <span class="inline-flex items-center justify-center w-5 h-5 rounded border border-shu/15 group-hover:border-shu/40 transition-colors">
                    <Trash2 size="10" />
                  </span>
                  刪除
                </button>
              </div>

              <!-- Title with decorative underline -->
              <h1 class="text-3xl md:text-5xl font-kai font-bold text-sumi-900 leading-tight mb-6 tracking-wide">{{ article.title }}</h1>

              <!-- Gradient underline from sakura to mizu -->
              <div class="relative h-1 mb-2">
                <div class="absolute inset-0 h-[2px] bg-gradient-to-r from-sakura/30 via-kin/10 to-mizu/25 rounded-full"></div>
                <div class="absolute left-0 top-0 h-[2px] w-8 bg-gradient-to-r from-sakura/50 to-sakura/20 rounded-full"></div>
              </div>

              <!-- Decorative diamonds under title -->
              <div class="flex items-center gap-2 mt-4">
                <div class="w-12 h-px bg-gradient-to-r from-sakura/15 to-transparent"></div>
                <svg class="w-1.5 h-1.5 text-kin/25" viewBox="0 0 10 10"><rect x="2" y="2" width="6" height="6" transform="rotate(45 5 5)" fill="currentColor"/></svg>
                <svg class="w-1 h-1 text-sakura/20" viewBox="0 0 10 10"><rect x="2" y="2" width="6" height="6" transform="rotate(45 5 5)" fill="currentColor"/></svg>
              </div>
            </div>
          </header>

          <!-- Article content with subtle side accents -->
          <div class="relative">
            <!-- Left decorative border -->
            <div class="absolute left-0 top-8 bottom-8 w-px bg-gradient-to-b from-transparent via-sakura/8 to-transparent"></div>

            <article id="article-content" class="prose prose-lg max-w-none p-8 md:p-12 pt-8 prose-headings:font-kai prose-pre:bg-sumi-900 text-sumi-700 leading-relaxed">
              <div v-html="article.html"></div>
            </article>
          </div>

          <!-- Comment section with ornamental header -->
          <div class="mx-8 md:mx-12 mb-8 md:mb-12 pt-8 border-t border-sakura/10">
            <!-- Ornamental comment header -->
            <div class="flex items-center gap-3 mb-8">
              <div class="flex items-center gap-2">
                <span class="inline-flex items-center justify-center w-7 h-7 rounded-full bg-gradient-to-br from-sakura/10 to-mizu/10">
                  <MessageCircle size="13" class="text-sakura/60" />
                </span>
                <h3 class="text-lg font-kai font-bold text-sumi-800">留言</h3>
              </div>
              <div class="flex-1 flex items-center gap-2">
                <div class="flex-1 h-px bg-gradient-to-r from-sakura/10 to-transparent"></div>
                <svg class="w-2 h-2 text-kin/20" viewBox="0 0 10 10"><rect x="2" y="2" width="6" height="6" transform="rotate(45 5 5)" fill="currentColor"/></svg>
              </div>
              <span class="text-[10px] text-sumi-300 font-song">{{ comments.length }} 則留言</span>
            </div>

            <!-- Comment input with ornamental frame -->
            <div class="mb-8">
              <div v-if="!authStore.isAuthenticated" class="flex items-center justify-between mb-3 px-1">
                <span class="text-xs text-sakura/50 font-kai flex items-center gap-1.5">
                  <svg class="w-3 h-3" viewBox="0 0 20 20"><circle cx="10" cy="10" r="8" fill="none" stroke="currentColor" stroke-width="1"/><circle cx="10" cy="7" r="2.5" fill="currentColor"/><path d="M4.5 16c0-3 2.5-5 5.5-5s5.5 2 5.5 5" fill="none" stroke="currentColor" stroke-width="1"/></svg>
                  訪客模式
                </span>
                <RouterLink to="/login" class="text-xs text-sumi-400 hover:text-sakura-deep transition-colors font-kai">登錄</RouterLink>
              </div>
              <div class="relative group/input">
                <!-- Ornamental corners on input -->
                <div class="absolute -top-1 -left-1 w-3 h-3 border-t border-l border-kin/10 rounded-tl-sm opacity-0 group-focus-within/input:opacity-100 transition-opacity"></div>
                <div class="absolute -top-1 -right-1 w-3 h-3 border-t border-r border-kin/10 rounded-tr-sm opacity-0 group-focus-within/input:opacity-100 transition-opacity"></div>
                <div class="absolute -bottom-1 -left-1 w-3 h-3 border-b border-l border-kin/10 rounded-bl-sm opacity-0 group-focus-within/input:opacity-100 transition-opacity"></div>
                <div class="absolute -bottom-1 -right-1 w-3 h-3 border-b border-r border-kin/10 rounded-br-sm opacity-0 group-focus-within/input:opacity-100 transition-opacity"></div>

                <textarea v-model="newComment" rows="3" class="sumi-input resize-none pr-14"
                  :placeholder="authStore.isAuthenticated ? '落筆留言...' : '路過留痕...'"></textarea>
                <button @click="submitComment"
                  class="absolute bottom-3 right-3 p-2.5 rounded-lg border border-sakura/20 text-sakura-deep hover:bg-sakura hover:text-white transition-all duration-300 disabled:opacity-30 hover:shadow-[0_0_12px_rgba(219,112,147,0.2)]"
                  :disabled="!newComment.trim()">
                  <Send size="13" />
                </button>
              </div>
            </div>

            <!-- Comments list -->
            <div class="space-y-4">
              <div v-for="c in comments" :key="c.id" class="flex gap-3 group/comment">
                <!-- Seal-style avatar -->
                <div class="relative w-10 h-10 shrink-0">
                  <div class="w-10 h-10 rounded-lg border-2 border-shu/15 bg-gradient-to-br from-red-50 to-orange-50/50 flex items-center justify-center text-xs font-kai font-bold text-shu/50 group-hover/comment:border-shu/30 transition-colors duration-300" style="transform: rotate(-2deg);">
                    {{ (c.user?.username || 'U')[0].toUpperCase() }}
                  </div>
                  <!-- Tiny seal dot -->
                  <div class="absolute -bottom-0.5 -right-0.5 w-2 h-2 rounded-full bg-shu/10"></div>
                </div>

                <!-- Comment bubble with warmer colors -->
                <div class="flex-1 bg-gradient-to-br from-white/80 to-orange-50/20 border border-sakura/10 rounded-lg rounded-tl-sm p-4 group-hover/comment:border-sakura/20 transition-colors duration-300">
                  <div class="flex justify-between items-center mb-2">
                    <span class="text-sm font-kai font-medium text-sumi-700">
                      {{ c.user?.username === 'Guest' ? '旅人' : c.user?.username }}
                      <span v-if="c.user?.role === 'admin'" class="ml-1.5 inline-flex items-center px-1.5 py-0.5 rounded border border-shu/20 bg-shu/5 text-shu/60 text-[8px] font-kai">主人</span>
                    </span>
                    <span class="text-[10px] text-sumi-300 font-song">{{ new Date(c.created_at).toLocaleDateString() }}</span>
                  </div>
                  <p class="text-sm text-sumi-600 leading-relaxed font-song">{{ c.content }}</p>
                </div>
              </div>

              <!-- Empty comments state -->
              <div v-if="comments.length === 0" class="text-center py-10">
                <div class="text-sumi-200 text-xs font-kai">尚無留言，成為第一位留言者吧</div>
              </div>
            </div>
          </div>
        </div>

        <!-- Loading state with brush character animation -->
        <div v-else class="washi-card p-20 text-center relative overflow-hidden">
          <!-- Background pattern -->
          <div class="absolute inset-0 opacity-[0.01]" style="background-image: url('data:image/svg+xml,<svg xmlns=%22http://www.w3.org/2000/svg%22 width=%2240%22 height=%2240%22><circle cx=%2220%22 cy=%2220%22 r=%221%22 fill=%22%23654321%22/></svg>');"></div>

          <div class="relative">
            <!-- Animated brush character -->
            <div class="relative inline-block mb-6">
              <div class="font-brush text-6xl text-sumi-100 animate-pulse leading-none">墨</div>
              <svg class="absolute -inset-4 w-[calc(100%+2rem)] h-[calc(100%+2rem)] text-sakura/8 animate-spin" style="animation-duration: 8s;" viewBox="0 0 100 100">
                <circle cx="50" cy="50" r="45" fill="none" stroke="currentColor" stroke-width="0.5" stroke-dasharray="4 8"/>
              </svg>
            </div>
            <div class="flex items-center justify-center gap-2 mb-3">
              <div class="w-8 h-px bg-gradient-to-r from-transparent to-sakura/15"></div>
              <div class="w-1.5 h-1.5 rounded-full bg-sakura/20 animate-pulse"></div>
              <div class="w-8 h-px bg-gradient-to-l from-transparent to-mizu/15"></div>
            </div>
            <p class="text-sumi-300 text-sm font-kai">載入中...</p>
          </div>
        </div>
      </div>

      <!-- TOC sidebar -->
      <div class="hidden lg:block lg:col-span-1 space-y-6">
        <TOC content-selector="#article-content" v-if="!loading && article" />
      </div>
    </div>
  </div>
</template>
