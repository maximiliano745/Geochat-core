import { createApp } from 'vue'
import App from './App.vue'
import router from './router' // <--- Esto busca automáticamente /index.ts

const app = createApp(App)

app.use(router) // <--- Aquí activamos el motor de rutas
app.mount('#app')