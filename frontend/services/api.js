// services/api.js
import axios from 'axios'

const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL, // ganti ini sesuai alamat backend Go kamu
})

export default api
