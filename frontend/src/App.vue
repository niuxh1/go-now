<script setup>
import { RouterView } from 'vue-router'
import NavBar from './components/Layout/NavBar.vue'
import BackToTop from './components/Layout/BackToTop.vue'

// Global styles are handled in style.css and tailwind config
</script>

<template>
  <div class="min-h-screen bg-gray-50 dark:bg-zinc-950 flex flex-col font-sans">
    <!-- Navigation -->
    <NavBar />

    <!-- Main Content -->
    <main class="flex-grow w-full">
      <RouterView v-slot="{ Component }">
        <transition name="fade" mode="out-in">
          <component :is="Component" />
        </transition>
      </RouterView>
    </main>

    <!-- Back To Top -->
    <BackToTop />

    <!-- Footer -->
    <footer class="bg-white dark:bg-zinc-900 border-t border-gray-100 dark:border-zinc-800 py-12 mt-20">
      <div class="max-w-7xl mx-auto px-6 text-center">
        <div class="flex justify-center gap-8 mb-8 text-zinc-500 text-sm font-medium uppercase tracking-wider">
           <a href="#" class="hover:text-fish-blue transition-colors">Home</a>
           <a href="#" class="hover:text-fish-blue transition-colors">About</a>
           <a href="#" class="hover:text-fish-blue transition-colors">Contact</a>
        </div>
        <p class="text-zinc-400 text-sm">
          © 2025 GoNow Blog. All rights reserved. <br/>
          <span class="text-xs mt-2 inline-block opacity-60">Design inspired by FishBlog</span>
        </p>
      </div>
    </footer>
  </div>
</template>

<style>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
