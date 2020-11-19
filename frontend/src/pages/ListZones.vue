<template>
  <div class="myDiv">
    <br>
    <div class="myBtn">
      <b-button variant="success">
        Logout
      </b-button>
    </div>
    <br><br><br>
    <b-jumbotron>
      <template slot="lead">
        <h3>
          <b>
            Zones
          </b>
        </h3>
        <hr class="my-2" />
        <b-list-group>
        <b-list-group-item>
          <h5>Zone A</h5>
          <b-progress :value="value+5" show-value :max="max" class="mb-3"></b-progress>
        </b-list-group-item>
        <b-list-group-item>
          <h5>Zone B</h5>
          <b-progress :value="value+20" show-value :max="max" class="mb-3"></b-progress>
        </b-list-group-item>
        <b-list-group-item>
          <h5>Zone C</h5>
          <b-progress :value="value+8" show-value :max="max" class="mb-3"></b-progress>
        </b-list-group-item>
      </b-list-group>
      </template>
    </b-jumbotron>
  </div>
</template>

<script>
export default {
  data(){
    return{
      value: 33.333,
      max:100
    }
  }
}
</script>

<style scoped>
.list-group {
  margin-bottom: 15px;
}
.list-group-item:hover {
  background: #EEE;
  cursor: pointer;
}
.myBtn{
  position: absolute; 
  right: 0;
  width: 38%;
  margin: auto;
}
.myDiv{
  width: 70%;
  align-content: center;
  margin: auto;
}
</style>