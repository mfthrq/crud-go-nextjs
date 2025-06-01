// services/api.js
import axios from 'axios'

const api = axios.create({
  baseURL: 'http://localhost:8080', // ganti ini sesuai alamat backend Go kamu
})

export default api
