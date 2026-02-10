<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import axios from 'axios'

// Contrato con el Backend (Soberanía de datos)
interface StatsResponse {
  saldo: number;
  nombreNodo: string;
  mensajes: string[];
}

// URL de tu Codespace (Asegúrate de que el puerto 8080 sea PUBLIC en la pestaña Ports)
const API_BASE = 'https://symmetrical-acorn-5xgw65wwgr6f4wrx-8080.app.github.dev'

const saldoOro = ref<number>(0)
const nombreNodo = ref<string>('Buscando señal del Pueblo...')
const mensajes = ref<string[]>([])
const modoTesla = ref<boolean>(false)
const cargando = ref<boolean>(true)
const vaultLocked = ref<boolean>(true) // Estado de la Bóveda E2E

// Computed para calcular el 15% simbólico del Modo Tesla [cite: 2026-02-09]
const aporteComunitario = computed(() => (saldoOro.value * 0.15).toFixed(4))

const actualizarDatos = async () => {
  try {
    cargando.value = true
    // Llamada al corazón de Go
    const res = await axios.get<StatsResponse>(`${API_BASE}/api/stats`)
    
    saldoOro.value = res.data.saldo
    nombreNodo.value = res.data.nombreNodo
    mensajes.value = res.data.mensajes
    vaultLocked.value = false // Si el nodo responde, la bóveda se sincroniza
  } catch (error) {
    console.error("⚠️ Error de conexión soberana:", error)
    nombreNodo.value = "Nodo Desconectado - Verificá el Backend en Go"
    vaultLocked.value = true
  } finally {
    cargando.value = false
  }
}

const toggleTesla = () => {
  modoTesla.value = !modoTesla.value
  // Notificación interna [cite: 2026-02-02]
  const nuevoMsg = modoTesla.value 
    ? "✨ Modo Tesla Activo: Compartiendo red y energía." 
    : "⏸️ Modo Tesla en espera.";
  mensajes.value.unshift(nuevoMsg);
}

onMounted(actualizarDatos)
</script>

<template>
  <div class="dashboard" :class="{ 'tesla-active': modoTesla }">
    <header class="header">
      <div class="logo-area">
        <h1>🌍 GeoChat <span class="core-tag">CORE</span></h1>
        <p class="subtitle">Soberanía Digital 60/40</p>
      </div>
      <div class="node-status-box">
        <span class="status-dot" :class="{ 'online': !cargando }"></span>
        <p class="node-name">{{ nombreNodo }}</p>
      </div>
    </header>

    <main class="content">
      <section class="card stats-card">
        <div class="card-header">
          <h3>Reserva de Oro</h3>
          <span class="vault-badge" :class="{ 'locked': vaultLocked }">
            {{ vaultLocked ? '🔒 VAULT LOCKED' : '🔓 E2E SECURE' }}
          </span>
        </div>
        <div class="gold-value">
          <span class="currency">PAXG</span>
          <span class="amount">{{ saldoOro.toFixed(4) }}</span>
        </div>
        <div class="trust-indicator">
          <p>Confianza del Pueblo: 100%</p>
          <div class="bar"><div class="fill" style="width: 100%"></div></div>
        </div>
      </section>

      <section class="card tesla-card">
        <h3>Modo Tesla</h3>
        <p class="description">Al activar, negocias un 15% de beneficio para el pueblo [cite: 2026-02-09].</p>
        
        <div v-if="modoTesla" class="community-impact">
          <small>Impacto estimado:</small>
          <strong>+{{ aporteComunitario }} PAXG</strong>
        </div>

        <button 
          @click="toggleTesla" 
          :class="modoTesla ? 'btn-tesla-on' : 'btn-tesla-off'"
        >
          {{ modoTesla ? 'FILANTROPÍA ACTIVA' : 'ACTIVAR FILANTROPÍA' }}
        </button>
      </section>

      <section class="card messages-card">
        <h3>Registro del Sistema (MIA)</h3>
        <div class="log-container">
          <ul v-if="mensajes.length">
            <li v-for="(msg, i) in mensajes" :key="i" class="message-item">
              <span class="timestamp">[{{ new Date().toLocaleTimeString() }}]</span> {{ msg }}
            </li>
          </ul>
          <p v-else class="empty-msg">Escaneando red en busca de señales...</p>
        </div>
      </section>
    </main>

    <footer class="footer">
      <div class="footer-grid">
        <p><strong>Estructura:</strong> 60% Operativo | 40% Capital</p>
        <p class="legal">La IA negocia con el usuario. El usuario es el dueño.</p>
      </div>
    </footer>
  </div>
