<template>
  <div class="dashboard" :class="{ 'tesla-active': modoTesla }">
    <BurbujaDinamica v-if="burbujaActiva" @posponer="manejarPosponer" />

    <header class="header">
      <div class="logo-area">
        <div class="brand">
          <span class="geo-icon">🌍</span>
          <h1>GeoChat <span class="edition">AVELLANEDA</span></h1>
        </div>
        <p class="tagline">Soberanía Digital & Energía Compartida</p>
      </div>
      
      <router-link to="/admin" class="admin-link">🔐 SOBERANO</router-link>

      <div class="node-status" :class="{ 'node-online': !cargando }">
        <div class="pulse-ring"></div>
        <span class="node-text">{{ nombreNodo }}</span>
      </div>
    </header>

    <main class="main-grid">
      <section v-if="pendientes.length > 0" class="bento-card golden-archive">
        <div class="card-label">TELAR DE DECISIONES (IA 5)</div>
        <div class="pendientes-grid">
          <div v-for="p in pendientes" :key="p.id" class="golden-seed" @click="abrirDecision(p)">
            <div class="seed-icon">🟡</div>
            <div class="seed-info">
              <h4>{{ p.id }}</h4>
              <p>{{ p.mensaje.substring(0, 30) }}...</p>
              <span class="status-pill">PENDIENTE</span>
            </div>
          </div>
        </div>
      </section>

      <section class="bento-card gold-vault">
        <div class="card-label">RESERVA DE ORO E2E</div>
        <div class="gold-display">
          <div class="gold-icon">🟡</div>
          <div class="value-wrap">
            <span class="amount">{{ saldoOro.toFixed(4) }}</span>
            <span class="unit">PAXG</span>
          </div>
        </div>
        <div class="vault-status">
          <span class="secure-tag">{{ vaultLocked ? '🔒 ENCRIPTADO' : '🔓 ACCESO SOBERANO' }}</span>
        </div>
      </section>

      <section class="bento-card tesla-module" :class="{ 'active': modoTesla }">
        <div class="card-content">
          <h3>Modo Tesla</h3>
          <p>Al compartir tu energía, el pueblo gana.</p>
          <div v-if="modoTesla" class="impact-indicator">
            <span class="impact-value">+{{ aporteComunitario }}</span>
            <span class="impact-label">PAXG aportados</span>
          </div>
        </div>
        <button @click="toggleTesla" class="tesla-toggle">
          {{ modoTesla ? 'FILANTROPÍA ACTIVA' : 'ACTIVAR FILANTROPÍA' }}
        </button>
      </section>

      <section class="bento-card social-impact">
        <div class="impact-header">
          <h3>Beneficio Local</h3>
          <span class="ia-badge">IA 1</span>
        </div>
        <div class="discount-circle">
          <span class="percent">-{{ porcentajeDescuento }}%</span>
          <span class="desc">en servicios de red</span>
        </div>
        <div class="fondo-pueblo">
          <p>Fondo Comunitario:</p>
          <strong>{{ fondoIA.toFixed(4) }} PAXG</strong>
        </div>
      </section>

      <section class="bento-card system-logs">
        <h3>Registro del Sistema</h3>
        <div class="log-viewport">
          <div v-for="(msg, i) in mensajes" :key="i" class="log-entry">
            <span class="log-time">[{{ new Date().toLocaleTimeString() }}]</span>
            <span class="log-msg">{{ msg }}</span>
          </div>
        </div>
      </section>
    </main>

    <footer class="dashboard-footer">
      <div class="fractal-info">Estructura Fractal 60/40 | Geochat (IA 5) CEO Autónomo</div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import axios from 'axios'
import BurbujaDinamica from '../components/admin/BurbujaDinamica.vue'

const API_BASE = 'https://symmetrical-acorn-5xgw65wwgr6f4wrx-8080.app.github.dev'

