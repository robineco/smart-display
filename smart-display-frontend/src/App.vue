<template>
  <div id="app">
    <Datetime />
    <div id="content">
      <TemperatureWidget id="temperature" />
      <div id="placeholder"></div>
      <CalenderWidget id="calender" />
    </div>
  </div>
</template>

<script>
import Datetime from "./components/Datetime.vue";
import TemperatureWidget from "./components/temperature/TemperatureWidget.vue";
import CalenderWidget from "./components/calender/CalenderWidget.vue";

export default {
  name: 'App',
  components: {
    Datetime,
    TemperatureWidget,
    CalenderWidget,
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
#content {
  margin-top: 2rem;
}
#temperature {
  display: inline-block;
  width: 25%;
}
#placeholder {
  display: inline-block;
  vertical-align: top;
  width: 50%;
  height: 30rem;
}
#calender {
 display: inline-block;
 vertical-align: top;
 width: 25%;
}
</style>
