import Vue from 'vue'
import VueGeolocation from 'vue-browser-geolocation'
import * as VueGoogleMaps from 'vue2-google-maps'

import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

import App from './App.vue'


Vue.use(BootstrapVue)
Vue.config.productionTip = false
Vue.use(VueGeolocation)
Vue.use(VueGoogleMaps,  {
  load: {
    key: 'AIzaSyDvzBG1YC-EqqOau_4BMMAgK7p-t9nrYjE',
    autobindAllEvents: true,
  },
  installComponents: true,
})

new Vue({
  render: h => h(App),
}).$mount('#app')
