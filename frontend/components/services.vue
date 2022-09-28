<script>
import { defineNuxtComponent } from '#app';

export default defineNuxtComponent({

    data: () => ({
        services: []
    }),
    created() {
        this.fetchServices()
    },
    methods: {
        fetchServices() {
            fetch('http://localhost:1337/api/services').then
                (response => response.json())
                .then(json => {
                    this.services = json.data
                })
        },
    }
})
</script>
<template>
    <div class="p-1 border-4 rounded-md">
        <div v-for="service in services">
            <div class="w-full h-12 my-3 bg-green-500 flex justify-between items-center rounded-sm ">
                <div class="pl-4 font-medium text-2xl text-stone-50">{{service.attributes.name}}</div>
                <div class=" pr-4 font-normal text-xl text-slate-700">{{service.attributes.status}}</div>
            </div>
        </div>
    </div>
</template>
