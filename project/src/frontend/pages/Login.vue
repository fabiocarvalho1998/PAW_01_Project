<template>
 <div>
   <form class="login">
     <h1>Sign in</h1>
     <label>User name</label>
     <input required v-model="username" type="text" placeholder="Username"/>
     <h1>{{username}}</h1>
     <hr/>
     <label>Password</label>
     <input required v-model="password" type="password" placeholder="Password"/>
     <h1>{{password}}</h1>
     <hr/>
     <button v-on:click.prevent="handlerSubmit">Login</button>
   </form>
 </div>
</template>

<script>

export default {
  name: 'Login',
  data(){
    return {
      username: this.username,
      password: this.password
    };
  },
 methods: {
    handlerSubmit: function(){
      this.$http.post('http://localhost:8081/login', {  //não aceita endereço docker pois é no browser. Como porta é diferente origina problema de cors
          username: this.username,
          password: this.password
        }).then(response => {
          let is_admin = response.data.user.is_admin
          localStorage.setItem('user',JSON.stringify(response.data.user))
          localStorage.setItem('jwt',response.data.token)

          if (localStorage.getItem('jwt') != null){
            if(is_admin== 1){
              this.$router.push('admin')
            }else {
              this.$router.push('map')
            }
          }
        }).catch(function (error) {
          console.error(error.response);
        });
    }
  }
}
</script>