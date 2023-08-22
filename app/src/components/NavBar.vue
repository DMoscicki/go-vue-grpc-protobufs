<template>
  <nav class="navbar navbar-light" style="background-color: #42b983;">
    <div class="container-fluid">
      <div class="navbar-text">
        <v-img :src="require('@/assets/Tux.svg')" width="50" height="25" class="d-inline-block align-text-top" />
        Mościcki
      </div>
      <div class="d-flex justify-space-around align-center flex-column flex-md-row fill-height">
        <v-btn rounded="pill" color="primary" to="/" class="mx-12">
          <v-img :src="require('@/assets/gopher-logo-svg-vector.svg')" width="28" height="35" class="ml-auto" />
          Main</v-btn>
        <!-- <v-btn rounded="pill" color="primary" to="/about" class="mx-12">
          <v-img :src="require('@/assets/postgresql-icon.svg')" width="30" height="24" style="padding-left: 40px"/>
          About</v-btn> -->
        <v-btn rounded="pill" color="primary" to="/todo" class="mx-12">
          <v-img :src="require('@/assets/python.svg')" width="28" height="35" style="padding-left: 40px"/>
          Todo
        </v-btn>
        <v-btn rounded="pill" color="primary" to="/diag" class="mx-12">
          <v-img :src="require('@/assets/rocket-312767.svg')" width="30" height="24" style="padding-left: 40px"/>
          Diag
        </v-btn>
      </div>
    </div>
  </nav>
  <v-container>
    <router-view class="animate__animated animate__fadeInUp" />
  </v-container>
</template>

<script>
export default {
  name: "NavBar"
}
</script>

<style scoped>


.navbar-text {
  color: black;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  font-size: large;
}

.navbar {
  min-height: 150px;
}
</style>