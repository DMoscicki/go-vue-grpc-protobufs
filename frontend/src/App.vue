<template>
  <div id="app">
    <NavBar />
  </div>
</template>

<style Lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;
}

nav a.router-link-exact-active {
  color: #42b983;
}
</style>
<script setup>
import NavBar from "@/components/NavBar.vue";

</script>