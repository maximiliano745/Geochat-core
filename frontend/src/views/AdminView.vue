//const API_BASE = 'https://symmetrical-acorn-5xgw65wwgr6f4wrx-8080.app.github.dev'

<template>
  <div class="admin-view">
    <header class="admin-header">
      <h2>🔐 Panel del Soberano</h2>
      <router-link to="/" class="btn-back">Volver al Dashboard</router-link>
    </header>

    <div class="admin-grid">
      <ControlSoberano />

      <section class="card logic-card">
        <h3>Próxima Evolución: {{ proximoHito }}</h3>
        <p>La IA 5 propone inyectar el código de <strong>Geocercas</strong>.</p>
        <div class="code-preview">
          <code>// backend/internal/geofence/model.go</code>
        </div>
      </section>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// @ts-ignore - Esto quita el error de "tipo any"
import ControlSoberano from '../components/admin/ControlSoberano.vue'

const proximoHito = ref('Módulo Geocercas')
</script>

<style scoped>
.admin-view {
  min-height: 100vh;
  background: #020617;
  color: #f8f9fa;
  padding: 40px;
  font-family: 'Space Mono', monospace;
}
.admin-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 30px; }
.btn-back { color: #10b981; text-decoration: none; border: 1px solid #10b981; padding: 5px 15px; border-radius: 8px; }
.admin-grid { display: grid; grid-template-columns: 1fr 1fr; gap: 20px; }
.code-preview { background: #000; padding: 10px; border-radius: 8px; font-size: 0.8rem; color: #10b981; }

/* Estilo para que se vea como el dashboard principal */
.card {
  background: #0f172a;
  padding: 25px;
  border-radius: 24px;
  border: 1px solid #1e293b;
}
</style>