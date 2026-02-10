<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'

const saldoOro = ref(0)
const nombreNodo = ref('Cargando...')
const mensajes = ref([])
const modoTesla = ref(false)

const actualizarDatos = async () => {
  try {
    // Nota: En Codespaces podrías necesitar la URL pública del puerto 3000
    const res = await axios.get('http://localhost:3000/api/stats')
    saldoOro.value = res.data.saldo
    nombreNodo.value = res.data.nombreNodo
    mensajes.value = res.data.mensajes
  } catch (error) {
    console.error("Error conectando al Nodo:", error)
  }
}

const toggleTesla = () => {
  modoTesla.value = !modoTesla.value
  // Activación manual para el sentimiento filantrópico [cite: 2026-02-02]
}

onMounted(actualizarDatos)
</script>

<template>
  <div class="dashboard" :class="{ 'tesla-active': modoTesla }">
    <header>
      <h1>🌍 GeoChat Node: {{ nombreNodo }}</h1>
      <div class="stats">
        <p><strong>Saldo Oro (PAXG):</strong> {{ saldoOro }}</p>
        <p><strong>Bóveda:</strong> E2E Encrypted 🔒</p>
      </div>
    </header>

    <main>
      <div class="card">
        <h3>Modo Tesla</h3>
        <p>Contribución voluntaria a la red mesh.</p>
        <button @click="toggleTesla" :class="modoTesla ? 'btn-active' : 'btn-off'">
          {{ modoTesla ? 'MODO FILANTRÓPICO ON' : 'ACTIVAR MODO TESLA' }}
        </button>
      </div>

      <div class="card">
        <h3>Mensajes del Sistema</h3>
        <ul>
          <li v-for="(msg, i) in mensajes" :key="i">{{ msg }}</li>
        </ul>
      </div>
    </main>

    <footer>
      <p>GeoChat Project | Reparto: 60% Operativo - 40% Capital [cite: 2026-02-01]</p>
    </footer>
  </div>
</template>

<style>
/* Estilos globales */
body {
  margin: 0;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  background-color: #f0f2f5;
  color: #1a1a1a;
  transition: all 0.5s ease;
}

.dashboard {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
}

/* El cambio de color del Modo Tesla */
.tesla-active {
  background-color: #1a2a6c !important; /* Azul profundo filantrópico */
  color: white !important;
}

.card {
  background: white;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0,0,0,0.1);
  margin: 10px;
  width: 300px;
  color: #333;
}

button {
  width: 100%;
  padding: 10px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-weight: bold;
}

.btn-off { background-color: #444; color: white; }
.btn-active { background-color: #00ff88; color: #1a2a6c; }

footer { margin-top: auto; font-size: 0.8rem; opacity: 0.7; }
</style>