const saldoOro = ref(1.2450)
const nombreNodo = ref('Nodo Avellaneda Principal')
const mensajes = ref<string[]>(['Sincronización inicial exitosa', 'Capa 2 activa'])
const modoTesla = ref(false)
const cargando = ref(false)
const vaultLocked = ref(false)
const porcentajeDescuento = ref(25)
const burbujaActiva = ref(true)

// LÓGICA DE PENDIENTES (Soberanía Fractal)
const pendientes = ref(JSON.parse(localStorage.getItem('decisiones_pendientes') || '[]'))

const fondoIA = computed(() => saldoOro.value * 0.15)
const aporteComunitario = computed(() => (saldoOro.value * 0.15).toFixed(4))

const toggleTesla = () => {
  modoTesla.value = !modoTesla.value
  mensajes.value.unshift(modoTesla.value ? "✨ Iniciando compartición de energía..." : "⏸️ Modo Tesla en espera.");
}

const manejarPosponer = () => {
  pendientes.value = JSON.parse(localStorage.getItem('decisiones_pendientes') || '[]')
  burbujaActiva.value = false
}

const abrirDecision = (p: any) => {
  // Aquí la IA 5 te recuerda el contexto antes de llevarte a firmar
  alert(`Contexto IA 5: ${p.mensaje}\nImpacto: ${p.impacto}`);
}

const actualizarDatos = async () => {
  try {
    const res = await axios.get(`${API_BASE}/api/stats`)
    saldoOro.value = res.data.saldo
    nombreNodo.value = res.data.nombreNodo
    mensajes.value = res.data.mensajes
  } catch (e) { console.log("Modo Offline: Usando datos de Vault local") }
}

onMounted(actualizarDatos)
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Space+Grotesk:wght@300;500;700&display=swap');

.dashboard {
  min-height: 100vh; background: #f1f5f9; color: #0f172a;
  padding: 30px; font-family: 'Space Grotesk', sans-serif; transition: 0.5s;
}
.tesla-active { background: #020617; color: #f8f9fa; }

/* TELAR DE DECISIONES */
.golden-archive { grid-column: span 3; background: #fffbeb !important; border: 1px solid #fde68a !important; }
.tesla-active .golden-archive { background: #1e1b10 !important; border-color: #d97706 !important; }

.pendientes-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 15px; margin-top: 15px;
}

.golden-seed { 
  background: white; padding: 15px; border-radius: 20px; 
  display: flex; align-items: flex-start; gap: 12px; cursor: pointer;
  box-shadow: 0 4px 6px -1px rgba(0,0,0,0.05); border: 1px solid #fde68a;
  transition: 0.3s;
}
.tesla-active .golden-seed { background: #0f172a; border-color: #d9770633; }
.golden-seed:hover { transform: translateY(-5px); border-color: #d97706; }

.seed-info h4 { margin: 0; font-size: 0.9rem; color: #d97706; }
.seed-info p { margin: 5px 0; font-size: 0.75rem; color: #64748b; }
.status-pill { font-size: 0.6rem; background: #d97706; color: white; padding: 2px 8px; border-radius: 10px; font-weight: bold; }

/* ESTRUCTURA BENTO */
.header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 40px; }
.admin-link { background: #0f172a; color: white; text-decoration: none; padding: 10px 20px; border-radius: 50px; font-weight: bold; font-size: 0.8rem; }
.tesla-active .admin-link { background: #10b981; color: #020617; }

.main-grid { display: grid; grid-template-columns: repeat(3, 1fr); gap: 20px; max-width: 1200px; margin: 0 auto; }
.bento-card { background: white; border-radius: 28px; padding: 30px; border: 1px solid rgba(0,0,0,0.05); position: relative; }
.gold-vault { grid-column: span 2; }
.amount { font-size: 4rem; font-weight: 700; color: #d97706; }
.system-logs { grid-column: span 2; }
.log-viewport { height: 120px; overflow-y: auto; font-family: monospace; font-size: 0.8rem; }

@media (max-width: 900px) { .main-grid { grid-template-columns: 1fr; } .gold-vault, .system-logs, .golden-archive { grid-column: span 1; } }
</style>