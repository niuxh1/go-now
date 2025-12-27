import { createRouter, createWebHistory } from 'vue-router'
import { jwtDecode } from 'jwt-decode'
import HomeView from '../views/HomeView.vue'
import ArticleView from '../views/ArticleView.vue'
import ArchiveView from '../views/ArchiveView.vue'
import LoginView from '../views/LoginView.vue'
import RegisterView from '../views/RegisterView.vue'
import ResetPasswordView from '../views/ResetPasswordView.vue'
import EditorView from '../views/EditorView.vue'
import SettingsView from '../views/SettingsView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/editor',
    name: 'editor',
    component: EditorView,
    meta: { requiresAuth: true, adminOnly: true }
  },
  {
    path: '/settings',
    name: 'settings',
    component: SettingsView,
    meta: { requiresAuth: true, adminOnly: true }
  },
  {
    path: '/archives',
    name: 'archives',
    component: ArchiveView
  },
  {
    path: '/articles/:id',
    name: 'article',
    component: ArticleView,
    props: true
  },
  {
    path: '/login',
    name: 'login',
    component: LoginView
  },
  {
    path: '/register',
    name: 'register',
    component: RegisterView
  },
  {
    path: '/reset-password',
    name: 'reset-password',
    component: ResetPasswordView
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior() {
    return { top: 0 }
  }
})

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  let userRole = null

  if (token) {
    try {
      const decoded = jwtDecode(token)
      userRole = decoded.role
    } catch (e) {
      // Invalid token
    }
  }

  if (to.meta.requiresAuth && !token) {
    next('/login')
  } else if (to.meta.adminOnly && userRole !== 'admin') {
    next('/') // Redirect non-admins to home
  } else {
    next()
  }
})

export default router
