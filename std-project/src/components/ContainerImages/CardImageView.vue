<template>
    <li class="">
        <img class=" w-full " @click="ShowPopUp" :src="urlImage" alt="">
        <div v-if="showPopUp"  @click="ShowPopUp" class=" z-50 w-full fixed top-0 left-0 min-h-screen flex justify-center items-center" style="background-color: rgba(0, 0, 0, 0.74);">
            <div class=" w-1/3">
                <img class=" w-full " :src="urlImage" alt="">
            </div>
        </div>
    </li>
</template>
<script>
export default {
    data() {
        return {
            showPopUp: false,
        }
    },
    methods:{
        ShowPopUp(){
            showPopUp =! showPopUp
        }
    },
    props: {
        urlImage: String
    },
    

}
</script>

<style scoped></style>