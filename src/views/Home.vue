<template>
    <div class="home">
        <div v-for="stream in streams" v-bind:key="stream.key">
            <Tail :stream="stream" />
        </div>
    </div>
</template>

<script lang="ts">
import Vue from "vue";
import Tail from "./Tail.vue";
import { query, Stream } from "./api";

export default Vue.extend({
    name: "Home",
    components: {
        Tail,
    },
    data(): {
        streams: Stream[];
    } {
        return {
            streams: [],
        };
    },
    async mounted(): Promise<void> {
        await query<{ streams: Stream[] }>("/v1/streams").then((reply) => {
            console.log(`reply: ${JSON.stringify(reply)}`);
            this.streams = reply.streams;
            return reply;
        });
    },
});
</script>
