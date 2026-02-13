import { createRouter, createWebHistory } from 'vue-router'
// Asegúrate de poner el .vue para que TypeScript no tire error
import Dashboard from '../views/Dashboard.vue' 
import AdminView from '../views/AdminView.vue'

const routes = [
  {
    path: '/',
    name: 'Dashboard',
    component: Dashboard
  },
  {
    path: '/soberano',
    name: 'Admin',
    component: AdminView
  }
]

const router = createRouter({
  history: createWebHistory(), // Esto permite usar URLs limpias sin el #
  routes
})

export default router