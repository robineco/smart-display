<template>
    <div id="news-widget">
        <div id="widget">
            <div id="widget-header">
                <p id="header-text">{{ placeholders.header }}</p>
            </div>
            <div id="widget-content">
                <p>{{ pages[currentHeadline].content }}</p>
            </div>
            <div id="icons">
                <span v-for="page in pages" :key="page.id">
                    <span v-if="page.isActive" class="dot active"></span>
                    <span v-else class="dot"></span>
                </span>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    name: "news-widget",
    created() {
        this.fetchNews();
        setInterval(this.switchHeadline, 10000)
    },
    data() {
        return {
            placeholders: {
                header: "News Headlines"
            },
            currentHeadline: 0,
            pages: [
                {
                    id: 0,
                    isActive: true,
                    content: ""
                },
                {
                    id: 1,
                    isActive: false,
                    content: ""
                },
                {
                    id: 2,
                    isActive: false,
                    content: ""
                }
            ],
            dotColors: {
                white: "dot active",
                back: "dot",
            },
            apiKey: process.env.VUE_APP_NEWS_API_KEY
        }
    },
    methods: {
        fetchNews: function() {
           const url = "https://newsapi.org/v2/top-headlines?country=de&apiKey=" + this.apiKey

            axios
                .get(url)
                .then((response) => {
                    const headlines = response.data.articles.slice(0, 3);
                    console.log(headlines);
                    for (let i = 0; i < headlines.length; i++) {
                        let description = headlines[i].description;
                        description = description.split(".");
                        this.pages[i].content = description[0] + ".";
                    }
                    
                })
                 .catch((response) => {
                    console.log(response)
                });
        },
        switchHeadline: function() {
           let nextHeadline = this.currentHeadline + 1;

           if(nextHeadline > 2) {
               nextHeadline = 0;
           }
           this.pages[nextHeadline].isActive = true
           this.pages[this.currentHeadline].isActive = false;
           this.currentHeadline = nextHeadline;
        }
    },
}
</script>

<style scoped>
#news-widget {
    width: 100%;
    text-align: center;
}
#widget {
    position: relative;
    width: 16rem;
    height: 16rem;
    display: inline-block;
    background-color: #381c3c;
    border-radius: 1rem;
}
#widget-header {
    width: 100%;
    height: 2.5rem;
    line-height: 2.5rem;
    background-color: #301c34;
    border-top-left-radius: 1rem;
    border-top-right-radius: 1rem;
}
#header-text {
    color: white;
    margin: 0;
}
#widget-content {
    width: 90%;
    height: 70%;
    display: inline-block;
    margin-top: 0.5rem;
    border-radius: 1rem;
    background-color: #452949;
    color: white;
    text-align: left;
}
#widget-content p{
    margin-left: 0.5rem;
}
#icons {
    margin-top: 0.3rem;
}
.dot {
    display: inline-block;
    height: 0.7rem;
    width: 0.7rem;
    background-image: url("../../assets/images/circle-black.png");
    background-size: 0.7rem;
    margin-right: 0.3rem;
}
.active {
    background-image: url("../../assets/images/circle-white.png");
}
</style>