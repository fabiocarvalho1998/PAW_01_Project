import Vue from 'vue'
import VueGeolocation from 'vue-browser-geolocation'
import * as VueGoogleMaps from 'vue2-google-maps'
import App from './App.vue'
//import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'

//import 'bootstrap/dist/css/bootstrap.css'
//import 'bootstrap-vue/dist/bootstrap-vue.css'

//Vue.use(BootstrapVue)
//Vue.use(IconsPlugin)

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
