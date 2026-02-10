import type { VoiceAuth, BackendResponse } from '@/types/auth';

const API_URL = "https://symmetrical-acorn-5xgw65wwgr6f4wrx-8080.app.github.dev";


export const authService = {
    async sendVoiceHash(data: VoiceAuth): Promise<BackendResponse> {
        const response = await fetch(`${API_URL}/auth/voice-hash`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data), // TS verifica que 'data' sea tipo VoiceAuth
        });

        if (!response.ok) {
            throw new Error('Error en la comunicación soberana');
        }

        // Aquí le decimos a TS que la respuesta tiene la forma de BackendResponse
        return await response.json() as BackendResponse;
    }
};
