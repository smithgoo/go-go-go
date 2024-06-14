<template>
  <div id="app">
    <h1>My App</h1>
    <UserRegister />
  </div>
</template>

<script>
import UserRegister from './components/UserRegister.vue';

export default {
  name: 'App',
  components: {
    UserRegister
  }
};
</script>

<style>
/* Add your styles here */
</style>
