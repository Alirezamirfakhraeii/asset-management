import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'
import DashboardView from '../views/DashboardView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'login',
      component: LoginView
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: DashboardView,
      // اضافه کردن این بخش برای شناسایی صفحات محافظت شده
      meta: { requiresAuth: true }
    }
  ]
})

// این تابع قبل از هر جابجایی بین صفحات اجرا می‌شود
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')

  if (to.meta.requiresAuth && !token) {
    next('/') // بفرستش به صفحه لاگین
  } else {
    next() // اجازه بده بره به صفحه مقصد
  }
})

export default router
