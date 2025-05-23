// src/main.ts
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { VueQueryPlugin } from "@tanstack/vue-query"
import App from './App.vue'
import router from './router'
import './assets/main.css' // Tailwind CSS / global styles
import { Toaster } from 'vue-sonner'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(VueQueryPlugin)

// Register vue-sonner Toaster globally or use it in App.vue
app.component('Toaster', Toaster)


app.mount('#app')