import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './assets/css/global.css'
import './plugins/element.js'
import axios from 'axios'

Vue.prototype.$http = axios
Vue.config.productionTip = false

axios.defaults.baseURL = `http://127.0.0.1:8888/api/`

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
