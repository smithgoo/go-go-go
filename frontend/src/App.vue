<template>
  <div id="app">
    <h1>My App</h1>
    <UserLogin />
  </div>
</template>

<script>
import UserLogin from './components/UserLogin.vue';

export default {
  name: 'App',
  components: {
    UserLogin
  }
};
</script>

<style>
/* Add your styles here */
</style>
