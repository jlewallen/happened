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
        ControlPanel,
        StreamViewer,
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
    flex-direction: column;
    flex: 1;
}

::v-deep .lower .control-panel {
    height: 130px;
}

::v-deep #scrolling {
    overflow-x: scroll;
    overflow-y: scroll;
    width: 100vw;
    height: calc(100vh - 185px);
}
</style>
