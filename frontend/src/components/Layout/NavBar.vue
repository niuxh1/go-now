<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { Menu, X, LogIn, LogOut, PenTool, Settings } from 'lucide-vue-next'
import { useAuthStore } from '../../stores/auth'

const isScrolled = ref(false)
const isMenuOpen = ref(false)
const authStore = useAuthStore()
const router = useRouter()

const handleScroll = () => {
  isScrolled.value = window.scrollY > 50
}

const handleLogout = () => {
  authStore.logout()
  router.push('/')
}

onMounted(() => {
  window.addEventListener('scroll', handleScroll)
})

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
})

const navLinks = [
  { name: '首页', path: '/', icon: 'home' },
  { name: '归档', path: '/archives', icon: 'archive' },
]
</script>

<template>
  <nav 
    class="fixed top-0 left-0 w-full z-50 transition-all duration-300"
    :class="[
      isScrolled ? 'bg-white/90 backdrop-blur-md shadow-md py-2 dark:bg-zinc-900/90' : 'bg-transparent py-4'
    ]"
  >
    <div class="max-w-7xl mx-auto px-6 flex items-center justify-between">
      <!-- Logo -->
      <RouterLink to="/" class="text-2xl font-bold tracking-tight transition-colors"
        :class="[isScrolled ? 'text-fish-blue' : 'text-white hover:text-white/80']"
      >
        GoNow
      </RouterLink>

      <!-- Desktop Nav -->
      <div class="hidden md:flex items-center gap-8">
        <RouterLink 
          v-for="link in navLinks" 
          :key="link.path" 
          :to="link.path"
          class="text-sm font-medium transition-colors relative group"
          :class="[isScrolled ? 'text-zinc-700 hover:text-fish-blue dark:text-zinc-200' : 'text-white/90 hover:text-white']"
        >
          {{ link.name }}
          <span class="absolute -bottom-1 left-0 w-0 h-0.5 bg-fish-blue transition-all duration-300 group-hover:w-full"></span>
        </RouterLink>
      </div>

      <!-- Actions -->
      <div class="flex items-center gap-4">
        <!-- Auth Actions (Desktop) -->
        <div class="hidden md:flex items-center gap-4">
          <template v-if="authStore.isAuthenticated">
             <RouterLink v-if="authStore.isAdmin" to="/editor" class="text-zinc-500 dark:text-zinc-300 hover:text-fish-blue dark:hover:text-fish-blue transition-colors" title="Write Article">
               <PenTool size="20" />
             </RouterLink>
             <RouterLink v-if="authStore.isAdmin" to="/settings" class="text-zinc-500 dark:text-zinc-300 hover:text-fish-blue dark:hover:text-fish-blue transition-colors" title="Settings">
               <Settings size="20" />
             </RouterLink>
             <button @click="handleLogout" class="text-zinc-500 dark:text-zinc-300 hover:text-red-500 dark:hover:text-red-400 transition-colors" title="Logout">
               <LogOut size="20" />
             </button>
          </template>
          <template v-else>
             <RouterLink to="/login" class="px-4 py-1.5 rounded-full border border-white/20 text-sm font-medium hover:bg-white hover:text-black transition-all"
               :class="[isScrolled ? 'text-zinc-700 border-zinc-300 hover:bg-zinc-100 hover:text-zinc-900 dark:text-zinc-200' : 'text-white']"
             >
               登录
             </RouterLink>
          </template>
        </div>

        <!-- Mobile Menu Button -->
        <button class="md:hidden" :class="[isScrolled ? 'text-zinc-800' : 'text-white']" @click="isMenuOpen = !isMenuOpen">
          <Menu v-if="!isMenuOpen" />
          <X v-else />
        </button>
      </div>
    </div>
    
    <!-- Mobile Menu -->
    <div v-if="isMenuOpen" class="md:hidden absolute top-full left-0 w-full bg-white dark:bg-zinc-900 shadow-xl border-t dark:border-zinc-800 animate-fade-in">
       <div class="flex flex-col p-4">
          <RouterLink 
            v-for="link in navLinks" 
            :key="link.path" 
            :to="link.path"
            class="py-3 px-4 text-zinc-700 dark:text-zinc-200 hover:bg-zinc-100 dark:hover:bg-zinc-800 rounded-lg"
            @click="isMenuOpen = false"
          >
            {{ link.name }}
          </RouterLink>
          <hr class="my-2 border-gray-100 dark:border-zinc-800" />
          <template v-if="authStore.isAuthenticated">
            <RouterLink to="/editor" class="py-3 px-4 text-zinc-700 dark:text-zinc-200 hover:bg-zinc-100 dark:hover:bg-zinc-800 rounded-lg flex items-center gap-2" @click="isMenuOpen = false">
               <PenTool size="16" /> 写文章
            </RouterLink>
             <RouterLink to="/settings" class="py-3 px-4 text-zinc-700 dark:text-zinc-200 hover:bg-zinc-100 dark:hover:bg-zinc-800 rounded-lg flex items-center gap-2" @click="isMenuOpen = false">
               <Settings size="16" /> 设置
            </RouterLink>
            <button @click="handleLogout(); isMenuOpen = false" class="py-3 px-4 text-red-500 hover:bg-red-50 dark:hover:bg-red-900/10 rounded-lg flex items-center gap-2 w-full text-left">
               <LogOut size="16" /> 退出登录
            </button>
          </template>
          <template v-else>
             <RouterLink to="/login" class="py-3 px-4 text-fish-blue font-bold hover:bg-blue-50 dark:hover:bg-blue-900/10 rounded-lg" @click="isMenuOpen = false">
               登录
            </RouterLink>
          </template>
       </div>
    </div>
  </nav>
</template>
