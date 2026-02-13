import axios from 'axios';

// URL de tu Codespace de GitHub (Asegúrate de que el puerto 8080 esté en "Public" en la pestaña Ports)
const API_BASE = 'https://symmetrical-acorn-5xgw65wwgr6f4wrx-8080.app.github.dev/api/admin';

export const evolutionService = {
  /**
   * Envía la Firma Soberana para autorizar que la IA 5 evolucione el ADN
   * @param key La Master Key ingresada por el usuario
   */
  async sendSovereignSignature(key: string) {
    try {
      // Usamos API_BASE que definimos arriba
      const response = await axios.post(`${API_BASE}/authorize`, {
        passphrase: key // Coincide con el struct 'input' de tu Go
      });
      
      return response.data; // Retorna: { mensaje: "Evolución iniciada, Jefe." }
    } catch (error: any) {
      console.error("Error en la firma:", error);
      
      // Manejo de errores específicos según la respuesta del backend
      if (error.response) {
        if (error.response.status === 401) {
          throw new Error("Firma inválida. Acceso Denegado.");
        }
        if (error.response.status === 404) {
          throw new Error("Endpoint no encontrado. Verifica la URL del Codespace.");
        }
      }
      throw new Error("Error de comunicación con el núcleo de GeoChat");
    }
  }
};