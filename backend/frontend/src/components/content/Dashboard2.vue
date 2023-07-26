<template lang="">


  <div class="p-4 my-12 sm:ml-64">
    <div id="cont1" class="bg-blue-100">
   this message is from the backend
   {{ mydata.message }}
   </div>

</div>

   
</template>
<script>
import { Collapse } from 'flowbite';

import axios from 'axios';

export default {
  name: 'Dashboard2',
  data() {
    return {
      mydata: {},
    }
},
    created() {
  this.populateProfile()
    },
methods: {
  populateProfile() {
    // Call the measurements endpoint and check if response is empty or not
    axios.get('http://127.0.0.1:8080/api/v1/test2')
      .then(response => {
        console.log(response.data)
        this.mydata = response.data;
        console.log(response.data)
      })
      .catch(error => {
        console.log(error);
      });
  }
}
  
    
    };



</script>
<style lang="">
    
</style>