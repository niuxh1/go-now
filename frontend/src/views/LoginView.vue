<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '../stores/auth';
import api from '../services/api';
import { LogIn, User, Lock, AlertCircle, ArrowRight } from 'lucide-vue-next';

const username = ref('');
const password = ref('');
const error = ref('');
const router = useRouter();
const authStore = useAuthStore();

const login = async () => {
  try {
    const response = await api.post('/auth/login', {
      username: username.value,
      password: password.value
    });
    authStore.setToken(response.data.token);
    router.push('/');
  } catch (err) {
    error.value = err.response?.data?.error || '登录失败，请检查账号密码';
  }
};
</script>

<template>
  <div class="min-h-screen pt-24 pb-12 flex justify-center items-center">
    <div class="w-full max-w-md animate-fade-in px-4">
      <div class="fish-card p-10 rounded-[2.5rem]">
        <div class="text-center mb-10">
          <div class="w-16 h-16 bg-zinc-900 text-white dark:bg-white dark:text-black rounded-3xl flex items-center justify-center mx-auto mb-6 rotate-3 shadow-lg">
             <User size="32" stroke-width="2.5" />
          </div>
          <h2 class="text-3xl font-bold text-zinc-900 dark:text-zinc-50 mb-2 tracking-tight">欢迎回来</h2>
          <p class="text-zinc-600 dark:text-zinc-400 text-sm font-medium">登录以管理您的数字花园</p>
        </div>
        
        <form @submit.prevent="login" class="space-y-5">
          <div class="space-y-2">
            <label class="text-xs font-bold text-zinc-600 dark:text-zinc-400 uppercase tracking-widest ml-1">用户名</label>
            <div class="relative group">
              <User size="18" class="absolute left-4 top-1/2 -translate-y-1/2 text-zinc-400 group-focus-within:text-fish-blue transition-colors" />
              <input 
                v-model="username" 
                type="text" 
                placeholder="您的账户名称" 
                class="w-full bg-gray-50 dark:bg-zinc-900 border border-gray-200 dark:border-zinc-700 text-zinc-800 dark:text-zinc-200 rounded-2xl py-4 pl-12 pr-4 focus:outline-none focus:border-fish-blue focus:ring-1 focus:ring-fish-blue transition-all"
              />
            </div>
          </div>
          
          <div class="space-y-2">
            <div class="flex justify-between items-center ml-1">
              <label class="text-xs font-bold text-zinc-600 dark:text-zinc-400 uppercase tracking-widest">密码</label>
              <router-link to="/reset-password" class="text-xs font-medium text-fish-blue hover:underline transition-colors">忘记密码?</router-link>
            </div>
            <div class="relative group">
              <Lock size="18" class="absolute left-4 top-1/2 -translate-y-1/2 text-zinc-400 group-focus-within:text-fish-blue transition-colors" />
              <input 
                v-model="password" 
                type="password" 
                placeholder="您的安全密码" 
                class="w-full bg-gray-50 dark:bg-zinc-900 border border-gray-200 dark:border-zinc-700 text-zinc-800 dark:text-zinc-200 rounded-2xl py-4 pl-12 pr-4 focus:outline-none focus:border-fish-blue focus:ring-1 focus:ring-fish-blue transition-all"
              />
            </div>
          </div>

          <div v-if="error" class="bg-red-50 dark:bg-red-900/10 border border-red-100 dark:border-red-900/20 rounded-2xl p-4 flex items-center gap-3 text-red-600 dark:text-red-400 text-sm">
            <AlertCircle size="18" />
            {{ error }}
          </div>

          <button type="submit" class="w-full bg-zinc-900 dark:bg-white text-white dark:text-black font-bold py-4 rounded-2xl shadow-xl shadow-zinc-200 dark:shadow-zinc-900/50 hover:bg-zinc-800 dark:hover:bg-zinc-200 transition-all flex items-center justify-center gap-2 mt-4 group">
            登录系统 <ArrowRight size="18" class="group-hover:translate-x-1 transition-transform" />
          </button>
        </form>
        
        <div class="mt-10 text-center">
          <p class="text-zinc-600 dark:text-zinc-400 text-sm">
            还没有账户? 
            <router-link to="/register" class="text-fish-blue font-bold hover:underline ml-1">立即加入</router-link>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>