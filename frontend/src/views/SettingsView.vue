<script setup>
import { ref, onMounted, reactive } from 'vue'
import { fetchSettings, updateSettings } from '../services/api'
import { Save, AlertCircle, CheckCircle, User, Link as LinkIcon } from 'lucide-vue-next'

const loading = ref(true); const saving = ref(false)
const message = ref({ type: '', text: '' })
const profile = reactive({ name: '', bio: '', avatar: '', location: '', social: { github: '', email: '' } })

onMounted(async () => {
  try { const { data } = await fetchSettings(); if (data.profile) Object.assign(profile, data.profile) }
  catch (e) { message.value = { type: 'error', text: '無法加載設置' } }
  finally { loading.value = false }
})

const handleSave = async () => {
  saving.value = true; message.value = { type: '', text: '' }
  try { await updateSettings({ profile }); message.value = { type: 'success', text: '設置已保存' }; setTimeout(() => message.value = { type: '', text: '' }, 3000) }
  catch (e) { message.value = { type: 'error', text: '保存失敗' } }
  finally { saving.value = false }
}
</script>

<template>
  <div class="max-w-4xl mx-auto px-6 pt-28 pb-16 animate-fade-in relative">
    <!-- Subtle background pattern -->
    <div class="absolute inset-0 pointer-events-none opacity-[0.012]" style="background-image: url('data:image/svg+xml,%3Csvg width=&quot;40&quot; height=&quot;40&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;%3E%3Ccircle cx=&quot;20&quot; cy=&quot;20&quot; r=&quot;0.8&quot; fill=&quot;%238B4513&quot;/%3E%3C/svg%3E'); background-size: 40px 40px;"></div>

    <!-- Header -->
    <div class="relative flex items-center justify-between mb-12">
      <div>
        <div class="flex items-center gap-2 mb-2">
          <span class="block w-10 h-[1px] bg-gradient-to-r from-shu/25 to-transparent"></span>
          <span class="text-shu/20 text-[8px]">&#9670;</span>
        </div>
        <h2 class="text-3xl font-kai font-bold text-sumi-800">站點設置</h2>
        <p class="text-sumi-400 text-sm font-kai mt-1">管理個人資料與站點配置</p>
        <div class="flex items-center gap-2 mt-3">
          <span class="block w-16 h-[1px] bg-gradient-to-r from-shu/20 to-transparent"></span>
          <span class="w-1 h-1 rounded-full bg-shu/15"></span>
        </div>
      </div>

      <!-- Seal/stamp style save button -->
      <button @click="handleSave"
        class="relative flex items-center gap-2 px-6 py-3 rounded-lg font-kai font-medium tracking-wider transition-all duration-300 bg-gradient-to-r from-shu/85 via-shu to-shu/85 text-white shadow-lg shadow-shu/20 hover:shadow-xl hover:shadow-shu/30 group overflow-hidden"
        :disabled="saving || loading">
        <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/10 to-transparent -translate-x-full group-hover:translate-x-full transition-transform duration-1000"></div>
        <Save size="16" class="relative z-10" />
        <span class="relative z-10">{{ saving ? '保存中...' : '保存更改' }}</span>
      </button>
    </div>

    <!-- Message -->
    <div v-if="message.text" :class="['mb-8 p-4 rounded-lg flex items-center gap-3 text-sm font-kai',
      message.type === 'error' ? 'bg-shu/8 border border-shu/15 text-shu' : 'bg-take/8 border border-take/15 text-take-deep']">
      <AlertCircle v-if="message.type === 'error'" size="16" /><CheckCircle v-else size="16" />
      {{ message.text }}
    </div>

    <div v-if="loading" class="text-center py-12">
      <div class="font-brush text-3xl text-shu/10 animate-pulse">載</div>
      <p class="text-sumi-300 font-kai text-sm mt-2">載入中...</p>
    </div>

    <div v-else class="space-y-8 relative z-10">
      <!-- Basic Info Section: Chinese frame double-border -->
      <div class="relative p-[2px] rounded-xl bg-gradient-to-br from-shu/10 via-sakura/8 to-kin/10">
        <div class="relative rounded-[10px] bg-kami overflow-hidden">
          <!-- Ornamental corners -->
          <div class="absolute top-2 left-2 w-4 h-4 border-t-2 border-l-2 border-shu/10 rounded-tl-sm"></div>
          <div class="absolute top-2 right-2 w-4 h-4 border-t-2 border-r-2 border-shu/10 rounded-tr-sm"></div>
          <div class="absolute bottom-2 left-2 w-4 h-4 border-b-2 border-l-2 border-shu/10 rounded-bl-sm"></div>
          <div class="absolute bottom-2 right-2 w-4 h-4 border-b-2 border-r-2 border-shu/10 rounded-br-sm"></div>

          <div class="p-8 relative z-10">
            <!-- Ornamental section header -->
            <h3 class="flex items-center gap-3 font-kai font-bold text-sumi-800 mb-6 pb-4 border-b border-shu/8">
              <User class="text-shu/40" size="18" />
              <span>基本信息</span>
              <span class="flex-1 flex items-center gap-2 justify-end">
                <span class="block w-8 h-[1px] bg-shu/10"></span>
                <span class="font-brush text-shu/10 text-lg">人</span>
              </span>
            </h3>
            <div class="grid gap-5 md:grid-cols-2">
              <div class="space-y-2">
                <label class="text-[10px] font-kai text-sumi-500 tracking-widest flex items-center gap-2">
                  <span class="w-2 h-[1px] bg-shu/20"></span>顯示名稱
                </label>
                <input v-model="profile.name" type="text" class="sumi-input focus:border-shu/15 focus:shadow-sm focus:shadow-shu/5" />
              </div>
              <div class="space-y-2">
                <label class="text-[10px] font-kai text-sumi-500 tracking-widest flex items-center gap-2">
                  <span class="w-2 h-[1px] bg-shu/20"></span>所在地
                </label>
                <input v-model="profile.location" type="text" class="sumi-input focus:border-shu/15 focus:shadow-sm focus:shadow-shu/5" />
              </div>
              <div class="space-y-2 md:col-span-2">
                <label class="text-[10px] font-kai text-sumi-500 tracking-widest flex items-center gap-2">
                  <span class="w-2 h-[1px] bg-shu/20"></span>頭像 URL
                </label>
                <input v-model="profile.avatar" type="text" class="sumi-input focus:border-shu/15 focus:shadow-sm focus:shadow-shu/5" />
              </div>
              <div class="space-y-2 md:col-span-2">
                <label class="text-[10px] font-kai text-sumi-500 tracking-widest flex items-center gap-2">
                  <span class="w-2 h-[1px] bg-shu/20"></span>個人簡介
                </label>
                <textarea v-model="profile.bio" rows="3" class="sumi-input resize-none focus:border-shu/15 focus:shadow-sm focus:shadow-shu/5"></textarea>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Social Links Section: Chinese frame double-border -->
      <div class="relative p-[2px] rounded-xl bg-gradient-to-br from-mizu/10 via-take/8 to-mizu-deep/10">
        <div class="relative rounded-[10px] bg-kami overflow-hidden">
          <!-- Ornamental corners -->
          <div class="absolute top-2 left-2 w-4 h-4 border-t-2 border-l-2 border-mizu/15 rounded-tl-sm"></div>
          <div class="absolute top-2 right-2 w-4 h-4 border-t-2 border-r-2 border-mizu/15 rounded-tr-sm"></div>
          <div class="absolute bottom-2 left-2 w-4 h-4 border-b-2 border-l-2 border-mizu/15 rounded-bl-sm"></div>
          <div class="absolute bottom-2 right-2 w-4 h-4 border-b-2 border-r-2 border-mizu/15 rounded-br-sm"></div>

          <div class="p-8 relative z-10">
            <!-- Ornamental section header -->
            <h3 class="flex items-center gap-3 font-kai font-bold text-sumi-800 mb-6 pb-4 border-b border-mizu/10">
              <LinkIcon class="text-mizu/40" size="18" />
              <span>社交鏈接</span>
              <span class="flex-1 flex items-center gap-2 justify-end">
                <span class="block w-8 h-[1px] bg-mizu/10"></span>
                <span class="font-brush text-mizu/10 text-lg">鏈</span>
              </span>
            </h3>
            <div class="grid gap-5 md:grid-cols-2">
              <div class="space-y-2">
                <label class="text-[10px] font-kai text-sumi-500 tracking-widest flex items-center gap-2">
                  <span class="w-2 h-[1px] bg-mizu/25"></span>GitHub
                </label>
                <div class="relative"><span class="absolute left-3 top-1/2 -translate-y-1/2 text-sumi-300 text-xs">github.com/</span><input v-model="profile.social.github" type="text" class="sumi-input pl-24 focus:border-mizu/15 focus:shadow-sm focus:shadow-mizu/5" /></div>
              </div>
              <div class="space-y-2">
                <label class="text-[10px] font-kai text-sumi-500 tracking-widest flex items-center gap-2">
                  <span class="w-2 h-[1px] bg-mizu/25"></span>郵箱地址
                </label>
                <input v-model="profile.social.email" type="text" class="sumi-input focus:border-mizu/15 focus:shadow-sm focus:shadow-mizu/5" />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
