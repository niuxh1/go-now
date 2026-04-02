<script setup>
import { ref, onMounted } from 'vue'
import { ChevronDown } from 'lucide-vue-next'

const text = ref('')
const showContent = ref(false)
const poetryAuthor = ref('')

const poems = [
  { text: '大漠孤烟直，长河落日圆', author: '王维《使至塞上》' },
  { text: '落霞与孤鹜齐飞，秋水共长天一色', author: '王勃《滕王阁序》' },
  { text: '海上生明月，天涯共此时', author: '张九龄《望月怀远》' },
  { text: '春风得意马蹄疾，一日看尽长安花', author: '孟郊《登科后》' },
  { text: '行到水穷处，坐看云起时', author: '王维《终南别业》' },
  { text: '会当凌绝顶，一览众山小', author: '杜甫《望岳》' },
  { text: '疏影横斜水清浅，暗香浮动月黄昏', author: '林逋《山园小梅》' },
  { text: '山重水复疑无路，柳暗花明又一村', author: '陆游《游山西村》' },
  { text: '采菊东篱下，悠然见南山', author: '陶渊明《饮酒》' },
  { text: '乱花渐欲迷人眼，浅草才能没马蹄', author: '白居易《钱塘湖春行》' },
  { text: '但愿人长久，千里共婵娟', author: '苏轼《水调歌头》' },
  { text: '竹外桃花三两枝，春江水暖鸭先知', author: '苏轼《惠崇春江晚景》' },
  { text: '长风破浪会有时，直挂云帆济沧海', author: '李白《行路难》' },
  { text: '空山新雨后，天气晚来秋', author: '王维《山居秋暝》' },
  { text: '随风潜入夜，润物细无声', author: '杜甫《春夜喜雨》' },
  { text: '接天莲叶无穷碧，映日荷花别样红', author: '杨万里《晓出净慈寺》' },
  { text: '日出江花红胜火，春来江水绿如蓝', author: '白居易《忆江南》' },
  { text: '千里莺啼绿映红，水村山郭酒旗风', author: '杜牧《江南春》' },
  { text: '醉后不知天在水，满船清梦压星河', author: '唐温如《题龙阳县青草湖》' },
  { text: '人间有味是清欢', author: '苏轼《浣溪沙》' },
  { text: '桃花潭水深千尺，不及汪伦送我情', author: '李白《赠汪伦》' },
  { text: '两岸猿声啼不住，轻舟已过万重山', author: '李白《早发白帝城》' },
  { text: '何当共剪西窗烛，却话巴山夜雨时', author: '李商隐《夜雨寄北》' },
  { text: '月落乌啼霜满天，江枫渔火对愁眠', author: '张继《枫桥夜泊》' },
  { text: '沉舟侧畔千帆过，病树前头万木春', author: '刘禹锡《酬乐天》' },
  { text: '此去经年，应是良辰好景虚设', author: '柳永《雨霖铃》' },
  { text: '水光潋滟晴方好，山色空蒙雨亦奇', author: '苏轼《饮湖上初晴后雨》' },
  { text: '小荷才露尖尖角，早有蜻蜓立上头', author: '杨万里《小池》' },
  { text: '春色满园关不住，一枝红杏出墙来', author: '叶绍翁《游园不值》' },
  { text: '忽如一夜春风来，千树万树梨花开', author: '岑参《白雪歌》' },
  { text: '人面不知何处去，桃花依旧笑春风', author: '崔护《题都城南庄》' },
]

const getDailyPoem = () => {
  const today = new Date()
  const dayIndex = (today.getFullYear() * 366 + today.getMonth() * 31 + today.getDate()) % poems.length
  return poems[dayIndex]
}

const typeText = async () => {
  const poem = getDailyPoem()
  poetryAuthor.value = poem.author
  const fullText = poem.text
  await new Promise(r => setTimeout(r, 1000))
  for (let i = 0; i <= fullText.length; i++) {
    text.value = fullText.slice(0, i)
    await new Promise(r => setTimeout(r, 100))
  }
}

const scrollDown = () => {
  window.scrollTo({ top: window.innerHeight, behavior: 'smooth' })
}

onMounted(() => { showContent.value = true; typeText() })
</script>

