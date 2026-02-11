<script setup lang="ts">

import { ref, onMounted } from 'vue'

import axios from 'axios'



// Estructura de la Propuesta de Código de la IA

interface CodeProposal {

  id: string;

  funcionalidad: string;

  codigoPropuesto: string;

  analisisImpacto: string;

  costoGas: number;

}



const API_BASE = 'https://symmetrical-acorn-5xgw65wwgr6f4wrx-8080.app.github.dev'



// Estado del Dashboard

const gasUnidades = ref(300) // Tus 300 unidades de "combustible"

const propuestaActiva = ref<CodeProposal | null>(null)

const firmaVocalActiva = ref(false)

const hashFirma = ref("")

const logsIA = ref<string[]>(["IA CEO: Sistema en espera de órdenes..."])



// 1. Obtener propuesta de la IA (Conexión con Vertex/Go)

const buscarPropuesta = async () => {

  try {

    logsIA.value.unshift("IA CEO: Analizando código en /internal/core...")

    const res = await axios.get(`${API_BASE}/api/ceo/laboratorio`)

    propuestaActiva.value = res.data

    logsIA.value.unshift(`IA CEO: Propuesta ${res.data.id} generada con éxito.`)

  } catch (e) {

    logsIA.value.unshift("⚠️ Error: El Nodo de IA no responde.")

  }

}



// 2. Protocolo de Validación (Tu Firma es la Ley)

const autorizarEjecucion = async () => {

  if (hashFirma.value !== "MI_FIRMA_ES_LA_LEY") {

    alert("Firma Inválida. Solo el Dueño puede autorizar.")

    return

  }



  try {

    logsIA.value.unshift("🚀 Ejecutando despliegue en /internal/core...")

    const res = await axios.post(`${API_BASE}/api/ceo/ejecutar`, {

      id: propuestaActiva.value?.id,

      firma: hashFirma.value

    })

    

    if (res.data.success) {

      gasUnidades.value -= propuestaActiva.value?.costoGas || 0

      logsIA.value.unshift("✅ Código aplicado. Reiniciando servidor core...")

      propuestaActiva.value = null

      hashFirma.value = ""

    }

  } catch (e) {

    logsIA.value.unshift("❌ Error crítico en el despliegue.")

  }

}

</script>



<template>

  <div class="ceo-panel">

    <header class="gas-tank">

      <div class="gas-meter">

        <span class="label">COMBUSTIBLE IA (GAS):</span>

        <div class="progress-bar">

          <div class="fill" :style="{ width: (gasUnidades / 300) * 100 + '%' }"></div>

        </div>

        <span class="value">{{ gasUnidades }} / 300 U</span>

      </div>

    </header>



    <main class="workspace">

      <section class="code-viewer" v-if="propuestaActiva">

        <div class="card-header">

          <h3>📂 Propuesta: {{ propuestaActiva.funcionalidad }}</h3>

          <span class="impact-tag">{{ propuestaActiva.analisisImpacto }}</span>

        </div>

        

        <pre class="code-block"><code>{{ propuestaActiva.codigoPropuesto }}</code></pre>



        <div class="auth-zone">

          <input 

            v-model="hashFirma" 

            placeholder="Escribe tu Firma de Mando..." 

            class="signature-input"

          />

          <button @click="autorizarEjecucion" class="btn-gold">

            ESTAMPAR FIRMA Y DESPLEGAR

          </button>

        </div>

      </section>



      <section class="empty-state" v-else>

        <p>No hay propuestas pendientes. ¿Quieres que la IA analice el sistema?</p>

        <button @click="buscarPropuesta" class="btn-scan">ESCANEAR NODO</button>

      </section>



      <aside class="system-logs">

        <h4>REGISTRO DE IA EMPRESARIA</h4>

        <div class="log-list">

          <p v-for="(log, i) in logsIA" :key="i">> {{ log }}</p>

        </div>

      </aside>

    </main>

  </div>

</template>



<style scoped>

.ceo-panel {

  background: #0a0a0a;

  color: #e0e0e0;

  padding: 30px;

  font-family: 'Space Mono', monospace;

}



.gas-tank {

  background: #1a1a1a;

  padding: 15px;

  border-radius: 10px;

  margin-bottom: 25px;

  border-left: 5px solid #d4af37;

}



.progress-bar {

  height: 10px;

  background: #333;

  width: 200px;

  border-radius: 5px;

  display: inline-block;

  margin: 0 15px;

}



.fill {

  height: 100%;

  background: linear-gradient(90deg, #d4af37, #f1c40f);

  box-shadow: 0 0 10px #d4af37;

}



.code-block {

  background: #000;

  padding: 20px;

  border-radius: 8px;

  border: 1px solid #333;

  color: #00ff41; /* Verde Matrix */

  overflow-x: auto;

  max-height: 400px;

}



.btn-gold {

  background: #d4af37;

  color: black;

  font-weight: bold;

  padding: 15px 30px;

  border: none;

  cursor: pointer;

  border-radius: 5px;

  transition: 0.3s;

}



.btn-gold:hover {

  background: #f1c40f;

  box-shadow: 0 0 15px #d4af37;

}



.signature-input {

  background: #222;

  border: 1px solid #d4af37;

  color: white;

  padding: 15px;

  margin-right: 10px;

  width: 300px;

}



.system-logs {

  margin-top: 30px;

  background: #111;

  padding: 15px;

  border-radius: 8px;

  font-size: 0.8rem;

}



.log-list {

  height: 100px;

  overflow-y: auto;

  color: #888;

}

</style>