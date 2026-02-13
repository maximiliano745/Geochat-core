<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import axios from 'axios'

// Contrato con el Backend (Soberanía de datos)
interface StatsResponse {
  saldo: number;
  nombreNodo: string;
  mensajes: string[];
}

// URL de tu Codespace
const API_BASE = 'https://symmetrical-acorn-5xgw65wwgr6f4wrx-8080.app.github.dev'
//https://symmetrical-acorn-5xgw65wwgr6f4wrx-8080.app.github.dev/

const saldoOro = ref<number>(0)
const nombreNodo = ref<string>('Buscando señal del Pueblo...')
const mensajes = ref<string[]>([])
const modoTesla = ref<boolean>(false)
const cargando = ref<boolean>(true)
const vaultLocked = ref<boolean>(true)

// --- LÓGICA DE IMPACTO AVELLANEDA (IA 1) ---
const porcentajeDescuento = ref<number>(25) // Subsidio negociado por IA 1
const fondoIA = computed(() => saldoOro.value * 0.15) // El 15% para el Pueblo [cite: 2026-02-10]

// Computed para calcular el aporte simbólico del Modo Tesla
const aporteComunitario = computed(() => (saldoOro.value * 0.15).toFixed(4))

const actualizarDatos = async () => {
  try {
    cargando.value = true
    const res = await axios.get<StatsResponse>(`${API_BASE}/api/stats`)
    
    saldoOro.value = res.data.saldo
    nombreNodo.value = res.data.nombreNodo
    mensajes.value = res.data.mensajes
    vaultLocked.value = false 
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

      <section class="card impact-card">
        <div class="card-header">
          <h3>Beneficio Avellaneda</h3>
          <span class="ia-tag">IA 1 ACTIVE</span>
        </div>
        <div class="subsidy-meter">
          <span class="discount-rate">-{{ porcentajeDescuento }}%</span>
          <p>Subsidio generado por IA 1 (Ads)</p>
        </div>
        <div class="progress-bar-small">
          <div class="fill-green" :style="{ width: porcentajeDescuento + '%' }"></div>
        </div>
        <p class="gas-info">Fondo acumulado (15%): <strong>{{ fondoIA.toFixed(4) }} PAXG</strong></p>
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

.tesla-active {
  background-color: #020617 !important;
  color: #f8f9fa !important;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 40px;
  border-bottom: 1px solid #e2e8f0;
  padding-bottom: 20px;
}

.core-tag {
  font-size: 0.8rem;
  background: #d97706;
  color: white;
  padding: 2px 8px;
  border-radius: 4px;
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

.status-dot { width: 10px; height: 10px; border-radius: 50%; background: #ef4444; }
.status-dot.online { background: #10b981; box-shadow: 0 0 12px #10b981; }

.content {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 25px;
  flex: 1;
}

.card {
  background: white;
  padding: 25px;
  border-radius: 24px;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease;
}

.tesla-active .card { background: #0f172a; border: 1px solid #1e293b; color: white; }

/* IMPACT CARD (AVELLANEDA) */
.impact-card { border-top: 4px solid #10b981; }
.ia-tag { font-size: 0.6rem; background: #10b981; color: white; padding: 2px 6px; border-radius: 4px; }
.discount-rate { font-size: 2.5rem; font-weight: 800; color: #10b981; font-family: 'Space Mono', monospace; display: block; }
.progress-bar-small { height: 8px; background: #e2e8f0; border-radius: 10px; margin: 15px 0; overflow: hidden; }
.fill-green { height: 100%; background: #10b981; transition: width 1s ease; }

/* ORO Y BOTONES */
.amount { font-size: 3rem; font-weight: bold; color: #d97706; }
.vault-badge { font-size: 0.7rem; padding: 4px 10px; border-radius: 6px; background: #dcfce7; color: #166534; }
.vault-badge.locked { background: #fee2e2; color: #991b1b; }

button { width: 100%; padding: 18px; border-radius: 16px; font-weight: bold; cursor: pointer; border: none; font-family: 'Space Mono', monospace; }
.btn-tesla-off { background: #1e293b; color: white; }
.btn-tesla-on { background: #10b981; color: #020617; animation: pulse 2s infinite; }

.log-container { background: #f8f9fa; padding: 15px; border-radius: 12px; height: 150px; overflow-y: auto; font-family: 'Space Mono', monospace; font-size: 0.8rem; }
.tesla-active .log-container { background: #020617; color: #10b981; }

@keyframes pulse { 0% { opacity: 1; } 50% { opacity: 0.7; } 100% { opacity: 1; } }
</style>