<template>
  <div class="relative h-screen w-full overflow-hidden">
    <!-- 底色：暖宣纸渐变 -->
    <div class="absolute inset-0" style="background: linear-gradient(170deg, #fdf9f5 0%, #f8f0e6 25%, #f3e8db 50%, #f0e4d6 70%, #f5ece3 90%, #faf7f2 100%);"></div>

    <!-- 水墨山水 Ink landscape (Hero特有) -->
    <svg class="absolute bottom-0 left-0 w-full" style="height: 50vh; opacity: 0.15;" viewBox="0 0 1440 600" preserveAspectRatio="none">
      <path d="M0 600 L0 350 Q100 180 200 280 Q300 140 400 240 Q500 100 600 220 Q700 80 800 200 Q900 120 1000 220 Q1100 100 1200 240 Q1300 160 1440 280 L1440 600Z"
        fill="url(#heroMtnFar)"/>
      <path d="M0 600 L0 400 Q200 280 400 360 Q550 260 700 340 Q850 240 1000 330 Q1150 270 1300 360 Q1380 320 1440 380 L1440 600Z"
        fill="url(#heroMtnMid)"/>
      <path d="M0 600 L0 470 Q200 380 400 430 Q600 370 800 420 Q1000 360 1200 430 Q1350 390 1440 450 L1440 600Z"
        fill="url(#heroMtnNear)"/>
      <defs>
        <linearGradient id="heroMtnFar" x1="0" y1="0" x2="0" y2="1">
          <stop offset="0%" stop-color="#6a8a9a" stop-opacity="0.7"/>
          <stop offset="100%" stop-color="#6a8a9a" stop-opacity="0.1"/>
        </linearGradient>
        <linearGradient id="heroMtnMid" x1="0" y1="0" x2="0" y2="1">
          <stop offset="0%" stop-color="#7a8a7a" stop-opacity="0.5"/>
          <stop offset="100%" stop-color="#7a8a7a" stop-opacity="0.08"/>
        </linearGradient>
        <linearGradient id="heroMtnNear" x1="0" y1="0" x2="0" y2="1">
          <stop offset="0%" stop-color="#5a6a5a" stop-opacity="0.3"/>
          <stop offset="100%" stop-color="#5a6a5a" stop-opacity="0.03"/>
        </linearGradient>
      </defs>
    </svg>

    <!-- 水彩晕染 -->
    <div class="absolute -top-32 -right-32 w-[800px] h-[800px] rounded-full opacity-[0.12] animate-drift"
      style="background: radial-gradient(circle, #e8a0b0 0%, transparent 60%);"></div>
    <div class="absolute top-[25%] -left-40 w-[700px] h-[700px] rounded-full opacity-[0.09] animate-drift"
      style="background: radial-gradient(circle, #6aacb8 0%, transparent 55%); animation-delay: -8s;"></div>
    <div class="absolute bottom-[10%] right-[10%] w-[500px] h-[500px] rounded-full opacity-[0.07] animate-drift"
      style="background: radial-gradient(circle, #9a88b8 0%, transparent 55%); animation-delay: -15s;"></div>
    <div class="absolute top-[60%] left-[30%] w-[400px] h-[400px] rounded-full opacity-[0.06]"
      style="background: radial-gradient(circle, #c9a959 0%, transparent 55%);"></div>

    <!-- 云雾 Mist band -->
    <div class="absolute top-[45%] left-0 right-0 h-[160px] opacity-[0.3]"
      style="background: linear-gradient(180deg, transparent 0%, #faf7f2 35%, #faf7f2 65%, transparent 100%);"></div>

    <!-- 宣纸纹理 -->
    <div class="absolute inset-0 opacity-[0.03]"
      style="background-image: url(&quot;data:image/svg+xml,%3Csvg viewBox='0 0 256 256' xmlns='http://www.w3.org/2000/svg'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.85' numOctaves='4' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23n)'/%3E%3C/svg%3E&quot;);"></div>

    <!-- 大背景漢字 -->
    <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-[55%] select-none pointer-events-none">
      <span class="font-brush text-[28vw] leading-none text-sumi-800/[0.06]">墨</span>
    </div>

    <!-- 枯枝梅花 Plum branch (hero specific, right side) -->
    <svg class="absolute top-[5%] right-[5%] w-[250px] h-[350px] hidden md:block" style="opacity: 0.3;" viewBox="0 0 250 350">
      <path d="M230 350 Q215 280 200 220 Q185 160 165 120 Q145 80 120 60 Q100 45 80 40"
        stroke="#5a3a2a" fill="none" stroke-width="2.5" stroke-linecap="round"/>
      <path d="M165 120 Q185 90 210 75" stroke="#5a3a2a" fill="none" stroke-width="1.8" stroke-linecap="round"/>
      <path d="M120 60 Q100 30 90 10" stroke="#5a3a2a" fill="none" stroke-width="1.5" stroke-linecap="round"/>
      <path d="M145 80 Q120 75 100 85" stroke="#5a3a2a" fill="none" stroke-width="1.2" stroke-linecap="round"/>
      <circle cx="80" cy="40" r="10" fill="#e8a0b0" opacity="0.5"/>
      <circle cx="83" cy="38" r="3.5" fill="#c9a959" opacity="0.4"/>
      <circle cx="210" cy="75" r="8" fill="#e8a0b0" opacity="0.4"/>
      <circle cx="213" cy="73" r="3" fill="#c9a959" opacity="0.35"/>
      <circle cx="90" cy="10" r="7" fill="#e8a0b0" opacity="0.4"/>
      <circle cx="100" cy="85" r="8" fill="#e8a0b0" opacity="0.35"/>
      <ellipse cx="130" cy="70" rx="3.5" ry="5.5" fill="#e8a0b0" opacity="0.3" transform="rotate(-15 130 70)"/>
      <ellipse cx="190" cy="100" rx="3" ry="5" fill="#e8a0b0" opacity="0.25" transform="rotate(10 190 100)"/>
    </svg>

    <!-- 飞鸟 Flying birds -->
    <svg class="absolute top-[15%] left-[25%] w-[160px] h-[50px]" style="opacity: 0.25;" viewBox="0 0 160 50">
      <path d="M15 25 Q25 12 38 20" stroke="#4a4a4a" fill="none" stroke-width="1.3" stroke-linecap="round"/>
      <path d="M15 25 Q25 34 38 30" stroke="#4a4a4a" fill="none" stroke-width="1.3" stroke-linecap="round"/>
      <path d="M55 20 Q63 10 73 16" stroke="#4a4a4a" fill="none" stroke-width="1.1" stroke-linecap="round"/>
      <path d="M55 20 Q63 28 73 24" stroke="#4a4a4a" fill="none" stroke-width="1.1" stroke-linecap="round"/>
      <path d="M90 28 Q96 20 104 25" stroke="#4a4a4a" fill="none" stroke-width="0.9" stroke-linecap="round"/>
      <path d="M90 28 Q96 34 104 31" stroke="#4a4a4a" fill="none" stroke-width="0.9" stroke-linecap="round"/>
    </svg>

    <!-- 落印 Red seal -->
    <div class="absolute bottom-[12%] left-[8%] w-12 h-12 hidden md:block" style="opacity: 0.35;">
      <div class="w-full h-full border-2 border-[#c86050] rounded-sm flex items-center justify-center"
        style="box-shadow: 1px 1px 3px rgba(200,96,80,0.12);">
        <span class="font-brush text-[#c86050] text-base leading-none">誌</span>
      </div>
    </div>

    <!-- 花瓣 Petals (sparse) -->
    <div class="petal animate-petal-fall" style="left: 10%; animation-delay: 0s; animation-duration: 14s;"></div>
    <div class="petal animate-petal-fall" style="left: 35%; animation-delay: 4s; animation-duration: 12s;"></div>
    <div class="petal animate-petal-fall" style="left: 60%; animation-delay: 8s; animation-duration: 15s;"></div>
    <div class="petal animate-petal-fall" style="left: 82%; animation-delay: 2s; animation-duration: 13s;"></div>
    <div class="petal-small animate-petal-fall" style="left: 22%; animation-delay: 6s; animation-duration: 16s;"></div>
    <div class="petal-small animate-petal-fall" style="left: 70%; animation-delay: 10s; animation-duration: 14s;"></div>

    <!-- ====== CONTENT ====== -->
    <div class="relative z-10 h-full flex flex-col items-center justify-center px-6 text-center">
      <transition name="hero">
        <div v-if="showContent" class="space-y-8">
          <!-- 装飾線 -->
          <div class="flex items-center justify-center gap-3 animate-fade-in" style="animation-delay: 0.2s;">
            <div class="w-20 h-px bg-gradient-to-r from-transparent to-sumi-800/10"></div>
            <span class="text-sumi-500/60 text-[10px] font-kai tracking-[0.6em]">數字庭園</span>
            <div class="w-20 h-px bg-gradient-to-l from-transparent to-sumi-800/10"></div>
          </div>

          <!-- 大標題 -->
          <div class="animate-ink-spread">
            <h1 class="font-brush text-7xl md:text-8xl lg:text-9xl text-sumi-800 leading-none tracking-wide"
              style="text-shadow: 0 2px 30px rgba(90,58,42,0.08);">
              GoNow
            </h1>
            <div class="flex justify-center items-center mt-6 gap-3">
              <div class="w-10 h-px bg-gradient-to-r from-transparent to-sumi-800/10"></div>
              <div class="w-10 h-10 rounded-full border border-shu/25 flex items-center justify-center">
                <span class="font-brush text-shu text-lg">誌</span>
              </div>
              <div class="w-10 h-px bg-gradient-to-l from-transparent to-sumi-800/10"></div>
            </div>
          </div>

          <!-- 每日古诗 -->
          <div class="animate-fade-in" style="animation-delay: 0.8s;">
            <div class="relative inline-block">
              <div class="absolute -left-5 -top-1 text-sumi-300/40 text-xl font-kai select-none">「</div>
              <div class="absolute -right-5 -bottom-1 text-sumi-300/40 text-xl font-kai select-none">」</div>
              <div class="text-lg md:text-xl font-kai text-sumi-600 min-h-[2rem] px-3">
                {{ text }}<span class="typed-cursor">|</span>
              </div>
            </div>
            <div v-if="poetryAuthor && text.length > 5" class="text-xs text-sumi-400 font-kai mt-3 animate-fade-in" style="animation-delay: 4s;">
              —— {{ poetryAuthor }}
            </div>
          </div>

          <!-- 分隔 -->
          <div class="flex items-center justify-center gap-3 animate-fade-in" style="animation-delay: 1s;">
            <div class="w-12 h-px bg-sumi-800/[0.06]"></div>
            <span class="text-shu/25 text-xs">❀</span>
            <div class="w-12 h-px bg-sumi-800/[0.06]"></div>
          </div>

          <!-- 导航关键词 -->
          <div class="flex items-center justify-center gap-6 animate-fade-in" style="animation-delay: 1.2s;">
            <span class="font-kai text-sm text-sumi-400 hover:text-shu transition-colors duration-500 cursor-default">技術</span>
            <span class="text-sumi-200">·</span>
            <span class="font-kai text-sm text-sumi-400 hover:text-shu transition-colors duration-500 cursor-default">生活</span>
            <span class="text-sumi-200">·</span>
            <span class="font-kai text-sm text-sumi-400 hover:text-shu transition-colors duration-500 cursor-default">思考</span>
          </div>
        </div>
      </transition>
    </div>

    <!-- 下スクロール -->
    <div class="absolute bottom-10 left-1/2 -translate-x-1/2 cursor-pointer z-20 group" @click="scrollDown">
      <div class="flex flex-col items-center gap-2 animate-fade-in" style="animation-delay: 2s;">
        <span class="text-sumi-400 text-[9px] tracking-[0.4em] font-kai group-hover:text-shu transition-colors">向下</span>
        <ChevronDown size="16" class="text-sumi-400 group-hover:text-shu animate-float transition-colors" />
      </div>
    </div>

    <!-- 底部过渡 -->
    <div class="absolute bottom-0 left-0 right-0 h-32 pointer-events-none">
      <div class="absolute inset-0" style="background: linear-gradient(to bottom, transparent 0%, rgba(250,247,242,0.5) 40%, #faf7f2 100%);"></div>
    </div>
  </div>
</template>

<style scoped>
.hero-enter-active { transition: all 1s cubic-bezier(0.22, 1, 0.36, 1); }
.hero-enter-from { opacity: 0; transform: translateY(30px); }
</style>
