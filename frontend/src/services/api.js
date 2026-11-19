const BASE = 'http://localhost:8080/api'
async function request(path, options = {}) {
  const response = await fetch(`${BASE}${path}`, {
    headers: { 'Content-Type': 'application/json', ...(options.headers || {}) },
    ...options
  })
  const data = response.status === 204 ? null : await response.json()
  if (!response.ok) throw new Error(data?.error || 'Error en la solicitud')
  return data
}
export const api = {
  autores: () => request('/autores'),
  crearAutor: (data) => request('/autores', { method: 'POST', body: JSON.stringify(data) }),
  actualizarAutor: (id, data) => request(`/autores/${id}`, { method: 'PUT', body: JSON.stringify(data) }),
  eliminarAutor: (id) => request(`/autores/${id}`, { method: 'DELETE' }),
  libros: () => request('/libros'),
  crearLibro: (data) => request('/libros', { method: 'POST', body: JSON.stringify(data) }),
  actualizarLibro: (id, data) => request(`/libros/${id}`, { method: 'PUT', body: JSON.stringify(data) }),
  eliminarLibro: (id) => request(`/libros/${id}`, { method: 'DELETE' }),
  resumen: () => request('/reportes/inventario')
}
