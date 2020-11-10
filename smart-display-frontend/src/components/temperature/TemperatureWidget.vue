<template>
    <div id="temperature-widget">
        <div id="widget">
            <div id="temp-header">
                <p id="name">{{ placeholders.outside }}</p>
            </div>
            <div id="guage">
                <p id="updated-time">{{ outside.updated }}</p>
                <VueSvgGauge
                    :start-angle="-110"
                    :end-angle="110"
                    :value="outside.temperature"
                    :separator-step="10"
                    :min="0"
                    :max="40"
                    :gauge-color="[{ offset: 0, color: '#32a852'}, { offset: 100, color: '#f71640'}]"
                    :scale-interval="2"
                />
                <span id="label">
                    {{ outside.temperature + placeholders.celcius }}
                </span>
            </div>
        </div>
        <div id="widget">
            <div id="temp-header">
                <p id="name">{{ placeholders.inside }}</p>
            </div>
            <div id="guage">
                <p id="updated-time">{{ inside.updated }}</p>
                <VueSvgGauge
                    :start-angle="-110"
                    :end-angle="110"
                    :value="inside.temperature"
                    :separator-step="10"
                    :min="0"
                    :max="40"
                    :gauge-color="[{ offset: 0, color: '#32a852'}, { offset: 100, color: '#f71640'}]"
                    :scale-interval="2"
                />
                <span id="label">
                    {{ inside.temperature + placeholders.celcius }}
                </span>
            </div>
        </div>
    </div>
</template>

<script>
import { VueSvgGauge } from 'vue-svg-gauge'
import axios from 'axios';

export default {
    name: "Temperature",
    components: {
    VueSvgGauge,
    },
    created() {
        this.fetchTemperatures();
        setInterval(this.fetchTemperatures, 5 * 60 * 1000);
    },
    data() {
        return {
            placeholders: {
                outside: "Temperature - Outside",
                inside: "Temperature - Inside",
                celcius: "°C"
            },
            outside: {
                updated: "11:40",
                temperature: 10.1
            },
            inside: {
                updated: "11:40",
                temperature: 24.3
            },
        }
    },
    methods: {
        fetchTemperatures: function() {
            const backendUrl = process.env.VUE_APP_BACKEND_URL;
            const locationOutside = process.env.VUE_APP_LOCATION_OUTSIDE;
            const locationInside = process.env.VUE_APP_LOCATION_INSIDE;
            axios
                .get(`http://${backendUrl}temperature/${locationOutside}/latest`)
                .then((response) => {
                    this.outside.temperature = response.data.Temp;
                    this.outside.updated = this.convertTimestampToTime(response.data.Time);
                })
                .catch((response) => {
                    console.log(response)
                });
            axios
                .get(`http://${backendUrl}temperature/${locationInside}/latest`)
                .then((response) => {
                    this.inside.temperature = response.data.Temp;
                    this.inside.updated = this.convertTimestampToTime(response.data.Time);
                })
                .catch((response) => {
                    console.log(response)
                });
        },
        convertTimestampToTime: function(timestamp) {
            const date = new Date(timestamp);
            let hour = date.getHours() >= 10 ? date.getHours() : `0${date.getHours()}`;
            let minutes = date.getMinutes() >= 10 ? date.getMinutes() : `0${date.getMinutes()}`;
            return hour + ":" + minutes;
        }
    },
}
</script>

<style scoped>
#temperature-widget {
    width: 100%;
    text-align: center;
}
#widget {
    width: 16rem;
    height: 16rem;
    display: inline-block;
    background-color: #381c3c;
    border-radius: 1rem;
    margin-bottom: 1rem;
}
#temp-header {
    width: 100%;
    height: 2.5rem;
    line-height: 2.5rem;
    background-color: #301c34;
    border-top-left-radius: 1rem;
    border-top-right-radius: 1rem;
}
#name {
    color: white;
    font-size: 1rem;
    margin: 0;
}
#updated-time {
    position: absolute;
    width: 100%;
    margin-top: 0.5rem;
    color: white;
}
#guage {
    position: relative;
    height: 14rem;
    width: 14rem;
    margin-left:  1rem;
    margin-right:  1rem;
}
#label {
    color: white;
    position: relative;
    top: -40%;
    font-size: 1.7rem;
}
</style>