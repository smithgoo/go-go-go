<template>
  <div>
    <h1>Register</h1>
    <form @submit.prevent="register">
      <van-cell-group>
        <van-field v-model="email" label="Email" type="email" placeholder="Enter your email" required />
        <van-field v-model="phone" label="Phone" type="tel" placeholder="Enter your phone" required />
        <van-field v-model="password" label="Password" type="password" placeholder="Enter your password" required />
        <van-field v-model="captchaAnswer" label="CaptchaAnswer" type="text" placeholder="Enter your captchaAnswer" required />
        
        <div id="captcha-container">
          <img :src="captchaImageSrc" alt="Captcha Image" v-if="captchaImageSrc">
        </div>
        <input type="hidden" v-model="captchaID" id="captchaID" name="captcha_id">
        <van-button type="primary" native-type="submit">Register</van-button>
      </van-cell-group>
    </form>
    <p v-if="message">{{ message }}</p>
    <p v-if="error">{{ error }}</p>
  </div>
</template>

<script>
import axios from './axios'; 
import { Toast } from 'vant';

export default {
  data() {
    return {
      email: '',
      phone: '',
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
    register() {
      const formData = new FormData();
      formData.append('email', this.email);
      formData.append('phone', this.phone);
      formData.append('password', this.password);
      formData.append('captcha_id', this.captchaID);
      formData.append('captcha_answer', this.captchaAnswer);

      axios.post('/user/register', formData)
        .then(response => {
          console.log(response.data);
          Toast.success('Register Successful!');
          console.log('Navigating to login...'); // Debug log
          this.$router.push('/login')
            .then(() => {
              console.log('Navigation successful'); // Debug log
            })
            .catch(err => {
              console.error('Navigation error:', err); // Debug log
            });
        })
        .catch(error => {
          this.message = ''; 
          console.error('Error registering:', error.message); 
          Toast.fail(error.message);
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
