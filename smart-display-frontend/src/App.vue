<template>
  <div id="app">
    <Datetime />
  </div>
</template>

<script>
import Datetime from "./components/Datetime.vue"

export default {
  name: 'App',
  components: {
    Datetime,
  }
}
</script>

<style>
body {
  background: url("assets/images/earth.jpg") no-repeat center center fixed;
  -webkit-background-size: cover;
  -moz-background-size: cover;
  -o-background-size: cover;
  background-size: cover;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
}
</style>
