<template>
    <div class=" bg-zinc-50 max-w-3xl w-4/5 p-4 flex-col justify-center h-full min-h-screen min-w-fit">
        <img class=" h-56 w-full rounded-sm"
            src="https://images.pexels.com/photos/531880/pexels-photo-531880.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
            alt="">
        <h1 class=" text-4xl text-slate-800 font-medium w-full flex justify-center border-b-4 border-neutral-300">
            StatOps
        </h1>
        <banner />
        <services />
        <div class="flex w-full justify-end p-2" v-if="false">
            <button
                class="font-medium text-2xl border-2 px-6 py-3 rounded-md border-green-500 text-slate-700">Subscribe</button>
        </div>
        <incidents />
    </div>
</template>

<script setup>

</script>
