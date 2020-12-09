<template>
    <div class="stream">
        <StreamViewer :stream="stream" v-if="stream" />
    </div>
</template>
<script lang="ts">
import Vue from "vue";
import StreamViewer from "./StreamViewer.vue";
import { Stream } from "@/store/model";

export default Vue.extend({
    name: "Home",
    components: {
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
.stream {
    padding: 1em;
}
</style>
