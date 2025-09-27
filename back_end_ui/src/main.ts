import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

import 'ant-design-vue/dist/reset.css'
import '@/assets/main.css'

const app = createApp(App)

const pinia = createPinia()
//持久化插件
pinia.use(piniaPluginPersistedstate)
app.use(pinia)
app.use(router)

app.mount('#app')
