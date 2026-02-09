# GeoChat Core 🌍🛰️

**GeoChat** es una plataforma de comunicación y gestión de activos soberana que distribuye el riesgo en lugar de concentrarlo. Diseñada bajo la premisa de "Tu llave, tus datos", GeoChat elimina el concepto de "Tesorería Central" y garantiza la privacidad total del usuario.

---

## 🛡️ Arquitectura de Seguridad
GeoChat ha sido diseñado con un enfoque en la soberanía digital:

* **Sovereign Digital Vault:** Un espacio seguro y cifrado de extremo a extremo (E2E) para almacenar semillas (seeds) y documentos. Solo el usuario posee la llave de descifrado.
* **Sin Custodia:** GeoChat nunca tiene acceso a las llaves privadas de los usuarios. Los fondos (PAXG y otros activos) residen directamente en la blockchain, controlados únicamente por el usuario.
* **E2E Encryption:** Todas las comunicaciones y documentos están protegidos contra ataques *Man-in-the-Middle*.
* **Account Recovery (DID/VC):** Recuperación de cuentas mediante Credenciales Verificables, mitigando el riesgo de pérdida sin depender de terceros centralizados.

---

## 🚀 Módulos del Sistema

La estructura del núcleo se divide en cuatro pilares fundamentales:

1.  **`src/vault/` (Seguridad y Custodia):** Gestión de la bóveda digital soberana y lógica de cifrado.
2.  **`src/mesh/` (Comunicación Híbrida):** Protocolos para comunicación vía Radio y redes Mesh, garantizando conectividad en cualquier entorno.
3.  **`src/ai/` (AI Friend):** Asistente inteligente con lógica local. Puede conocer el perfil inversor del usuario, pero tiene restringido el acceso a las llaves privadas o documentos del Vault.
4.  **`src/tesla/` (Energía y Filantropía):** Gestión de recursos y el **"Modo Tesla"**.

---

## 🔋 Modo Tesla (Filantropía Activa)
El **Modo Tesla** es una funcionalidad única de GeoChat que el usuario activa de forma **manual**. 
* **Propósito:** Permitir que el usuario actúe de forma filantrópica dentro del ecosistema.
* **Control:** No es automático; requiere la voluntad expresa del usuario para ser activado, reforzando el compromiso personal con la red y la comunidad.

---

## 🛠️ Estructura del Proyecto
```text
Geochat-core/
├── src/
│   ├── vault/     # Seguridad E2E
│   ├── mesh/      # Comunicación Radio/Híbrida
│   ├── tesla/     # Gestión de Energía y Pagos
│   ├── ai/        # AI Friend local
│   └── index.ts   # Punto de entrada
├── data/          # DB local encriptada
└── .env           # Configuración sensible (Ignorado en Git)
