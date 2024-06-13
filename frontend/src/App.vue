<template>
  <div>
    <h1>Login</h1>
    <form @submit.prevent="login">
      <label>Email:</label>
      <input type="email" v-model="email" required><br>
      <label>Password:</label>
      <input type="password" v-model="password" required><br>
      <button type="submit">Login</button>
    </form>
    <p v-if="message">{{ message }}</p>
    <p v-if="error">{{ error }}</p>
  </div>
</template>

<script>
import axios from './axios';

export default {
  data() {
    return {
      email: '',
      password: '',
      message: '',
      error: ''
    };
  },
  methods: {
    login() {
      axios.post('/user/login', {
        email: this.email,
        password: this.password
      })
      .then(response => {
        console.log(response.data);
        this.message = response.data.token;
        this.error = '';
      })
      .catch(error => {
        this.message = '';
        this.error = 'Login failed. Please check your credentials.';
        console.error('Error logging in:', error);
        console.error(this.email);
        console.error(this.password);
      });
    }
  }
};
</script>
