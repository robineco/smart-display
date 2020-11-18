<template>
  <div id="app">
    <Datetime />
    <div id="content">
      <TemperatureWidget id="temperature" />
      <div id="placeholder"></div>
        <div id="widget-wrapper">
          <CalenderWidget id="calender" />
          <NewsWidget id="news" />
        </div>
    </div>
  </div>
</template>

<script>
import Datetime from "./components/Datetime.vue";
import TemperatureWidget from "./components/temperature/TemperatureWidget.vue";
import CalenderWidget from "./components/calender/CalenderWidget.vue";
import NewsWidget from "./components/news/NewsWidget.vue";

export default {
  name: 'App',
  components: {
    Datetime,
    TemperatureWidget,
    CalenderWidget,
    NewsWidget,
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
#widget-wrapper {
  display: inline-block;
  vertical-align: top;
  width: 25%;
  text-align: center;
}
#calender {
  width: 100%;
  display: inline-block;
  margin-bottom: 2rem;
}
#news {
  width: 100%;
  display: inline-block;
}

</style>
