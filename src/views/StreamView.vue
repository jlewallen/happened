<template>
    <div class="stream-view">
        <Header />
        <div v-if="stream" v-bind:key="stream.key" class="lower">
            <ControlPanel :stream="stream" />
            <StreamViewer :stream="stream" />
        </div>
    </div>
</template>
<script lang="ts">
import Vue from "vue";
import Header from "./Header.vue";
import StreamViewer from "./StreamViewer.vue";
import ControlPanel from "./ControlPanel.vue";
import { Stream } from "@/store/model";

export default Vue.extend({
    name: "Home",
    components: {
        Header,
        StreamViewer,
        ControlPanel,
    },
    computed: {
        stream(): Stream | null {
            return this.$store.state.streams.find((stream: Stream) => stream.key == this.$route.params.key);
        },
    },
});
</script>
<style lang="scss" scoped>
.stream-view {
    display: flex;
    min-height: 100vh;
    flex-direction: column;
}

.stream-view .lower {
    display: flex;
    flex-direction: row;
    flex: 1;
}

::v-deep .lower .control-panel {
    flex: 0 0 200px;
}

::v-deep .lower .stream {
    height: calc(100vh - 55px);
    width: 100vh;
    flex: 1;
}
</style>
