<template>
    <footer>
        <div class="footer">
            <div class="map">
                <iframe
                    src="https://www.google.com/maps/embed?pb=!1m14!1m8!1m3!1d15612.986052326949!2d108.426631!3d11.957427!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x317113b07111e0b3%3A0x4d7fb786e7ac7e20!2zUXVheSBwaGltIMSQw6AgTOG6oXQgVEFNIFBST0RVQ1RJT04!5e0!3m2!1svi!2sus!4v1687840722907!5m2!1svi!2sus"
                    width="600" height="250" style="border:0;"  loading="lazy"
                    referrerpolicy="no-referrer-when-downgrade"></iframe>
            </div>
            <div class="about-us relative">
                <div class="context context-left">
                    <h1>TAM PRODUCTION</h1>
                    <p>Địa chỉ: 217/24/12 Ngô Quyền, Phường 6, Tp Đà Lạt </p>
                    <p>Hotline: <a href="tel:+0908747742"  style="color: #fff;"> 0908747742 (Mr Tâm)</a>
                    </p>
                    <p>Email: <a href="mailto:tamproduction102@gmail.com"
                            style="color: #fff;">tamproduction102@gmail.com</a>
                    </p>
                    <p>Chụp ảnh Đà Lạt, quay phim Đà Lạt</p>
                    <div class="line"></div>
                    <div class="icon">
                        <a href="https://www.facebook.com/ngothanhtam.photodalat" target="_blank">
                           
                        </a>
                        <a href="https://www.youtube.com/@quayphimdalat102" target="_blank">
                           

                        </a>
                        <a href="https://www.tiktok.com/@ngothanhtam.dalat" target="_blank">
                          
                        </a>
                        <a href="https://tamproduction.com/" target="_blank">
                           
                        </a>
                    </div>

                </div>
            </div>
            <div class="box-icon-copyright">
                <p>Copyright &copy; by Tam Production</p>
            </div>
        </div>
    </footer>
</template>

<style scoped>
.footer {
    position: relative;
    bottom: 0;
    overflow: hidden;
    width: 100%;
    color: #fff;
}

.about-us {
    background: url("./logo/pngkey.com-world-map-black-and-1707126.png"), #1a5b34e0;
    background-repeat: no-repeat;
    background-position: center;
    width: 100%;
    height: 450px;
    border-top: 2px solid #fff;
}

.about-us p {
    font-size: 15px;
    text-align: center;
    margin-bottom: 10px;
}

.about-us h1 {
    font-size: 3.5em;
    font-weight: bold;
    text-align: center;
}

.context-left {
    transform: translateY(35%);
}

.line {
    width: 45%;
    border-bottom: 1px solid #fff;
    margin: auto;
}

.box-icon-copyright {
    display: flex;
    height: 2.5em;
    background-color: #fff;
    justify-content: center;
}

.box-icon-copyright p {
    color: #858585d4;
    line-height: 4;
    font-size: 10px;
    text-align: left;
}


.map iframe {
    width: 100%;
}

.map {
    line-height: 0;
    border-top: 2px solid #fff;
}
</style>