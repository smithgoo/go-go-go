<template>
  <div>
    <h1>Login</h1>
    <form @submit.prevent="login">
      <label>Email:</label>
      <input type="email" v-model="email" required><br>
      <label>Password:</label>
      <input type="password" v-model="password" required><br>
      <div id="captcha-container">
        <img :src="captchaImageSrc" alt="Captcha Image" v-if="captchaImageSrc">
      </div>
      <label for="captcha">Captcha:</label>
      <input type="text" v-model="captchaAnswer" id="captchaAnswer" name="captcha_answer" required><br><br>
      <input type="hidden" v-model="captchaID" id="captchaID" name="captcha_id">
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
      captchaImageSrc: '',
      captchaID: '',
      captchaAnswer: '',
      message: '',
      error: ''
    };
  },
  created() {
    console.log('Component created. Fetching captcha...');
    this.fetchCaptcha();
  },
  methods: {
    login() {
      const formData = new FormData();
      formData.append('email', this.email);
      formData.append('password', this.password);
      formData.append('captcha_id', this.captchaID);
      formData.append('captcha_answer', this.captchaAnswer);

      axios.post('/user/login', formData)
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
    },
    fetchCaptcha() {
      axios.get('/user/captcha')
        .then(response => {
          console.log('Captcha data:', response.data); // Debug log
          this.captchaImageSrc = response.data.captcha_image;
          this.captchaID = response.data.captcha_id;
        })
        .catch(error => {
          console.error('Error fetching captcha:', error);
          this.error = 'Error fetching captcha. Please try again later.';
        });
    }
  }
};
</script>

<style scoped>
/* Add your styles here */
</style>
