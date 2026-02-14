<template>
  <div class="dashboard-container" style="padding: 20px; font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background: #0f0f0f; color: white; min-height: 100vh;">
    
    <header style="border-bottom: 2px solid #42b983; padding-bottom: 15px; margin-bottom: 30px; display: flex; justify-content: space-between; align-items: center;">
      <div>
        <h1 style="margin: 0; color: #42b983;">🛡️ GeoChat - Panel del Soberano</h1>
        <p style="margin: 5px 0 0 0;">Nodo Activo: <span style="font-weight: bold; color: #3498db;">{{ nombreNodo }}</span></p>
      </div>
      <div style="text-align: right;">
        <button @click="toggleTesla" :style="{ background: modoTesla ? '#f39c12' : '#2c3e50' }" style="color: white; border: none; padding: 10px 20px; border-radius: 20px; cursor: pointer; transition: 0.3s;">
          {{ modoTesla ? '✨ Modo Tesla Activo' : '🔋 Activar Modo Tesla' }}
        </button>
      </div>
    </header>

    <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 20px;">
      
      <div style="background: #1e1e1e; padding: 25px; border-radius: 15px; border: 1px solid #333; text-align: center;">
        <h3 style="color: #888; text-transform: uppercase; font-size: 0.8rem; letter-spacing: 1px;">Saldo Comunitario (15%)</h3>
        <p style="font-size: 3rem; font-weight: bold; color: #f39c12; margin: 10px 0;">{{ aporteComunitario }} <span style="font-size: 1rem;">PAXG</span></p>
        <p style="color: #555; font-size: 0.9rem;">Total en Red: {{ saldoOro }} PAXG</p>
      </div>

      <div style="background: #1e1e1e; padding: 20px; border-radius: 15px; border: 1px solid #333; max-height: 300px; overflow-y: auto;">
        <h3 style="margin-top: 0; font-size: 1rem; border-bottom: 1px solid #333; padding-bottom: 10px;">📜 Logs del CEO Autónomo</h3>
        <ul style="list-style: none; padding: 0; font-size: 0.85rem; line-height: 1.6;">
          <li v-for="(log, idx) in mensajes" :key="idx" style="margin-bottom: 8px; border-left: 3px solid #42b983; padding-left: 12px; color: #ccc;">
            {{ log }}
          </li>
        </ul>
      </div>
    </div>

    <div v-if="pendientes.length > 0" style="margin-top: 30px; background: #2c3e50; padding: 15px; border-radius: 10px;">
      <h4 style="margin: 0 0 10px 0;">⚠️ Decisiones de Infraestructura Pendientes:</h4>
      <div style="display: flex; gap: 10px;">
        <button v-for="p in pendientes" :key="p.id" @click="abrirDecision(p)" style="background: #34495e; color: white; border: 1px solid #42b983; padding: 8px 15px; border-radius: 5px; cursor: pointer;">
          📍 {{ p.mensaje }}
        </button>
      </div>
    </div>

    <div style="margin-top: 50px; text-align: center; border-top: 1px solid #333; padding-top: 30px;">
      <p style="color: #666; margin-bottom: 15px;">La IA propone, pero sin tu firma, no hay movimiento económico.</p>
      <button @click="autorizar" style="background: #42b983; color: white; border: none; padding: 20px 50px; font-size: 1.4rem; font-weight: bold; border-radius: 10px; cursor: pointer; box-shadow: 0 4px 15px rgba(66, 185, 131, 0.3);">
        🚀 EJECUTAR FIRMA SOBERANA
      </button>
    </div>

    <BurbujaDinamica v-if="burbujaActiva" @posponer="manejarPosponer" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import axios from 'axios'
import BurbujaDinamica from '../components/admin/BurbujaDinamica.vue'

// 🔗 URL de tu Núcleo (Configurada por entorno o manual)
const API_BASE = 'https://symmetrical-acorn-5xgw65wwgr6f4wrx-8080.app.github.dev'

// Variables de Estado
const saldoOro = ref(0)
const nombreNodo = ref('Conectando...')
const mensajes = ref<string[]>([])
const modoTesla = ref(false)
const cargando = ref(true)
const burbujaActiva = ref(true)

// LÓGICA DE SOBERANÍA (Pendientes)
const pendientes = ref(JSON.parse(localStorage.getItem('decisiones_pendientes') || '[]'))

// Cálculo del 15% para el Pueblo
const aporteComunitario = computed(() => (saldoOro.value * 0.15).toFixed(4))

const toggleTesla = () => {
  modoTesla.value = !modoTesla.value
  const msg = modoTesla.value ? "✨ [MODO TESLA]: Compartición filantrópica activada." : "⏸️ Modo Tesla en espera.";
  mensajes.value.unshift(msg);
}

const manejarPosponer = () => {
  burbujaActiva.value = false
}

const abrirDecision = (p: any) => {
  alert(`Contexto IA 5: ${p.mensaje}\nImpacto: ${p.impacto}\nRequiere Validación SOBERANA.`);
}

// 🎯 SINCRONIZACIÓN CON EL MOTOR GO
const actualizarDatos = async () => {
  try {
    cargando.value = true
    const res = await axios.get(`${API_BASE}/api/dashboard/stats`)
    
    // Mapeo de datos del Backend a la Interfaz
    saldoOro.value = res.data.saldo_total || 0
    nombreNodo.value = res.data.nodo_nombre || 'Nodo GeoChat'
    mensajes.value = res.data.logs_recientes || []
    
    cargando.value = false
  } catch (e) { 
    console.error("❌ [DASHBOARD]: Error de enlace. Revisa Codespaces Puerto 8080.");
    cargando.value = false;
  }
}

// 👑 FUNCIÓN DE FIRMA: Mi Firma es la Ley
const autorizar = () => {
  const clave = prompt("Introduce tu MASTER_KEY (Tu Firma es la Ley):");
  if (clave === "Dale Hacelo") {
    mensajes.value.unshift("👑 FIRMA VALIDADA: Ejecutando órdenes en la red...");
    alert("Firma validada. La IA CEO inicia la expansión.");
  } else {
    alert("🚫 Acceso denegado. Solo el Soberano puede autorizar.");
  }
};

onMounted(() => {
  actualizarDatos()
  // Sincronización cada 30 segundos
  setInterval(actualizarDatos, 30000)
})
</script>