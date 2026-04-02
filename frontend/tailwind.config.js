/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        // 紙 - Paper tones (light base)
        kami: {
          DEFAULT: '#faf7f2',
          warm: '#f5efe4',
          cream: '#ede5d8',
          dark: '#d9cfc0',
        },
        // 墨 - Ink (text tones)
        sumi: {
          950: '#1a1815',
          900: '#2a2725',
          800: '#3a3535',
          700: '#4a4545',
          600: '#5a5555',
          500: '#7a7570',
          400: '#9a9590',
          300: '#b5b0a8',
          200: '#d0ccc5',
          100: '#e8e4dd',
        },
        // 桜 - Sakura pink
        sakura: {
          DEFAULT: '#e8a0b0',
          light: '#f5c8d4',
          pale: '#fde8ee',
          deep: '#c97888',
          muted: 'rgba(232, 160, 176, 0.12)',
        },
        // 碧 - Azure water
        mizu: {
          DEFAULT: '#6aacb8',
          light: '#8acad5',
          pale: '#daf0f4',
          deep: '#4a8c98',
          muted: 'rgba(106, 172, 184, 0.10)',
        },
        // 翠 - Bamboo green
        take: {
          DEFAULT: '#7aaa80',
          light: '#a0cca5',
          pale: '#e0f0e2',
          deep: '#5a8a60',
          muted: 'rgba(122, 170, 128, 0.10)',
        },
        // 藤 - Wisteria purple
        fuji: {
          DEFAULT: '#9a88b8',
          light: '#baacd0',
          pale: '#e8e0f0',
          deep: '#7a68a0',
        },
        // 朱 - Vermillion (accent)
        shu: {
          DEFAULT: '#c86050',
          light: '#e08070',
          pale: '#f8e0dc',
          deep: '#a04840',
        },
        // 金 - Gold
        kin: {
          DEFAULT: '#c9a959',
          light: '#e0c878',
          muted: 'rgba(201, 169, 89, 0.12)',
        },
      },
      fontFamily: {
        kai: ['"Noto Serif SC"', '"ZCOOL XiaoWei"', 'serif'],
        song: ['"Noto Sans SC"', 'sans-serif'],
        brush: ['"Ma Shan Zheng"', 'cursive'],
        mono: ['monospace'],
      },
      animation: {
        'fade-up': 'fadeUp 0.6s cubic-bezier(0.22, 1, 0.36, 1)',
        'fade-in': 'fadeIn 0.5s ease-out',
        'slide-down': 'slideDown 0.4s cubic-bezier(0.22, 1, 0.36, 1)',
        'float': 'float 6s ease-in-out infinite',
        'float-slow': 'float 10s ease-in-out infinite',
        'drift': 'drift 25s ease-in-out infinite',
        'sway': 'sway 4s ease-in-out infinite',
        'petal-fall': 'petalFall 8s ease-in-out infinite',
        'ink-spread': 'inkSpread 1.2s cubic-bezier(0.22, 1, 0.36, 1)',
        'ripple': 'ripple 3s ease-out infinite',
        'float-gentle': 'floatGentle 6s ease-in-out infinite',
        'shimmer': 'shimmer 2s ease-in-out infinite',
        'ink-drop': 'inkDrop 0.6s cubic-bezier(0.22, 1, 0.36, 1)',
        'pulse-soft': 'pulseSoft 3s ease-in-out infinite',
        'slide-up': 'slideUp 0.5s cubic-bezier(0.22, 1, 0.36, 1)',
      },
      keyframes: {
        fadeUp: {
          '0%': { opacity: '0', transform: 'translateY(24px)' },
          '100%': { opacity: '1', transform: 'translateY(0)' },
        },
        fadeIn: {
          '0%': { opacity: '0' },
          '100%': { opacity: '1' },
        },
        slideDown: {
          '0%': { opacity: '0', transform: 'translateY(-10px)' },
          '100%': { opacity: '1', transform: 'translateY(0)' },
        },
        float: {
          '0%, 100%': { transform: 'translateY(0)' },
          '50%': { transform: 'translateY(-8px)' },
        },
        drift: {
          '0%, 100%': { transform: 'translate(0, 0) rotate(0deg)' },
          '25%': { transform: 'translate(30px, -20px) rotate(3deg)' },
          '50%': { transform: 'translate(-10px, 15px) rotate(-2deg)' },
          '75%': { transform: 'translate(20px, 5px) rotate(1deg)' },
        },
        sway: {
          '0%, 100%': { transform: 'rotate(-3deg)' },
          '50%': { transform: 'rotate(3deg)' },
        },
        petalFall: {
          '0%': { transform: 'translate(0, -20px) rotate(0deg)', opacity: '0' },
          '10%': { opacity: '1' },
          '90%': { opacity: '1' },
          '100%': { transform: 'translate(40px, 100vh) rotate(360deg)', opacity: '0' },
        },
        inkSpread: {
          '0%': { opacity: '0', transform: 'scale(0.8)', filter: 'blur(10px)' },
          '100%': { opacity: '1', transform: 'scale(1)', filter: 'blur(0)' },
        },
        ripple: {
          '0%': { transform: 'scale(1)', opacity: '0.3' },
          '100%': { transform: 'scale(2.5)', opacity: '0' },
        },
        floatGentle: {
          '0%, 100%': { transform: 'translateY(0)' },
          '50%': { transform: 'translateY(-6px)' },
        },
        shimmer: {
          '0%': { backgroundPosition: '200% 0' },
          '100%': { backgroundPosition: '-200% 0' },
        },
        inkDrop: {
          '0%': { transform: 'scale(0)', opacity: '0' },
          '60%': { opacity: '1' },
          '100%': { transform: 'scale(1)', opacity: '1' },
        },
        pulseSoft: {
          '0%, 100%': { opacity: '1' },
          '50%': { opacity: '0.7' },
        },
        slideUp: {
          '0%': { opacity: '0', transform: 'translateY(16px)' },
          '100%': { opacity: '1', transform: 'translateY(0)' },
        },
      },
      typography: {
        DEFAULT: {
          css: {
            '--tw-prose-body': '#4a4545',
            '--tw-prose-headings': '#2a2725',
            '--tw-prose-links': '#4a8c98',
            '--tw-prose-bold': '#2a2725',
            '--tw-prose-code': '#c86050',
            '--tw-prose-quotes': '#7a7570',
            '--tw-prose-quote-borders': '#e8a0b0',
            '--tw-prose-pre-bg': '#2a2725',
            '--tw-prose-pre-code': '#d0ccc5',
            '--tw-prose-hr': '#e8e4dd',
          },
        },
      },
    },
  },
  plugins: [
    require('@tailwindcss/typography'),
  ],
}
