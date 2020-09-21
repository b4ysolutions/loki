<template>
  <div id="app">
    <Menu />
    <Main msg="Welcome to Loki Webservice Documentation"/>
    <notifications group="loki" position="top center" />
  </div>
</template>

<script>
import Menu from './components/Menu.vue'
import Main from './components/Main.vue'

export default {
  name: 'App',
  components: {
    Main, Menu
  }
}
</script>

<style scoped>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>

<style  lang="scss">
:root {
    --color-fun-green: #008b3a;
    --color-crusoe: #024205;
    --color-mine-shaft: #252525;
    --color-wild-sand: #f5f5f5;
  }
.main,
body {
  padding-top: 80px;
}
.loki-info {
  min-height: calc(100vh - 280px);
}
.top-title {
  color: var(--color-fun-green);
  margin: 60px auto 80px auto;
}
.main-title {
  color: var(--color-mine-shaft) !important;
  font-weight: bold;
  font-size: 23px;
  padding: 14px 0;
  margin: 140px 0 0 0;
}
.container {
  h3 {
    color: var(--color-mine-shaft);
    font-weight: bold;
    font-size: 18px;
    text-align: left;
    margin: 70px auto 20px auto;
  }
  h4 {
    color: var(--color-mine-shaft);
    font-weight: bold;
    font-size: 23px;
    text-align: left;
    margin: 70px auto 20px auto;
  }
  h5 {
    color: var(--color-mine-shaft);
    font-weight: bold;
    font-size: 16px;
    text-align: left;
    margin: 70px auto 20px auto;
  }
  p {
    color: var(--color-mine-shaft);
    font-size: 15px;
    text-align: left;
  }
}
div.code {
  text-align: left;
  border-left: 5px solid var(--color-crusoe);
  color: var(--color-crusoe);
  padding: 8px 8px 8px 10px;
  margin-bottom: 20px;
  background: var(--color-wild-sand);

  .code-json {
    min-height: 140px;
  }
}
.loki-short-description {
  margin-bottom: 150px;
}
.loki-short-description p {
  text-align: center;
}
.example-col {
  text-align: left;
}
.examples-row {
  margin-bottom: 20px;
}
.copy-btn {
  border: 0;
  background: #ffffff;
  margin: -4px 20px 0 0;
  padding: 2px 6px;
  color: var(--color-crusoe);
}
p {
  margin: 0;
}

</style>
