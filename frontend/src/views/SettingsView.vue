<script setup>
import { ref, onMounted, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { fetchSettings, updateSettings } from '../services/api'
import { Save, AlertCircle, CheckCircle, User, Link as LinkIcon } from 'lucide-vue-next'

const router = useRouter()
const loading = ref(true)
const saving = ref(false)
const message = ref({ type: '', text: '' })

// Profile Data
const profile = reactive({
  name: '',
  bio: '',
  avatar: '',
  location: '',
  social: {
    github: '',
    email: ''
  }
})

// Fetch settings
onMounted(async () => {
  try {
    const { data } = await fetchSettings()
    if (data.profile) {
      Object.assign(profile, data.profile)
    }
  } catch (error) {
    console.error("Failed to load settings", error)
    message.value = { type: 'error', text: '无法加载设置' }
  } finally {
    loading.value = false
  }
})

// Save settings
const handleSave = async () => {
  saving.value = true
  message.value = { type: '', text: '' }
  
  try {
    await updateSettings({ profile })
    message.value = { type: 'success', text: '设置已保存' }
    setTimeout(() => message.value = { type: '', text: '' }, 3000)
  } catch (error) {
    console.error("Failed to save", error)
    message.value = { type: 'error', text: '保存失败: ' + (error.response?.data?.error || error.message) }
  } finally {
    saving.value = false
  }
}
</script>

<template>
  <div class="max-w-4xl mx-auto py-12 px-6 animate-fade-in">
    <div class="flex items-center justify-between mb-12">
      <div>
        <h2 class="text-3xl font-bold text-zinc-800 dark:text-white tracking-tight">站点设置</h2>
        <p class="text-zinc-500 text-sm mt-1">管理您的个人资料与站点配置</p>
      </div>
      <button 
        @click="handleSave" 
        class="flex items-center gap-2 bg-zinc-900 dark:bg-white text-white dark:text-black px-6 py-2 rounded-xl font-bold hover:opacity-90 transition-all shadow-lg"
        :disabled="saving || loading"
      >
        <Save size="18" />
        {{ saving ? '保存中...' : '保存更改' }}
      </button>
    </div>

    <!-- Message -->
    <div v-if="message.text" :class="[
      'mb-8 p-4 rounded-xl flex items-center gap-3 text-sm font-medium',
      message.type === 'error' ? 'bg-red-50 text-red-600 border border-red-100' : 'bg-green-50 text-green-600 border border-green-100'
    ]">
      <AlertCircle v-if="message.type === 'error'" size="18" />
      <CheckCircle v-else size="18" />
      {{ message.text }}
    </div>

    <div v-if="loading" class="text-center py-12 text-zinc-500">加载中...</div>

    <div v-else class="space-y-8">
      <!-- Basic Profile -->
      <div class="fish-card p-8">
        <h3 class="flex items-center gap-2 font-bold text-zinc-800 dark:text-white mb-6 border-b border-gray-100 dark:border-zinc-700 pb-4">
          <User class="text-fish-blue" size="20" /> 基本信息
        </h3>
        
        <div class="grid gap-6 md:grid-cols-2">
          <div class="space-y-2">
            <label class="text-xs font-bold text-zinc-500 uppercase">显示名称</label>
            <input v-model="profile.name" type="text" class="w-full bg-gray-50 dark:bg-zinc-900 border border-gray-200 dark:border-zinc-700 rounded-lg p-3 focus:border-fish-blue focus:outline-none transition-colors" />
          </div>
          
          <div class="space-y-2">
            <label class="text-xs font-bold text-zinc-500 uppercase">所在地</label>
            <input v-model="profile.location" type="text" class="w-full bg-gray-50 dark:bg-zinc-900 border border-gray-200 dark:border-zinc-700 rounded-lg p-3 focus:border-fish-blue focus:outline-none transition-colors" />
          </div>

          <div class="space-y-2 md:col-span-2">
            <label class="text-xs font-bold text-zinc-500 uppercase">头像 URL</label>
            <input v-model="profile.avatar" type="text" class="w-full bg-gray-50 dark:bg-zinc-900 border border-gray-200 dark:border-zinc-700 rounded-lg p-3 focus:border-fish-blue focus:outline-none transition-colors" />
          </div>

          <div class="space-y-2 md:col-span-2">
            <label class="text-xs font-bold text-zinc-500 uppercase">个人简介</label>
            <textarea v-model="profile.bio" rows="3" class="w-full bg-gray-50 dark:bg-zinc-900 border border-gray-200 dark:border-zinc-700 rounded-lg p-3 focus:border-fish-blue focus:outline-none transition-colors resize-none"></textarea>
          </div>
        </div>
      </div>

      <!-- Social Links -->
      <div class="fish-card p-8">
        <h3 class="flex items-center gap-2 font-bold text-zinc-800 dark:text-white mb-6 border-b border-gray-100 dark:border-zinc-700 pb-4">
          <LinkIcon class="text-fish-blue" size="20" /> 社交链接
        </h3>
        
        <div class="grid gap-6 md:grid-cols-2">
          <div class="space-y-2">
            <label class="text-xs font-bold text-zinc-500 uppercase">GitHub 用户名</label>
            <div class="relative">
               <span class="absolute left-3 top-1/2 -translate-y-1/2 text-zinc-400 text-xs">github.com/</span>
               <input v-model="profile.social.github" type="text" class="w-full bg-gray-50 dark:bg-zinc-900 border border-gray-200 dark:border-zinc-700 rounded-lg p-3 pl-24 focus:border-fish-blue focus:outline-none transition-colors" />
            </div>
          </div>
          
          <div class="space-y-2">
            <label class="text-xs font-bold text-zinc-500 uppercase">邮箱地址</label>
            <input v-model="profile.social.email" type="text" class="w-full bg-gray-50 dark:bg-zinc-900 border border-gray-200 dark:border-zinc-700 rounded-lg p-3 focus:border-fish-blue focus:outline-none transition-colors" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>