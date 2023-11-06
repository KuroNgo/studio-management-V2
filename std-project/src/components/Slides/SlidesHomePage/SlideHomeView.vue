<template>
    <div class="slides-inner relative">
        <!-- <button @click="ClickChangeSlides(1)"><font-awesome-icon id="left" class="a font-awesome-icon-left"
                 style="color: #4678ce;" /></button> -->

        <svg @click="ClickChangeSlides(1)" class="font-awesome-icon-left" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
            <path id="Icon_awesome-arrow-alt-circle-left" data-name="Icon awesome-arrow-alt-circle-left"
                d="M12.563,24.563a12,12,0,1,1,12-12A12,12,0,0,1,12.563,24.563Zm5.613-14.129H12.563V7a.581.581,0,0,0-.992-.411L6.04,12.151a.575.575,0,0,0,0,.818l5.531,5.56a.581.581,0,0,0,.992-.411V14.692h5.613a.582.582,0,0,0,.581-.581v-3.1A.582.582,0,0,0,18.175,10.433Z"
                transform="translate(-0.563 -0.563)" fill="white" />
        </svg>
        
        <ul class="cardousel scroll-smooth grid grid-flow-col mb-14 gap-4 overflow-hidden ">
            <ImageHome class="card snap-start" v-for="(x) in dataImage" draggable="flase" :srcUrl="x" />
        </ul>

        <svg @click="ClickChangeSlides(-1)" class="font-awesome-icon-right" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
            <path id="Icon_awesome-arrow-alt-circle-right" data-name="Icon awesome-arrow-alt-circle-right"
                d="M12.563.563a12,12,0,1,1-12,12A12,12,0,0,1,12.563.563ZM6.95,14.692h5.613v3.431a.581.581,0,0,0,.992.411l5.531-5.56a.575.575,0,0,0,0-.818L13.554,6.592A.581.581,0,0,0,12.563,7v3.431H6.95a.582.582,0,0,0-.581.581v3.1A.582.582,0,0,0,6.95,14.692Z"
                transform="translate(-0.563 -0.563)" fill="white"/>
        </svg>

        <!-- <button @click="ClickChangeSlides(-1)"><font-awesome-icon id="right" class="a "
                 style="color: #316cd3;" /></button> -->
    </div>
</template>

<script>
import ImageHome from './ImageHome.vue'
export default {
    data() {
        return {
            dataImage: ['../../../../public/assets/images/Homepage/ImageSlider.png', '../../../../public/assets/images/Homepage/ImageSlider.png']
        }
    },
    mounted() {
        this.MouseMoveChangeSlides()
    },
    components: {
        ImageHome
    },
    methods: {
        //Click Button Change Slides
        ClickChangeSlides(x) {
            const cardousel = document.querySelector('.cardousel')
            const firstCardWidth = cardousel.querySelector('.card').offsetWidth;
            if (x == 1) {
                cardousel.scrollLeft += -firstCardWidth;

            }
            else if (x == -1) {
                cardousel.scrollLeft += firstCardWidth;

            }

        },

        //Mouse Move Change Slides
        MouseMoveChangeSlides() {
            const cardousel = document.querySelector('.cardousel')
            let isDragging = false, startX, startScrollLeft;
            const dragStart = (e) => {
                isDragging = true;
                cardousel.classList.add("dragging");
                startX = e.pageX;
                startScrollLeft = cardousel.scrollLeft;
            }
            const dragStop = () => {
                isDragging = false;
                cardousel.classList.remove("dragging");
            }
            const dragging = (e) => {
                if (!isDragging) return;
                cardousel.scrollLeft = startScrollLeft - (e.pageX - startX);
            }
            cardousel.addEventListener("mousemove", dragging);
            cardousel.addEventListener("mousedown", dragStart);
            cardousel.addEventListener("mouseup", dragStop);
        }
    },
}
</script>

<style scoped>
.cardousel {
    grid-auto-columns: 100%;
    /* overflow-x: auto; */
    scroll-snap-type: x mandatory;
    scrollbar-width: 0;
    height: 90vh;
    transition: 2s ease-in;
}

.dragging {
    user-select: none;
    scroll-behavior: auto;
    scroll-snap-type: none;
}

.font-awesome-icon-left {
    left: 40px;
    position: absolute;
    font-size: 42px;
    top: 50%;

}

.font-awesome-icon-right {
    right: 40px;
    position: absolute;
    font-size: 42px;
    top: 50%;


}
</style>