<template>
  <div v-if="isOpen" class="workspace-overlay">
    <div class="workspace-modal">
      <header class="workspace-header">
        <div class="title-section">
          <span class="badge">PROYECTO ACTIVO</span>
          <h2>{{ tarea.titulo }}</h2>
        </div>
        <button @click="$emit('cerrar')" class="btn-close">×</button>
      </header>

      <div class="fases-container">
        <div v-for="(fase, index) in fases" :key="index" class="fase-item">
          <div class="fase-header" @click="toggleFase(index)">
            <div class="fase-info">
              <span class="step-num">{{ index + 1 }}</span>
              <h3>{{ fase.nombre }}</h3>
            </div>
            <span class="status-icon">{{ fase.completada ? '✅' : '🧠' }}</span>
          </div>

          <transition name="slide">
            <div v-if="fase.abierta" class="fase-detalle">
              <div class="grid-detalle">
                <div class="logica">
                  <h4>🧠 Estrategia del Arquitecto:</h4>
                  <p>{{ fase.explicacion }}</p>
                </div>

                <div class="codigo-propuesto">
                  <div class="code-header">
                    <h4>💻 Código Propuesto (Go):</h4>
                    <span class="file-path">backend/internal/logic/{{ fase.archivo }}</span>
                  </div>
                  <pre><code>{{ fase.codigoSnippet }}</code></pre>
                  <button @click="inyectarEnLab(fase)" class="btn-inyectar">
                    INYECTAR EN ADN (LAB)
                  </button>
                </div>
              </div>
            </div>
          </transition>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router' // Usamos el hook oficial de Vue Router

const props = defineProps<{
  tarea: any,
  isOpen: boolean
}>()

const emit = defineEmits(['cerrar'])
const router = useRouter() // Inicializamos el router aquí

// Fases simuladas que la IA 5 generará según tu idea cruda
const fases = ref([
  {
    nombre: "Definición del Contrato (PAXG)",
    explicacion: "Establecer la bóveda de garantía para el intercambio de energía.",
    codigoSnippet: "func IniciarContratoEnergia() { ... }",
    archivo: "energy_vault.go",
    abierta: true,
    completada: false
  },
  {
    nombre: "Lectura de Capa 2 (Hardware)",
    explicacion: "Conexión con los sensores de voltaje del nodo local.",
    codigoSnippet: "type SensorNode struct { ID string }",
    archivo: "hw_monitor.go",
    abierta: false,
    completada: false
  }
])

const toggleFase = (index: number) => {
  fases.value[index].abierta = !fases.value[index].abierta
}

const inyectarEnLab = (fase: any) => {
  // 1. Guardamos en el almacenamiento local qué es lo que vamos a autorizar
  localStorage.setItem('pending_evolution', JSON.stringify({
    modulo: fase.nombre,
    archivo: fase.archivo,
    snippet: fase.codigoSnippet
  }))

  // 2. Notificación visual en consola
  console.log(`🧬 IA 5: Preparando inyección de ${fase.archivo}...`)
  
  // 3. Salto automático al Panel de Firma usando la instancia del hook
  // Usamos '/admin' que es la ruta que definimos en tu index.ts
  router.push('/admin') 
}
</script>

<style scoped>
/* Mantengo tus estilos que están excelentes */
.workspace-overlay {
  position: fixed;
  top: 0; left: 0; width: 100%; height: 100%;
  background: rgba(0, 0, 0, 0.85);
  backdrop-filter: blur(8px);
  display: flex; align-items: center; justify-content: center;
  z-index: 1000;
}

.workspace-modal {
  background: #0f172a;
  width: 90%; max-width: 1000px;
  max-height: 85vh;
  border-radius: 24px;
  border: 1px solid #1e293b;
  padding: 30px;
  overflow-y: auto;
}

.workspace-header { display: flex; justify-content: space-between; margin-bottom: 30px; }
.badge { color: #10b981; font-size: 0.7rem; letter-spacing: 2px; }
.btn-close { background: none; border: none; color: white; font-size: 2rem; cursor: pointer; }

.fase-item {
  background: #1e293b55;
  border-radius: 16px;
  margin-bottom: 15px;
  border: 1px solid #334155;
}

.fase-header {
  padding: 20px; cursor: pointer;
  display: flex; justify-content: space-between; align-items: center;
}

.step-num {
  background: #10b981; color: #020617;
  width: 30px; height: 30px; display: inline-flex;
  align-items: center; justify-content: center;
  border-radius: 50%; font-weight: bold; margin-right: 15px;
}

.grid-detalle {
  display: grid; grid-template-columns: 1fr 1fr;
  gap: 20px; padding: 20px; border-top: 1px solid #334155;
}

pre {
  background: #020617; padding: 15px; border-radius: 12px;
  color: #10b981; font-size: 0.85rem; overflow-x: auto;
}

.btn-inyectar {
  width: 100%; margin-top: 10px;
  background: linear-gradient(90deg, #10b981, #059669);
  color: #020617; border: none; padding: 12px;
  border-radius: 10px; font-weight: bold; cursor: pointer;
}

.btn-inyectar:hover {
  filter: brightness(1.2);
  transform: translateY(-2px);
}

/* Transición para el despliegue de fases */
.slide-enter-active, .slide-leave-active { transition: all 0.3s ease; }
.slide-enter-from, .slide-leave-to { opacity: 0; transform: translateY(-10px); }
</style>