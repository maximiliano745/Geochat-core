<template>
  <div class="dashboard-ceo">
    <header class="header-soberano">
      <h1>🛸 Panel de Validación: IA CEO</h1>
      <div class="stats-bar">
        <span>⚡ Energía Vital (Tokens): <strong>{{ tokens }}/300</strong></span>
        <span>💰 Fondo del Pueblo (15%): <strong>{{ fondo }} PAXG</strong></span>
      </div>
    </header>

    <div v-if="propuestas.length === 0" class="empty-state">
      La IA está analizando oportunidades en la Red Mesh...
    </div>

    <div v-for="p in propuestas" :key="p.id" class="card-propuesta">
      <h3>🛠️ Módulo Propuesto: {{ p.modulo }}</h3>
      <div class="code-box">
        <code>{{ p.arquitectura }}</code>
      </div>
      <p class="cost">Costo de Energía: {{ p.costoTokens }} tokens</p>
      
      <div class="actions">
        <button @click="autorizar(p.id)" class="btn-approve">MI FIRMA ES LA LEY (EJECUTAR)</button>
        <button @click="rechazar(p.id)" class="btn-reject">RECHAZAR</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';

const tokens = ref(300); // Energía vital de la IA [cite: 2026-02-10]
const fondo = ref(0.00); // El 15% para el pueblo [cite: 2026-02-09]
const propuestas = ref([]);

// 1. Cargar la realidad del Nodo desde el Backend
const cargarDatos = async () => {
  try {
    const response = await fetch('http://localhost:3000/api/proposals');
    const data = await response.json();
    propuestas.value = data.propuestas || [];
    tokens.value = data.tokensGratis;
    fondo.value = data.fondoGas;
  } catch (error) {
    console.error("Error conectando con la Capa 5:", error);
  }
};

// 2. Ejecución Soberana: "Mi Firma es la Ley"
const autorizar = async (id) => {
  console.log("Iniciando protocolo de inyección...");

  try {
    const response = await fetch('http://localhost:3000/api/autorizar', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      // Enviamos el SignedHash que valida tu autoridad absoluta [cite: 2026-02-10]
      body: JSON.stringify({ 
        firma: "MI_FIRMA_ES_LA_LEY", 
        id: id 
      })
    });

    if (response.ok) {
      alert("✅ ¡Soberanía confirmada! El código ha sido inyectado en el ADN de GeoChat.");
      // Limpiamos la pantalla y refrescamos balances
      await cargarDatos(); 
    } else {
      alert("❌ Error: La firma no fue validada por la Capa de Seguridad.");
    }
  } catch (error) {
    alert("🚨 Error de conexión con el Nodo.");
  }
};

const rechazar = (id) => {
  // Aquí la IA aprende que ese camino no te gusta [cite: 2026-02-09]
  propuestas.value = propuestas.value.filter(p => p.id !== id);
};

// Arrancamos el panel pidiendo datos al organismo vivo
onMounted(cargarDatos);
</script>