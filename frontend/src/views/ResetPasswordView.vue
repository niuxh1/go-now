<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import api from '../services/api';
import { KeyRound, User, Mail, Lock, AlertCircle, ArrowRight, CheckCircle2 } from 'lucide-vue-next';

const username = ref('');
const email = ref('');
const newPassword = ref('');
const message = ref('');
const error = ref('');
const router = useRouter();

const resetPassword = async () => {
  error.value = '';
  message.value = '';
  try {
    await api.post('/auth/reset-password', {
      username: username.value,
      email: email.value,
      new_password: newPassword.value
    });
    message.value = '密码重置成功，请重新登录';
    setTimeout(() => {
      router.push('/login');
    }, 2000);
  } catch (err) {
    error.value = err.response?.data?.error || '重置失败，请检查信息是否正确';
  }
};
</script>

<template>
  <div class="min-h-[70vh] flex justify-center items-center">
    <div class="w-full max-w-md animate-fade-in px-4">
      <div class="fish-card p-10 rounded-[2.5rem] shadow-2xl">
        <div class="text-center mb-10">
          <div class="w-16 h-16 bg-zinc-900 text-white dark:bg-white dark:text-black rounded-3xl flex items-center justify-center mx-auto mb-6">
             <KeyRound size="32" stroke-width="2.5" />
          </div>
          <h2 class="text-3xl font-bold text-zinc-900 dark:text-zinc-50 mb-2 tracking-tight">找回密码</h2>
          <p class="text-zinc-600 dark:text-zinc-400 text-sm font-medium">验证您的信息以重新获得访问权限</p>
        </div>
        
        <form @submit.prevent="resetPassword" class="space-y-5">
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
            <label class="text-xs font-bold text-zinc-600 dark:text-zinc-400 uppercase tracking-widest ml-1">电子邮箱</label>
            <div class="relative group">
              <Mail size="18" class="absolute left-4 top-1/2 -translate-y-1/2 text-zinc-400 group-focus-within:text-fish-blue transition-colors" />
              <input 
                v-model="email" 
                type="email" 
                placeholder="注册时使用的邮箱" 
                class="w-full bg-gray-50 dark:bg-zinc-900 border border-gray-200 dark:border-zinc-700 text-zinc-800 dark:text-zinc-200 rounded-2xl py-4 pl-12 pr-4 focus:outline-none focus:border-fish-blue focus:ring-1 focus:ring-fish-blue transition-all"
              />
            </div>
          </div>
          
          <div class="space-y-2">
            <label class="text-xs font-bold text-zinc-600 dark:text-zinc-400 uppercase tracking-widest ml-1">新密码</label>
            <div class="relative group">
              <Lock size="18" class="absolute left-4 top-1/2 -translate-y-1/2 text-zinc-400 group-focus-within:text-fish-blue transition-colors" />
              <input 
                v-model="newPassword" 
                type="password" 
                placeholder="设置您的新密码" 
                class="w-full bg-gray-50 dark:bg-zinc-900 border border-gray-200 dark:border-zinc-700 text-zinc-800 dark:text-zinc-200 rounded-2xl py-4 pl-12 pr-4 focus:outline-none focus:border-fish-blue focus:ring-1 focus:ring-fish-blue transition-all"
              />
            </div>
          </div>

          <div v-if="error" class="bg-red-50 dark:bg-red-900/10 border border-red-100 dark:border-red-900/20 rounded-2xl p-4 flex items-center gap-3 text-red-600 dark:text-red-400 text-sm">
            <AlertCircle size="18" />
            {{ error }}
          </div>

          <div v-if="message" class="bg-emerald-50 dark:bg-emerald-900/10 border border-emerald-100 dark:border-emerald-900/20 rounded-2xl p-4 flex items-center gap-3 text-emerald-600 dark:text-emerald-400 text-sm">
            <CheckCircle2 size="18" />
            {{ message }}
          </div>

          <button type="submit" class="w-full bg-zinc-900 dark:bg-white text-white dark:text-black font-bold py-4 rounded-2xl shadow-xl shadow-zinc-200 dark:shadow-zinc-900/50 hover:bg-zinc-800 dark:hover:bg-zinc-200 transition-all flex items-center justify-center gap-2 mt-4 group">
            重置并登录 <ArrowRight size="18" class="group-hover:translate-x-1 transition-transform" />
          </button>
        </form>
        
        <div class="mt-10 text-center border-t border-gray-100 dark:border-zinc-800 pt-8">
          <p><router-link to="/login" class="text-zinc-600 dark:text-zinc-400 text-sm hover:text-fish-blue transition-colors">记起密码了? 返回登录</router-link></p>
        </div>
      </div>
    </div>
  </div>
</template>