import { createApp } from 'vue'
import App from './App.vue'
import RegisterForm from './components/RegisterForm.vue'
import LoginForm from './components/LoginForm.vue'
import HomePage from './components/HomePage.vue'
import { createRouter, createWebHistory } from 'vue-router'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap'

// set routes
const routes = [
  { path: '/register', component: RegisterForm },
  { path: '/login', component: LoginForm },
  { path: '/home', component: HomePage },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

const app = createApp(App)
app.use(router)
app.mount('#app')
