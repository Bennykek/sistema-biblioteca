import { useEffect, useState } from 'react'
import { api } from './services/api'

const emptyAutor = { nombre: '', nacionalidad: '' }
const emptyLibro = { titulo: '', isbn: '', anio: new Date().getFullYear(), autor_id: '' }

export default function App() {
  const [tab, setTab] = useState('libros')
  const [autores, setAutores] = useState([])
  const [libros, setLibros] = useState([])
  const [autorForm, setAutorForm] = useState(emptyAutor)
  const [libroForm, setLibroForm] = useState(emptyLibro)
  const [editAutor, setEditAutor] = useState(null)
  const [editLibro, setEditLibro] = useState(null)
  const [error, setError] = useState('')
  const [resumen, setResumen] = useState(null)

  async function cargar() {
    try {
      const [a, l, r] = await Promise.all([api.autores(), api.libros(), api.resumen()])
      setAutores(a); setLibros(l); setResumen(r); setError('')
    } catch (e) { setError(e.message) }
  }
  useEffect(() => { cargar() }, [])

  async function guardarAutor(e) {
    e.preventDefault()
    try {
      if (editAutor) await api.actualizarAutor(editAutor, autorForm)
      else await api.crearAutor(autorForm)
      setAutorForm(emptyAutor); setEditAutor(null); await cargar()
    } catch (e) { setError(e.message) }
  }
  async function guardarLibro(e) {
    e.preventDefault()
    const payload = { ...libroForm, anio: Number(libroForm.anio), autor_id: Number(libroForm.autor_id) }
    try {
      if (editLibro) await api.actualizarLibro(editLibro, payload)
      else await api.crearLibro(payload)
      setLibroForm(emptyLibro); setEditLibro(null); await cargar()
    } catch (e) { setError(e.message) }
  }
  async function borrar(tipo, id) {
    if (!confirm('¿Seguro que deseas eliminar este registro?')) return
    try { tipo === 'libro' ? await api.eliminarLibro(id) : await api.eliminarAutor(id); await cargar() }
    catch (e) { setError(e.message) }
  }
  function editarAutor(a) { setEditAutor(a.id); setAutorForm({ nombre: a.nombre, nacionalidad: a.nacionalidad }) }
  function editarLibro(l) { setEditLibro(l.id); setLibroForm({ titulo: l.titulo, isbn: l.isbn, anio: l.anio, autor_id: String(l.autor_id) }) }

  return <div className="app">
    <header><div><span className="brand">B</span><div><h1>Biblioteca</h1><p>Sistema de gestión de libros y autores</p></div></div></header>
    <main>
      {error && <div className="error">{error}<button onClick={() => setError('')}>×</button></div>}
      <section className="stats">
        <article><span>Libros registrados</span><strong>{libros.length}</strong></article>
        <article><span>Autores registrados</span><strong>{autores.length}</strong></article>
        <article><span>Año más antiguo</span><strong>{resumen?.anio_mas_antiguo || '—'}</strong></article>
      </section>
      <nav><button className={tab === 'libros' ? 'active' : ''} onClick={() => setTab('libros')}>Libros</button><button className={tab === 'autores' ? 'active' : ''} onClick={() => setTab('autores')}>Autores</button></nav>

      {tab === 'libros' ? <section className="layout">
        <form className="panel" onSubmit={guardarLibro}>
          <h2>{editLibro ? 'Editar libro' : 'Registrar libro'}</h2>
          <label>Título<input required value={libroForm.titulo} onChange={e => setLibroForm({...libroForm, titulo:e.target.value})}/></label>
          <label>ISBN<input required value={libroForm.isbn} onChange={e => setLibroForm({...libroForm, isbn:e.target.value})}/></label>
          <label>Año de publicación<input required type="number" min="1000" max="2100" value={libroForm.anio} onChange={e => setLibroForm({...libroForm, anio:e.target.value})}/></label>
          <label>Autor<select required value={libroForm.autor_id} onChange={e => setLibroForm({...libroForm, autor_id:e.target.value})}><option value="">Selecciona un autor</option>{autores.map(a => <option key={a.id} value={a.id}>{a.nombre}</option>)}</select></label>
          <button className="primary">{editLibro ? 'Guardar cambios' : 'Registrar libro'}</button>
          {editLibro && <button type="button" className="secondary" onClick={() => {setEditLibro(null);setLibroForm(emptyLibro)}}>Cancelar</button>}
          {autores.length === 0 && <small>Primero registra un autor en la pestaña Autores.</small>}
        </form>
        <div className="panel table-panel"><div className="panel-title"><h2>Listado de libros</h2><span>{libros.length} registros</span></div>
          <div className="table-wrap"><table><thead><tr><th>Título</th><th>ISBN</th><th>Año</th><th>Autor</th><th>Acciones</th></tr></thead>
            <tbody>{libros.map(l => <tr key={l.id}><td>{l.titulo}</td><td>{l.isbn}</td><td>{l.anio}</td><td>{l.autor}</td><td className="actions"><button onClick={() => editarLibro(l)}>Editar</button><button className="danger" onClick={() => borrar('libro',l.id)}>Eliminar</button></td></tr>)}
            {libros.length === 0 && <tr><td colSpan="5" className="empty">No hay libros registrados.</td></tr>}</tbody></table></div>
        </div>
      </section> : <section className="layout">
        <form className="panel" onSubmit={guardarAutor}>
          <h2>{editAutor ? 'Editar autor' : 'Registrar autor'}</h2>
          <label>Nombre completo<input required value={autorForm.nombre} onChange={e => setAutorForm({...autorForm,nombre:e.target.value})}/></label>
          <label>Nacionalidad<input value={autorForm.nacionalidad} onChange={e => setAutorForm({...autorForm,nacionalidad:e.target.value})}/></label>
          <button className="primary">{editAutor ? 'Guardar cambios' : 'Registrar autor'}</button>
          {editAutor && <button type="button" className="secondary" onClick={() => {setEditAutor(null);setAutorForm(emptyAutor)}}>Cancelar</button>}
        </form>
        <div className="panel table-panel"><div className="panel-title"><h2>Listado de autores</h2><span>{autores.length} registros</span></div>
          <div className="table-wrap"><table><thead><tr><th>Nombre</th><th>Nacionalidad</th><th>Acciones</th></tr></thead>
            <tbody>{autores.map(a => <tr key={a.id}><td>{a.nombre}</td><td>{a.nacionalidad || '—'}</td><td className="actions"><button onClick={() => editarAutor(a)}>Editar</button><button className="danger" onClick={() => borrar('autor',a.id)}>Eliminar</button></td></tr>)}
            {autores.length === 0 && <tr><td colSpan="3" className="empty">No hay autores registrados.</td></tr>}</tbody></table></div>
          <small>Un autor asociado a uno o más libros no puede eliminarse. Primero reasigna o elimina sus libros.</small>
        </div>
      </section>}
    </main>
  </div>
}
