import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '@/views/LoginView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: LoginView,
    },
    {
      path: '/',
      name: 'DashBoard',
      component: () => import('../views/DashBoardView.vue'),
      children: [
        {
          path: '/dict',
          name: 'Dict',
          component: () => import('../views/system/DictView.vue'),
        },
        {
          path: '/generate',
          name: 'Generate',
          component: () => import('../views/system/GenerateView.vue'),
        },
      ],
    },
  ],
})

// 路由拦截
// router.beforeEach((to, from, next) => {
//   if (to.path === '/login') {
//     next()
//     return
//   }
//   const token = localStorage.getItem('token')
//   if (token) {
//     next()
//     return
//   } else {
//     next('/login')
//     return
//   }
// })

export default router
