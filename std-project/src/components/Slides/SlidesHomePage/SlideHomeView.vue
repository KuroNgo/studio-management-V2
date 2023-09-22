<template>
    <div class="slides-inner relative">
        <!-- <button @click="ClickChangeSlides(1)"><font-awesome-icon id="left" class="a font-awesome-icon-left"
                 style="color: #4678ce;" /></button> -->
        <ul class="cardousel scroll-smooth grid grid-flow-col mb-14 gap-4 overflow-hidden ">
            <ImageHome class="card snap-start" v-for="(x) in dataImage" draggable="flase" :srcUrl="x" />
        </ul>
        <!-- <button @click="ClickChangeSlides(-1)"><font-awesome-icon id="right" class="a font-awesome-icon-right"
                 style="color: #316cd3;" /></button> -->
    </div>
</template>

<script>
import ImageHome from './ImageHome.vue'
export default {
    data() {
        return {
            dataImage : ['../../../../public/assets/images/BusLogin.png','../../../../public/assets/images/viewinbus.png']
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
}

.dragging {
    cursor: grab;
    user-select: none;
    scroll-behavior: auto;
    scroll-snap-type: none;
}

.font-awesome-icon-left {
    left: -80px;
    position: absolute;
    font-size: 42px;
    top: 50%;
}

.font-awesome-icon-right {
    right: -80px;
    position: absolute;
    font-size: 42px;
    top: 50%;
}
</style>