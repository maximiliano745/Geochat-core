<template>
  <div class="ia-control-panel">
    <h3>🧬 Centro de Mando Geochat </h3>
    <p>Estado: <span class="status">Esperando Firma Soberana</span></p>
    
    <input 
      v-model="passphrase" 
      type="password" 
      placeholder="Type Master Key..." 
      @keyup.enter="autorizar"
    />
    
    <button @click="autorizar" :disabled="loading">
      {{ loading ? 'Evolucionando...' : 'Autorizar Evolución' }}
    </button>

    <div v-if="mensaje" class="log">{{ mensaje }}</div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const passphrase = ref('')
const mensaje = ref('')
const loading = ref(false)

const autorizar = async () => {
  loading.value = true
  try {
    const response = await fetch('http://localhost:8080/api/admin/autorizar', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ passphrase: passphrase.value })
    })
    const data = await response.json()
    mensaje.value = data.mensaje || data.error
  } catch (e) {
    mensaje.value = "Error de conexión con el Core"
  } finally {
    loading.value = false
    passphrase.value = ''
  }
}
</script>
