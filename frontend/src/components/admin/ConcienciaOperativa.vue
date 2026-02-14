<template>
  <section class="conciencia-container card">
    <header class="terminal-header">
      <div class="status-indicator">
        <span class="pulse-dot"></span>
        <span class="status-text">{{ conciencia.status_red }}</span>
      </div>
      <div class="timestamp">{{ currentTime }}</div>
    </header>

    <div class="conciencia-grid">
      <div class="panel pasado">
        <h3><span class="icon">✅</span> LOG DE EVOLUCIÓN</h3>
        <ul>
          <li v-for="(item, i) in conciencia.hecho" :key="i">
            <span class="arrow">></span> {{ item }}
          </li>
        </ul>
      </div>

      <div class="panel presente">
        <h3><span class="icon">⚙️</span> CONSCIENCIA ACTIVA</h3>
        <div class="process-box">
          <p class="typing-text">{{ conciencia.haciendo }}</p>
          <div class="purpose-tag">
             <strong>POR QUÉ:</strong> {{ conciencia.proposito }}
          </div>
          <div class="progress-bar-container">
            <div class="progress-fill"></div>
          </div>
        </div>
      </div>

      <div class="panel futuro">
        <h3><span class="icon">🚀</span> PRÓXIMOS PASOS</h3>
        <ul>
          <li v-for="(item, i) in conciencia.por_hacer" :key="i" class="pending-item">
            <span class="bullet">□</span> {{ item }}
          </li>
        </ul>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

const currentTime = ref(new Date().toLocaleTimeString())

const conciencia = ref({
  hecho: [
    "Capa 4: Billetera multifirma validada",
    "Capa 1: Sincronización de ADN con GitHub realizada",
    "IA 1: Distribución de beneficios en Avellaneda optimizada"
  ],
  haciendo: "Procesando flujo de datos del Telar de Ideas...",
  por_hacer: [
    "Inyectar módulo de Geocercas (Esperando tu Firma)",
    "Actualizar base de datos MongoDB"
  ],
  status_red: "CAPA 5: ONLINE",
  proposito: "Garantizar que cada línea de código respete el esquema 60/40 y la soberanía del usuario."
})

onMounted(() => {
  setInterval(() => {
    currentTime.value = new Date().toLocaleTimeString()
  }, 1000)
})
</script>

<style scoped>
.conciencia-container {
  background: #020617 !important;
  border: 1px solid #1e293b;
  font-family: 'Space Mono', monospace;
  color: #10b981;
  padding: 20px;
}

.terminal-header {
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid #1e293b;
  padding-bottom: 10px;
  margin-bottom: 20px;
}

.pulse-dot {
  height: 8px; width: 8px;
  background-color: #10b981;
  border-radius: 50%;
  display: inline-block;
  box-shadow: 0 0 10px #10b981;
  animation: pulse 1.5s infinite;
}

.conciencia-grid {
  display: grid;
  grid-template-columns: 1fr 1.5fr 1fr;
  gap: 20px;
}

.panel h3 {
  font-size: 0.75rem;
  color: #94a3b8;
  margin-bottom: 15px;
  letter-spacing: 1px;
}

ul { list-style: none; padding: 0; font-size: 0.8rem; }
li { margin-bottom: 8px; color: #f8f9fa; }
.arrow { color: #10b981; font-weight: bold; }

.process-box {
  background: #0f172a;
  padding: 15px;
  border-radius: 12px;
  border: 1px solid #10b98133;
}

.typing-text {
  color: #10b981;
  font-size: 0.9rem;
  margin-bottom: 10px;
}

.purpose-tag {
  font-size: 0.7rem;
  color: #d97706;
  background: #d9770611;
  padding: 8px;
  border-radius: 6px;
  margin-bottom: 10px;
}

.progress-bar-container {
  height: 4px; background: #1e293b; border-radius: 2px; overflow: hidden;
}
.progress-fill {
  height: 100%; background: #10b981; width: 60%;
  animation: loading 2s infinite ease-in-out;
}

@keyframes loading {
  0% { transform: translateX(-100%); }
  100% { transform: translateX(250%); }
}

@keyframes pulse {
  0% { transform: scale(1); opacity: 1; }
  70% { transform: scale(1.5); opacity: 0; }
  100% { transform: scale(1); opacity: 1; }
}
</style>