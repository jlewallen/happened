<template>
    <div class="home">
        <div v-for="stream in streams" v-bind:key="stream.key">
            <router-link :to="{ name: 'stream', params: { key: stream.key } }">{{ stream.key }}</router-link>
        </div>
    </div>
</template>
<script lang="ts">
import Vue from "vue";
import { Stream } from "@/store/model";

export default Vue.extend({
    name: "Home",
    data(): {} {
        return {};
    },
    computed: {
        streams(): Stream[] {
            return this.$store.state.streams;
        },
    },
});
</script>
<style lang="scss">
.home {
    padding: 1em;
}
</style>
