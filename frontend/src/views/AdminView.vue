<template>
  <div class="admin-view">
    <BurbujaDinamica v-if="propuestaActiva" :propuesta="propuestaActiva" @posponer="manejarPosponer" />

    <header class="admin-header">
      <div class="title-wrap">
        <span class="sovereign-tag">NIVEL 5: CEO ESTRATÉGICO</span>
        <h2>🔐 Panel de Gestión Soberana</h2>
      </div>
      <router-link to="/" class="btn-back">← Volver a Vista Pueblo</router-link>
    </header>

    <div class="admin-grid">
      <div class="control-column">
        <ControlSoberano />
        
        <section class="card status-card">
          <div class="card-header">
            <h3>Geochat: CEO Autónomo</h3>
            <span class="pulse-dot"></span>
          </div>
          <div class="status-content">
            <p>Estado: <span class="online">SINCRO OK</span></p>
            <p class="small">El 15% del fondo ({{ fondo15 }} PAXG) está listo para reinversión en infraestructura.</p>
          </div>
        </section>

        <section v-if="pendientes.length > 0" class="card pending-card">
          <h3>🗄️ Decisiones en Archivo</h3>
          <div class="golden-grid">
            <div v-for="p in pendientes" :key="p.id" class="golden-item" @click="recuperarPropuesta(p)">
              <span class="gold-icon">🟡</span>
              <div class="gold-info">
                <p class="gold-id">{{ p.id }}</p>
                <small>{{ p.modulo }}</small>
              </div>
            </div>
          </div>
        </section>
      </div>

      <div class="telar-column">
        <section class="card ideas-card">
          <div class="card-header">
            <h3>🧵 Telar de Ideas</h3>
            <span class="ia-tag">IA 5 ARCHITECT</span>
          </div>
          <p class="description">Seleccioná una semilla de idea para desglosar su ADN técnico y autorizar su inyección.</p>
          
          <div class="ideas-grid">
            <div 
              v-for="tarea in tareas" 
              :key="tarea.id" 
              class="idea-item" 
              @click="abrirTarea(tarea)"
            >
              <div class="seed-icon">{{ tarea.icon }}</div>
              <div class="idea-text">
                <h4>{{ tarea.titulo }}</h4>
                <span class="status-pill">{{ tarea.fase }}</span>
              </div>
            </div>
          </div>
        </section>
      </div>
    </div>

    <WorkspaceTask 
      v-if="showWorkspace" 
      :tarea="tareaSeleccionada" 
      :isOpen="showWorkspace"
      @cerrar="showWorkspace = false" 
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import ControlSoberano from '../components/admin/ControlSoberano.vue'
import WorkspaceTask from '../components/admin/WorkspaceTask.vue'
import BurbujaDinamica from '../components/admin/BurbujaDinamica.vue' // Asegurate de tenerlo creado

const showWorkspace = ref(false)
const tareaSeleccionada = ref(null)
const propuestaActiva = ref(null)
const pendientes = ref<any[]>([])
const fondo15 = ref("1.420") // Simulado, vendrá del backend

const tareas = ref([
  { id: '1', titulo: 'Geocercas Avellaneda', fase: 'Lógica', icon: '📍' },
  { id: '2', titulo: 'Nodo Energía Solar', fase: 'Hardware', icon: '☀️' },
  { id: '3', titulo: 'Billetera Soberana', fase: 'Seguridad', icon: '🔑' }
])

onMounted(() => {
  // Cargar decisiones pospuestas del storage
  const guardadas = localStorage.getItem('decisiones_pendientes')
  if (guardadas) pendientes.value = JSON.parse(guardadas)

  // Simulación: La IA CEO lanza una propuesta al entrar
  setTimeout(() => {
    propuestaActiva.value = {
      id: 'OP-PAXG-01',
      mensaje: "Jefe, liquidez ociosa detectada. ¿Compramos PAXG para reserva?",
      impacto: "+1.2% Estabilidad",
      modulo: "Finanzas"
    }
  }, 3000)
})

const abrirTarea = (tarea: any) => {
  tareaSeleccionada.value = tarea
  showWorkspace.value = true
}

const manejarPosponer = (propuesta: any) => {
  pendientes.value.push(propuesta)
  localStorage.setItem('decisiones_pendientes', JSON.stringify(pendientes.value))
  propuestaActiva.value = null
}

const recuperarPropuesta = (p: any) => {
  propuestaActiva.value = p
  pendientes.value = pendientes.value.filter(item => item.id !== p.id)
  localStorage.setItem('decisiones_pendientes', JSON.stringify(pendientes.value))
}
</script>

<style scoped>
.admin-view {
  min-height: 100vh; background: #020617; color: #f8f9fa;
  padding: 40px; font-family: 'Space Mono', monospace;
}
.admin-header { display: flex; justify-content: space-between; margin-bottom: 40px; }
.sovereign-tag { color: #d97706; font-size: 0.7rem; border: 1px solid #d97706; padding: 2px 8px; border-radius: 4px; }
.btn-back { color: #10b981; text-decoration: none; border: 1px solid #10b98144; padding: 10px 20px; border-radius: 12px; transition: 0.3s; }
.btn-back:hover { background: #10b98111; border-color: #10b981; }

.admin-grid { display: grid; grid-template-columns: 400px 1fr; gap: 30px; }

/* TELAR Y IDEAS */
.ideas-grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); gap: 15px; margin-top: 20px; }
.idea-item {
  background: #1e293b44; border: 1px solid #334155;
  padding: 20px; border-radius: 18px; cursor: pointer; transition: 0.3s;
  display: flex; align-items: center; gap: 15px;
}
.idea-item:hover { border-color: #10b981; background: #1e293b88; transform: translateY(-3px); }

/* CARDS DORADAS (PENDIENTES) */
.golden-grid { display: grid; grid-template-columns: 1fr 1fr; gap: 10px; margin-top: 15px; }
.golden-item {
  background: #d9770611; border: 1px solid #d9770644;
  padding: 10px; border-radius: 12px; cursor: pointer;
  display: flex; align-items: center; gap: 10px;
}
.gold-id { font-size: 0.7rem; font-weight: bold; color: #d97706; margin: 0; }
.gold-info small { font-size: 0.6rem; color: #94a3b8; }

.card { background: #0f172a; padding: 25px; border-radius: 24px; border: 1px solid #1e293b; margin-bottom: 20px; }
.ia-tag { font-size: 0.7rem; background: #10b981; color: #020617; padding: 2px 8px; border-radius: 4px; font-weight: bold; }
.online { color: #10b981; font-weight: bold; }
.pulse-dot { width: 8px; height: 8px; background: #10b981; border-radius: 50%; box-shadow: 0 0 10px #10b981; }
</style>