<template>
    <div id="datetime">
        <p id="time">{{ time }}</p>
        <p id="date">{{ date }}</p>
    </div>
</template>

<script>
export default {
    name: "Datetime",
    data() {
        return {
            time: "",
            date: ""
        }
    },
    created() {
        this.getNow();
        this.getDate();
        setInterval(this.getNow, 1000);
        setInterval(this.getDate, 1000);
    },
    methods: {
        getNow: function() {
            const today = new Date();
            let hours = today.getHours() >= 10 ? today.getHours() : `0${today.getHours()}`;
            let minutes = today.getMinutes() >= 10 ? today.getMinutes() : `0${today.getMinutes()}`;
            let seconds = today.getSeconds() >= 10 ? today.getSeconds() : `0${today.getSeconds()}`;
            const time = hours + ":" + minutes+ ":" + seconds;
            this.time = time;
        },
        getDate: function() {
            const today = new Date();

            const week = ['So', 'Mo', 'Di', 'Mi', 'Do', 'Fr', 'Sa'];
            const day = week[today.getDay()];

            let month = "";
            switch (today.getMonth() + 1) {
                case 1:
                    month = "Januar"
                    break;
                case 2:
                    month = "Februar"
                    break;
                case 3:
                    month = "März"
                    break;
                case 4:
                    month = "April"
                    break;
                case 5:
                    month = "Mai"
                    break;
                case 6:
                    month = "Juni"
                    break;
                case 7:
                    month = "Juli"
                    break;
                case 8:
                    month = "August"
                    break;
                case 9:
                    month = "September"
                    break;
                case 10:
                    month = "Oktober"
                    break;
                case 11:
                    month = "November"
                    break;
                case 12:
                    month = "Dezember"
                    break;
            }

            this.date = day + " " + today.getDate() + " " + month + " " + today.getFullYear();
        }
    }
}
</script>

<style scoped>
#datetime {
    margin-top: 4rem;
}
#time {
    color: white;
    font-size: 5rem;
    text-decoration: underline;
    margin: 0;
}
#date {
    color: white;
    font-size: 1.5rem;
    margin-top: -0.7rem;
}
</style>