</template>

<style scoped>
/* FUENTES Y BASE */
@import url('https://fonts.googleapis.com/css2?family=Space+Mono:wght@400;700&family=Inter:wght@300;600&display=swap');

.dashboard {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  padding: 40px;
  transition: all 0.5s cubic-bezier(0.4, 0, 0.2, 1);
  background-color: #f1f5f9;
  font-family: 'Inter', sans-serif;
  color: #1e293b;
}

/* MODO TESLA - TRANSFORMACIÓN VISUAL */
.tesla-active {
  background-color: #020617 !important;
  color: #f8f9fa !important;
  background-image: radial-gradient(circle at top right, #1e293b, #020617);
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 40px;
  border-bottom: 1px solid #e2e8f0;
  padding-bottom: 20px;
}

.tesla-active .header { border-color: #1e293b; }

.core-tag {
  font-size: 0.8rem;
  background: #d97706;
  color: white;
  padding: 2px 8px;
  border-radius: 4px;
  vertical-align: middle;
}

.node-status-box {
  background: white;
  padding: 10px 20px;
  border-radius: 50px;
  display: flex;
  align-items: center;
  gap: 12px;
  box-shadow: 0 4px 6px -1px rgba(0,0,0,0.1);
}

.tesla-active .node-status-box { background: #1e293b; color: white; }

.status-dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: #ef4444;
}

.status-dot.online {
  background: #10b981;
  box-shadow: 0 0 12px #10b981;
}

/* CARTAS */
.content {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
  gap: 25px;
  flex: 1;
}

.card {
  background: white;
  padding: 25px;
  border-radius: 24px;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.05);
  transition: transform 0.3s ease;
}

.tesla-active .card {
  background: #0f172a;
  border: 1px solid #1e293b;
  color: white;
}

.card:hover { transform: translateY(-5px); }

/* ORO Y BÓVEDA */
.gold-value {
  margin: 20px 0;
  font-family: 'Space Mono', monospace;
}

.amount {
  font-size: 3rem;
  font-weight: bold;
  color: #d97706;
}

.vault-badge {
  font-size: 0.7rem;
  padding: 4px 10px;
  border-radius: 6px;
  background: #dcfce7;
  color: #166534;
}

.vault-badge.locked { background: #fee2e2; color: #991b1b; }

/* BOTONES */
button {
  width: 100%;
  padding: 18px;
  border-radius: 16px;
  font-weight: bold;
  cursor: pointer;
  border: none;
  transition: all 0.3s ease;
  font-family: 'Space Mono', monospace;
}

.btn-tesla-off { background: #1e293b; color: white; }
.btn-tesla-on {
  background: #10b981;
  color: #020617;
  box-shadow: 0 0 20px rgba(16, 185, 129, 0.4);
  animation: pulse 2s infinite;
}

/* LOGS */
.log-container {
  background: #f8f9fa;
  padding: 15px;
  border-radius: 12px;
  height: 150px;
  overflow-y: auto;
  font-family: 'Space Mono', monospace;
  font-size: 0.8rem;
}

.tesla-active .log-container { background: #020617; color: #10b981; }

.message-item { list-style: none; margin-bottom: 8px; border-bottom: 1px solid #eee; }
.tesla-active .message-item { border-color: #1e293b; }

.footer {
  margin-top: 40px;
  padding-top: 20px;
  border-top: 1px solid #e2e8f0;
  text-align: center;
}

@keyframes pulse {
  0% { opacity: 1; }
  50% { opacity: 0.7; }
  100% { opacity: 1; }
}
</style>
