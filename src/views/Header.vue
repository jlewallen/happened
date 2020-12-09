<template>
    <div class="header">
        <div v-for="stream in streams" v-bind:key="stream.key" :class="{ stream: true, selected: selected(stream) }">
            <router-link :to="{ name: 'stream', params: { key: stream.key } }">{{ stream.key }}</router-link>
        </div>
    </div>
</template>
<script lang="ts">
import Vue from "vue";
import { Stream } from "@/store/model";

export default Vue.extend({
    name: "Header",
    components: {},
    computed: {
        streams(): Stream[] {
            return this.$store.state.streams;
        },
    },
    methods: {
        selected(stream: Stream): boolean {
            return stream.key == this.$route.params.key;
        },
    },
});
</script>
<style lang="scss" scoped>
.header {
    display: flex;
    justify-content: default;

    div {
        padding: 0.5em;
        margin: 0.5em;
        border: 1px black solid;
        border-radius: 4px;
        background: #6495ed;
        a {
            color: black;
        }
    }

    div.selected {
        font-weight: bold;
        background: #5f9ea0;
    }
}
</style>
