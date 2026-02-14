<template>
  <div class="geochat-container">
    <router-view v-slot="{ Component }">
      <transition name="fade" mode="out-in">
        <component :is="Component" @autorizar="autorizarInversion" />
      </transition>
    </router-view>
  </div>
</template>

<script setup lang="ts">
import { provide } from 'vue';

const autorizarInversion = async (propuestaId: string) => {
  // --- GRITO 1: GATILLO ---
  console.log("🚀 [FRONTEND]: ¡Gatillo presionado! Intentando firmar:", propuestaId);
  
  try {
    const masterKey = "Dale Hacelo"; 
    const backendURL = "https://symmetrical-acorn-5xgw65wwgr6f4wrx-8080.app.github.dev/api/ceo/autorizar";

    // --- GRITO 2: CONEXIÓN ---
    console.log("📡 [FRONTEND]: Conectando con el núcleo en:", backendURL);
    console.log("📦 [FRONTEND]: Enviando datos:", { nombre_archivo: propuestaId, firma: masterKey });

    const respuesta = await fetch(backendURL, {
      method: 'POST',
      headers: { 
        'Content-Type': 'application/json',
        'Accept': 'application/json'
      },
      body: JSON.stringify({
        nombre_archivo: propuestaId || "test_evolucion.go", 
        firma: masterKey 
      })
    });

    // --- GRITO 3: RECEPCIÓN ---
    console.log("📥 [FRONTEND]: Respuesta recibida del núcleo. Status:", respuesta.status);

    const resultado = await respuesta.json();
    console.log("📄 [FRONTEND]: Contenido del resultado:", resultado);

    if (respuesta.ok) {
      alert("✅ ÉXITO SOBERANO: " + (resultado.status || "ADN Actualizado"));
    } else {
      alert("❌ NÚCLEO DICE NO: " + (resultado.error || "Firma Inválida"));
    }

  } catch (error) {
    // --- GRITO 4: FALLO TOTAL ---
    console.error("🚨 [ERROR CRÍTICO EN RED]:", error);
    alert("FALLO DE COMUNICACIÓN: Mira la consola (F12) para ver el error técnico.");
  }
};

// Pasamos la función a los hijos para que el Dashboard la pueda usar
provide('autorizarInversion', autorizarInversion);
</script>

<style>
:root {
  --background: #020617;
  --accent: #3b82f6;
  --text: #f8fafc;
}

body {
  margin: 0;
  padding: 0;
  background-color: var(--background);
  color: var(--text);
  font-family: 'Inter', sans-serif;
}

.geochat-container {
  min-height: 100vh;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>