import axios from 'axios'
const API_URL = process.env.VUE_APP_API_HOST

const securedAxiosInstance = axios.create({
  baseURL: API_URL,
  // withCredentials: true,
  headers: {
    'Content-Type': `application/json`,
    // 'Access-Control-Allow-Origi': '*',
  },
})

const methodSetup = (config) => {
  const method = config.method.toUpperCase()
  if(method !== 'OPTIONS') {
    config.headers = {
      ...config.headers,
      // 'Authorization': localStorage.csrf
    }
  }
  return config
}
securedAxiosInstance.interceptors.request.use(methodSetup)

// const errorhandler = (error) => {
//   if(error.response && error.response.config && error.response.status === 401){
//       console.error(error)
//   } else {
//
//   }
// }

export { securedAxiosInstance }
