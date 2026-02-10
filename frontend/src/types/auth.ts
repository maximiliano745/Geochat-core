// src/types/auth.ts

export interface VoiceAuth {
    user_hash: string;
    device_id: string;
}

export interface BackendResponse {
    status: string;
    info: string;
    tesla_mode?: boolean; // El '?' significa que puede o no venir del backend
}
