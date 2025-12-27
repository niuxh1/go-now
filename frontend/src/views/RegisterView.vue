<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import api from '../services/api';
import { UserPlus, User, Mail, Lock, AlertCircle, ArrowRight } from 'lucide-vue-next';

const username = ref('');
const email = ref('');
const password = ref('');
const error = ref('');
const router = useRouter();

const register = async () => {
  try {
    await api.post('/auth/register', {
      username: username.value,
      email: email.value,
      password: password.value
    });
    router.push('/login');
  } catch (err) {
    error.value = err.response?.data?.error || '注册失败，请检查输入';
  }
};
</script>

<template>
  <div class="min-h-screen pt-24 pb-12 flex justify-center items-center">
    <div class="w-full max-w-md animate-fade-in px-4">
      <div class="fish-card p-10 rounded-[2.5rem]">
        <div class="text-center mb-10">
          <div class="w-16 h-16 bg-fish-blue text-white rounded-3xl flex items-center justify-center mx-auto mb-6 -rotate-3 shadow-lg">
             <UserPlus size="32" stroke-width="2.5" />
          </div>
          <h2 class="text-3xl font-bold text-zinc-800 dark:text-white mb-2 tracking-tight">创建账户</h2>
          <p class="text-zinc-500 text-sm font-medium">加入我们，开始记录您的创作之旅</p>
        </div>
        
        <form @submit.prevent="register" class="space-y-5">
          <div class="space-y-2">
            <label class="text-xs font-bold text-zinc-400 uppercase tracking-widest ml-1">用户名</label>
            <div class="relative group">
              <User size="18" class="absolute left-4 top-1/2 -translate-y-1/2 text-zinc-400 group-focus-within:text-fish-blue transition-colors" />
              <input 
                v-model="username" 
                type="text" 
                placeholder="设置您的用户名" 
                class="w-full bg-gray-50 dark:bg-zinc-900 border border-gray-200 dark:border-zinc-700 text-zinc-800 dark:text-white rounded-2xl py-4 pl-12 pr-4 focus:outline-none focus:border-fish-blue focus:ring-1 focus:ring-fish-blue transition-all"
              />
            </div>
          </div>

          <div class="space-y-2">
            <label class="text-xs font-bold text-zinc-400 uppercase tracking-widest ml-1">电子邮箱</label>
            <div class="relative group">
              <Mail size="18" class="absolute left-4 top-1/2 -translate-y-1/2 text-zinc-400 group-focus-within:text-fish-blue transition-colors" />
              <input 
                v-model="email" 
                type="email" 
                placeholder="您的联系邮箱" 
                class="w-full bg-gray-50 dark:bg-zinc-900 border border-gray-200 dark:border-zinc-700 text-zinc-800 dark:text-white rounded-2xl py-4 pl-12 pr-4 focus:outline-none focus:border-fish-blue focus:ring-1 focus:ring-fish-blue transition-all"
              />
            </div>
          </div>
          
          <div class="space-y-2">
            <label class="text-xs font-bold text-zinc-400 uppercase tracking-widest ml-1">密码</label>
            <div class="relative group">
              <Lock size="18" class="absolute left-4 top-1/2 -translate-y-1/2 text-zinc-400 group-focus-within:text-fish-blue transition-colors" />
              <input 
                v-model="password" 
                type="password" 
                placeholder="设置安全密码" 
                class="w-full bg-gray-50 dark:bg-zinc-900 border border-gray-200 dark:border-zinc-700 text-zinc-800 dark:text-white rounded-2xl py-4 pl-12 pr-4 focus:outline-none focus:border-fish-blue focus:ring-1 focus:ring-fish-blue transition-all"
              />
            </div>
          </div>

          <div v-if="error" class="bg-red-50 dark:bg-red-900/10 border border-red-100 dark:border-red-900/20 rounded-2xl p-4 flex items-center gap-3 text-red-500 dark:text-red-400 text-sm">
            <AlertCircle size="18" />
            {{ error }}
          </div>

          <button type="submit" class="w-full bg-fish-blue text-white font-bold py-4 rounded-2xl shadow-xl shadow-fish-blue/20 hover:bg-fish-orange transition-all flex items-center justify-center gap-2 mt-4 group">
            立即注册 <ArrowRight size="18" class="group-hover:translate-x-1 transition-transform" />
          </button>
        </form>
        
        <div class="mt-10 text-center border-t border-gray-100 dark:border-zinc-800 pt-8">
          <p class="text-zinc-500 text-sm">
            已有账户? 
            <router-link to="/login" class="text-fish-blue font-bold hover:underline ml-1">直接登录</router-link>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>