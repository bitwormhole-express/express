import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'

const routes = [{
    path: '/',
    name: 'Home',
    component: Home
},
{
    path: '/about',
    component: () =>
        import('../views/About.vue')
},
{
    path: '/forget-password',
    component: () =>
        import('../views/ForgetPassword.vue'),
},
{
    path: '/login',
    component: () =>
        import('../views/Login.vue'),
},
{
    path: '/packages/:id',
    component: () => import('../views/PackageDetail.vue')
},
{
    path: '/query',
    component: () =>
        import('../views/Query.vue'),
},
{
    path: '/receive',
    component: () =>
        import('../views/Receive.vue'),
},
{
    path: '/send',
    component: () =>
        import('../views/Send.vue'),
},
{
    path : '/settings' , 
    component : () =>import('../views/Settings.vue') ,
},
{
    path : '/settings/buckets' , 
    component : () =>import('../views/settings/buckets.vue') ,
},
{
    path: '/sign-up',
    component: () =>
        import('../views/SignUp.vue'),
},
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