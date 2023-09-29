<template>
    <li class="relative" :style="{background : 'url('+ srcUrl +') center/cover no-repeat' }" style="padding-top: 50%;">
        <!-- <img class="w-full object-cover" :src="srcUrl" alt="" style="height: 90vh;"> -->
    </li>
</template>

<script>
export default {
    props: {
        srcUrl: String
    }
}
</script>

<style scoped>

</style>