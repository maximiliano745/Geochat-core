<template>
  <div class="dashboard-soberano">
    <h2>📋 Propuesta de Evolución Diaria</h2>
    <p>ID: {{ propuesta.id_transaccion }}</p>

    <section class="stats-sociales">
      <h3>⚡ Energía Social</h3>
      <ul>
        <li>Buena Onda: {{ propuesta.resumen_energia.usuarios_buena_onda }}</li>
        <li>Bloqueos (Mala Onda): {{ propuesta.resumen_energia.intentos_bloqueados_mala_onda }}</li>
        <li>Plasticidad: +{{ propuesta.resumen_energia.puntos_plasticidad_ganados }} pts</li>
      </ul>
    </section>

    <section class="finanzas">
      <h3>💰 Fondo del Pueblo (15%)</h3>
      <p>Balance: <strong>{{ propuesta.estado_financiero.fondo_15_pueblo }}</strong></p>
      <p>Reinversión sugerida: {{ propuesta.estado_financiero.propuesta_reinversion }}%</p>
      <p>Destino: {{ propuesta.estado_financiero.destino_sugerido }}</p>
    </section>

    <button 
      @click="autorizarEvolucion(propuesta)" 
      class="btn-firmar"
      :disabled="!propuesta.requiere_firma"
    >
      ✍️ Firmar y Ejecutar Logros
    </button>
  </div>
</template>

<script setup>
import axios from 'axios';
import { checkBiometrico } from '@/services/security'; // Simula FaceID/Huella

const props = defineProps(['propuesta']);

const autorizarEvolucion = async (jsonIA) => {
  // 1. "Mi Firma es la Ley": Validación biométrica real
  const confirmacion = await checkBiometrico();
  
  if (confirmacion) {
    try {
      // 2. Enviamos la orden al Backend en Go
      const response = await axios.post('/api/v1/ceo/validar', {
        id_transaccion: jsonIA.id_transaccion,
        decision: 'APROBADO',
        firma_digital: 'MI_FIRMA_ES_LA_LEY' // Tu llave de validación
      });

      if (response.status === 200) {
        alert("🚀 Éxito: Logros grabados en el Vault y fondos ejecutados.");
        // Aquí podrías disparar una animación de "Crecimiento de Red"
      }
    } catch (error) {
      console.error("Error en la validación soberana:", error);
      alert("Fallo en la comunicación con el Core.");
    }
  } else {
    alert("Autorización cancelada. La IA permanece a la espera.");
  }
};
</script>

<style scoped>
.btn-firmar {
  background: #00ff88;
  color: #000;
  padding: 15px 30px;
  border-radius: 8px;
  font-weight: bold;
  cursor: pointer;
}
.btn-firmar:disabled {
  background: #555;
}
</style>