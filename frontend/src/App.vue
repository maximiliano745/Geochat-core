<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'

const saldoOro = ref(0)
const nombreNodo = ref('Cargando...')
const mensajes = ref([])
const modoTesla = ref(false)

const actualizarDatos = async () => {
  try {
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
  // Aquí es donde el usuario activa manualmente para sentirse "filantrópico" [cite: 2026-02-02]
}

onMounted(actualizarDatos)
</script>
