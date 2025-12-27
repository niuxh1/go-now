import axios from 'axios'

const api = axios.create({
  // 在生产环境(Nginx)下使用相对路径 '/api'，开发环境使用 localhost:8080
  baseURL: import.meta.env.PROD ? '/api' : 'http://localhost:8080/api',
  timeout: 5000
})

api.interceptors.request.use(config => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export const fetchArticles = (category) => api.get('/articles', { params: { category } })
export const fetchArticle = (id) => api.get(`/articles/${id}`)
export const createArticle = (data) => api.post('/articles', data)
export const deleteArticle = (id) => api.delete(`/articles/${id}`)
export const fetchComments = (articleId) => api.get(`/articles/${articleId}/comments`)
export const postComment = (articleId, content) => api.post(`/articles/${articleId}/comments`, { content })
export const fetchTags = () => api.get('/tags')
export const fetchSettings = () => api.get('/settings')
export const updateSettings = (data) => api.put('/settings', data)

export default api