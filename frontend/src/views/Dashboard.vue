<template>
  <div class="dashboard-container" style="padding: 20px; font-family: 'Segoe UI', sans-serif; background: #0f0f0f; color: white; min-height: 100vh;">
    
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

      <div style="background: #1e1e1e; padding: 20px; border-radius: 15px; border: 1px solid #333; max-height: 250px; overflow-y: auto;">
        <h3 style="margin-top: 0; font-size: 1rem; border-bottom: 1px solid #333; padding-bottom: 10px;">📜 Logs del CEO Autónomo</h3>
        <ul style="list-style: none; padding: 0; font-size: 0.85rem; line-height: 1.6;">
          <li v-for="(log, idx) in mensajes" :key="idx" style="margin-bottom: 8px; border-left: 3px solid #42b983; padding-left: 12px; color: #ccc;">
            {{ log }}
          </li>
        </ul>
      </div>
    </div>

    <div class="mapa-vision" style="margin-top: 40px;">
      <h2 style="text-align: center; color: #42b983;">🗺️ Mapa de Visión: Del Lab al Core</h2>
      
      <div class="contenedor-nodos">
        <div v-for="nodo in mapaVision" :key="nodo.id" 
             :class="['elemento', nodo.estado === 'REAL' ? 'cuadro-dorado' : 'burbuja-gris']">
          
          <div class="icono-estado">{{ nodo.estado === 'REAL' ? '💎' : '🧠' }}</div>
          <h3 style="font-size: 0.9rem; text-align: center; margin: 10px 5px;">{{ nodo.nombre }}</h3>
          <p class="tag" style="font-size: 0.7rem; background: rgba(0,0,0,0.3); padding: 2px 8px; border-radius: 10px;">{{ nodo.categoria }}</p>

          <button v-if="nodo.estado === 'IDEAL'" @click="abrirFirma(nodo)" class="btn-firmar">
            ✍️ FIRMAR ADN
          </button>
          
          <div v-else style="color: #f1c40f; font-weight: bold; font-size: 0.7rem; margin-top: 10px;">ACTIVO EN CORE</div>
        </div>
      </div>
    </div>

    <BurbujaDinamica v-if="burbujaActiva" @posponer="manejarPosponer" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import axios from 'axios'
import BurbujaDinamica from '../components/admin/BurbujaDinamica.vue'

// 🔗 URL de tu Núcleo (Sincronizada con tu puerto 8080)
const API_BASE = 'https://symmetrical-acorn-5xgw65wwgr6f4wrx-8080.app.github.dev'

const saldoOro = ref(0)
const nombreNodo = ref('Conectando...')
const mensajes = ref<string[]>([])
const modoTesla = ref(false)
const burbujaActiva = ref(true)

// MOCK DE DATOS PARA EL MAPA (Luego vendrán del Backend)
const mapaVision = ref([
  { id: 1, nombre: 'Optimización Ancho Banda', estado: 'IDEAL', categoria: 'Infraestructura' },
  { id: 2, nombre: 'Bóveda E2E Soberana', estado: 'REAL', categoria: 'Seguridad' },
  { id: 3, nombre: 'Nodo Rural Chaco', estado: 'IDEAL', categoria: 'Expansión' }
])

const aporteComunitario = computed(() => (saldoOro.value * 0.15).toFixed(4))

const toggleTesla = () => {
  modoTesla.value = !modoTesla.value
  mensajes.value.unshift(modoTesla.value ? "✨ [MODO TESLA]: Activado." : "⏸️ Modo Tesla en espera.");
}

// ✅ SOLUCIÓN AL ERROR: Función para cerrar la burbuja de la IA
const manejarPosponer = () => {
  burbujaActiva.value = false;
  mensajes.value.unshift("⏳ [SISTEMA]: Análisis dinámico pospuesto por el Soberano.");
}

const abrirFirma = (nodo: any) => {
  const clave = prompt(`¿Deseas inyectar el ADN de "${nodo.nombre}"?\nIntroduce tu MASTER_KEY (Tu Firma es la Ley):`);
  if (clave === "Dale Hacelo") {
    nodo.estado = 'REAL'; // Transmutación: De Burbuja Gris a Cuadrado Dorado
    mensajes.value.unshift(`👑 FIRMA VALIDADA: ${nodo.nombre} ahora es Realidad.`);
    alert("ADN Inyectado con éxito.");
  } else {
    alert("🚫 Firma rechazada.");
  }
}

const actualizarDatos = async () => {
  try {
    const res = await axios.get(`${API_BASE}/api/dashboard/stats`)
    saldoOro.value = res.data.saldo_total || 0
    nombreNodo.value = res.data.nodo_nombre || 'Nodo GeoChat'
    mensajes.value = res.data.logs_recientes || []
  } catch (e) { 
    console.error("❌ Error de enlace con el CEO.");
  }
}

onMounted(() => {
  actualizarDatos()
  setInterval(actualizarDatos, 30000)
})
</script>

<style scoped>
.contenedor-nodos {
  display: flex;
  flex-wrap: wrap;
  gap: 25px;
  justify-content: center;
  padding: 20px;
}

/* 🧠 BURBUJA GRIS: La Idea en el Lab */
.burbuja-gris {
  width: 180px;
  height: 180px;
  background: rgba(100, 100, 100, 0.1);
  border: 2px dashed #666;
  border-radius: 50%; /* FORMA CIRCULAR */
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  transition: all 0.6s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  filter: grayscale(100%);
}

.burbuja-gris:hover {
  transform: scale(1.05);
  background: rgba(100, 100, 100, 0.2);
}

/* 💎 CUADRO DORADO: La Realidad inyectada */
.cuadro-dorado {
  width: 180px;
  height: 180px;
  background: linear-gradient(145deg, #f39c12, #d35400);
  border: 3px solid #fff;
  border-radius: 15px; /* FORMA GEOMÉTRICA */
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  box-shadow: 0 0 25px rgba(243, 156, 18, 0.6);
  transform: scale(1.05);
  animation: shine 2s infinite alternate;
}

@keyframes shine {
  from { box-shadow: 0 0 15px rgba(243, 156, 18, 0.4); }
  to { box-shadow: 0 0 30px rgba(243, 156, 18, 0.8); }
}

.btn-firmar {
  margin-top: 10px;
  padding: 5px 12px;
  border-radius: 20px;
  border: none;
  background: #42b983;
  color: white;
  font-size: 0.7rem;
  cursor: pointer;
  font-weight: bold;
}

.icono-estado {
  font-size: 2.2rem;
}
</style>