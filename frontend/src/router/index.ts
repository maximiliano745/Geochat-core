import { createRouter, createWebHistory } from 'vue-router'
// Importamos los componentes. 
// Nota: Verifica que los nombres de los archivos en src/views sean EXACTAMENTE estos.
import Dashboard from '../views/Dashboard.vue' 
import AdminView from '../views/AdminView.vue'

const routes = [
  {
    path: '/',
    name: 'Dashboard',
    component: Dashboard
  },
  {
    // Cambiamos el path a /admin para que coincida con tus router-links
    path: '/admin', 
    name: 'Admin',
    component: AdminView
  }
]

const router = createRouter({
  // En Codespaces/Web, esto asegura que las rutas funcionen correctamente
  history: createWebHistory(import.meta.env.BASE_URL), 
  routes
})

export default router