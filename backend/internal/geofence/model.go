package geofence

type PuntoInteres struct {
    // 1. LUGAR (La identidad que la IA descubre)
    ID_IA        string  `json:"id_ia"` 
    
    // 2. NOMBRE (Lo que el usuario ve y pone)
    NombreTag    string  `json:"nombre_tag"` 
    
    // 3. CÓDIGO (Coordenadas y Radio - La base técnica)
    Latitud      float64 `json:"lat"`
    Longitud     float64 `json:"lng"`
    RadioMetros  int     `json:"radio"`
    
    // El "Sensor": Cuántos usuarios hay aquí ahora
    ContadorReal int     `json:"presencia_actual"`
}