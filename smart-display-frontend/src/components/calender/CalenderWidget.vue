<template>
    <div id="calender-widget">
        <div id="widget">
            <div id="widget-header">
                <p id="header-text">{{ placeholders.header }}</p>
            </div>
            <div id="widget-content">
                <div id="calender-event" v-for="event in events" :key="event.id" >
                    <div id="event-date">
                        <p>{{ event.date }}</p>
                    </div>
                    <div id="event-name">
                        <p>{{ event.name }}</p>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "calender-widget",
    data() {
        return {
            placeholders: {
                header: "Calender"
            },
            events: [
                {
                    id: 0,
                    date: "10.11",
                    name: "Meeting Work"
                },
                {
                    id: 1,
                    date: "11.11",
                    name: "Meeting Uni"
                },
                {
                    id: 2,
                    date: "12.11",
                    name: "Arzt Termin"
                },
                {
                    id: 3,
                    date: "13.11",
                    name: "Nach Hause fahren"
                }
            ]
        }
    }
}
</script>

<style scoped>
#calender-widget {
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
    font-size: 1rem;
    margin: 0;
}
#calender-event {
    width: 100%;
    height: 2.5rem;
    margin-top: 0.5rem;
    
    text-align: left;
    line-height: 2.5rem;
}
#event-date p, #event-name p {
    margin: 0;
}
#event-date {
    width: 15%;
    height: 2.5rem;
    margin-left: 5%;
    display: inline-block;
    color: white;

}
#event-name {
    color: white;
    display: inline-block;
    width: 70%;
    margin-left: 5%;
    background-color: #452949;
}
#event-name p {
    margin-left: 5%;
}

</style>