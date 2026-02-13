<template>
  <transition name="pop">
    <div v-if="visible" class="burbuja-flotante" :class="roleClass">
      <div class="burbuja-header">
        <div class="ia-status">
          <span class="pulso-ia"></span>
          <h4>{{ headerTitle }}</h4>
        </div>
        <button @click="visible = false" class="close-mini">×</button>
      </div>
      
      <div class="burbuja-body">
        <p class="mensaje">{{ propuesta.mensaje }}</p>
        <div class="proyeccion">
          <small>Impacto estimado:</small>
          <span class="impacto-tag">{{ propuesta.impacto }}</span>
        </div>
      </div>

      <div class="burbuja-actions">
        <button class="btn-primary" @click="ejecutar">
          {{ isSovereign ? 'FIRMAR AHORA' : 'ACEPTAR' }}
        </button>
        <button class="btn-secondary" @click="posponer">PENSAR LUEGO</button>
      </div>
    </div>
  </transition>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';

const router = useRouter();
const route = useRoute();
const visible = ref(true);

// Detectamos el rol según la URL (Fractalidad)
const isSovereign = computed(() => route.path.includes('admin'));
const roleClass = computed(() => isSovereign.value ? 'role-sovereign' : 'role-citizen');
const headerTitle = computed(() => isSovereign.value ? 'GeoChat CEO: Estrategia' : 'IA Amiga: Beneficio');

// Propuesta dinámica (esto vendrá del Backend via API/WebSocket)
const propuesta = ref({
  id: isSovereign.value ? 'DEC-CEO-001' : 'BEN-USR-001',
  mensaje: isSovereign.value 
    ? "Baja en PAXG detectada. ¿Reforzamos reserva con el 15% del fondo?" 
    : "Hay energía solar disponible en tu manzana. ¿Quieres comprar a precio reducido?",
  impacto: isSovereign.value ? "+1.2% Liquidez" : "-15% Costo Eléctrico",
  modulo: isSovereign.value ? "Finanzas" : "Energía Local"
});

const ejecutar = () => {
  if (isSovereign.value) {
    // Si eres tú, guardamos el "contrato" y vamos a firmar
    localStorage.setItem('pending_evolution', JSON.stringify({
      modulo: propuesta.value.modulo,
      archivo: 'logic_update.go',
      snippet: '// Inyección autorizada por Soberano'
    }));
    router.push('/admin');
  } else {
    // Si es un usuario, simplemente ejecuta la acción de la IA 1
    alert("¡Beneficio aceptado! La IA 1 está gestionando tu ahorro.");
    visible.value = false;
  }
};

const posponer = () => {
  const key = isSovereign.value ? 'decisiones_pendientes' : 'notificaciones_usuario';
  const pendientes = JSON.parse(localStorage.getItem(key) || '[]');
  pendientes.push({ ...propuesta.value, fecha: new Date().toISOString() });
  localStorage.setItem(key, JSON.stringify(pendientes));
  visible.value = false;
  console.log(`Guardado en: ${key}`);
};
</script>

<style scoped>
.burbuja-flotante {
  position: fixed; bottom: 30px; right: 30px; width: 340px;
  background: #0f172a; border-radius: 24px; padding: 22px; 
  z-index: 9999; box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.5);
  border: 1px solid #334155;
}

/* Diferenciación visual Fractal */
.role-sovereign { border: 2px solid #d97706; } /* Dorado para el Soberano */
.role-citizen { border: 2px solid #10b981; }   /* Verde para el Pueblo */

.burbuja-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 15px; }
.ia-status { display: flex; align-items: center; gap: 10px; }
.ia-status h4 { font-size: 0.85rem; margin: 0; font-family: 'Space Mono', monospace; }

.pulso-ia {
  width: 10px; height: 10px; border-radius: 50%;
  animation: pulse 2s infinite;
}
.role-sovereign .pulso-ia { background: #d97706; }
.role-citizen .pulso-ia { background: #10b981; }

.mensaje { font-size: 0.95rem; line-height: 1.5; color: #f1f5f9; margin-bottom: 15px; }

.proyeccion { background: #020617; padding: 10px; border-radius: 12px; margin-bottom: 20px; }
.impacto-tag { display: block; font-weight: bold; color: #10b981; }

.burbuja-actions { display: grid; grid-template-columns: 1.2fr 1fr; gap: 12px; }

button { padding: 12px; border-radius: 14px; border: none; font-weight: bold; cursor: pointer; transition: 0.3s; font-size: 0.8rem; }
.btn-primary { background: #10b981; color: #020617; }
.role-sovereign .btn-primary { background: #d97706; }

.btn-secondary { background: #1e293b; color: #94a3b8; }
.btn-secondary:hover { color: white; background: #334155; }

.close-mini { background: none; border: none; color: #475569; font-size: 1.5rem; cursor: pointer; padding: 0; }

@keyframes pulse {
  0% { transform: scale(0.95); opacity: 1; }
  70% { transform: scale(1.1); opacity: 0.5; }
  100% { transform: scale(0.95); opacity: 1; }
}

/* Animación de entrada */
.pop-enter-active { transition: all 0.3s ease-out; }
.pop-leave-active { transition: all 0.2s ease-in; }
.pop-enter-from { transform: scale(0.8) translateY(20px); opacity: 0; }
</style>