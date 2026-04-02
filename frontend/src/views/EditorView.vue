<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { createArticle } from '../services/api'
import { Save, FileText, ArrowLeft, AlertCircle } from 'lucide-vue-next'

const router = useRouter()
const title = ref(''); const content = ref(''); const loading = ref(false); const errorMsg = ref('')

const handlePublish = async () => {
  if (!title.value || !content.value) return
  loading.value = true; errorMsg.value = ''
  try {
    const { data } = await createArticle({ title: title.value, content: content.value, date: new Date().toISOString().split('T')[0] })
    router.push(`/articles/${data.id}`)
  } catch (e) { errorMsg.value = e.response?.data?.error || '發佈失敗' }
  finally { loading.value = false }
}
</script>

<template>
  <div class="max-w-4xl mx-auto px-6 pt-28 pb-16 animate-fade-in relative">
    <!-- Subtle background pattern -->
    <div class="absolute inset-0 pointer-events-none opacity-[0.012]" style="background-image: url('data:image/svg+xml,%3Csvg width=&quot;60&quot; height=&quot;60&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;%3E%3Cpath d=&quot;M30 0v60M0 30h60&quot; stroke=&quot;%238B4513&quot; stroke-width=&quot;0.3&quot; fill=&quot;none&quot;/%3E%3C/svg%3E'); background-size: 60px 60px;"></div>

    <!-- Header with ornamental decoration -->
    <div class="relative flex flex-col md:flex-row items-center justify-between gap-6 mb-12">
      <div class="flex items-center gap-5">
        <button @click="router.back()" class="w-10 h-10 rounded-full border border-shu/10 text-sumi-400 hover:text-shu/70 hover:border-shu/25 hover:bg-shu/[0.03] flex items-center justify-center transition-all duration-300">
          <ArrowLeft size="18" />
        </button>
        <div>
          <div class="flex items-center gap-2 mb-1">
            <span class="block w-8 h-[1px] bg-gradient-to-r from-shu/25 to-transparent"></span>
            <span class="text-shu/20 text-[8px]">&#9670;</span>
          </div>
          <h2 class="text-3xl font-kai font-bold text-sumi-800">撰 文</h2>
          <p class="text-sumi-400 text-sm font-kai mt-1">捕捉靈感，落筆成章</p>
          <div class="flex items-center gap-2 mt-2">
            <span class="block w-12 h-[1px] bg-gradient-to-r from-shu/20 to-transparent"></span>
            <span class="w-1 h-1 rounded-full bg-shu/15"></span>
          </div>
        </div>
      </div>

      <!-- Red seal/stamp style publish button -->
      <button @click="handlePublish"
        class="relative flex items-center gap-2 px-6 py-3 rounded-lg font-kai font-medium tracking-wider transition-all duration-300 bg-gradient-to-r from-shu/85 via-shu to-shu/85 text-white shadow-lg shadow-shu/20 hover:shadow-xl hover:shadow-shu/30 group overflow-hidden"
        :disabled="loading">
        <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/10 to-transparent -translate-x-full group-hover:translate-x-full transition-transform duration-1000"></div>
        <Save size="16" class="relative z-10" />
        <span class="relative z-10">{{ loading ? '發佈中...' : '發佈文章' }}</span>
      </button>
    </div>

    <div v-if="errorMsg" class="mb-8 p-4 bg-shu/8 border border-shu/15 rounded-lg flex items-center gap-3 text-shu text-sm font-kai">
      <AlertCircle size="16" /> {{ errorMsg }}
    </div>

    <div class="space-y-6 relative z-10">
      <!-- Title input: calligraphy-style with decorative bottom border -->
      <div class="relative bg-kami/80 border border-sumi-100/40 rounded-xl overflow-hidden">
        <input v-model="title" type="text" placeholder="題目..."
          class="w-full bg-transparent border-none text-sumi-800 text-3xl md:text-4xl font-kai font-bold p-6 focus:ring-0 placeholder-sumi-200" />
        <div class="absolute bottom-0 left-6 right-6 h-[2px] bg-gradient-to-r from-transparent via-shu/15 to-transparent"></div>
      </div>

      <!-- Markdown editor: Chinese frame card -->
      <div class="relative p-[2px] rounded-xl bg-gradient-to-br from-shu/10 via-sumi-100/20 to-kin/10">
        <div class="relative rounded-[10px] bg-kami overflow-hidden flex flex-col min-h-[500px]">
          <!-- Ornamental corners -->
          <div class="absolute top-1.5 left-1.5 w-4 h-4 border-t border-l border-shu/10 rounded-tl-sm pointer-events-none z-20"></div>
          <div class="absolute top-1.5 right-1.5 w-4 h-4 border-t border-r border-shu/10 rounded-tr-sm pointer-events-none z-20"></div>
          <div class="absolute bottom-1.5 left-1.5 w-4 h-4 border-b border-l border-shu/10 rounded-bl-sm pointer-events-none z-20"></div>
          <div class="absolute bottom-1.5 right-1.5 w-4 h-4 border-b border-r border-shu/10 rounded-br-sm pointer-events-none z-20"></div>

          <!-- Ornamental toolbar header -->
          <div class="flex items-center justify-between px-6 py-3 border-b border-shu/8 bg-gradient-to-r from-shu/[0.02] via-transparent to-kin/[0.02]">
            <div class="flex items-center gap-3 text-sumi-400 text-xs font-kai tracking-widest">
              <FileText size="13" />
              <span>Markdown 編輯器</span>
            </div>
            <div class="flex items-center gap-2">
              <span class="block w-6 h-[1px] bg-shu/15"></span>
              <span class="text-shu/20 text-[8px]">&#9753;</span>
              <span class="block w-6 h-[1px] bg-shu/15"></span>
            </div>
          </div>

          <textarea v-model="content" class="flex-grow w-full bg-transparent border-none text-sumi-700 text-base leading-relaxed p-6 focus:ring-0 resize-none font-mono placeholder-sumi-200"
            placeholder="開始書寫..."></textarea>
        </div>
      </div>
    </div>
  </div>
</template>
