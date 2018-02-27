import Vue from 'vue'
import 'bulma/css/bulma.css'
import ElementUI from 'element-ui'
import locale from 'element-ui/lib/locale/lang/en'
import 'element-ui/lib/theme-chalk/index.css'
import Vuex from 'vuex'
import App from './App.vue'
import store from './store/store.js'

Vue.use(ElementUI, {
  locale
})

new Vue({
  el: '#app',
  store,
  render: h => h(App)
})
