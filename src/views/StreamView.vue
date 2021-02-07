<template>
    <div class="stream-view">
        <Header :expanded="expanded" @expanded="onExpandedToggle" />
        <div v-if="stream" v-bind:key="stream.key" class="lower">
            <ControlPanel :stream="stream" v-if="expanded" />
            <StreamViewer :stream="stream" v-bind:class="{ expanded: expanded }" @changed="onChanged" @scrolled="onScrolled" />
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
    data(): {
        expanded: boolean;
        bottom: boolean;
    } {
        return {
            expanded: false,
            bottom: true,
        };
    },
    computed: {
        stream(): Stream | null {
            return this.$store.state.streams.find((stream: Stream) => stream.key == this.$route.params.key);
        },
    },
    methods: {
        onExpandedToggle(expanded: boolean): void {
            this.expanded = expanded;
            this.onChanged();
        },
        onScrolled(args: { bottom: boolean }): void {
            if (this.bottom != args.bottom) {
                console.log("bottom", args.bottom);
            }
            this.bottom = args.bottom;
        },
        onChanged(): void {
            if (this.bottom) {
                this.$nextTick(() => {
                    const el = this.$el.querySelector("#scrolling");
                    el.scrollTop = el.scrollHeight;
                });
            }
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
    height: calc(100vh - 55px);
}

::v-deep #scrolling.expanded {
    height: calc(100vh - 185px);
}
</style>
