<template>
    <div>
        <h1>Welcome to the Home Page</h1>
    </div>
</template>

<script>
    import axios from './axios';
    // import storageService from './storage';
export default {

    created(){
        this.getAllList();
    },
    methods: {
       getAllList() {
        axios.get('/v1/products')
        .then(response => {
            console.log(response.data);
        })
        .catch(error => {
            console.log(error.message);
        })
       }
    }
   
}
 
</script>

<style scoped>
</style>