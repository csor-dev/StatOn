<script>
import { defineNuxtComponent } from '#app';
import moment from 'moment'

export default defineNuxtComponent({

    data: () => ({
        incidents: []
    }),
    created() {
        this.fetchIncidents()
    },
    methods: {
        convertDate(value) {
            return moment(String(value)).format('MMM Do YY, h:mm a')
        },
        fetchIncidents() {
            fetch('http://localhost:1337/api/incidents').then
                (response => response.json())
                .then(json => {
                    this.incidents = json.data
                })
        },

    }
})
</script>
<template>
    <div class="w-full flex-col min-h-fit">
        <div v-if="incidents.length>0" class="w-full flex-col h-fit ">
            <h2 class="text-2xl font-normal py-4 mb-4 border-b-4 w-fit border-neutral-300">Incidents</h2>
            <div v-for="incident in incidents" :key="`incident-id-${incident.id}`">
                <div class="border-2 border-slate-200 pl-2 py-3">
                    <h2 class="text-xl pb-1">{{incident.attributes.issue_name}} - {{incident.attributes.status}}</h2>
                    <p class="text-lg text-neutral-700 pl-2">{{incident.attributes.incident_description}}</p>
                    <p class="text-neutral-500 font-normal pl-2">{{convertDate(incident.attributes.issue_start)}}</p>
                </div>
            </div>
        </div>
        <div v-else="incidents.length>0"
            class=" h-20 bg-lime-200 text-2lg font-medium flex w-full justify-center items-center">
            Loading</div>
    </div>
</template>
<!-- <template>
    <div>
        <h2 class=" text-2xl text-slate-800 font-medium py-2 my-2 border-b-2 border-blue-500 w-fit">Incidents</h2>
        <div class="incidents w-full h-64 bg-violet-200 rounded-md flex justify-center">
            <div v-if="error">
                <p class=" text-2xl w-full flex justify-center1"> No Incidents reported </p>
            </div>
<ul>
    <li v-for="incident in incidents" :key="incident.id">
        {{ incident.name }}
    </li>
</ul>
        </div>
    </div>
</template> -->

<!-- <script>

export default {
    name: 'App',
    data() {
        return {
            incidents: [],
            error: null,
            headers: { 'Content-Type': 'application/json' }
        }
    },
    methods: {
        parseJSON: function (resp) {
            return (resp.json ? resp.json() : resp);
        },
        checkStatus: function (resp) {
            if (resp.status >= 200 && resp.status < 300) {
                return resp;
            }
            return this.parseJSON(resp).then((resp) => {
                throw resp;
            });
        }
    },
    async mounted() {
        try {
            const response = await fetch("http://localhost:1337/api/incidents", {
                method: 'GET',
                headers: this.headers,
            }).then(this.checkStatus)
                .then(this.parseJSON);
            this.incidents = response
        } catch (error) {
            this.error = error
        }
    }
}
</script>
 -->
