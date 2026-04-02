<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { Menu, X, LogOut, PenTool, Settings } from 'lucide-vue-next'
import { useAuthStore } from '../../stores/auth'

const isScrolled = ref(false)
const isMenuOpen = ref(false)
const authStore = useAuthStore()
const router = useRouter()

const handleScroll = () => { isScrolled.value = window.scrollY > 50 }
const handleLogout = () => { authStore.logout(); router.push('/') }

onMounted(() => window.addEventListener('scroll', handleScroll))
onUnmounted(() => window.removeEventListener('scroll', handleScroll))

const navLinks = [
  { name: '首頁', path: '/' },
  { name: '歸檔', path: '/archives' },
]
</script>

<template>
  <nav class="fixed top-0 left-0 w-full z-50 transition-all duration-500"
    :class="[isScrolled
      ? 'bg-gradient-to-r from-white/75 via-[#fdf8f4]/80 to-white/75 backdrop-blur-xl shadow-sm shadow-sakura/5 py-2'
      : 'bg-transparent py-4']"
    :style="!isScrolled ? 'text-shadow: 0 1px 3px rgba(255,255,255,0.6)' : ''">

    <!-- Decorative bottom border when scrolled: sakura -> kin -> mizu gradient -->
    <div v-if="isScrolled" class="absolute bottom-0 left-0 w-full">
      <div class="h-[2px] bg-gradient-to-r from-transparent via-sakura/25 to-transparent"></div>
      <div class="h-px mt-px bg-gradient-to-r from-sakura/10 via-kin/20 to-mizu/10"></div>
    </div>

    <div class="max-w-5xl mx-auto px-6 flex items-center justify-between">
      <!-- Logo with ornamental ring -->
      <RouterLink to="/" class="flex items-center gap-3 group">
        <div class="relative">
          <div class="w-9 h-9 rounded-full bg-gradient-to-br from-sakura/30 via-kin/10 to-mizu/30 group-hover:from-sakura/50 group-hover:to-mizu/50 flex items-center justify-center transition-all duration-300">
            <span class="font-brush text-sakura-deep text-sm">記</span>
          </div>
          <!-- Ornamental ring around logo -->
          <div class="absolute inset-[-3px] rounded-full border border-kin/[0.15] group-hover:border-kin/30 transition-all duration-300"></div>
          <div class="absolute inset-[-6px] rounded-full border border-dashed border-sakura/[0.08] group-hover:border-sakura/15 transition-all duration-500 group-hover:rotate-45"></div>
        </div>
        <div class="flex items-center gap-1.5">
          <span class="font-kai text-lg font-bold text-sumi-800 tracking-wide">GoNow</span>
          <span class="text-kin/40 text-[10px] font-kai hidden sm:inline">庭園</span>
        </div>
      </RouterLink>

      <!-- Desktop Nav with decorative separators -->
      <div class="hidden md:flex items-center gap-0">
        <template v-for="(link, index) in navLinks" :key="link.path">
          <!-- Decorative diamond separator between links -->
          <span v-if="index > 0" class="text-kin/25 text-[8px] mx-1 select-none">&#9670;</span>
          <RouterLink :to="link.path"
            class="px-4 py-2 text-sm font-kai tracking-wider text-sumi-500 hover:text-sakura-deep transition-all duration-300 relative group rounded-lg hover:bg-sakura/5">
            {{ link.name }}
            <!-- Ink-spread underline on hover -->
            <span class="absolute bottom-0.5 left-1/2 -translate-x-1/2 w-0 h-[2px] rounded-full transition-all duration-500 group-hover:w-6 overflow-hidden"
              style="background: linear-gradient(90deg, transparent, rgba(232,160,176,0.6), rgba(201,169,89,0.4), transparent);"></span>
            <!-- Subtle dot under active -->
            <span class="absolute -bottom-0.5 left-1/2 -translate-x-1/2 w-1 h-1 rounded-full bg-sakura/0 group-hover:bg-sakura/30 transition-all duration-300"></span>
          </RouterLink>
        </template>
      </div>

      <!-- Actions -->
      <div class="flex items-center gap-3">
        <div class="hidden md:flex items-center gap-2">
          <template v-if="authStore.isAuthenticated">
            <RouterLink v-if="authStore.isAdmin" to="/editor" class="p-2 text-sumi-400 hover:text-sakura-deep hover:bg-sakura/5 rounded-lg transition-all duration-300 relative group" title="撰文">
              <PenTool size="17" />
              <span class="absolute -bottom-0.5 left-1/2 -translate-x-1/2 w-0 h-px bg-sakura/30 group-hover:w-3 transition-all duration-300"></span>
            </RouterLink>
            <RouterLink v-if="authStore.isAdmin" to="/settings" class="p-2 text-sumi-400 hover:text-mizu-deep hover:bg-mizu/5 rounded-lg transition-all duration-300 relative group" title="設置">
              <Settings size="17" />
              <span class="absolute -bottom-0.5 left-1/2 -translate-x-1/2 w-0 h-px bg-mizu/30 group-hover:w-3 transition-all duration-300"></span>
            </RouterLink>
            <button @click="handleLogout" class="p-2 text-sumi-400 hover:text-shu hover:bg-shu/5 rounded-lg transition-all duration-300" title="退出">
              <LogOut size="17" />
            </button>
          </template>
          <template v-else>
            <RouterLink to="/login" class="sumi-btn text-xs py-1.5 px-4 rounded-full border border-kin/10 hover:border-kin/20 transition-all duration-300">登錄</RouterLink>
          </template>
        </div>
        <button class="md:hidden p-2 text-sumi-500" @click="isMenuOpen = !isMenuOpen">
          <Menu v-if="!isMenuOpen" size="20" />
          <X v-else size="20" />
        </button>
      </div>
    </div>

    <!-- Mobile Menu with pattern background -->
    <transition name="page">
      <div v-if="isMenuOpen" class="md:hidden absolute top-full left-0 w-full border-b border-sakura/10 shadow-lg shadow-sakura/5 overflow-hidden">
        <!-- Warm frosted background -->
        <div class="absolute inset-0 bg-gradient-to-b from-[#fdf8f4]/97 to-white/95 backdrop-blur-xl"></div>
        <!-- Pattern overlay -->
        <div class="absolute inset-0 seigaiha-pattern opacity-30"></div>
        <div class="absolute inset-0 asanoha-pattern opacity-20"></div>
        <!-- Watercolor wash -->
        <div class="absolute top-0 right-0 w-[200px] h-[150px] rounded-full opacity-[0.06]"
          style="background: radial-gradient(circle, #e8a0b0, transparent 70%);"></div>

        <div class="flex flex-col p-4 gap-1 relative z-10">
          <RouterLink v-for="link in navLinks" :key="link.path" :to="link.path"
            class="py-3 px-4 text-sumi-600 hover:text-sakura-deep hover:bg-sakura/5 font-kai rounded-lg transition-all"
            @click="isMenuOpen = false">{{ link.name }}</RouterLink>
          <div class="h-px bg-gradient-to-r from-sakura/10 via-kin/10 to-mizu/10 my-2"></div>
          <template v-if="authStore.isAuthenticated">
            <RouterLink to="/editor" class="py-3 px-4 text-sumi-600 hover:text-sakura-deep hover:bg-sakura/5 font-kai rounded-lg flex items-center gap-3" @click="isMenuOpen = false">
              <PenTool size="15" /> 撰文</RouterLink>
            <RouterLink to="/settings" class="py-3 px-4 text-sumi-600 hover:text-mizu-deep hover:bg-mizu/5 font-kai rounded-lg flex items-center gap-3" @click="isMenuOpen = false">
              <Settings size="15" /> 設置</RouterLink>
            <button @click="handleLogout(); isMenuOpen = false" class="py-3 px-4 text-shu hover:bg-shu/5 font-kai rounded-lg flex items-center gap-3 w-full text-left">
              <LogOut size="15" /> 退出</button>
          </template>
          <template v-else>
            <RouterLink to="/login" class="py-3 px-4 text-sakura-deep font-kai font-medium hover:bg-sakura/5 rounded-lg" @click="isMenuOpen = false">登錄</RouterLink>
          </template>
        </div>
      </div>
    </transition>
  </nav>
</template>
