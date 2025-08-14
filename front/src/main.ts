import { createApp } from 'vue'
import { createPinia } from 'pinia'
import './style.css'
import App from './App.vue'
import router from './router'
import naive from 'naive-ui'
import { useAuthStore } from './stores/auth.store'

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)

const authStore = useAuthStore()
authStore.loadFormStorage()

app.use(router)
app.use(naive)
app.mount('#app')