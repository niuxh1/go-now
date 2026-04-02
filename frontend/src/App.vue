<script setup>
import { RouterView } from 'vue-router'
import NavBar from './components/Layout/NavBar.vue'
import BackToTop from './components/Layout/BackToTop.vue'
</script>

<template>
  <div class="min-h-screen bg-kami flex flex-col font-song relative">
    <!-- ====== 水墨山水意境背景 Ink Landscape Background ====== -->
    <div class="fixed inset-0 pointer-events-none z-0 overflow-hidden">

      <!-- 底层：暖宣纸渐变 Base warm gradient -->
      <div class="absolute inset-0" style="background: linear-gradient(170deg, #faf7f2 0%, #f6ede3 20%, #f2e8db 40%, #f5ebe0 60%, #faf6f0 80%, #fdf9f5 100%);"></div>

      <!-- 墨色山峦剪影 Ink mountain silhouettes (bottom) -->
      <svg class="absolute bottom-0 left-0 w-full" style="height: 45vh; opacity: 0.18;" viewBox="0 0 1440 600" preserveAspectRatio="none">
        <!-- 远山 Distant mountains -->
        <path d="M0 600 L0 380 Q120 200 240 320 Q360 180 480 280 Q540 200 600 240 Q720 120 840 260 Q960 160 1080 240 Q1200 180 1320 300 Q1380 260 1440 320 L1440 600Z"
          fill="url(#mountainFar)"/>
        <!-- 中景山 Mid mountains -->
        <path d="M0 600 L0 420 Q180 300 360 380 Q480 280 600 350 Q720 250 840 340 Q960 280 1080 360 Q1200 300 1440 400 L1440 600Z"
          fill="url(#mountainMid)"/>
        <!-- 近景山 Near mountains -->
        <path d="M0 600 L0 480 Q240 380 480 440 Q600 380 720 430 Q840 370 960 420 Q1100 380 1200 440 Q1320 400 1440 460 L1440 600Z"
          fill="url(#mountainNear)"/>
        <defs>
          <linearGradient id="mountainFar" x1="0" y1="0" x2="0" y2="1">
            <stop offset="0%" stop-color="#6aacb8" stop-opacity="0.6"/>
            <stop offset="100%" stop-color="#6aacb8" stop-opacity="0.15"/>
          </linearGradient>
          <linearGradient id="mountainMid" x1="0" y1="0" x2="0" y2="1">
            <stop offset="0%" stop-color="#7a8a7a" stop-opacity="0.5"/>
            <stop offset="100%" stop-color="#7a8a7a" stop-opacity="0.1"/>
          </linearGradient>
          <linearGradient id="mountainNear" x1="0" y1="0" x2="0" y2="1">
            <stop offset="0%" stop-color="#5a6a5a" stop-opacity="0.35"/>
            <stop offset="100%" stop-color="#5a6a5a" stop-opacity="0.05"/>
          </linearGradient>
        </defs>
      </svg>

      <!-- 云雾层 Cloud/mist layers -->
      <div class="absolute top-[35%] left-0 right-0 h-[200px] opacity-[0.25]"
        style="background: linear-gradient(180deg, transparent 0%, #faf7f2 30%, #faf7f2 70%, transparent 100%);"></div>
      <div class="absolute top-[55%] left-0 right-0 h-[120px] opacity-[0.15]"
        style="background: linear-gradient(180deg, transparent 0%, #f0e8de 40%, #f0e8de 60%, transparent 100%);"></div>

      <!-- 大面积水彩晕染 Watercolor washes (ink painting feel) -->
      <div class="absolute -top-32 -right-32 w-[800px] h-[800px] rounded-full opacity-[0.20] animate-drift"
        style="background: radial-gradient(circle, #e8a0b0 0%, transparent 60%);"></div>
      <div class="absolute top-[25%] -left-48 w-[700px] h-[700px] rounded-full opacity-[0.16] animate-drift"
        style="background: radial-gradient(circle, #6aacb8 0%, transparent 55%); animation-delay: -8s;"></div>
      <div class="absolute bottom-[15%] right-[5%] w-[600px] h-[600px] rounded-full opacity-[0.14] animate-drift"
        style="background: radial-gradient(circle, #9a88b8 0%, transparent 55%); animation-delay: -15s;"></div>
      <div class="absolute top-[60%] left-[25%] w-[500px] h-[500px] rounded-full opacity-[0.12]"
        style="background: radial-gradient(circle, #c9a959 0%, transparent 55%);"></div>

      <!-- 宣纸纤维纹理 Paper fiber texture -->
      <div class="absolute inset-0 opacity-[0.03]"
        style="background-image: url(&quot;data:image/svg+xml,%3Csvg viewBox='0 0 256 256' xmlns='http://www.w3.org/2000/svg'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.85' numOctaves='4' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23n)'/%3E%3C/svg%3E&quot;);"></div>

      <!-- 枯枝梅花 Plum branch (top-right corner, painting style) -->
      <svg class="absolute top-[5%] right-[3%] w-[300px] h-[400px] hidden lg:block" style="opacity: 0.3;" viewBox="0 0 300 400">
        <!-- 主枝干 Main branch -->
        <path d="M280 400 Q260 320 240 260 Q220 200 200 160 Q180 120 160 100 Q140 80 120 70"
          stroke="#5a3a2a" fill="none" stroke-width="3" stroke-linecap="round"/>
        <!-- 分支 Sub branches -->
        <path d="M200 160 Q220 130 250 110" stroke="#5a3a2a" fill="none" stroke-width="2" stroke-linecap="round"/>
        <path d="M160 100 Q140 60 130 30" stroke="#5a3a2a" fill="none" stroke-width="1.8" stroke-linecap="round"/>
        <path d="M240 260 Q270 240 290 230" stroke="#5a3a2a" fill="none" stroke-width="1.5" stroke-linecap="round"/>
        <path d="M180 120 Q150 110 130 120" stroke="#5a3a2a" fill="none" stroke-width="1.2" stroke-linecap="round"/>
        <!-- 梅花 Plum blossoms -->
        <circle cx="120" cy="70" r="12" fill="#e8a0b0" opacity="0.5"/>
        <circle cx="125" cy="68" r="4" fill="#c9a959" opacity="0.4"/>
        <circle cx="250" cy="110" r="10" fill="#e8a0b0" opacity="0.4"/>
        <circle cx="253" cy="108" r="3.5" fill="#c9a959" opacity="0.35"/>
        <circle cx="130" cy="30" r="9" fill="#e8a0b0" opacity="0.45"/>
        <circle cx="133" cy="28" r="3" fill="#c9a959" opacity="0.35"/>
        <circle cx="290" cy="230" r="8" fill="#e8a0b0" opacity="0.35"/>
        <circle cx="130" cy="120" r="10" fill="#e8a0b0" opacity="0.4"/>
        <circle cx="133" cy="118" r="3.5" fill="#c9a959" opacity="0.3"/>
        <!-- 花蕾 Buds -->
        <ellipse cx="155" cy="95" rx="4" ry="6" fill="#e8a0b0" opacity="0.3" transform="rotate(-20 155 95)"/>
        <ellipse cx="225" cy="135" rx="3.5" ry="5.5" fill="#e8a0b0" opacity="0.25" transform="rotate(10 225 135)"/>
        <ellipse cx="265" cy="250" rx="3" ry="5" fill="#e8a0b0" opacity="0.25" transform="rotate(-15 265 250)"/>
      </svg>

      <!-- 竹 Bamboo (left side, painting style) -->
      <svg class="absolute top-[15%] left-[2%] w-[180px] h-[500px] hidden lg:block" style="opacity: 0.25;" viewBox="0 0 180 500">
        <!-- 竹竿 Bamboo stalk -->
        <path d="M80 500 Q82 400 78 300 Q76 200 80 100 Q82 50 80 0" stroke="#4a6a4a" fill="none" stroke-width="4" stroke-linecap="round"/>
        <!-- 竹节 Nodes -->
        <line x1="72" y1="100" x2="88" y2="100" stroke="#4a6a4a" stroke-width="2.5"/>
        <line x1="72" y1="200" x2="88" y2="200" stroke="#4a6a4a" stroke-width="2.5"/>
        <line x1="72" y1="300" x2="88" y2="300" stroke="#4a6a4a" stroke-width="2.5"/>
        <line x1="72" y1="400" x2="88" y2="400" stroke="#4a6a4a" stroke-width="2.5"/>
        <!-- 竹叶 Leaves -->
        <path d="M78 95 Q50 70 30 50" stroke="#4a6a4a" fill="none" stroke-width="2.5" stroke-linecap="round"/>
        <path d="M78 90 Q55 55 45 25" stroke="#4a6a4a" fill="none" stroke-width="2" stroke-linecap="round"/>
        <path d="M82 105 Q110 80 140 65" stroke="#4a6a4a" fill="none" stroke-width="2.5" stroke-linecap="round"/>
        <path d="M82 195 Q110 170 135 160" stroke="#4a6a4a" fill="none" stroke-width="2" stroke-linecap="round"/>
        <path d="M78 205 Q50 185 25 180" stroke="#4a6a4a" fill="none" stroke-width="2.2" stroke-linecap="round"/>
        <path d="M82 295 Q115 275 145 270" stroke="#4a6a4a" fill="none" stroke-width="2" stroke-linecap="round"/>
        <path d="M78 305 Q45 290 20 295" stroke="#4a6a4a" fill="none" stroke-width="2" stroke-linecap="round"/>
        <path d="M78 395 Q50 375 30 365" stroke="#4a6a4a" fill="none" stroke-width="1.8" stroke-linecap="round"/>
        <path d="M82 405 Q105 390 130 385" stroke="#4a6a4a" fill="none" stroke-width="1.8" stroke-linecap="round"/>
      </svg>

      <!-- 飞鸟 Flying birds (ink style) -->
      <svg class="absolute top-[12%] left-[30%] w-[200px] h-[60px]" style="opacity: 0.25;" viewBox="0 0 200 60">
        <path d="M20 30 Q30 15 45 25" stroke="#4a4a4a" fill="none" stroke-width="1.5" stroke-linecap="round"/>
        <path d="M20 30 Q30 40 45 35" stroke="#4a4a4a" fill="none" stroke-width="1.5" stroke-linecap="round"/>
        <path d="M70 25 Q80 12 93 20" stroke="#4a4a4a" fill="none" stroke-width="1.2" stroke-linecap="round"/>
        <path d="M70 25 Q80 34 93 30" stroke="#4a4a4a" fill="none" stroke-width="1.2" stroke-linecap="round"/>
        <path d="M110 35 Q118 24 128 30" stroke="#4a4a4a" fill="none" stroke-width="1" stroke-linecap="round"/>
        <path d="M110 35 Q118 42 128 38" stroke="#4a4a4a" fill="none" stroke-width="1" stroke-linecap="round"/>
        <path d="M155 20 Q162 12 170 17" stroke="#4a4a4a" fill="none" stroke-width="0.8" stroke-linecap="round"/>
        <path d="M155 20 Q162 26 170 23" stroke="#4a4a4a" fill="none" stroke-width="0.8" stroke-linecap="round"/>
      </svg>

      <!-- 落印 Red seal stamp (bottom-left) -->
      <div class="absolute bottom-[8%] left-[6%] w-14 h-14 hidden lg:block" style="opacity: 0.35;">
        <div class="w-full h-full border-2 border-[#c86050] rounded-sm flex items-center justify-center"
          style="box-shadow: 1px 1px 4px rgba(200,96,80,0.15);">
          <span class="font-brush text-[#c86050] text-lg leading-none">雅</span>
        </div>
      </div>

      <!-- 第二枚印章 Second seal (top area) -->
      <div class="absolute top-[18%] right-[12%] w-10 h-10 hidden lg:block" style="opacity: 0.3;">
        <div class="w-full h-full border-[1.5px] border-[#c86050] rounded-sm flex items-center justify-center">
          <span class="font-brush text-[#c86050] text-sm leading-none">墨</span>
        </div>
      </div>

      <!-- 飘落花瓣 Falling petals (sparse, artistic) -->
      <div class="petal animate-petal-fall" style="left: 8%; animation-delay: 0s; animation-duration: 12s;"></div>
      <div class="petal animate-petal-fall" style="left: 25%; animation-delay: 3s; animation-duration: 15s;"></div>
      <div class="petal animate-petal-fall" style="left: 45%; animation-delay: 6s; animation-duration: 11s;"></div>
      <div class="petal animate-petal-fall" style="left: 65%; animation-delay: 2s; animation-duration: 14s;"></div>
      <div class="petal animate-petal-fall" style="left: 85%; animation-delay: 8s; animation-duration: 13s;"></div>
      <div class="petal-small animate-petal-fall" style="left: 18%; animation-delay: 5s; animation-duration: 16s;"></div>
      <div class="petal-small animate-petal-fall" style="left: 55%; animation-delay: 10s; animation-duration: 14s;"></div>
      <div class="petal-small animate-petal-fall" style="left: 78%; animation-delay: 7s; animation-duration: 17s;"></div>

      <!-- 縦書き装飾 Vertical calligraphy -->
      <div class="absolute right-5 top-1/4 tategaki text-sumi-800/[0.15] font-brush text-xl tracking-[0.6em] select-none hidden xl:block">
        山高水長
      </div>
      <div class="absolute left-5 bottom-1/4 tategaki text-sumi-800/[0.12] font-brush text-xl tracking-[0.6em] select-none hidden xl:block">
        行雲流水
      </div>
    </div>

    <NavBar />

    <main class="flex-grow w-full relative z-10">
      <RouterView v-slot="{ Component }">
        <transition name="page" mode="out-in">
          <component :is="Component" />
        </transition>
      </RouterView>
    </main>

    <BackToTop />

    <!-- Footer -->
    <footer class="relative z-10 mt-24">
      <!-- Ink brush stroke divider -->
      <div class="h-px bg-gradient-to-r from-transparent via-sumi-800/10 to-transparent"></div>

      <div class="relative overflow-hidden" style="background: linear-gradient(180deg, #faf7f2, #f5ece3);">
        <!-- Footer mountain echo -->
        <svg class="absolute bottom-0 left-0 w-full h-24 opacity-[0.025]" viewBox="0 0 1440 100" preserveAspectRatio="none">
          <path d="M0 100 L0 60 Q180 20 360 50 Q540 10 720 40 Q900 15 1080 45 Q1260 20 1440 55 L1440 100Z" fill="#5a6a5a"/>
        </svg>

        <div class="py-14 relative z-10">
          <div class="max-w-5xl mx-auto px-6">
            <div class="flex flex-col md:flex-row items-center justify-between gap-8">
              <div class="flex items-center gap-4">
                <div class="relative">
                  <div class="w-11 h-11 rounded-full bg-gradient-to-br from-sakura/15 to-mizu/15 flex items-center justify-center">
                    <span class="font-brush text-shu text-lg">記</span>
                  </div>
                  <div class="absolute inset-[-3px] rounded-full border border-sumi-800/[0.06]"></div>
                </div>
                <div>
                  <span class="font-kai text-lg font-bold text-sumi-800">GoNow</span>
                  <span class="text-sumi-400 text-xs ml-2 font-kai">数字庭園</span>
                </div>
              </div>

              <div class="flex items-center gap-6 text-sumi-400 text-sm font-kai">
                <RouterLink to="/" class="hover:text-shu transition-colors duration-300">首頁</RouterLink>
                <span class="text-sumi-200">·</span>
                <RouterLink to="/archives" class="hover:text-shu transition-colors duration-300">歸檔</RouterLink>
              </div>
            </div>

            <div class="flex items-center gap-4 my-8">
              <div class="flex-1 h-px bg-gradient-to-r from-transparent via-sumi-800/[0.06] to-transparent"></div>
              <span class="text-shu/30 text-xs">❀</span>
              <div class="flex-1 h-px bg-gradient-to-r from-transparent via-sumi-800/[0.06] to-transparent"></div>
            </div>

            <p class="text-center text-sumi-300 text-xs tracking-widest font-kai">
              &copy; 二〇二五 GoNow Blog · 花開見文 · 水到渠成
            </p>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>
