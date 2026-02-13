<template>
  <section class="card control-card">
    <h3>✍️ Firma Soberana</h3>
    
    <div v-if="pendingTask" class="pending-task-box">
      <p class="label">Módulo a inyectar:</p>
      <div class="task-info">
        <span class="task-name">{{ pendingTask.modulo }}</span>
        <code class="task-file">{{ pendingTask.archivo }}</code>
      </div>
    </div>

    <p class="description">
      {{ pendingTask ? 'Confirma la inyección de este módulo al ADN.' : 'Ingresa la Master Key para autorizar la evolución.' }}
    </p>
    
    <div class="input-group">
      <input 
        v-model="masterKey" 
        type="password" 
        placeholder="••••••••" 
        :disabled="loading"
        @keyup.enter="autorizar"
      />
      <button @click="autorizar" :disabled="loading || !masterKey">
        {{ loading ? 'Sincronizando...' : 'Firmar y Ejecutar' }}
      </button>
    </div>

    <p v-if="mensaje" :class="['status-msg', isError ? 'error' : 'success']">
      {{ mensaje }}
    </p>
  </section>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { evolutionService } from '../../services/evolution.service';

const masterKey = ref('');
const mensaje = ref('');
const loading = ref(false);
const isError = ref(false);
const pendingTask = ref<any>(null);

onMounted(() => {
  // Leemos si el usuario viene desde una propuesta del Workspace
  const data = localStorage.getItem('pending_evolution');
  if (data) {
    pendingTask.value = JSON.parse(data);
  }
});

const autorizar = async () => {
  loading.value = true;
  mensaje.value = '';
  isError.value = false;

  try {
    const res = await evolutionService.sendSovereignSignature(masterKey.value);
    mensaje.value = res.mensaje;
    masterKey.value = ''; 
    
    // Si tuvo éxito, limpiamos la tarea pendiente
    localStorage.removeItem('pending_evolution');
    setTimeout(() => { pendingTask.value = null; }, 2000);

  } catch (err: any) {
    isError.value = true;
    mensaje.value = err.message;
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
.control-card { border: 1px solid #10b98166; }
.pending-task-box {
  background: #10b98115;
  border-left: 4px solid #10b981;
  padding: 15px;
  margin-bottom: 20px;
  border-radius: 4px 12px 12px 4px;
}
.label { font-size: 0.7rem; color: #10b981; margin: 0; text-transform: uppercase; }
.task-info { display: flex; justify-content: space-between; align-items: center; margin-top: 5px; }
.task-name { font-weight: bold; color: #f8f9fa; }
.task-file { font-size: 0.7rem; color: #94a3b8; background: #020617; padding: 2px 6px; border-radius: 4px; }

.description { font-size: 0.9rem; color: #94a3b8; margin-bottom: 20px; }
.input-group { display: flex; flex-direction: column; gap: 15px; }

input {
  background: #020617; border: 1px solid #334155;
  padding: 12px; border-radius: 12px; color: white; outline: none;
}
input:focus { border-color: #10b981; }

button {
  background: #10b981; color: #020617; font-weight: bold;
  padding: 12px; border-radius: 12px; border: none; cursor: pointer; transition: 0.3s;
}
button:hover:not(:disabled) { background: #34d399; transform: scale(1.02); }
button:disabled { background: #1e293b; color: #64748b; cursor: not-allowed; }

.status-msg { margin-top: 15px; font-size: 0.8rem; }
.error { color: #ef4444; }
.success { color: #10b981; }
</style>