import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'
import router from './router'
import store from './store'

import MainFrame from './components/frames/mainframe/'
import VeriCodeInput from './components/accounts/verification-code.vue'

const app = createApp(App)
app.component("MainFrame", MainFrame)
app.component("VeriCodeInput", VeriCodeInput)
app.use(ElementPlus)
app.use(store).use(router)
app.mount('#app')
