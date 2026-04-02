<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import api from '../services/api'
import { UserPlus, User, Mail, Lock, AlertCircle, ArrowRight } from 'lucide-vue-next'

const username = ref(''); const email = ref(''); const password = ref(''); const error = ref('')
const router = useRouter()

const register = async () => {
  try { await api.post('/auth/register', { username: username.value, email: email.value, password: password.value }); router.push('/login') }
  catch (e) { error.value = e.response?.data?.error || '註冊失敗' }
}
</script>

<template>
  <div class="min-h-screen pt-28 pb-16 flex justify-center items-start relative">
    <!-- Floating background watermark -->
    <div class="fixed inset-0 pointer-events-none select-none flex items-center justify-center overflow-hidden">
      <span class="font-brush text-[20rem] text-sumi-800/[0.012] leading-none transform rotate-6">新</span>
    </div>

    <div class="w-full max-w-md animate-fade-up px-4 relative z-10">
      <!-- Chinese frame card: double border with mizu/take theme -->
      <div class="relative p-[3px] rounded-xl bg-gradient-to-br from-mizu/25 via-take/15 to-mizu-deep/20">
        <div class="relative rounded-[10px] border border-mizu/10 bg-kami overflow-hidden">
          <!-- Inner pattern background -->
          <div class="absolute inset-0 opacity-[0.015]" style="background-image: url('data:image/svg+xml,%3Csvg width=&quot;40&quot; height=&quot;40&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;%3E%3Ccircle cx=&quot;20&quot; cy=&quot;20&quot; r=&quot;1&quot; fill=&quot;%23366B6B&quot;/%3E%3C/svg%3E'); background-size: 40px 40px;"></div>

          <!-- Ornamental corners -->
          <div class="absolute top-2 left-2 w-5 h-5 border-t-2 border-l-2 border-mizu/20 rounded-tl-sm"></div>
          <div class="absolute top-2 right-2 w-5 h-5 border-t-2 border-r-2 border-mizu/20 rounded-tr-sm"></div>
          <div class="absolute bottom-2 left-2 w-5 h-5 border-b-2 border-l-2 border-mizu/20 rounded-bl-sm"></div>
          <div class="absolute bottom-2 right-2 w-5 h-5 border-b-2 border-r-2 border-mizu/20 rounded-br-sm"></div>

          <div class="relative z-10 p-10">
            <!-- Header with seal character -->
            <div class="text-center mb-10">
              <div class="w-16 h-16 rounded-full bg-gradient-to-br from-mizu/20 via-take/15 to-mizu-deep/10 flex items-center justify-center mx-auto mb-5 border border-mizu/10 shadow-inner">
                <span class="font-brush text-mizu-deep text-3xl">友</span>
              </div>
              <!-- Decorative lines -->
              <div class="flex items-center justify-center gap-3 mb-4">
                <span class="block w-12 h-[1px] bg-gradient-to-r from-transparent to-mizu/25"></span>
                <span class="text-mizu/30 text-xs">&#9670;</span>
                <span class="block w-12 h-[1px] bg-gradient-to-l from-transparent to-mizu/25"></span>
              </div>
              <h2 class="text-3xl font-kai font-bold text-sumi-800 mb-2">創建賬戶</h2>
              <p class="text-sumi-400 text-sm font-kai">加入我們，開始書寫之旅</p>
            </div>

            <form @submit.prevent="register" class="space-y-5">
              <div class="space-y-2">
                <label class="text-[10px] font-kai text-sumi-500 tracking-widest ml-1 flex items-center gap-2">
                  <span class="w-3 h-[1px] bg-mizu/25"></span>用戶名
                </label>
                <div class="relative group">
                  <User size="15" class="absolute left-4 top-1/2 -translate-y-1/2 text-sumi-300 group-focus-within:text-mizu-deep transition-colors" />
                  <input v-model="username" type="text" placeholder="設置用戶名" class="sumi-input pl-11 focus:border-mizu/20 focus:shadow-sm focus:shadow-mizu/5" />
                </div>
              </div>
              <div class="space-y-2">
                <label class="text-[10px] font-kai text-sumi-500 tracking-widest ml-1 flex items-center gap-2">
                  <span class="w-3 h-[1px] bg-mizu/25"></span>電子郵箱
                </label>
                <div class="relative group">
                  <Mail size="15" class="absolute left-4 top-1/2 -translate-y-1/2 text-sumi-300 group-focus-within:text-mizu-deep transition-colors" />
                  <input v-model="email" type="email" placeholder="聯繫郵箱" class="sumi-input pl-11 focus:border-mizu/20 focus:shadow-sm focus:shadow-mizu/5" />
                </div>
              </div>
              <div class="space-y-2">
                <label class="text-[10px] font-kai text-sumi-500 tracking-widest ml-1 flex items-center gap-2">
                  <span class="w-3 h-[1px] bg-mizu/25"></span>密碼
                </label>
                <div class="relative group">
                  <Lock size="15" class="absolute left-4 top-1/2 -translate-y-1/2 text-sumi-300 group-focus-within:text-mizu-deep transition-colors" />
                  <input v-model="password" type="password" placeholder="設置安全密碼" class="sumi-input pl-11 focus:border-mizu/20 focus:shadow-sm focus:shadow-mizu/5" />
                </div>
              </div>

              <div v-if="error" class="bg-shu/8 border border-shu/15 rounded-lg p-4 flex items-center gap-3 text-shu text-sm font-kai">
                <AlertCircle size="15" /> {{ error }}
              </div>

              <!-- Ornamental submit button with shimmer -->
              <button type="submit"
                class="relative w-full py-4 rounded-lg font-kai font-medium tracking-wider transition-all duration-300 bg-gradient-to-r from-mizu via-mizu-deep to-mizu text-white shadow-lg shadow-mizu/20 hover:shadow-xl hover:shadow-mizu/30 flex items-center justify-center gap-2 mt-4 group overflow-hidden">
                <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/10 to-transparent -translate-x-full group-hover:translate-x-full transition-transform duration-1000"></div>
                <span class="relative z-10 flex items-center gap-2">立即註冊 <ArrowRight size="15" class="group-hover:translate-x-1 transition-transform" /></span>
              </button>
            </form>

            <div class="flex items-center gap-3 my-8">
              <span class="flex-1 h-[1px] bg-gradient-to-r from-transparent to-sumi-200/40"></span>
              <span class="text-sumi-200 text-[10px] font-kai">或</span>
              <span class="flex-1 h-[1px] bg-gradient-to-l from-transparent to-sumi-200/40"></span>
            </div>

            <p class="text-center text-sumi-400 text-sm font-kai">
              已有賬戶? <router-link to="/login" class="text-mizu-deep font-medium hover:text-mizu transition-colors ml-1">直接登錄</router-link>
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
