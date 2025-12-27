import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import { jwtDecode } from 'jwt-decode';

export const useAuthStore = defineStore('auth', () => {
  const token = ref(localStorage.getItem('token') || null);
  const user = ref(null);

  if (token.value) {
    try {
      const decoded = jwtDecode(token.value);
      user.value = { id: decoded.user_id, role: decoded.role };
    } catch (e) {
      token.value = null;
      localStorage.removeItem('token');
    }
  }

  const isAuthenticated = computed(() => !!token.value);
  const isAdmin = computed(() => user.value?.role === 'admin');

  function setToken(newToken) {
    token.value = newToken;
    localStorage.setItem('token', newToken);
    try {
      const decoded = jwtDecode(newToken);
      user.value = { id: decoded.user_id, role: decoded.role };
    } catch (e) {
      console.error('Invalid token');
    }
  }

  function logout() {
    token.value = null;
    user.value = null;
    localStorage.removeItem('token');
  }

  return { token, user, isAuthenticated, isAdmin, setToken, logout };
});
