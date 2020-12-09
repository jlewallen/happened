<template>
    <div class="stream-view">
        <Header />
        <StreamViewer :stream="stream" v-if="stream" v-bind:key="stream.key" />
    </div>
</template>
<script lang="ts">
import Vue from "vue";
import Header from "./Header.vue";
import StreamViewer from "./StreamViewer.vue";
import { Stream } from "@/store/model";

export default Vue.extend({
    name: "Home",
    components: {
        Header,
        StreamViewer,
    },
    computed: {
        stream(): Stream | null {
            return this.$store.state.streams.find((stream: Stream) => stream.key == this.$route.params.key);
        },
    },
});
</script>
<style lang="scss">
.stream-view {
    padding: 1em;
}
</style>
