<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>

export default {
  name: 'App',
  watch: {
    $route() {
      window.scrollTo(0, 0);
    },
  },
};
</script>

<style>
/* Add your styles here */
</style>
