<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
   <Login />
   <!--<ListZones />-->
   <!-- <Map /> Só para testar enquanto não há rotas-->
  </div>
</template>

<script>
import Login from './frontend/pages/Login.vue'
//import Map from './frontend/pages/Map.vue' //Só para testar enquanto não há rotas
//import ListZones from './frontend/pages/ListZones.vue'

export default {
  name: 'App',
  components: {
    Login
    //ListZones
    //Map //Só para testar enquanto não há rotas
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
