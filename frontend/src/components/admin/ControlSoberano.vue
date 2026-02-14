<script setup lang="ts">
import { ref, onMounted, inject } from 'vue';

// 🔑 INYECTAMOS LA FUNCIÓN SOBERANA DESDE APP.VUE
const autorizarGlobal = inject<(id: string) => Promise<void>>('autorizarInversion');

const masterKeyInput = ref(''); // La clave que el usuario escribe
const mensaje = ref('');
const loading = ref(false);
const isError = ref(false);
const pendingTask = ref<any>(null);

onMounted(() => {
  const data = localStorage.getItem('pending_evolution');
  if (data) {
    pendingTask.value = JSON.parse(data);
  }
});

const autorizar = async () => {
  if (!autorizarGlobal) {
    mensaje.value = "Error: Sistema de firmas no inyectado";
    return;
  }

  loading.value = true;
  mensaje.value = '';
  
  try {
    // 🚀 USAMOS LA FUNCIÓN GLOBAL (App.vue)
    // Pasamos el ID del archivo de la tarea pendiente
    await autorizarGlobal(pendingTask.value?.archivo || "evolucion_manual.go");
    
    mensaje.value = "✅ Orden ejecutada correctamente";
    masterKeyInput.value = ''; 
    localStorage.removeItem('pending_evolution');
    setTimeout(() => { pendingTask.value = null; }, 2000);

  } catch (err: any) {
    isError.value = true;
    mensaje.value = "❌ Error en la firma";
  } finally {
    loading.value = false;
  }
};
</script>