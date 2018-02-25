<template>
    <div>
        <date-picker
            v-model="day"
            format="yyyy-MM-dd"
            lang="en"
            confirm
        >
        </date-picker>
        <a class="button is-rounded">Submit</a>
    </div>
</template>

<script>
export default {
    props: ['type', 'placeholder'],
    computed: {

    }
}
</script>

<style>

</style>
