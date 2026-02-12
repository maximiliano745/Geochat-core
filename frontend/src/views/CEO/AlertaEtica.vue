<template>
  <div v-if="mostrarAlerta" class="modal-reflexion">
    <div class="alerta-content">
      <div class="icono-ia">🤖⚠️</div>
      <h3>Advertencia de Trayectoria</h3>
      <p class="mensaje-ia">{{ mensajeIA }}</p>
      
      <div class="check-box">
        <input type="checkbox" v-model="confirmarEvolucion" id="evolucion">
        <label for="evolucion">Entiendo que esta decisión cambia el ADN de GeoChat</label>
      </div>

      <div class="botones">
        <button @click="cancelar" class="btn-abortar">Revisar Decisión</button>
        <button 
          @click="procederConReflexion" 
          :disabled="!confirmarEvolucion" 
          class="btn-soberano"
        >
          Confirmar Evolución
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const props = defineProps(['mensajeIA']);
const emit = defineEmits(['continuar', 'cancelar']);

const mostrarAlerta = ref(true);
const confirmarEvolucion = ref(false);

const procederConReflexion = () => {
  // Aquí es donde se activaría el segundo factor de autenticación
  emit('continuar');
  mostrarAlerta.value = false;
};

const cancelar = () => {
  emit('cancelar');
  mostrarAlerta.value = false;
};
</script>

<style scoped>
.modal-reflexion {
  background: rgba(0, 0, 0, 0.9);
  border: 2px solid #ff4444; /* Rojo alerta */
  padding: 2rem;
  color: white;
  border-radius: 15px;
}
.mensaje-ia {
  font-style: italic;
  color: #ffcc00; /* Color advertencia */
}
</style>