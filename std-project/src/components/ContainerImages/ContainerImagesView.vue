<template>
    <div >
        <ul class=" flex justify-between gap-x-4">
            <CardImageView class=" w-1/3"  v-for="(x) in dataImages" :url-image="x"   />
        </ul>
    </div>
</template>

<script>
import CardImageView from './CardImageView.vue'

export default{
    components:{
        CardImageView
    },
    data(){
        return{
            dataImage : ['../../../../public/assets/images/Homepage/ImageSlider.png','../../../../public/assets/images/Homepage/ImageSlider.png','../../../../public/assets/images/Homepage/ImageSlider.png']

        }   
    },
    props:{
        dataImages: Array
    },
    methods: {
    shouldDisplay(index) {
      return index < 3;
    }
  }
}
</script>

<style scoped></style>