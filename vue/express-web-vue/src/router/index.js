import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/send', component: () => import('../views/Send.vue'),
  },
  {
    path: '/receive', component: () => import('../views/Receive.vue'),
  },
  {
    path: '/query', component: () => import('../views/Query.vue'),
  },
  {
    path: '/login', component: () => import('../views/Login.vue'),
  },
  {
    path: '/sign-up', component: () => import('../views/SignUp.vue'),
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  }
]

function initNames(list) {
  for (let i in list) {
    let item = list[i]
    let name = item.name
    if (name == null) {
      item.name = item.path
    }
  }
}
initNames(routes)

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
